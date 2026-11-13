# CLAUDE.md — pokyny pro AI agenty

## Projekt

ZTČ-009 = český řízený jazyk pro technickou dokumentaci. Struktura 1:1 s ASD-STE100 Issue 9: 9 sekcí, 53 pravidel (14/2/7/5/5/6/3/7/4), doporučení D-1 až D-8 (= STE GR-1 až GR-8), slovník.

## Soubory

| Cesta | Účel |
|---|---|
| `spec/ztc-009.md` | specifikace pro lidi (zdroj pravdy) |
| `llm/ztc-009.llm.md` | system prompt pro LLM, musí odpovídat spec |
| `slovnik/ztc-009.csv` | slovník, schéma v `slovnik/README.md` |
| `tools/` | Go validátor `ztcvalidate` (slovník, číslování, shoda verzí) |

## Tvrdá pravidla

- Nikdy nekopíruj ani nepřekládej text z dokumentu ASD-STE100 (copyright ASD). Přebírá se jen struktura, čísla pravidel a číselné limity. Text, vysvětlení a příklady piš nově.
- Číslování 1.1–9.4 je pevné. Nová čísla nepřidávej. České úpravy patří do odstavce **[CZ]** u existujícího pravidla.
- Změna spec ⇒ uprav i `llm/ztc-009.llm.md`. Změna verze ⇒ obě místa + `CHANGELOG.md`.
- Prózu v repozitáři (spec, README, CONTRIBUTING, šablony) piš v ZTČ-009.
- Nástroje: Go. Python jen na jednorázové prototypy, ne do repozitáře.

## Kontrola před commitem

```sh
cd tools && gofmt -l . && go vet ./... && go test ./... && go run ./cmd/ztcvalidate -root ..
markdownlint-cli2 "**/*.md"
```
