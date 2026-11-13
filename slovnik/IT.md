# IT a vývoj softwaru

Tato počáteční sada doplňuje společný slovník o 24 hesel.
Definice a příklady jsou v [CSV](ztc-009.csv). Sada je návrh pro odborné posouzení před přijetím PR.
Nejde o úplný slovník IT.

## Pokrytí

- Soubory a programy: software, zdrojový kód, soubor, adresář.
- Správa verzí: repozitář, větev, commit, požadavek na sloučení.
- Programování: závislost, knihovna, rozhraní, funkce, proměnná, parametr, argument.
- Data: databáze, záloha.
- Ověření a provoz: automatizovaný test, sestavení, nasazení, testovací prostředí, produkční prostředí, spustit, spouštět.

## Hranice významů

- `větev` označuje historii změn, nikoli část podmínky v programu.
- `funkce` označuje část programu, nikoli vlastnost produktu.
- `parametr` je součást definice funkce. `argument` je hodnota předaná při volání.
- `rozhraní` zde označuje komunikaci softwarových součástí. Pro ovládání aplikace použijte název „uživatelské rozhraní“.
- Stávající `nainstalovat` má mechanický význam: dát na správné místo a připevnit.
  Jeho schválení se automaticky nevztahuje na instalaci softwaru. Technická slovesa posuzujte podle sekce 1 specifikace.
- Názvy Git příkazů a identifikátory kódu zachovejte přesně. Heslo `commit` je podstatné jméno, nikoli schválení slovesa „commitnout“.

## Další rozšíření

Síťové protokoly, bezpečnost, cloud, kontejnery, mobilní vývoj a strojové učení zatím nemají samostatné sady hesel.
Přidávejte je podle skutečných textů s uživatelem podle [postupu](GUIDELINES.md) a [pravidel](RULEBOOK.md).
Sloupec `en_ste` zůstává u nových hesel prázdný: shoda s hesly STE nebyla ověřena.
