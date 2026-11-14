# Audit instalace: ChatGPT a Codex

Datum ověření: 2026-11-14. Výchozí místní commit: `28c6d67`.
Rozsah: OpenAI prostředí; bez instalace, změny účtu nebo vzdáleného zápisu.

## Závěr

Běžný ChatGPT má připravený kopírovatelný prompt. Codex má lokální komponenty, ale chybí instalační katalog.
Nelze tvrdit, že je připravená univerzální instalace jedním kliknutím.
Rodič ověřil veřejný `main` na `1b6c1f2727942444ff5854dc0bbe3df48ba41ee1`; nové pluginové soubory jsou zatím pouze lokální.

## ChatGPT: chat a instrukce

- `docs/CHAT.md` obsahuje samostatný prompt o 1480 znacích, plný režim odkazuje na LLM profil a CSV.
- Text lze vložit do zprávy bez pluginu. Samotné „odpovídej mi v TZČ-009“ nepředá neznámému modelu pravidla.
- Pro další chat nelze spoléhat na předchozí kontext. Návod tuto mez správně uvádí.
- Nemáme ověřený konkrétní průchod ukládáním vlastních či projektových instrukcí ani automatické připnutí profilu.
- Nejmenší cesta: nabídnout celý prompt jako snadno kopírovatelný text; nepřidávat kvůli tomu server.

## ChatGPT plugin a Work

Aktuální dokumentace popisuje společný katalog pro ChatGPT a Codex.
Pluginy lze používat v Chat a Work na webu, desktopu a mobilu podle dostupnosti účtu.
Instalace přes kartu Plugins a tlačítko plus je podporovaná cesta pro existující katalogovou položku.
Náš balíček takovou publikovanou položku ani ověřený instalační odkaz zatím nemá.

Hooky nefungují v cloudově řízeném ChatGPT Work. Synchronizovaný Work má omezenou podporu správcem určených MCP hooků.
Samotná webová instalace nedoručí shellový skript do výpočetního prostředí.
Proto běžný ChatGPT nesmí záviset na našem shellovém hooku.
Zdroj: [Plugins](https://learn.chatgpt.com/docs/plugins).

## Codex desktop a CLI

- `.codex-plugin/plugin.json` má skills i hook; relativní odkazy potřebují celý kořen balíčku.
- Soubor `marketplace.json` v repozitáři chybí. To je konkrétní mezera distribuce.
- Dokumentované minimum: `.agents/plugins/marketplace.json` s položkou `ztc-009` a místním `source.path: "./"`.
- Položka potřebuje `policy.installation`, `policy.authentication` a kategorii; katalog vlastní název a zobrazovaný název.
- Kompatibilní `.codex-plugin/plugin.json` je stále podporovaný. Kořenový `plugin.json` je doporučený nový formát, ne nutná migrace tohoto minima.
- Aktuální dokumentace připouští pluginové lifecycle hooks v Codex runtime; neplatí starý plošný závěr „jen ruční desktop“.
- Instalace není důvěra hooku. Uživatel musí hook přezkoumat a povolit.
- Veřejné publikování vyžaduje samostatný publikační proces; nepřidávat MCP službu jen kvůli distribuci stylu.
Zdroj: [Package your plugin](https://developers.openai.com/plugins/build/plugins).

Lokálně ověřeno pomocí `--help`: Codex CLI 0.167.0 podporuje následující cestu po přidání a zveřejnění katalogu:

```sh
codex plugin marketplace add michaelpetrik/ztc-009
codex plugin add ztc-009@ztc-009
```

Druhý příkaz předpokládá název katalogu `ztc-009`; před publikací jej sjednoťte se skutečným JSON.
Alternativou je `/plugins` v CLI a nová relace po instalaci.
IDE rozšíření pluginy podle aktuální dokumentace nepodporuje.

## Tlačítko

Dokumentovaný odkaz `codex://plugins/install/ztc-009?marketplace=ztc-009` otevře instalační tok jen pro již známý katalog.
Odkaz sám nepřidá GitHub katalog. Při neznámém katalogu skončí na stránce Plugins.
Vydávat jej za univerzální instalaci na první kliknutí by bylo chybné.
Zdroj: [Plugin links](https://learn.chatgpt.com/docs/reference/commands#plugins).

## Opravy a důkazní mez

1. Přidat katalog a ověřit načtení v izolovaném testovacím prostředí.
2. Zveřejnit aktuální balíček před uvedením GitHub instalačních příkazů jako funkčních.
3. Nahradit neurčité pokyny Codex v `docs/INSTALL.md` konkrétními cestami a aktuálním omezením Work.
4. Zachovat promptovou cestu jako nezávislou základní variantu.

Při testu hooků použít `/hooks`; nová nebo změněná definice se bez důvěry přeskočí.
Zdroj: [Hooks](https://learn.chatgpt.com/docs/hooks#review-and-trust-hooks).

Neověřeno: živá instalace, dostupnost pluginu v konkrétním účtu, mobilní použití a trvalé obnovení témat.
Úspěšná validace JSON nebo CLI nápověda nejsou důkazem skutečného načtení skills ani správné volby stylu.
