# Audit obejití kontroly odpovědí

Posouzený commit: `3541b7c`.
Výsledek: kontrolor vynucuje část formálních pravidel u spolupracujícího modelu.
Není izolovanou ani neobejitelnou hranicí mezi modelem a uživatelem.
Audit níže zaznamenává původní stav. Následné opravy popisuje [uzavření auditu](2026-11-15-enforcement-fixes.md).

## Jak kontrola pracuje

1. `UserPromptSubmit` načte stav podle ID relace a předá modelu příkaz kontroloru.
2. Model sestaví odpověď a může zvolit výjimku pro daný tah.
3. `Stop` zkontroluje poslední odpověď, pokud je kontrola pro tah aktivní.
4. Při chybě vrátí `decision: block`, konkrétní chyby a požadavek na opravu.
5. Po třech neúspěšných opravách ukončí pokračování s oznámením selhání.

Mechanismus neopravuje obsah sám. Opravu provádí model.
Stav `continue: false` neodstraňuje chybný text, který již hostitel zobrazil.

## Potvrzené možnosti obejití

### P1: Kontrolovaný model ovládá výjimky a stav

Místo: `tools/cmd/ztcguard/main.go:145`, `:174` a `:207`.

Příkaz `outside` neověřuje téma ani souhlas uživatele. Režim `off` vypne relaci.
Model zná potřebné ID a číslo tahu, protože je dostává v kontextu hooku.
Před zahájením oprav tak může zvolit výjimku a vrátit libovolný obsah.

Reprodukce v testu: aktivní stav, `handle("outside", ...)`, potom `Stop` s textem `CI; např. chyba.`.
Výsledek: prázdné rozhodnutí bez chyby, tedy povolené ukončení.

Ani zákaz přepnutí během opravy není autorizační hranice.
Akce `prompt` přijímá JSON ze standardního vstupu bez ověření, že jej poslal hostitel.
Podvržený nový prompt resetoval počet oprav ze tří na nulu; následná akce `off` prošla.
Agent se zápisem do stejného uživatelského účtu má také přístup k souboru stavu a kódu kontroloru.
Práva `0600` chrání proti jiným účtům, nikoli proti tomuto agentovi.

Náprava: rozhodnutí uživatele a stav spravovat mimo oprávnění kontrolovaného agenta.
Výjimky nesmí schvalovat tentýž model bez nezávislého ověření.
Samotný tajný token uložený ve stejném čitelném prostředí tuto hranici nevytvoří.

### P1: Vlastní prózu lze vydávat za kód nebo citaci

Místo: `tools/cmd/ztcguard/check.go:25`, `:36`, `:39` a `:42`.

Kontrolor neumí rozlišit skutečně převzatou citaci od nově vytvořené odpovědi.
Následující vstupy byly skutečně vyhodnoceny:

| Vstup | Nálezy |
| --- | --- |
| `CI; např. chyba.` | Středník, obecná zkratka a nevysvětlené CI |
| Tentýž text v inline kódu | Žádné |
| `> CI; např. chyba.` | Žádné |
| `„CI; např. chyba.“` | Žádné |

Zachování legitimního kódu je potřebné. Chybí ale ověření původu chráněného obsahu.
Náprava: chránit jen konkrétní úseky doložené zadáním nebo výstupem nástroje.
Novou vysvětlující prózu kontrolovat i při pokusu o změnu jejího obalu.

### P1: Spouštěč neověřuje odpověď kontroloru

Místo: `hooks/ztc-guard.sh:8`.

Spouštěč důvěřuje návratovému kódu nula. Neověřuje JSON ani identitu/verzi binárního souboru.
Reprodukce:

```sh
ZTC_GUARD_BIN=/usr/bin/true sh hooks/ztc-guard.sh stop
```

Výsledek: návratový kód nula a pouze prázdný řádek, bez rozhodnutí nebo hlášení selhání.
Chybná konfigurace nebo záměna kontroloru tedy může kontrolu tiše vyřadit.
Přesný výsledek v klientovi závisí na zpracování prázdného výstupu hostitelem.

Náprava: validovat odpověď a verzi protokolu; prázdný nebo neplatný výstup považovat za selhání.
Pro odolnost proti agentovi musí být i spouštěč a binární soubor mimo jeho zápisová oprávnění.

### P2: Parser a počítání slov mají mezery

Místo: `tools/cmd/ztcguard/check.go:14`, `:16` a `:79`.

- Znak `#` na začátku řádku vypne kontrolu délky i bez mezery požadované pro běžný nadpis Markdownu.
  Test `#` následovaný čtyřiceti slovy prošel.
- Celá závorka se počítá jako jedno slovo. Osmdesát slov uvnitř závorky prošlo.
  Počítání závorky odpovídá konvenci projektu; chybí ověření, že jde o krátké povolené vysvětlení.
- Zkratka `např.` s vloženým U+200B mezi `a` a `p` prošla.
  Kontrolor nenormalizuje neviditelné znaky a kontroluje jen doslovné podoby zkratek.

Náprava: správně parsovat Markdown, normalizovat kontrolní kopii textu a ověřovat rozsah výjimek.
Normalizace nesmí svévolně změnit původní kód nebo technické identifikátory.

### P2: Úspěch není potvrzení jazykové správnosti

Test `Ventil musí být uzavřen.` prošel bez nálezu.
Přitom obsahuje konstrukci zakázanou LLM profilem i neschválený tvar hesla.
Kontrolor nekontroluje českou morfologii, všechna slovníková hesla, význam ani všech 53 pravidel.
Sedm technických zkratek je vložených přímo v kódu; nové oborové zkratky automaticky nepřebírá.

Jde o známou hranici rozsahu, nikoli důkaz, že regex může garantovat úplnou shodu s jazykem.
Náprava: rozlišovat formální validaci, jazykové posouzení a odborné ověření.

## Omezení hostitele a neověřené scénáře

- `Stop` běží po generování; již streamovaný chybný návrh nezmizí.
- Limit tří oprav ukončí cyklus s chybou. Nevytvoří automaticky vyhovující odpověď.
- Zakázané hooky, nepodporovaný hostitel nebo ignorované rozhodnutí kontrolu vyřadí.
- Zpracování timeoutu a výpadku procesu v přihlášených klientech nebylo živě ověřeno.
- Větvení `Stop` neporovnává hostitelské `turn_id` s aktuálním uloženým tahem.
  Zpožděná událost může použít stav jiného tahu; skutečné pořadí takových událostí nebylo reprodukováno v klientovi.

## Ověření auditu

Kontrolní vstup bez obalu byl odmítnut třemi pravidly.
Obaly, falešný nadpis, dlouhá závorka, neviditelný znak, výjimka `outside` a reset oprav byly ověřeny cíleným Go testem.
Prázdný úspěch spouštěče byl ověřen skutečným shellovým příkazem.
Reprodukce používaly syntetický obsah; produkční stav relací nebyl měněn.
Dočasný testovací soubor byl odstraněn. Audit neprováděl opravy kódu.

## Doporučený další krok

Nejprve oddělit ovládání kontroly od oprávnění modelu.
Potom opravit parser a validaci protokolu, rozšířit regresní testy a provést živý test obou klientů.
Pro záruku, že uživatel nikdy neuvidí neplatný text, je nutná kontrola před zobrazením v hostiteli.
Samotný následný `Stop` hook tuto vlastnost neposkytne.
