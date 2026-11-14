# Postup pro nový oborový slovník

Použijte tento postup společně s [pravidly](RULEBOOK.md).

1. S uživatelem určete obor, čtenáře a jeden skutečný text, na kterém budete pracovat.
2. Porovnejte výrazy z textu s [přehledem slovníků](README.md#obsah-a-tvorba).
   Kontrolujte význam, nejen shodu slova.
3. Zapište chybějící výrazy do pracovního seznamu. Začněte malou sadou potřebnou pro daný text.
4. Pro každý výraz připravte základní tvar, slovní druh, význam a příklad.
   U slovesa doplňte vid. Zapište také obor a důvod přidání.
5. Požádejte uživatele o ověření významu a příkladu. Nejasné výrazy ponechte jako návrhy.
6. Použijte dohodnuté výrazy při úpravě textu. Nové případy postupně přidávejte do pracovního seznamu.
7. Nabídněte [příspěvek do projektu](../CONTRIBUTING.md).
   Se souhlasem uživatele založte issue „Nové slovo“. Pro každý návrh uveďte obor a výsledek ověření.
8. Pro PR doplňte hesla do příslušného CSV slovníku a aktualizujte dokument oboru.
   Dodržte přesné pořadí sloupců. Pole s čárkou uzavřete do dvojitých uvozovek.
9. Z kořene repozitáře spusťte kontroly:

   ```sh
   cd tools && go test ./... && go run ./cmd/ztcvalidate -root ..
   ```

   Markdown zkontrolujte z kořene repozitáře:

   ```sh
   npx --yes markdownlint-cli2 "**/*.md"
   ```

10. V PR uveďte rozsah, příklady a výsledek kontrol. Počkejte na posouzení správce.

## Příklad pracovního návrhu

- Obor: vývoj softwaru.
- Výraz: `zdrojový kód`.
- Slovní druh: `N`.
- Význam: zápis programu v programovacím jazyce určený k úpravám vývojářem.
- Příklad: „Uložte zdrojový kód do repozitáře.“
- Důvod: přesné rozlišení od spustitelného souboru.
- Ověření: doplňte, kdo význam posoudil a na jakém textu.
- Stav návrhu: čeká na odborné posouzení. Tento stav nepatří do sloupce `status`.
