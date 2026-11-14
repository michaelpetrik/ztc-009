# Pravidla pro tvorbu slovníku

Tato pravidla platí pro nové obory i změny existujících hesel.
Doplňují [specifikaci](../spec/ztc-009.md) a [formát CSV](README.md#sloupce).

## Rozsah a uložení

- Určete obor, čtenáře a typy textů, které má slovník pokrýt.
- Popište pokryté i nepokryté oblasti v dokumentu oboru.
- Obecná hesla ukládejte do `ztc-009.csv`, softwarová do `software-development.csv`. Dokument oboru popisuje rozsah.
- Nový CSV soubor zaregistrujte v `tools/cmd/ztcvalidate/main.go`. CI používá tento validátor.
- Náhrady musí být schválené ve stejném CSV souboru. Shody a střety významů mezi soubory ověřte při revizi.
- Technické názvy a technická slovesa posuzujte podle sekce 1 specifikace. Chybějící heslo samo o sobě není zákaz slova.

## Heslo a význam

- Nejdříve hledejte existující heslo se stejným významem. Nevytvářejte zbytečná synonyma.
- Použijte základní tvar, jeden slovní druh a jeden přesně vymezený význam.
- Existujícímu heslu nepřidávejte druhý význam pro nový obor. Použijte přesnější technický název.
- U slovesa vyplňte vid. Dokonavý a nedokonavý tvar zapisujte jako samostatná hesla.
- V příkazech používejte dokonavý vid. V zákazech používejte nedokonavý vid.
- Víceslovný technický název zapisujte celý. Názvy produktů, příkazy a identifikátory kódu nepřekládejte svévolně.

## Stav a příklady

- `approved` označuje povolený výraz s vyplněným významem a příkladem `ok`.
- `unapproved` označuje nevhodný výraz s náhradou a příklady `ok` a `bad`. Neoznačuje návrh čekající na posouzení.
- Náhrada musí být schválené lemma nebo odkaz na existující pravidlo `@X.Y`.
- Náhradu lze použít jen tam, kde zachová význam. Kontextovou náhradu vysvětlete v návrhu.
- Příklady pište vlastními slovy a ověřte jejich odbornou správnost. Odstraňte důvěrné údaje.
- `en_ste` vyplňte jen pro ověřené heslo STE. Běžný anglický překlad sem nepatří.
- Návrhy čekající na posouzení držte v pracovních poznámkách nebo issue, mimo schválený CSV slovník.

## Přijetí změny

- Význam a příklady ověřte s člověkem, který zná obor.
- Dohoda s uživatelem platí pro pracovní slovník. Přijetí do projektu potvrzuje správce sloučením PR.
- Uveďte obor, důvod, zdroj ověření a případné střety se stávajícími významy.
- Slovníková data musí být slučitelná s CC0. Nekopírujte cizí definice ani obsah ASD-STE100.
- Spusťte validátor a kontrolu Markdownu. Kontroly nenahrazují odborné posouzení významu.
- Změnu zaznamenejte do nevydané části CHANGELOGu. Číslo vydání měňte až při přípravě vydání.
