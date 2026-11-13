# Slovník ZTČ-009

Soubor `ztc-009.csv` obsahuje slovník ZTČ-009.
Licence slovníku je CC0 1.0. Slovník můžete použít bez omezení.

## Obsah a tvorba

Projekt používá jeden společný CSV slovník.
Původní základ obsahuje 27 hesel pro obecné technické postupy, převážně s příklady z mechaniky.
Počáteční sada [IT a vývoj softwaru](IT.md) přidává 24 hesel. Celkem je v CSV 51 hesel.
Samostatné CSV slovníky pro jednotlivé obory zatím nejsou podporované.

- [Pravidla pro tvorbu slovníku](RULEBOOK.md): významy, stavy, příklady a přijetí změn.
- [Postup pro nový obor](GUIDELINES.md): společná práce agenta a uživatele, příprava příspěvku a kontroly.
- [IT a vývoj softwaru](IT.md): rozsah počáteční sady a hranice významů.

## Sloupce

| Sloupec | Popis | Povinné |
| --- | --- | --- |
| `lemma` | základní tvar slova | ano |
| `pos` | slovní druh: `V`, `N`, `ADJ`, `ADV`, `PREP`, `CONJ`, `PRON`, `NUM` | ano |
| `vid` | `pf` (dokonavý) nebo `impf` (nedokonavý) | jen pro `V` |
| `status` | `approved` nebo `unapproved` | ano |
| `meaning` | jeden schválený význam | jen pro `approved` |
| `replace_with` | náhrady oddělené znakem `\|` | jen pro `unapproved` |
| `en_ste` | odpovídající anglické heslo STE | ne |
| `ok` | správný příklad | ano |
| `bad` | nesprávný příklad | jen pro `unapproved` |

## Náhrady

Každá náhrada v `replace_with` musí být jedna z těchto hodnot:

- lemma, které má ve slovníku `status=approved`,
- odkaz na pravidlo ve tvaru `@X.Y` (například `@3.7`). Odkaz použijte, když náhrada závisí na kontextu.

## Kontrola

CI zkontroluje slovník při každé změně. Kontrolu můžete spustit také lokálně:

```sh
cd tools && go run ./cmd/ztcvalidate -root ..
```
