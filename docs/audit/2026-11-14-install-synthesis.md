# Rozhodnutí o snadné instalaci ZTČ-009

- Stav k 2026-11-14: připravit distribuční minimum; neslibovat veřejnou instalaci ani plnou podporu všech hostitelů.
- Nejmenší úplný výsledek: prompt pro chat, celý plugin ZIP pro Claude, katalogy pro Code/Codex a Git balíček pro pi.
- Pi nabídnout příkaz ke kopírování; doložené univerzální instalační tlačítko chybí.

## Důkazy a jejich meze

- Lokální základ: `28c6d67`, ověřeno při syntéze.
- Veřejný `main`: `1b6c1f2727942444ff5854dc0bbe3df48ba41ee1`, převzato z ověření rodiče přes `gh`.
  Nový plugin nebyl publikován; GitHub příkazy níže jsou cílové, nikoli dnes ověřeně funkční návody.
- Lokální kontrola: `cd tools && go test ./internal/plugincheck` prošla.
  Ověřuje manifesty a shellový hook, včetně cesty s mezerami a absence zápisů; ne instalaci v hostiteli.
- Dokumentovaná podpora: doložena třemi [audity Claude][A], [OpenAI][O] a [pi][P].
  Sporné tvrzení o hooks OpenAI bylo navíc přímo porovnáno s primární dokumentací [B], [H] a [W].
- Živá instalace: neověřena ve všech šesti hostitelích; dostupnost účtu a trvalé načítání profilu neověřeny.
- Veřejná dostupnost: chybí vydaný nový balíček; vlastní Git katalog není zápis v oficiálním veřejném katalogu.

## Matice hostitelů

Všechny řádky mají stav „živě neověřeno“; cesty uvádějí dokumentovanou možnost, nikoli dosažený výsledek.

| Hostitel | Nejkratší cesta | Co zbývá | Automatická připomínka |
| --- | --- | --- | --- |
| Claude chat | Prompt; celý plugin ZIP [A] | ZIP, upload návod, test účtu | Hook neběží |
| Cowork | Customize → Plugins → Add → Upload plugin [A] | ZIP a test nové úlohy | Dle podpory hostitele |
| Claude Code | Lokálně `--plugin-dir`; poté Git marketplace [A] | Katalog, publikace, test | Hook s POSIX `sh` |
| ChatGPT | Prompt; nativně publikovaná položka Plugins [O] | Prompt lze použít; položka chybí | Náš shell není základ |
| Codex CLI/desktop | Přidat marketplace a plugin [O] | Katalog, publikace, test verze | Hook až po důvěře |
| pi | Git balíček, pak `/skill:ztc-install` [P] | Manifest; extension pro připomínku | `before_agent_start` |

- Textová cesta v [CHAT.md](../CHAT.md) funguje bez instalace, ale nezaručuje trvalost mezi chaty.
- Cílový Code příkaz: `claude plugin install ztc-009 --marketplace michaelpetrik/ztc-009`.
  Audit [A] uvádí minimum CLI 2.1.292; interaktivní varianta má odlišné minimum a potvrzení zdroje.
- Cílový Codex: `codex plugin marketplace add michaelpetrik/ztc-009`, pak `codex plugin add ztc-009@ztc-009`.
  Audit [O] ověřil nápovědu CLI 0.167.0; nikoli instalaci. Název katalogu musí odpovídat JSON.
- Cílové pi: `pi install git:github.com/michaelpetrik/ztc-009`; v auditu [P] nebylo pi dostupné.

## Rozpory a rozhodnutí

- Výchozí [INSTALL.md](../INSTALL.md) omezuje Codex na ruční desktop a vylučuje pluginové hooks z veřejného katalogu.
  Jako plošný zákaz je to zastaralé: [B] popisuje pluginové hooks a [H] jejich kontrolu přes `/hooks` v CLI.
- Širší formulaci „ChatGPT Work a Codex“ v [B] omezuje konkrétní runtime v [W].
  Cloudově řízený Work nepodporuje tyto pluginové hooks; synchronizovaný Work má pouze správcovské MCP hooks.
  Webová instalace sama nedodá shellový skript. Důvěra hooku a jeho přítomnost jsou samostatné podmínky.
- `CLAUDE_PLUGIN_ROOT` není sama o sobě chyba v Codex adaptéru: [B] potvrzuje kompatibilní proměnnou.
- Staré tvrzení „Claude chat umí jen prompt“ odporuje [A] a [C]; nově dokumentovaný upload ale není živý test.
- [Oficiální katalog pi][I] podle [P] nabízí **Copy**, nikoli doložené `pi://` pro instalaci.
  Ani odkaz Codex `codex://plugins/install/...` nepřidá dosud neznámý katalog [O], [L].

## Minimální opravy pro rodiče

1. Přidat `.claude-plugin/marketplace.json` a `.agents/plugins/marketplace.json` pro stejný kořen balíčku.
   Dva malé katalogy jsou přijatelné pro jasné nativní návody; OpenAI podporuje i kompatibilní Claude katalog [B].
   Druhý katalog tedy není univerzální podmínkou formátu. Sjednotit názvy, verze, cesty a povinné policy položky.
2. Zabalit celý plugin přes stávající Go modul a `archive/zip`; bez nové knihovny a závislostí pro příjemce.
   Výslovný seznam obsahu: manifesty, skills, hook, pravidla, slovníky, potřebné návody a licence.
   Zachovat relativní odkazy včetně `docs`, `spec` a `CONTRIBUTING.md`; vynechat `.git`, `.ztc` a audity.
   Samostatné ZIPy skills nepřidávat: jejich odkazy vedou mimo složku skillu a celý ZIP řeší stejnou potřebu.
3. Pro pi přidat `package.json` a malou JS extension bez balíčkových závislostí, jak navrhuje [P].
   Extension je potřebná pro připomínku, nikoli pro samotné skills; musí zachovat původní system prompt.
   Připomínat potvrzená témata, vypnutí a výjimku triviality; nespouštět onboarding, zápis ani upload.
4. Opravit [INSTALL.md](../INSTALL.md) a [CHAT.md](../CHAT.md): konkrétní postupy, rozsah účtu, trvalost a stav vydání.
   Zkrátit popis `ztc-install` pod 200 znaků pro kompatibilitu dle [A]; nevydávat to za prokázaný blok celého ZIPu.
   „Bez Node.js“ vztáhnout na obsah skills; samotné pi má vlastní runtime. Go potřebuje pouze tvůrce ZIPu.
5. V existujícím webu použít tlačítko „Kopírovat příkaz“, pokud tam patří; pro první vydání stačí blok příkazu.
   Nový web, server, MCP, npm publikace a migrace na nový kořenový OpenAI manifest nejsou nutné pro tento cíl.
6. Publikaci řešit odděleně; teprve dostupný commit a ZIP dovolují nabízet veřejný instalační postup.
   Tyto opravy samy nedokončí nativní distribuci v běžném ChatGPT; jeho minimální cestou zůstává celý prompt.

## Kritéria přijetí

- [ ] „Lokálně validováno“: parsovat katalogy, ověřit deklarované soubory a odkazy také v rozbaleném čistém ZIPu.
- [ ] Ověřit obsah ZIPu proti povolenému seznamu; žádný soukromý profil, únik přes symlink ani cestu mimo balíček.
- [ ] Pi test se stubem `pi.on`: zachovaný prompt, cesta s mezerami, správný skill, žádné vedlejší účinky.
- [ ] „Instalace ověřena“ uvést zvlášť pro hostitele, verzi, OS, způsob instalace a datum skutečného průchodu.
- [ ] V každém hostiteli načíst skill, potvrdit témata a ověřit profesní dotaz, dotaz mimo rozsah a trivialitu.
- [ ] Ověřit vypnutí, opakované nastavení bez duplicit a novou relaci; jinak označit profil jako relační či neověřený.
- [ ] S hookem i bez něj zachovat funkční ruční cestu; soukromý profil nesmí být podmínkou veřejného balíčku.
- [ ] „Veřejně dostupné“: čisté stažení z remote obsahuje katalogy a zdroje; instalační příkaz míří na stejné vydání.
- [ ] Nativní ChatGPT nabídnout až s dostupnou katalogovou položkou; jedno kliknutí tvrdit jen po testu čistého účtu.

## Zdroje a nejistota

- Anthropic a pi přebírají primární odkazy a závěry analytiků; tato syntéza jejich živé hostitele nezkoušela.
- Výsledek platí pro uvedený základ; plánované změny rodiče nejsou tímto memo ověřeny.
- Vynecháno: provozní instalace, publikace a commity. Riziko: rozdíly verzí, plánů a správcovských politik.

[A]: 2026-11-14-install-claude.md
[O]: 2026-11-14-install-openai.md
[P]: 2026-11-14-install-pi.md
[B]: https://developers.openai.com/plugins/build/plugins
[H]: https://learn.chatgpt.com/docs/hooks#review-and-trust-hooks
[W]: https://learn.chatgpt.com/docs/plugins
[C]: https://support.claude.com/en/articles/13837440-use-plugins-in-claude
[I]: https://pi.dev/packages
[L]: https://learn.chatgpt.com/docs/reference/commands#plugins
