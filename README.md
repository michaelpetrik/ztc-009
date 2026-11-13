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
| [`tools/`](tools) | vývojáři | validátor v jazyce Go |

## Jak použít ZTČ-009 s jazykovým modelem

1. Zkopírujte obsah souboru `llm/ztc-009.llm.md`.
2. Vložte obsah do system promptu.
3. Napište zadání.

Režimy:

- Normální zadání: model napíše text v ZTČ-009.
- Slovo „kontrola": model vrátí tabulku chyb a opravený text.
- Slovo „překlad": model přepíše váš text do ZTČ-009.

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
