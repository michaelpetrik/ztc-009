# Pravidla pro tvorbu slovníku

Tato pravidla platí pro nové obory i změny existujících hesel.
Doplňují [specifikaci](../spec/ztc-009.md) a [formát CSV](README.md#sloupce).

## Rozsah a uložení

- Určete obor, čtenáře a typy textů, které má slovník pokrýt.
- Popište pokryté i nepokryté oblasti v dokumentu oboru.
- Obecná hesla ukládejte do `ztc-009.csv`, softwarová do `software-development.csv`. Dokument oboru popisuje rozsah.
- Nový publikovaný slovník uložte jako `slovnik/<obor>.csv`. Validátor načte všechny CSV v tomto adresáři.
  Pro distribuci přidejte CSV a průvodní dokument také do seznamu `files` v `tools/cmd/ztcpack/main.go`.
- Pracovní slovníky ukládejte mimo publikovaný adresář, standardně do místního `.ztc/dictionaries/`.
- Náhrady musí být schválené ve stejném CSV souboru. Shody a střety významů mezi soubory ověřte při revizi.
- Technické názvy a technická slovesa posuzujte podle sekce 1 specifikace. Chybějící heslo samo o sobě není zákaz slova.

## Heslo a význam

- Nejdříve hledejte existující heslo se stejným významem. Nevytvářejte zbytečná synonyma.
- Použijte základní tvar, jeden slovní druh a jeden přesně vymezený význam.
- Existujícímu heslu nepřidávejte druhý význam pro nový obor. Použijte přesnější technický název.
- U slovesa vyplňte vid. Dokonavý a nedokonavý tvar zapisujte jako samostatná hesla.
- V příkazech používejte dokonavý vid. V zákazech používejte nedokonavý vid.
- Víceslovný technický název zapisujte celý. Názvy produktů, příkazy a identifikátory kódu nepřekládejte svévolně.

## Technické zkratky

- Běžnou oborovou zkratku přidejte jako samostatné heslo vedle plného názvu.
- Do `lemma` napište samotnou zkratku. Zachovejte obvyklou velikost písmen.
- Do `meaning` napište plný název a jeho význam. Význam musí odpovídat heslu plného názvu.
- Do `ok` vložte příklad ve tvaru `CI (průběžná integrace)` nebo `VM (virtuální stroj)`.
- Při prvním použití v samostatném textu vysvětlete zkratku tímto způsobem. Dále lze používat samotnou zkratku.
- Víceznačnou zkratku vymezte pro daný obor. Nespojujte odlišné významy do jednoho hesla.
- Technické zkratky nemění zákaz běžných zkratek podle doporučení D-6.

## Stav a příklady

- `approved` označuje povolený výraz s vyplněným významem a příkladem `ok`.
- `unapproved` označuje nevhodný výraz s náhradou a příklady `ok` a `bad`. Neoznačuje návrh čekající na posouzení.
- Náhrada musí být schválené lemma nebo odkaz na existující pravidlo `@X.Y`.
- Náhradu lze použít jen tam, kde zachová význam. Kontextovou náhradu vysvětlete v návrhu.
- Příklady pište vlastními slovy a ověřte jejich odbornou správnost. Odstraňte důvěrné údaje.
- `en_ste` vyplňte jen pro ověřené heslo STE. Běžný anglický překlad sem nepatří.
- `status` popisuje povolené použití slova, nikoli stav příspěvku. Neoznačuje odbornou certifikaci.
- Pracovní CSV smí navrhovat `approved`. Do potvrzení uživatelem se heslo považuje za neověřený návrh.
- Po potvrzení lze pracovní slovník místně používat. Publikace nastává až přijetím PR správcem.
- Stav návrhu a způsob odborného ověření evidujte mimo CSV, v místním profilu a popisu PR.
- V kombinaci slovníků ověřte jedinečný význam a slovní druh každého lemmatu ručně.
  Validátor kontroluje pouze duplicitu kombinace lemma, slovní druh a vid uvnitř jednoho CSV.

## Přijetí změny

- Význam a příklady ověřte s člověkem, který zná obor.
- Dohoda s uživatelem platí pro pracovní slovník. Neopravňuje k jeho nahrání.
- Před uploadem ukažte sanitizované soubory, cíl a licenci CC0. Vyžádejte souhlas s tímto konkrétním příspěvkem.
- Jeden oborový slovník může být jeden PR bez předchozího issue. Přijetí potvrzuje správce sloučením PR.
- Uveďte obor, důvod, zdroj ověření a případné střety se stávajícími významy.
- Slovníková data musí být slučitelná s CC0. Nekopírujte cizí definice ani obsah ASD-STE100.
- Spusťte validátor a kontrolu Markdownu. Kontroly nenahrazují odborné posouzení významu.
- Změnu zaznamenejte do nevydané části CHANGELOGu. Číslo vydání měňte až při přípravě vydání.
