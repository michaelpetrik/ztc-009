# Seznam změn

Formát: [Keep a Changelog](https://keepachangelog.com/cs/1.1.0/). Verze: [SemVer](https://semver.org/lang/cs/).

## [Nevydáno]

### Přidáno

- Pokyny pro agenty k postupné tvorbě oborového slovníku s uživatelem.
- Pravidla tvorby slovníku a postup pro nové obory.
- Počáteční sada IT a vývoj softwaru: 24 hesel ve společném CSV slovníku.

## [1.0.0-rc1] — 2026-11-02

### Přidáno

- Specifikace pro lidi (`spec/ztc-009.md`): 9 sekcí, 53 pravidel, doporučení D-1 až D-8.
- LLM profil (`llm/ztc-009.llm.md`).
- Slovník (`slovnik/ztc-009.csv`): 27 hesel.
- Validátor v jazyce Go (`tools/`): kontrola slovníku, číslování pravidel a shody verzí.
- CI: markdownlint, kontrola odkazů, testy a validace.
