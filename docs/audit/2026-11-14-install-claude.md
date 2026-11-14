# Instalační audit: Claude, Cowork a Claude Code

Datum: 2026-11-14. Výchozí lokální commit: `28c6d67`.
Rozsah: statická kontrola balíčku a aktuální oficiální dokumentace. Bez instalace do účtu a bez publikace.

## Výsledek

Snadná veřejná instalace zatím připravená není.
Lokální balíček má základ pro všechny tři hostitele, ale chybí distribuce a návod pro nativní upload.

| Hostitel | Současný místní stav | Nejkratší podporovaná cesta | Mezera |
| --- | --- | --- | --- |
| Claude chat | Funkční textová cesta je popsaná; plugin má manifest a skills. | Vložit instrukci; pro instalaci přidat celý plugin ZIP v Customize. | Není připravený distribuční ZIP ani tento postup v návodu. Hook v chatu neběží. |
| Claude Cowork | Stejný plugin lze zabalit; chybí vlastní návod. | Customize → Plugins → Add → Upload plugin. | Nutno ověřit skutečný upload, kontext skillů a profil v nové úloze. |
| Claude Code | Je popsán lokální `--plugin-dir`; manifest a hook existují. | Lokální složka dnes; po publikaci marketplace jeden instalační příkaz. | Chybí marketplace.json a publikované aktuální soubory. |

Rodičovský agent ověřil GitHub `main` na `1b6c1f2727942444ff5854dc0bbe3df48ba41ee1`.
Nové pluginové soubory jsou pouze v lokální historii. Klonování veřejného main nyní nový plugin neposkytne.
Tento vzdálený údaj je převzatý z jeho přímého API ověření, nikoli z mého samostatného čtení remotu.

## Aktuální podpora Anthropic

Claude nyní podporuje pluginy v běžném chatu, Cowork i Claude Code.
Skills a příkazy fungují ve všech třech; hooks běží v Cowork a Code, nikoli v chatu.
Pluginy přidané přes web či desktop se ukládají k účtu.
Do Code se synchronizují při přihlášení stejným Claude účtem od verze 2.1.273; API-key přihlášení tuto synchronizaci neposkytuje.
Dostupnost pluginů a firemní politiky závisejí na plánu a správci.
[Oficiální nápověda](https://support.claude.com/en/articles/13837440-use-plugins-in-claude).

Nativní upload jde přes Customize → Plugins → Add → Upload plugin.
Přijímá `.zip` nebo `.plugin`; archiv může obsahovat složku pluginu nebo její obsah.
Uvnitř musí být jeden `.claude-plugin/plugin.json`.
Celý balíček zachová existující relativní odkazy mezi `skills`, `llm`, `slovnik` a `docs`.
Toto je vhodnější minimální cesta než samostatné ZIPy každého dnešního skillu.
[Oficiální dokumentace pluginů](https://claude.com/docs/plugins/overview).

## Claude Code jedním příkazem

Po přidání marketplace a zveřejnění souborů lze dokumentovat:

```sh
claude plugin install ztc-009 --marketplace michaelpetrik/ztc-009
```

Shellová varianta vyžaduje verzi 2.1.292 nebo novější.
V interaktivní relaci odpovídá `/plugin install ztc-009 --marketplace michaelpetrik/ztc-009`, od verze 2.1.275.
Interaktivní postup zahrnuje potvrzení zdroje a volbu rozsahu; není to neviditelná instalace bez rozhodnutí uživatele.
[Instalace pluginů](https://code.claude.com/docs/en/discover-plugins).

Nejmenší katalog je `.claude-plugin/marketplace.json` s vlastním názvem, vlastníkem a jednou položkou `ztc-009`, `source: "./"`.
Zdroj v kořeni vlastního pluginového repozitáře je výslovně podporovaný.
Při dalších vydáních se musí zvýšit manifestová verze, jinak může updater ponechat starou kopii.
[Publikace pluginu](https://code.claude.com/docs/en/plugins/publish).

## Samostatné skill ZIPy

Dnešní `skills/ztc-install` ani `skills/ztc-response` nejsou samostatné přenosné uploady.
Odkazují mimo vlastní složku (`../../llm`, `../../slovnik`, druhý skill).
ZIP jednoho skillu musí mít kořenovou složku skillu a všechny potřebné zdroje uvnitř.
Oficiální návod navíc uvádí limit popisu 200 znaků a požadavek zapnutého code execution.
`ztc-install` má 207 znaků, `ztc-response` 195. Pro bezpečnější kompatibilitu zkraťte první popis.
[Formát a upload custom skills](https://support.claude.com/en/articles/12512198-how-to-create-custom-skills).

Generování samostatných skill ZIPů nyní nepřidávat, pokud celý plugin ZIP řeší stejnou potřebu.
Textová instrukce zůstává cesta bez pluginového runtime a bez nativní instalace.

## Minimální opravy

1. Přidat marketplace.json s pluginem v kořeni.
2. Vytvořit čistý plugin ZIP z výslovně vybraných distribučních souborů; nezahrnout `.git`, `.ztc` ani pracovní audity.
3. Zkrátit popis instalačního skillu pod 200 znaků.
4. Doplnit konkrétní upload postup pro Claude chat a Cowork; vysvětlit, že chat přeskočí hooks.
5. Publikovat aktuální změny a teprve potom nabízet veřejný instalační příkaz nebo odkaz ke stažení.
6. Ověřit lokální marketplace v izolovaném nastavení a celý ZIP staticky; živý test účtu přiznat jako samostatný stav.

## Nejistoty a protiváha

- Podpora podle dokumentace není důkaz, že aktuální účet má povolený upload nebo hooks.
- Nativní plugin pro Claude chat je dnes dokumentovaný; starší závěr „chat umí pouze prompt“ by byl nesprávný.
- Přenesení pluginu neznamená přenesení soukromého `.ztc/profile.md` mezi zařízeními nebo hostiteli.
- Jeden příkaz není jedno tlačítko v libovolném hostiteli. Pro katalogové tlačítko je nutná publikace či registrace katalogu.
- Nezkoušel jsem živé nahrání ZIPu, nový chat, Cowork task ani synchronizaci účtu do Code.
