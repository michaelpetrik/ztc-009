# Místní profil ZTČ-009

Při vytvoření profilu vyplňte skutečné hodnoty. Nevyplněný profil není aktivní nastavení.
Relativní cesty pracovních souborů jsou vůči projektu; balíčkové cesty jsou vůči `package_root`.
Tento profil je místní preference, nikoli součást příspěvku do repozitáře.

```yaml
enabled: false
scope_confirmed: false
scope: project
topics: []
exclusions: []
package_root: ""
dictionaries: []
local_dictionaries: []
dictionary_reviews: []
offer_contributions: true
declined_uploads: []
```

- `enabled`: aktivace stylu v potvrzeném rozsahu.
- `scope_confirmed`: uživatel potvrdil témata a rozsah; prázdný seznam témat neznamená všechna témata.
- `scope`: projekt nebo hostitelem podporovaný uživatelský rozsah, který uživatel výslovně zvolil.
- `topics`, `exclusions`: uživatelem zvolené oblasti použití a výjimky.
- `package_root`: skutečné umístění balíčku. Při jeho přesunu opravte cesty.
- `dictionaries`: cesty k vybraným balíčkovým CSV souborům.
- `local_dictionaries`: cesty k místním CSV souborům s ověřenými hesly.
- `dictionary_reviews`: cesta návrhu, stav revize a případné nejistoty; oddělené od sloupce CSV `status`.
- `offer_contributions`: zda lze nabídnout příspěvek; hodnota `true` není souhlas s uploadem.
- `declined_uploads`: cesta konkrétního exportu, identifikace jeho obsahu a cílový repozitář při odmítnutí.

Pro identifikaci exportu použijte dostupný hash nebo jiný jednoznačný místní záznam jeho přesného obsahu.
Záznam odmítnutí neobsahuje tokeny ani soukromé příklady. Nový obsah automaticky neznamená souhlas.
Souhlas s dokončeným uploadem nepoužívejte jako trvalé povolení dalších uploadů.
