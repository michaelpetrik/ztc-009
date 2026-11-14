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
| [`llm/ztc-009.llm.md`](llm/ztc-009.llm.md) | jazykové modely | profil s výběrem rozsahu, pravidla jako kontrolní seznam |
| [`slovnik/ztc-009.csv`](slovnik/ztc-009.csv) | lidé a programy | slovník schválených a neschválených slov |
| [`slovnik/software-development.csv`](slovnik/software-development.csv) | lidé a programy | frontend, backend, databáze a infrastruktura |
| [`tools/`](tools) | vývojáři | validátor v jazyce Go |

## Začněte v Claude nebo ChatGPT

Použijte [instrukci pro běžný chat](docs/CHAT.md#rychlé-použití). Stačí ji vložit do zprávy.
Plugin, GitHub ani terminál nejsou potřeba. Model nejprve zjistí profesi a témata, pokud je ještě nezná.
Pro všechna pravidla vložte [úplný LLM profil](llm/ztc-009.llm.md) a potom napište „Odpovídej mi v TZČ-009“.
Pouhý název stylu bez pravidel nezaručuje znalost ZTČ-009. Triviální jednověté odpovědi zůstávají přirozené.

## Instalace podle prostředí

[Claude, Cowork, Claude Code, ChatGPT, Codex a Pi](docs/HARNESSES.md): podporované cesty, příkazy a stav ověření.
Pro Pi je připravená jednopříkazová cesta, nikoli ověřené univerzální instalační tlačítko.
Veřejný GitHub při auditu ještě neobsahoval nový balíček; instalační příkazy z Gitu čekají na publikaci.

## Volitelný plugin pro agenty

Použijte [instalační návod](docs/INSTALL.md) a [skill ztc-install](skills/ztc-install/SKILL.md).
Agent nejprve zjistí profesi a témata, pro která chcete ZTČ používat.
Pro chybějící obor vytvoří pracovní slovník a nabídne jeho zveřejnění přes PR až po vašem souhlasu.
Před publikací ověří GitHub connector, `gh` nebo jiný nástroj schopný vytvořit PR.
Příspěvek pomůže dalším instalacím ve stejném oboru. GitHub není podmínkou místního použití.

Balíček obsahuje skills a hooky pro Claude Code a Codex.
Pro jiné agenty použijte skills a trvalé instrukce podle možností hostitele.
Triviální dotazy s úplnou jednovětou odpovědí používají běžný styl.

## Použití bez pluginu

Pro jednorázovou úpravu textu dejte modelu [LLM profil](llm/ztc-009.llm.md) a příslušné [slovníky](slovnik/README.md).
Výslovně určete text, který má model upravit. Pro použití v celém chatu postupujte podle [návodu pro Claude a ChatGPT](docs/CHAT.md).

## Tvorba slovníků s agentem

Použijte [pravidla](slovnik/RULEBOOK.md) a [postup pro nový obor](slovnik/GUIDELINES.md).
Agent postupně sbírá potřebné výrazy, ověří s vámi jejich význam a zachová jednotné názvy.
Pracovní slovník zůstává místní. Jeden ucelený profesní slovník lze nabídnout jedním PR.

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
