# Instalace podle prostředí

Ověření dokumentace: 2026-11-14.
Textová instrukce z [CHAT.md](CHAT.md) zůstává cestou bez instalace.
Pluginový balíček neznamená automatické zapnutí stylu: po instalaci spusťte `ztc-install` a potvrďte témata.

## Stav distribuce

Při auditu byl veřejný `main` na commitu `1b6c1f2`, bez nového pluginu.
Níže uvedené GitHub příkazy jsou připravené pro publikaci, nikoli potvrzená instalace ze současného veřejného repozitáře.
Lokální složku a vytvořený ZIP lze použít bez čekání na publikaci.
Živý průchod přihlášeným účtem zatím nebyl ověřen v žádném z uvedených hostitelů.

| Prostředí | Nejkratší cesta | Automatické připomenutí | Meze |
| --- | --- | --- | --- |
| Claude chat | Textová instrukce nebo upload celého plugin ZIPu | Bez shellového hooku | Dostupnost uploadu podle účtu; zachování témat ověřit v nové relaci |
| Claude Cowork | Upload celého plugin ZIPu | Host podporuje hooks | Neověřený živý upload a důvěra hooku |
| Claude Code | Lokální `--plugin-dir` nebo marketplace příkaz | `UserPromptSubmit` | GitHub cesta čeká na publikaci |
| ChatGPT | Textová instrukce; plugin přes dostupný katalog | Textová cesta bez shellu | Vlastní plugin není publikovaný v katalogu; cloudový Work nepoužije náš shell |
| Codex | Lokální katalog nebo GitHub katalog | Podle runtime a důvěry hooks | Ověřená syntaxe CLI, nikoli živá instalace |
| Pi | Místní balíček nebo `pi install` z Gitu | `before_agent_start` extension | Pi není zde nainstalované; GitHub cesta čeká na publikaci |

## Claude chat a Cowork

V rozhraní s dostupnými pluginy otevřete Customize → Plugins → Add → Upload plugin.
Nahrajte celý `ztc-009.zip`. Archiv obsahuje jeden `.claude-plugin/plugin.json` a všechny podklady skills.
Nerozbalujte jej na samostatné skill ZIPy: jejich relativní odkazy potřebují celý balíček.

V Claude chatu hooks neběží. Aktivujte styl instrukcí nebo vyvoláním skillu.
Cowork a Claude Code hooks podporují, ale jejich povolení a účinek ověřte samostatně.
Synchronizace pluginů do Code vyžaduje podporovanou verzi a stejný Claude účet; API-key přihlášení není stejná cesta.
[Podpora pluginů](https://support.claude.com/en/articles/13837440-use-plugins-in-claude) a [upload ZIPu](https://claude.com/docs/plugins/overview).

## Claude Code

Pro lokální checkout:

```sh
claude --plugin-dir /cesta/k/ztc-009
```

Po publikaci katalogu `.claude-plugin/marketplace.json`:

```sh
claude plugin install ztc-009 --marketplace michaelpetrik/ztc-009
```

Shellová jednopříkazová cesta je dokumentovaná od verze 2.1.292.
Potom použijte `/ztc-009:ztc-install`. Instalace může vyžadovat potvrzení zdroje a rozsahu.
[Oficiální postup](https://code.claude.com/docs/en/discover-plugins).

## ChatGPT

Pro běžný chat vložte [samostatnou instrukci](CHAT.md#rychlé-použití).
Model nemusí mít přístup k souborům, shellu ani našemu katalogu. Samotný název TZČ-009 mu nepředá pravidla.

Dokumentace popisuje také pluginy v Chat a Work podle dostupnosti účtu.
Náš vlastní balíček však není publikovanou položkou tohoto katalogu.
Nepoužívejte instalační odkaz Codexu jako univerzální tlačítko pro ChatGPT.
Cloudově řízený Work nepodporuje náš shellový hook; textové instrukce a skills nesmějí záviset na jeho spuštění.
[Pluginy ChatGPT a Codex](https://learn.chatgpt.com/docs/plugins).

## Codex

Pro místní checkout zaregistrujte lokální katalog:

```sh
codex plugin marketplace add /cesta/k/ztc-009
codex plugin add ztc-009@ztc-009
```

Po publikaci lze první příkaz nahradit tímto:

```sh
codex plugin marketplace add michaelpetrik/ztc-009
```

Katalog je v `.agents/plugins/marketplace.json` a odkazuje na kořen balíčku.
Po instalaci otevřete novou relaci, vyvolejte `ztc-install` a v podporovaném runtime zkontrolujte `/hooks`.
Instalace neznamená důvěru hooku. Podpora desktopu, CLI a cloudových režimů není totožná.

Odkaz `codex://plugins/install/ztc-009?marketplace=ztc-009` vyžaduje, aby klient katalog již znal.
Není náhradou registrace katalogu na první kliknutí.
[Balení pluginu](https://developers.openai.com/plugins/build/plugins), [důvěra hooks](https://learn.chatgpt.com/docs/hooks#review-and-trust-hooks).

## Pi

Pi čte `package.json`, skills a vlastní extension. Nepoužívá Claude `hooks.json`.
Projektová instalace místní složky:

```sh
pi install --local /cesta/k/ztc-009
```

Po publikaci celého balíčku bude stačit jeden instalační příkaz:

```sh
pi install git:github.com/michaelpetrik/ztc-009
```

Potom spusťte `/skill:ztc-install` a potvrďte témata.
Pro dočasný test bez trvalé instalace použijte `pi -e /cesta/k/ztc-009`.
Ověřte načtení přes `pi list`, `pi config` nebo `/reload` podle typu instalace.

Extension používá běhové prostředí Pi a vestavěný modul Node.js. Balíček nepřidává žádné externí závislosti.
Nepublikuje data a nenahrazuje celý systémový prompt. Připojí krátkou připomínku výběru stylu.

Oficiální katalog nabízí tlačítko Copy pro příkaz. Doložené univerzální instalační tlačítko `pi://` jsme nenašli.
Nabídka „zkopírovat příkaz“ není totéž jako „nainstalovat jedním kliknutím“.
[Pi Packages](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/docs/packages.md),
[Pi Extensions](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/docs/extensions.md#before_agent_start),
[katalog Pi](https://pi.dev/packages).

## Vytvoření distribučního ZIPu

Z adresáře `tools` spusťte:

```sh
go run ./cmd/ztcpack -root .. -out ../dist/ztc-009.zip
```

Go potřebuje pouze správce, který archiv vytváří. Uživatel při uploadu Go nepotřebuje.
Archiv vybírá distribuční soubory; neobsahuje `.git`, `.ztc`, pracovní audity ani osobní profil.
Před vydáním ověřte manifesty, odkazy, obsah archivu a živou instalaci zvolenou cestou.

## Co ještě zbývá pro veřejné vydání

1. Pushnout změny přes PR a nechat projít CI.
2. Zveřejnit verzovaný ZIP jako release asset a aktualizovat skutečný odkaz ke stažení.
3. Ověřit stažení nebo instalaci z čistého prostředí a následné vyvolání `ztc-install`.
4. Ověřit jeden profesní, jeden triviální a jeden vyloučený dotaz; potom novou relaci a vypnutí stylu.
5. Katalogovou dostupnost ChatGPT a účetní omezení ostatních hostitelů ověřit odděleně.
