# Instalace ZTČ-009 pro agenty

Pro běžný Claude a ChatGPT začněte [textovou instrukcí](CHAT.md). Plugin není podmínkou použití ZTČ.

Pro prostředí s podporou skills začněte načtením [instalačního skillu](../skills/ztc-install/SKILL.md).
Instalace nejprve zjistí profesi, témata a výjimky. Styl se nezapíná pro všechny dotazy automaticky.

## Přenosný balíček

Stáhněte celý repozitář, nebo jej naklonujte:

```sh
git clone https://github.com/michaelpetrik/ztc-009.git
```

Zachovejte složky `skills`, `llm`, `slovnik` a `spec` ve společném kořeni.
Samotné zkopírování souboru `SKILL.md` rozbije odkazy na pravidla a slovníky.
Textové instrukce a skills nepřidávají závislost na Go, Node.js ani Pythonu.
Hostitel má vlastní provozní požadavky; Pi používá své běhové prostředí Node.js.
Stažení balíčku vyžaduje síť, místní načtení již staženého obsahu nikoli. Účet GitHub není potřeba pro textovou cestu.
Git je potřeba jen pro klonování nebo gitový příspěvek. Balíček lze získat také jako archiv.

## Volba prostředí

Konkrétní příkazy a dostupnost najdete v [návodu pro šest prostředí](HARNESSES.md).
Obsahuje Claude chat, Cowork, Claude Code, ChatGPT, Codex a Pi.
Pro Claude chat a Cowork lze připravit celý plugin ZIP. Pro Code a Codex jsou přibalené katalogy.
Pi používá vlastní balíček a extension. GitHub příkazy vyžadují nejprve publikaci aktuálních souborů.

## Jiný agent nebo prostředí bez hooků

Dejte agentovi celý balíček a požádejte jej o použití `skills/ztc-install/SKILL.md`.
Agent vloží odkaz na `skills/ztc-response/SKILL.md` a potvrzený profil do podporovaných trvalých instrukcí hostitele.
Před úpravou zachová existující obsah a použije uživatelem zvolený rozsah instalace.
Opakovaná instalace aktualizuje stejný odkaz, nevytváří další kopii instrukcí.
Pokud hostitel nemá trvalé instrukce, platí nastavení jen v aktuální relaci.

Skills tvoří přenosné jádro. Automatické spouštění závisí na hostiteli.
Hook pro Claude Code a Codex vyžaduje POSIX `sh`. Na Windows bez tohoto shellu použijte skills bez hooku.
Hook nepřepisuje odpovědi, neposílá data a nevytváří PR. Jen připomíná pravidla výběru stylu.

## Výběr témat a chybějící slovník

1. Agent se zeptá: „Pro jakou profesi a témata chcete používat ZTČ? Která témata chcete vynechat?“
2. Agent vybere základní slovník a vhodné oborové slovníky.
3. Pro nepokrytý obor vytvoří místní pracovní slovník a ukáže významy a příklady uživateli.
4. Agent uloží potvrzené nastavení. Výchozí projektová cesta je `.ztc/profile.md`.
5. Agent nabídne konkrétní slovník k příspěvku. Odmítnutí nebrání místnímu použití.

Místní `.ztc/` nepatří do veřejného příspěvku. V cílovém projektu ověřte ignorování tohoto adresáře před prací s Gitem.
Profil ukládá pouze témata, výjimky, cesty ke slovníkům a stav nastavení. Neukládá tokeny ani pracovní dokumenty.
Pro instalaci na úrovni uživatele použijte úložiště instrukcí podporované daným prostředím a potvrzené uživatelem.
Pokud existuje společný soubor globálních instrukcí, respektujte jej. Nevytvářejte konkurenční globální nastavení.

## Dobrovolný příspěvek přes PR

Agent ukáže přesné soubory, které chce zveřejnit, a cíl `michaelpetrik/ztc-009`.
Odstraní osobní údaje, interní názvy a neveřejné příklady. Připomene licenci slovníku CC0.
Vysvětlí: „Tento slovník pomůže dalším instalacím ve vašem oboru a omezí opakované vytváření stejných hesel.“

Agent ověří dostupnou cestu bez zveřejnění dat:

- GitHub connector: dostupné operace pro čtení identity, zápis souborů do větve a vytvoření PR.
- `gh`: dostupnost CLI, přihlášený účet a přístup k cíli nebo možnost příspěvku přes fork.
- Jiný nástroj: stejné schopnosti prostřednictvím API nebo podporovaného rozhraní.

Samotný Git umí vytvořit commit a pushnout větev. PR vytváří GitHub API, connector, CLI nebo webové rozhraní.
Pouhá přítomnost connectoru či CLI nedokazuje oprávnění k zápisu.
Při ověření nezobrazujte přístupové údaje. Neinstalujte nástroje ani neměňte účet bez požadavku uživatele.

Před prvním uploadem agent položí otázku:
„Souhlasíte se zveřejněním těchto souborů pod CC0 v repozitáři michaelpetrik/ztc-009 prostřednictvím PR?“
Teprve souhlas opravňuje vytvořit veřejný příspěvek, včetně potřebné větve, pushnutí a PR.
Souhlas platí jen pro ukázaný obsah a cíl. Změna obsahu nebo cíle vyžaduje nové potvrzení.
Jeden ucelený profesní slovník patří do jednoho PR. Předchozí issue není povinné.

Předchozí odmítnutí všech příspěvků respektujte bez další nabídky a bez kontroly GitHubu.
Bez souhlasu se nic nepublikuje. Bez vhodného nástroje agent připraví místní soubory a návod k ručnímu PR.
Místní instalace zůstane použitelná. Agent nesmí oznámit vytvoření PR bez odkazu na skutečný PR.

## Použití, změna a vypnutí

- V potvrzených tématech používejte ZTČ pro věcné odpovědi, které potřebují více než triviální jednu větu.
- Triviální dotaz s úplnou přirozenou jednovětou odpovědí zodpovězte běžně.
- Krátký dotaz může být složitý. Délka zadání není test triviality.
- Výslovný požadavek na ZTČ nebo na vypnutí stylu pro danou odpověď má přednost před výchozím výběrem.
- Pro změnu témat znovu vyvolejte `ztc-install`. Známé odpovědi se nemají znovu vyžadovat.
- Pro vypnutí změňte `enabled` v místním profilu na `false` nebo požádejte agenta o vypnutí.
- Pro odinstalování odstraňte pouze odkazy a nastavení patřící ZTČ; soukromé slovníky zachovejte, dokud uživatel nežádá opak.

## Ověření

Ověřte běžnou profesní odpověď, dotaz mimo vybraná témata a triviální otázku „Kolik je 2 + 2?“.
Zkontrolujte také vypnutí stylu a další relaci. Pokud se nastavení neobnoví, instalace není trvalá.
Validátor kontroluje strukturu slovníku, nikoli odbornou správnost definic nebo automatickou shodu odpovědí se ZTČ.

## Přehled rozhodování

Viz [diagram](ztc-flow.svg) a jeho [Mermaid zdroj](ztc-flow.mmd).

## Podklady k hostitelům

- [Agent Skills](https://agentskills.io/specification).
- [Claude Code pluginy](https://code.claude.com/docs/en/plugins-reference) a [hooks](https://code.claude.com/docs/en/hooks).
- [Codex pluginy](https://developers.openai.com/plugins/build/plugins) a [hooks](https://learn.chatgpt.com/docs/hooks).
