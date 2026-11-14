# Kritická revize připravenosti instalací

Datum: 2026-11-14.
Podklady: [syntéza](2026-11-14-install-synthesis.md), tři dílčí audity, [návod](../HARNESSES.md) a aktuální implementace.
Revize je statická. Nepotvrzuje dostupnost funkcí konkrétního účtu ani živou instalaci.

## Verdikt

Syntéza odpovídá zadání a nezaměňuje dokumentovanou podporu za živé ověření.
Veřejná snadná instalace stále čeká na publikaci nových souborů, ZIPu a test hostitelů.
Pi má doloženou cestu jedním příkazem; univerzální instalační tlačítko doložené není.

## Nálezy k opravě

- **P1 — ZIP vynechává Pi extension.**
  Při revizi seznam přípon v `tools/cmd/ztcpack/main.go`, funkce `distributable`, neobsahoval `.js`.
  Manifest `package.json` přitom odkazuje na `extensions/ztc.js`.
  Přidat `.js` a ověřit přítomnost deklarované extension přímo v dokončeném archivu.
  Úspěšný test extension ve zdrojovém repozitáři tuto chybu balení neodhalí.

- **P2 — publikovaný příkaz na vytvoření ZIPu selže na čistém checkoutu.**
  `docs/HARNESSES.md` uvádí výstup `../dist/ztc-009.zip`, ale adresář `dist` není verzovaný.
  `pack` volá `os.CreateTemp` v cílovém adresáři bez jeho vytvoření.
  Vytvořit cílový adresář v nástroji, nebo přidat příkaz do návodu; otestovat přesně uvedený průchod.

- **P2 — příliš široké tvrzení o nulových požadavcích.**
  `docs/INSTALL.md` uvádí, že instalace ani odpovědi nepotřebují Node.js ani síť.
  Pi extension používá Node runtime hostitele; stažení balíčku z Gitu potřebuje síť.
  Tvrzení o nulových dodatečných závislostech omezit na textové jádro a odlišit runtime hostitele.

Nálezy byly předány rodiči během práce. Tento dokument zachycuje stav při revizi, nikoli výsledek následných oprav.

## Co je přijatelné

- Katalogy odkazují na celý kořen balíčku a zachovávají podklady odkazované ze skills.
- Pi extension nepřebírá obsah uživatelského dotazu a připojuje připomínku k původnímu systémovému promptu.
- Pi příkaz je označený jako použitelný až po publikaci; obecné tlačítko není slibované.
- Codex deep link je podmíněný předem známým katalogem, nikoli prezentovaný jako registrace katalogu.
- Claude chat nemá přislíbený shellový hook. ChatGPT má samostatnou textovou cestu.
- Rozsah témat a souhlas se zveřejněním slovníku zůstávají oddělené od instalace pluginu.
- Publikační blokátor je uvedený na začátku návodu a ve všech dílčích závěrech.

## Podmínky pro závěrečné tvrzení

Před označením balíčku jako lokálně připraveného zkontrolovat hotový ZIP, nikoli pouze zdrojové soubory.
Ověřit manifestové cesty, Markdown odkazy, vyloučení soukromých souborů a přesný příkaz z návodu.

„Instalace ověřena“ vyžaduje samostatný živý průchod pro hostitele, verzi, účet a instalační cestu.
„Veřejně dostupné“ vyžaduje odpovídající remote a dostupný distribuční artefakt.
Ani publikovaný Git katalog sám nezaručuje katalogovou dostupnost v běžném ChatGPT.

Bez živých instalací, publikace a nového ověření externích stránek; primární zdroje byly posouzené prostřednictvím auditů.
Zbývající riziko: rozdíly hostitelů, účtů a verzí; ty nelze odstranit samotným testem JSON nebo stubem extension.

## Vypořádání a ověření rodičem

- Packer nyní zahrnuje `.js` a vytváří cílový adresář. Přesný příkaz z návodu úspěšně vytvořil `dist/ztc-009.zip`.
- Tvrzení v INSTALL.md rozlišuje textové jádro, runtime hostitele a síť potřebnou ke stažení.
- Go formátování, vet, race testy a validátor prošly. Pi testy prošly ve třech scénářích.
- Rozbalený ZIP obsahuje 40 souborů včetně extension a obou katalogů; lokální Markdown odkazy se překládají na existující cesty.
- Archiv neobsahuje `.git`, `.ztc`, audity ani testovací složku. Filtr není kontrolou tajných údajů uvnitř povolených souborů.
- Živá instalace, veřejná publikace a obnova profilu v dalším chatu nebyly provedeny.
