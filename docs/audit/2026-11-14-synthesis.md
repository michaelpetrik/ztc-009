# Syntéza auditu ZTČ-009

- Datum: 2026-11-14.
- Podklady: dodané závěry analytiků A, B a C, kontrola repozitáře a dokumentace hostů.
- Výchozí commit: `0e05720056bd626262675ff73c557d4fc810b954`.
- Nálezy a vstupní Go kontroly popisují stav před souběžnou implementací rodičovského agenta.
- Výsledek: přijmout níže uvedené minimum; funkčnost pluginu zatím není ověřena.
- Tento audit mění pouze tento dokument.
  Každou dokončenou změnu včetně této syntézy commituje rodičovský agent.

## Prioritizované nálezy

- **P1 — rozsah odpovědí a zachování významu (B).**
  `llm/ztc-009.llm.md:10–11` vyžaduje ZTČ pro každou větu a nadřazuje pravidla přirozenosti.
  To odporuje požadovanému výběru témat, výlukám a výjimce pro triviální odpověď.
  Řádek 115 navíc mění „měl by“ na „musí“, tedy doporučení na povinnost.
  Rozhodnout o použití ZTČ před načtením přísného profilu; odstranit význam měnící substituci.
  Doporučení zůstane doporučením, například „Doporučujeme restartovat službu.“

- **P1 — další obor unikne CI (A).**
  `tools/cmd/ztcvalidate/main.go:20–25` registruje přesně dva CSV soubory.
  Třetí soubor nezakazuje, ale vůbec jej nekontroluje.
  `slovnik/RULEBOOK.md:11` proto vyžaduje změnu Go kódu pro každý další obor.
  Automaticky načíst publikované `slovnik/*.csv`; chybějící `slovnik/ztc-009.csv` musí zůstat chybou.
  Pracovní návrhy uložit mimo tento adresář, aby je objevování nezahrnulo mezi publikovaná data.

- **P1 — falešně platné odkazy na pravidla (A).**
  `tools/internal/dictionary/rules.go:35,103` kontroluje jen tvar `@X.Y`.
  Staticky z něj plyne přijetí neexistujícího `@9.99`, také `@1.0`.
  Po parsování ověřit sekci a rozsah `1..spec.ExpectedRules[sekce]`.
  Použít existující mapu z `tools/internal/spec/spec.go:16`, nepřidávat další seznam pravidel.

- **P1 — souhlas a možnosti zveřejnění nejsou dostatečně určeny (B).**
  `README.md:49–54` a `slovnik/GUIDELINES.md:13–15` vyžadují souhlas a odstranění důvěrných údajů.
  Chybí postup pro přesný obsah, cílový repozitář, odmítnutí a přístup pouze pro čtení.
  To je mezera návrhu budoucího instalátoru, nikoli důkaz již provedeného úniku dat.
  Před souhlasem ukázat přesný očištěný slovník, cílový repozitář a zamýšlený PR.
  Změna obsahu nebo cíle vyžaduje nové schválení této změny.

- **P2 — lexikální stav se plete se stavem příspěvku (A).**
  `slovnik/RULEBOOK.md:36–47` správně odděluje lexikální stav od přijetí PR.
  `slovnik/IT.md:5` však označuje nová hesla v CSV za návrhy před přijetím PR.
  `approved` znamená povolený výraz v daném významu; `unapproved` znamená nevhodný výraz s náhradou.
  Ani jeden stav neoznačuje čekání na revizi ani sloučení do upstreamu.
  Lokální odborné ověření a přijetí správcem evidovat odděleně, bez nového sloupce v CSV.

- **P2 — pravidla si odporují u zkratek a náhrad (A).**
  `slovnik/RULEBOOK.md:26–28` žádá zkratku vedle plného názvu se stejným významem.
  `CONTRIBUTING.md:44` a `spec/ztc-009.md:638` nové heslo se stejným významem vylučují.
  Doplnit úzkou výjimku pro zavedenou technickou zkratku; zákaz běžných zkratek D-6 zachovat.
  `CONTRIBUTING.md:43` a `spec/ztc-009.md:637` také opomíjejí povolenou náhradu odkazem na pravidlo.
  Sjednotit je s formátem CSV; při změně specifikace upravit odpovídající část LLM profilu.

- **P2 — příspěvek za celý obor nemá jednotný postup (A, B).**
  `CONTRIBUTING.md:17`, `README.md:50` a `slovnik/GUIDELINES.md:14` vedou přes issue „Nové slovo“.
  Texty výslovně nenařizují jedno issue pro každé heslo, ale dávkový příspěvek nevysvětlují.
  Přijmout jeden oborový slovník v jednom PR; samostatné issue pro každé slovo nevyžadovat.
  Upravit také `.github/PULL_REQUEST_TEMPLATE.md`, kde se nyní předpokládá související issue.

- **P2 — automatická kontrola nepotvrzuje významy (A).**
  `tools/internal/dictionary/checker.go:37` vytváří kontext pro jeden soubor.
  Náhrady zůstanou uvnitř svého CSV; střety významů napříč soubory vyžadují odbornou revizi.
  Opakované počty hesel v `slovnik/README.md:10–11` a `slovnik/IT.md:3–4` odstranit.
  Nevytvářet generátor počtů jen kvůli těmto dvěma textům.

## Přijaté minimum realizace

- Přidat dva Markdown skilly: instalátor a odpovědi.
  Každý použije `SKILL.md` s povinným `name` a `description` podle [Agent Skills][skills].
  Nedělat nový běhový program, MCP server ani vlastní systém nastavení.
- Instalátor nejprve využije známé preference.
  Doplní jen chybějící profesi, témata, výluky a rozsah uložení: konverzace, projekt nebo uživatel.
  Profil uloží rozsah použití, cesty ke slovníkům a odmítnutí zveřejnění.
  Neukládá tokeny ani trvalé oprávnění zveřejňovat další obsah.
- Chybějící slovník vytvoří lokálně z potřeb skutečného textu a nabídne odbornou revizi uživateli.
  Pracovní soubory a profil uloží mimo publikované `slovnik/` a mimo instalační balíček pluginu.
  Existující soubory zachová; opakované spuštění znovu využije známé volby.
  Globální instrukce v tomto prostředí zapisuje pouze do `/Users/michal/.agents/AGENTS.md`.
- Výslovný požadavek na styl má přednost před výchozími preferencemi pluginu.
  Jinak použije přísný profil pouze uvnitř zvoleného rozsahu a mimo výluky.
  Mimo rozsah použije přirozenou češtinu.
  Triviální odpověď bez výslovného požadavku na ZTČ tvoří jedna úplná přirozená věta.
  Výjimka neznamená nucené zkrácení složité odpovědi mimo obor.
- Zveřejnění provede jen po konkrétním souhlasu.
  Ověří zapisovací schopnosti konektoru, nebo přihlášení `gh` a možnost zápisu do větve či forku.
  Samotný Git nevytvoří PR; přihlášení samo nedokazuje právo zapisovat.
  Bez sítě, konektoru či oprávnění zůstane lokální slovník použitelný a PR pouze připravený.
- Balit od kořene repozitáře, aby zůstaly dostupné `llm/`, `slovnik/` a licence.
  `.claude-plugin/plugin.json` a `.codex-plugin/plugin.json` použít jako tenké kompatibilní adaptéry.
  Nedržet v nich další kopie pravidel.
  Claude automaticky nenačítá kořenový `CLAUDE.md` pluginu; potřebné instrukce patří do skillů.
  Viz [Claude plugin reference][claude-plugin].
- Sdílený `hooks/hooks.json` použije `UserPromptSubmit` a krátkou statickou připomínku na stdout.
  Připomínka pouze odkáže na rozhodnutí o použitelnosti; neklasifikuje vstup a neprovádí upload.
  Volitelný příkaz vystačí s `sh` a `cat`, bez Node a Pythonu.
  Prostý stdout této události podporují [Claude hooks][claude-hooks] i [Codex hooks][codex-hooks].
  Cesty musí být uvozené, včetně instalace do adresáře s mezerami.
- Upravit stávající validátor a jeho testy, nepřidávat druhý validátor.
  Zahrnout volání `checkers`, existující test softwarového CSV a testy náhrad.
  Chyba čtení ani chybějící základní slovník se nesmí změnit na úspěch s prázdným seznamem.
  Sjednotit README, CONTRIBUTING, RULEBOOK, GUIDELINES, IT, specifikaci, LLM profil a šablonu PR.
  Pro implementaci doplnit nevydaný CHANGELOG; verzi automaticky nezvyšovat.

## Rozpory a zbývající nejistoty

- **Upřesnění A:** pevná registrace neblokuje vytvoření oboru, blokuje jeho automatické ověření.
  Odbornou revizi nenahrazuje ani automatické hledání CSV, ani kontrola shodných lemmat.
- **Upřesnění B:** přísný profil může zůstat přísný uvnitř zvoleného režimu.
  Musí být oddělen od rozhodnutí, zda jej vůbec použít; protichůdné instrukce nenačítat současně.
- **Aktualizace C:** [OpenAI plugin reference][openai-plugin] nyní doporučuje kořenový `plugin.json`.
  `.codex-plugin/plugin.json` zůstává podporovaný kompatibilní formát, nikoli univerzální standard.
  Pro zadané minimum zachovat oba adaptéry; netvrdit tím shodu s novým formátem Agent Plugins.
  Dokumentace omezuje pluginové hooks na ruční instalaci v Codex desktop a vylučuje je z veřejného katalogu.
  Obecná dokumentace hooks popisuje také CLI; to samo nepotvrzuje instalaci tohoto balíčku v CLI.
- Codex dokumentuje kompatibilní `CLAUDE_PLUGIN_ROOT` a vyžaduje kontrolu a důvěru hooks.
  Změna hooku či přibaleného skriptu může vyžadovat novou kontrolu.
  Viz [Codex hooks][codex-hooks].
  Zapnutí hooku není souhlas se zveřejněním slovníku.
- Agent Skills definuje formát skillu, nikoli podporu shellu, hooks nebo perzistence na každém hostu.
  Kopie samotných dvou složek skillů bez jejich podkladů nestačí.
  Host bez hooks nebo shellu použije dostupné Markdown skilly ručně.
- Statická připomínka nezaručuje, že model vždy správně určí téma nebo triviálnost odpovědi.
  Načtení skillů, vložení kontextu a důvěra hooks čekají na test hostů.
  Totéž platí pro zachování profilu po aktualizaci.

## Akceptační scénáře pro implementaci

1. Třetí publikovaný CSV se ověří bez změny Go registrace; neplatný řádek způsobí neúspěch kontroly.
2. Chybějící základní CSV i chyba čtení vedou k neúspěchu; pracovní slovník mimo `slovnik/` se nenačte.
3. Platné `@1.14`, `@3.7`, `@9.4` projdou; `@1.0`, `@1.15`, `@9.99` a `@0.1` neprojdou.
4. Náhrada lemmatem pouze z jiného CSV neprojde; střet významů je uveden pro ruční revizi.
5. Instalátor využije známou profesi a ptá se jen na chybějící témata, výluky a rozsah uložení.
6. Nový obor funguje lokálně bez sítě; neověřené významy zůstanou označené jako pracovní návrhy.
7. Oborový postup použije ZTČ; vyloučené téma přirozenou češtinu; triviální dotaz jednu úplnou větu.
8. Výslovný jiný styl převáží nad profilem; výslovné ZTČ převáží nad triviální výjimkou.
9. Přepis věty „Služba by se měla restartovat“ nepřidá povinnost restartovat službu.
10. Odmítnutý upload se zapamatuje; profil jej později nepovolí bez nového konkrétního souhlasu.
11. Před uploadem uživatel vidí přesné očištěné CSV a cíl.
    Konektor pouze pro čtení nezpůsobí falešný úspěch.
12. Jeden obor vytvoří jeden PR bez povinných issue pro každé heslo; lokální schválení se nevydává za merge.
13. Balíček zkopírovaný do cesty s mezerami zachová odkazy na profil a slovníky.
14. Na každém podporovaném hostu se samostatně ověří načtení obou skillů a doručení připomínky hooku.
15. Nedůvěryhodný či vypnutý hook nic nespustí; ruční použití skillu zůstane dostupné.
16. Opakovaná instalace ani aktualizace neztratí profil, odmítnutí uploadu nebo pracovní slovník.

## Provedené ověření

- Statická kontrola souborů a pěti odkazovaných specifikací provedena 2026-11-14.
- `go test ./...` v `tools/`: 20 testů prošlo v pěti balíčcích.
- `go run ./cmd/ztcvalidate -root ..` v `tools/`: bez nálezů pro specifikaci, oba CSV a verzi.
- `markdownlint-cli2 docs/audit/2026-11-14-synthesis.md`: bez nálezů.
- Úspěch stávajících testů nepotvrzuje uvedené nové scénáře ani správnost významů.
- Neprovedeno: instalace pluginu, běh hooks v hostech, skutečný upload a vytvoření PR.
  Tyto výsledky nelze označit za ověřené chování za běhu.

[skills]: https://agentskills.io/specification
[claude-plugin]: https://code.claude.com/docs/en/plugins-reference
[claude-hooks]: https://code.claude.com/docs/en/hooks
[openai-plugin]: https://developers.openai.com/plugins/build/plugins
[codex-hooks]: https://learn.chatgpt.com/docs/hooks
