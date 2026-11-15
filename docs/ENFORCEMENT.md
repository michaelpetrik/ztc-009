# Kontrola odpovědí ZTČ-009

Claude Code a lokální Codex mohou kontrolovat hotovou odpověď pomocí `Stop` hooku.
Kontrolor `ztcguard` při chybě vrátí `decision: block` a konkrétní seznam porušení.
Hostitel potom pokračuje a model opraví odpověď. Nejde jen o připomínku pravidel.

[Diagram kontroly](ztc-enforcement.svg).

## Instalace kontroloru

Pro sestavení potřebujete Go podle `tools/go.mod`. Kontrolor používá jen standardní knihovnu.
Z adresáře `tools` spusťte:

```sh
mkdir -p "$HOME/.local/bin"
go build -o "$HOME/.local/bin/ztcguard" ./cmd/ztcguard
```

Sestavený soubor při běhu Go nepotřebuje. Nová verze zdrojů vyžaduje nové sestavení.
Jinou cestu ke spustitelnému souboru lze nastavit proměnnou `ZTC_GUARD_BIN` v prostředí hostitele.
Plugin s hooky bez tohoto souboru zastaví tah a oznámí nefunkční kontrolu.
Neschovávejte tuto chybu přepnutím na instrukční režim bez vědomí uživatele.

Po instalaci pluginu schvalte jeho hooky v hostiteli a otevřete novou relaci.
Ověřte `UserPromptSubmit` i `Stop`. Instalace souborů sama nepotvrzuje důvěru ani běh hooku.
Claude chat, cloudové režimy bez shellových hooků a současné rozšíření Pi tuto kontrolu neposkytují.
V těchto prostředích je dostupný jen textový režim; plugin tam nesmí slibovat vynucení.

## Stav a rozsah

`UserPromptSubmit` načte stav podle skutečného `session_id` hostitele.
Nová relace převezme pouze `enabled: true` a `scope_confirmed: true` z projektového `.ztc/profile.md`, pokud existuje.
Jiné umístění profilu zatím kontrolor automaticky nečte; instalační agent musí tuto mez oznámit.
Příkaz `/ztc` nebo vyvolaný skill přepne stav relace. Dvojí zpracování v témže tahu jej nepřepne zpět.

Stav se ukládá pod uživatelský adresář cache do `ztc-009/sessions/<hash>.json` s právy `0600`.
Adresář používá `0700`. Obsahuje stav, číslo tahu a počty oprav; neobsahuje uživatelské zadání ani odpověď.
Stav zůstává při obnovení stejné relace i při shrnutí kontextu. Nové ID znamená nezávislou relaci.
Soubory stavu se automaticky nemažou, aby obnovená relace neztratila nastavení.
Nepřidávají se do projektu ani distribučního ZIPu.

Model dostane stav, nikoli oprávnění vypnout kontrolu. Změnové akce CLI `off`, `outside` a další jsou odmítnuté.
Přepnutí zpracuje `UserPromptSubmit` pouze z přesného uživatelského příkazu.
Pokud klient rozbalí slash příkaz před hookem, použijte běžnou zprávu `ztc toggle`.

Na prvním řádku uživatelského zadání lze použít:

- `/ztc skip`: výjimka pro jednu odpověď, například pro téma mimo zvolený obor.
- `/ztc once`: jednorázový převod se zapnutou kontrolou i při vypnuté relaci.
- `/ztc code`: povolení nového kódu v této odpovědi; okolní próza zůstává kontrolovaná.

Podporovaný textový zápis je také `ztc skip`, `ztc once` a `ztc code` bez lomítka.
Samostatné `/ztc on` a `/ztc off` nastaví relaci jednoznačně, bez obracení stavu.
Automatická výjimka pro trivialitu platí pouze pro rozpoznané pozdravy a poděkování.
Model nemůže svévolně schválit jiný triviální dotaz ani téma mimo rozsah.
To je záměrně přísnější než instrukční režim běžného chatu.

Výjimky platí jen pro aktuální tah. Opravné pokračování je zachovává včetně rozpočtu tří oprav.
`Stop` s nesouhlasícím hostitelským ID tahu je odmítnutý.
Bez ID tahu, například u Claude, závisí přiřazení na pořadí událostí dané relace.

## Co kontrolor blokuje

- Středník v próze podle pravidla 8.1.
- Obecné zkratky `tj.`, `tzn.`, `např.`, `atd.`, `apod.`, `resp.`, `cca`, `tzv.`, `max.`, `min.`.
- Více než 20 slov ve větě. Jde o přísnější společný limit pro postup i popis.
- Více než šest vět v odstavci.
- První použití `CI`, `VM`, `FE`, `BE`, `UI`, `DB` nebo `PR` bez odpovídajícího plného názvu v závorce.
- Neuzavřený blok kódu nebo neplatný formát výjimky `trivial`.
- Konstrukci „musí být“ následovanou rozpoznaným krátkým trpným příčestím.

Kontrolor odstraní neviditelné formátovací znaky z kontrolní kopie a dekóduje HTML entity.
Nadpisy a závorky neobcházejí limit: jejich obsah se počítá stejně jako jiná próza.
Samotný kódový nebo citační obal není důvod k vynechání kontroly.
Chráněné jsou pouze stejné úseky kódu nebo citací doložené aktuálním zadáním uživatele.
Kontrolor uchovává jejich otisky, nikoli jejich obsah. Upravený úsek shodu ztratí.
Nový kód potřebuje uživatelské povolení `/ztc code`. Model jej nesmí udělit sám.
Odkazy kontroluje podle viditelného textu; jejich URL nemění.
Nejde o úplný parser Markdownu ani morfologický analyzátor češtiny.

Přesný význam, činný rod, vid sloves, odbornou pravdivost a všech 53 pravidel nelze tímto kontrolorem certifikovat.
Stejně tak nekontroluje všechny skloňované tvary slovníku ani nové oborové zkratky.
Úspěch kontroly znamená splnění uvedené měřitelné části, nikoli důkaz úplné shody s jazykovou specifikací.

## Opravy a selhání

Neplatný výstup vyvolá nejvýše tři opravné pokusy.
Opakované pokračování vyvolané kontrolorem nesmí resetovat počet pokusů.
Po třetí neúspěšné opravě hook vrátí `continue: false` a viditelně označí neověřenou odpověď.
Chybějící text, stav, binární soubor nebo poškozený stav také znamenají oznámené selhání, nikoli úspěšnou kontrolu.
Spouštěč ověřuje neprázdný JSON, povolené schéma rozhodnutí a verzi výstupního protokolu.
Program `true`, prázdný výstup ani neplatný JSON nejsou úspěšná kontrola.

`Stop` se spouští po generování. Již streamovaný text nemůže odebrat z obrazovky.
Kontrola nezaručuje, že uživatel nikdy neuvidí chybný návrh, a nenahrazuje bezpečnostní hranici hostitele.
Zakázaný hook, pád procesu nebo hostitelem ignorovaný výsledek ruší záruku tohoto opravného cyklu.

## Hranice důvěry

Plugin není bezpečnostní izolace vůči procesu se stejnými systémovými právy.
Agent s neomezeným shellem může změnit binární soubor nebo stav, nahradit spouštěč či podvrhnout celý hook vstup.
Ověření protokolu odhaluje chybný výstup; neověřuje kryptograficky identitu procesu ani původ události.
Takový zásah nelze spolehlivě zakázat dalším pravidlem nebo tokenem uloženým ve stejném účtu.

Pro odolnost proti takovému agentovi musí hostitel provozovat kontrolor a stav mimo jeho zápisová oprávnění.
Musí také posílat události odděleným důvěryhodným kanálem a validovat text před zobrazením.
Tato změna hostitele není součástí přenosného pluginu a není označená za hotovou.
V podporovaném provozu jsou hook vstupy důvěryhodné a model nesmí upravovat řídicí soubory.

## Ověření instalace

1. Zapněte ZTČ a potvrďte témata. Zkontrolujte skutečný stav kontroloru.
2. Zadejte test s výslovně požadovaným středníkem v próze. V událostech ověřte zamítnutí a opravu.
3. Ověřte triviální dotaz, vyloučené téma, vypnutí a druhou nezávislou relaci.
4. Ověřte chybu při nedostupném kontroloru. Nesmí se zobrazit úspěšná kontrola.

Automatické testy simulují skutečný protokol hooků, chybný a opravený výstup, stav i limit opakování.
Neprokazují běh hooků v přihlášeném účtu uživatele; tento průchod musí proběhnout v cílovém hostiteli.

[Claude Stop](https://code.claude.com/docs/en/hooks#stop),
[Codex Stop](https://learn.chatgpt.com/docs/hooks#stop).
