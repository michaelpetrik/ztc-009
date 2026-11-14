# Postup pro nový oborový slovník

Použijte tento postup společně s [pravidly](RULEBOOK.md).
Pro instalaci stylu nejprve použijte [instalační skill](../skills/ztc-install/SKILL.md).

1. S uživatelem určete profesi, témata, čtenáře a účel odpovědí. Použijte již potvrzené informace.
2. Z [přehledu](README.md#obsah-a-tvorba) vyberte základní a vhodné oborové slovníky.
   Kontrolujte význam, nejen shodu slova. Znovu negenerujte pokrytou oblast.
3. Pro nepokrytý obor připravte malou použitelnou sadu v `.ztc/dictionaries/<obor>.csv`.
   Vycházejte z běžných úkolů profese nebo z ukázky textu. Nevyžadujte pevný počet hesel.
4. Uveďte základní tvar, slovní druh, jeden význam a příklad. U slovesa doplňte vid.
   Běžné technické zkratky přidejte vedle plného názvu. Nejistoty zapište do pracovních poznámek.
5. Ukažte slovník uživateli k potvrzení významů. Generování ani úspěšná validace nejsou odborným schválením.
   Bez potvrzení nepředstírejte schválení. Při psaní lze dál používat technické názvy podle specifikace.
6. Uložte potvrzený slovník do místního profilu a používejte jej jednotně. Nová hesla přidávejte podle potřeby.
7. Připravte sanitizované soubory k příspěvku a ukažte cílový veřejný repozitář.
   Vysvětlete, že sdílení pomůže příštím instalacím stejného oboru. Připomeňte licenci CC0.
8. Ověřte dostupný GitHub connector, `gh` nebo jiný nástroj podle [instalačního návodu](../docs/INSTALL.md#dobrovolný-příspěvek-přes-pr).
   Před prvním uploadem se zeptejte na souhlas s přesným obsahem a cílem.
9. Se souhlasem připravte jeden PR pro ucelený slovník. Do `slovnik/<obor>.csv` vložte pouze určená hesla.
   Doplňte přehled, rozsah, hranice významů a CHANGELOG. Předchozí issue není povinné.
10. Bez souhlasu zůstaňte u místního použití. Bez publikačního nástroje předejte soubory a návod k ručnímu PR.

## Kontroly

Z adresáře `tools` zkontrolujte pracovní CSV, pokud máte Go:

```sh
go run ./cmd/ztcvalidate -dictionary /absolutni/cesta/k/pracovnimu.csv
```

Pro příspěvek zkontrolujte z adresáře `tools` celý repozitář:

```sh
go test ./...
go run ./cmd/ztcvalidate -root ..
```

Markdown zkontrolujte z kořene repozitáře:

```sh
npx --yes markdownlint-cli2 "**/*.md"
```

Chybějící vývojové nástroje neblokují instalaci ani místní používání slovníku.
V PR uveďte přesně, které kontroly proběhly. CI musí projít před sloučením.
Validátor kontroluje strukturu, povinná pole, odkazy a duplicity uvnitř CSV.
Významy, odborné příklady, zkratky a střety mezi slovníky posuzuje člověk.

## Záznam pracovního návrhu

Zapište obor, výrazy, důvod doplnění, nejistoty a výsledek ověření uživatelem.
Stav příspěvku držte mimo sloupec `status`.
Nevkládejte do veřejného PR místní profil, pracovní dokumenty ani neveřejné příklady.
