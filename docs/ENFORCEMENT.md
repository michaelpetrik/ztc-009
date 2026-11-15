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

Agent dostane příkaz s přesným ID a číslem tahu. Staré číslo tahu kontrolor odmítne.
Aktivní relace má kontrolu zapnutou jako výchozí stav každého tahu.
Agent smí před odpovědí vybrat jen tyto odůvodněné výjimky:

- `--action outside`: téma mimo potvrzený rozsah nebo samotný instalační rozhovor.
- `--action trivial`: skutečně triviální dotaz. Výstup musí mít jednu větu, jeden řádek a nejvýše 25 slov.
- `--action enforce`: výslovný převod konkrétního textu, i při vypnuté relaci.

Instalační skill může po potvrzení rozsahu použít `--action on` nebo `--action off` pro jednoznačnou změnu relace.
Samotné přepsání profilu již běžící relaci nezmění.
Výjimky `outside`, `trivial` a `enforce` nemění stav celé relace. Další uživatelský tah výjimku zruší.
Rozpoznání tématu a odborné triviality stále provádí model; není to deterministická jazyková klasifikace.
Během oprav nelze příkazem změnit výjimku nebo přepnout relaci.
Uživatel může opravný cyklus přerušit a v novém tahu kontrolu vypnout.

## Co kontrolor blokuje

- Středník v próze podle pravidla 8.1.
- Obecné zkratky `tj.`, `tzn.`, `např.`, `atd.`, `apod.`, `resp.`, `cca`, `tzv.`, `max.`, `min.`.
- Více než 20 slov ve větě. Jde o přísnější společný limit pro postup i popis.
- Více než šest vět v odstavci.
- První použití `CI`, `VM`, `FE`, `BE`, `UI`, `DB` nebo `PR` bez odpovídajícího plného názvu v závorce.
- Neuzavřený blok kódu nebo neplatný formát výjimky `trivial`.

Kontrolor zpracuje běžný Markdown. Kód, blokové a doslovné citace nekontroluje jako vlastní prózu.
Odkazy kontroluje podle viditelného textu; jejich URL nemění. Závorku počítá jako jedno slovo.
Nadpisy nejsou věty. Odstavce odděluje prázdný řádek, položky seznamu posuzuje odděleně.
Nejde o úplný parser Markdownu ani morfologický analyzátor češtiny.

Přesný význam, činný rod, vid sloves, odbornou pravdivost a všech 53 pravidel nelze tímto kontrolorem certifikovat.
Stejně tak nekontroluje všechny skloňované tvary slovníku ani nové oborové zkratky.
Úspěch kontroly znamená splnění uvedené měřitelné části, nikoli důkaz úplné shody s jazykovou specifikací.

## Opravy a selhání

Neplatný výstup vyvolá nejvýše tři opravné pokusy.
Opakované pokračování vyvolané kontrolorem nesmí resetovat počet pokusů.
Po třetí neúspěšné opravě hook vrátí `continue: false` a viditelně označí neověřenou odpověď.
Chybějící text, stav, binární soubor nebo poškozený stav také znamenají oznámené selhání, nikoli úspěšnou kontrolu.

`Stop` se spouští po generování. Již streamovaný text nemůže odebrat z obrazovky.
Kontrola nezaručuje, že uživatel nikdy neuvidí chybný návrh, a nenahrazuje bezpečnostní hranici hostitele.
Zakázaný hook, pád procesu nebo hostitelem ignorovaný výsledek ruší záruku tohoto opravného cyklu.

## Ověření instalace

1. Zapněte ZTČ a potvrďte témata. Zkontrolujte skutečný stav kontroloru.
2. Zadejte test s výslovně požadovaným středníkem v próze. V událostech ověřte zamítnutí a opravu.
3. Ověřte triviální dotaz, vyloučené téma, vypnutí a druhou nezávislou relaci.
4. Ověřte chybu při nedostupném kontroloru. Nesmí se zobrazit úspěšná kontrola.

Automatické testy simulují skutečný protokol hooků, chybný a opravený výstup, stav i limit opakování.
Neprokazují běh hooků v přihlášeném účtu uživatele; tento průchod musí proběhnout v cílovém hostiteli.

[Claude Stop](https://code.claude.com/docs/en/hooks#stop),
[Codex Stop](https://learn.chatgpt.com/docs/hooks#stop).
