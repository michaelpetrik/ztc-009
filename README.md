# ZTČ-009 — Zjednodušená technická čeština

ZTČ-009 je řízený jazyk pro technickou dokumentaci v češtině.
ZTČ-009 má stejnou strukturu jako ASD-STE100 Simplified Technical English (Issue 9).
ZTČ-009 má 9 sekcí a 53 pravidel.

> **POZNÁMKA**
> ZTČ-009 není oficiální produkt organizace ASD. Pravidla, vysvětlení a příklady jsou nový text.
> Originál ASD-STE100 získáte zdarma na <https://www.asd-ste100.org>. Podrobnosti jsou v souboru [NOTICE](NOTICE).

## Obsah

| Soubor | Pro koho | Popis |
| --- | --- | --- |
| [`spec/ztc-009.md`](spec/ztc-009.md) | lidé | úplná specifikace s vysvětlením a příklady |
| [`llm/ztc-009.llm.md`](llm/ztc-009.llm.md) | jazykové modely | system prompt, pravidla jako kontrolní seznam |
| [`slovnik/ztc-009.csv`](slovnik/ztc-009.csv) | lidé a programy | slovník schválených a neschválených slov |
| [`slovnik/software-development.csv`](slovnik/software-development.csv) | lidé a programy | frontend, backend, databáze a infrastruktura |
| [`tools/`](tools) | vývojáři | validátor v jazyce Go |

## Jak použít ZTČ-009 s jazykovým modelem

1. Zkopírujte obsah souboru `llm/ztc-009.llm.md`.
2. Vložte obsah do system promptu.
3. Napište zadání.

Režimy:

- Normální zadání: model napíše text v ZTČ-009.
- Slovo „kontrola": model vrátí tabulku chyb a opravený text.
- Slovo „překlad": model přepíše váš text do ZTČ-009.

## Pokyny pro agenty: slovník pro další obory

Při tvorbě obsahu ověřte, zda příslušný [slovník](slovnik/README.md#obsah-a-tvorba) obsahuje potřebné výrazy a jejich významy.
Pokud obor uživatele není pokrytý nebo je pokrytý jen částečně, nabídněte společnou tvorbu oborového slovníku.

Při tvorbě postupujte podle [pravidel slovníku](slovnik/RULEBOOK.md) a [návodu pro nový obor](slovnik/GUIDELINES.md).
Aktuální pokrytí uvádí [přehled slovníku](slovnik/README.md#obsah-a-tvorba), včetně sady [IT a vývoj softwaru](slovnik/IT.md).

1. S uživatelem určete obor a účel textu.
2. Při práci na konkrétním textu sbírejte chybějící výrazy. Slovník doplňujte postupně podle potřeby.
3. Nejprve ověřte, zda lze použít existující schválené slovo se stejným významem.
4. Pro nový výraz navrhněte základní tvar, slovní druh, jeden význam a příklad použití v oboru.
   U slovesa doplňte vid. Údaje připravte podle [formátu slovníku](slovnik/README.md).
5. Význam a příklad ověřte s uživatelem. Neověřené návrhy označte v pracovních poznámkách.
   Dohodu s uživatelem nepovažujte za schválení hesla v tomto projektu.
6. Dohodnuté výrazy používejte v dalším obsahu jednotně. Při změně významu upravte také pracovní slovník.
7. Nabídněte uživateli [návod k přispění](CONTRIBUTING.md#pravidla-pro-slovník).
   Pomozte připravit issue „Nové slovo“ s oborem, významem a příkladem.
   Pro pull request připravte změnu CSV a spusťte kontrolu podle [CONTRIBUTING.md](CONTRIBUTING.md#pull-request).

Návrhy zveřejněte až se souhlasem uživatele. Z příkladů odstraňte důvěrné údaje.

## Jak přispět

Návrhy jsou vítané. Postup je v souboru [CONTRIBUTING.md](CONTRIBUTING.md).

- Nové slovo do slovníku: [založte issue „Nové slovo"](../../issues/new?template=nove-slovo.yml).
- Změna pravidla: [založte issue „Změna pravidla"](../../issues/new?template=zmena-pravidla.yml).
- Chyba v textu: [založte issue „Chyba"](../../issues/new?template=chyba.yml).
- Diskuse: použijte záložku Discussions.

## Licence

| Část | Licence |
| --- | --- |
| Specifikace, LLM profil, dokumentace | [CC BY-SA 4.0](LICENSE) |
| Slovník (`slovnik/`) | [CC0 1.0](slovnik/LICENSE) |
| Nástroje (`tools/`) | [MIT](tools/LICENSE) |

## Verze

Aktuální verze: **1.0.0-rc1**. Změny jsou v souboru [CHANGELOG.md](CHANGELOG.md).
