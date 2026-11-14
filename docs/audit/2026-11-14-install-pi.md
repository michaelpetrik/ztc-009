# Instalace do Pi: ověření připravenosti

Datum: 2026-11-14. Výchozí commit: `28c6d67`.
Rozsah: Pi coding agent, původně `badlogic/pi-mono`, nyní [earendil-works/pi](https://github.com/earendil-works/pi).
V tomto shellu není příkaz `pi` dostupný. Následující závěry vycházejí z dokumentace a kontroly souborů, nikoli živé instalace.

## Verdikt

Základní skills lze z existujícího lokálního balíčku nabídnout Pi.
Připravená instalace s automatickou připomínkou ale chybí: repo nemá Pi manifest ani Pi extension.
`hooks/hooks.json` a `hooks/ztc-context.sh` jsou adaptéry jiných hostitelů, nikoli Pi.
Stav publikovaného remote nebyl v tomto dílčím auditu ověřen.

## Podporovaná instalační cesta

Pi podporuje balíčky z Gitu, npm i místního adresáře.
Dokáže objevit konvenční `skills/` i bez pole `pi`; explicitní manifest je přehlednější pro distribuci.
Viz [Pi Packages](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/docs/packages.md).

Po zveřejnění změn by instalační příkaz byl:

```sh
pi install git:github.com/michaelpetrik/ztc-009
```

Pro test lokálního checkoutu bez změny uživatelských nastavení:

```sh
pi -e /Users/michal/Projects/ztc-009/ztc-009-repo
```

Trvalá projektová instalace:

```sh
pi install --local /cesta/k/ztc-009
```

Následuje `/skill:ztc-install` v Pi a potvrzení témat.
Příkaz instalace standardně ukládá osobní konfiguraci; `--local` zvolí projekt a vyžaduje důvěru projektu.
Pro ověření slouží `pi list`, `pi config` a `/reload`.
Viz [CLI](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/docs/cli.md#package-commands).

## Minimální balíček

Navržený `package.json` nevyžaduje externí závislosti ani npm publikaci:

```json
{
  "name": "ztc-009",
  "version": "0.1.0",
  "type": "module",
  "keywords": ["pi-package"],
  "pi": {
    "skills": ["./skills"],
    "extensions": ["./extensions/ztc.js"]
  }
}
```

Celý gitový balíček zachová `llm/`, `slovnik/`, `spec/` a odkazy relativní ke skills.
Pouhé kopírování samotné složky skillu by tyto odkazy rozbilo.
Pi zveřejní jména a popisy skills modelu; úplný text načte až při použití.
Příkaz `/skill:ztc-install` vyvolá konkrétní skill bez závislosti na správném automatickém výběru modelu.
Viz [Skills](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/docs/skills.md).

## Automatická připomínka

Extension je ES modul s default exportem přijímajícím `pi`.
Nejmenší doložený bod je `pi.on("before_agent_start", handler)`.
Handler může vrátit `{ systemPrompt: event.systemPrompt + "\n\n" + reminder }`.
Tím zachová dosavadní prompt a přidá krátkou připomínku rozhodování o ZTČ pro aktuální běh.

Připomínka má obsahovat cestu k odpovědnímu skillu, potvrzená témata, výjimku pro triviální dotaz a respekt k vypnutí.
Nemá načítat přísný profil plošně, číst tajné údaje, zapisovat konfiguraci ani spouštět upload.
Absolutní cestu lze odvodit přes `new URL("../skills/ztc-response/SKILL.md", import.meta.url)` a vestavěné `node:url`.
Není nutný import Pi SDK; JavaScript bez typových importů nemá další runtime závislost.

Dokumentace také nabízí strukturované úpravy promptu šetrnější ke cache.
Pro krátký neměnný reminder je doložený návrat `systemPrompt` nejmenší varianta, ale živou kompatibilitu je nutné ověřit.
Viz [Extensions](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/docs/extensions.md#before_agent_start).

## Tlačítko versus jeden příkaz

[Oficiální katalog](https://pi.dev/packages) zobrazuje příkaz `pi install ...` a tlačítko **Copy**.
V prozkoumané oficiální dokumentaci jsem nenašel doložené instalační URL `pi://` ani univerzální one-click install.
Tlačítko „Kopírovat příkaz“ lze poctivě nabídnout; není to automatická instalace.
Npm katalog vyžaduje publikovaný balíček s klíčovým slovem `pi-package`; pro první verzi stačí Git.

## Ověření před tvrzením „připraveno“

- Ověřit JSON, existenci deklarovaných skills, extension a relativních podkladů.
- V testu importovat extension se stubem `pi.on`; vyvolat handler a ověřit zachování původního promptu.
- Ověřit cestu balíčku s mezerami a absenci vedlejších efektů.
- Na skutečném Pi ověřit načtení přes `-e`, `/skill:ztc-install`, další dotaz a `/reload`.
- Gitový příkaz uvádět jako použitelný teprve po ověření odpovídajících souborů na remote.

Bez živého Pi testu a bez ověření remote nelze deklarovat hotovou instalaci jedním příkazem.
