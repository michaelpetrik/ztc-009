# Ověření pluginu a opraveného postupu

Datum: 2026-11-14.
Podklady: [syntéza](2026-11-14-synthesis.md), [kritická revize](2026-11-14-review.md) a nezávislá simulace sedmi scénářů.

## Spuštěné kontroly

- Go: formátování, `go vet ./...`, `go test -race ./...` a validace celého repozitáře prošly.
- Testy ověřují nový oborový CSV bez registrace, povinný základ, samostatný pracovní CSV a návratové kódy.
- Testy odmítají neexistující pravidla a přijímají platné mezní odkazy.
- Testy hooku ověřují JSON manifesty, cestu s mezerami, statický výstup nezávislý na promptu a nezměněné soubory.
- Chybějící kořen balíčku nebo odpovědní skill vede k tichému ukončení hooku bez chyby.
- Oba skills prošly `quick_validate.py`; celý Markdown prošel lintem.
- `claude plugin validate .` prošlo. Upozornění: kořenový CLAUDE.md není kontext pluginu, chybí katalogová ikona a marketplace instalační řádek.
- CLAUDE.md zůstává pro vývoj repozitáře. Plugin používá skills. Ikona ani marketplace nejsou nutné pro lokální balíček.

## Nezávislá simulace chování

1. Nová instalace: první krok je otázka na profesi, témata a výjimky.
2. Elektrikář bez slovníku a bez souhlasu s uploadem: místní CSV a profil, žádný upload ani instalace gh.
3. Aktivní softwarový profil a „Kolik je 2 + 2?“: přirozená jedna věta bez přísného profilu.
4. Migrace databáze bez výpadku: odborný postup v ZTČ, nikoli umělé zkrácení na jednu větu.
5. Vypnutý profil a výslovný převod do ZTČ: jednorázový převod při zachování doporučení oproti povinnosti.
6. Connector pouze pro čtení a nová důvěrná data mimo schválený export: žádný upload, očištěný místní export.
7. Host bez shellu a trvalých instrukcí: použití skills jen v relaci, bez slibu obnovení.

Simulace zjistila nadbytečnou přípravu příspěvku při předchozím odmítnutí všech uploadů.
Oprava přidala časné přeskočení celé nabídky i kontroly GitHubu a místní preferenci `offer_contributions`.

## Meze důkazů

Simulace není živý test každého hostitele ani záruka rozhodování modelu.
Neproběhla instalace do uživatelského prostředí, ověření další relace, schválení hooků ani upload na GitHub.
Nebyl vytvořen PR ani publikován plugin do katalogu. Odborné definice slovníků čekají na lidskou revizi.
