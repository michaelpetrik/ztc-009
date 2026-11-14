# Vývoj softwaru: fullstack

[Slovník vývoje softwaru](software-development.csv) obsahuje 89 hesel.
Používejte jej spolu se [základním slovníkem](ztc-009.csv), který obsahuje 27 hesel.
Definice a příklady jsou v CSV. Nová hesla jsou návrhem pro odborné posouzení před přijetím PR.

## Pokrytí

- Společný základ: soubory, zdrojový kód, Git, závislosti, knihovny, funkce, testy, sestavení a nasazení.
- Frontend: komponenty, formuláře, vstupní pole, tlačítka, přístupnost, responzivní rozvržení, události a styly.
- Backend: klient a server, koncové body, požadavky, odpovědi, autentizace, autorizace, relace a tokeny.
  Dále validace vstupu, serializace, fronty zpráv, mezipaměť a časové limity.
- Databáze: tabulky, řádky, sloupce, klíče, indexy, dotazy, transakce, schéma a migrace.
  Dále replikace, zámky, fond připojení, zálohy a obnovení dat.
- Infrastruktura: kontejnery, obrazy, virtuální stroje, clustery, uzly, síťové porty a doménová jména.
  Dále vyvažování zátěže, reverzní proxy, certifikáty, konfigurace, tajné údaje, metriky a protokoly událostí.
  Dodání softwaru pokrývají průběžná integrace, pipeline, artefakty, nasazení a návrat verze.

## Hranice významů

- `větev` označuje historii změn, nikoli část podmínky v programu.
- `funkce` označuje část programu, nikoli vlastnost produktu.
- `parametr` je součást definice funkce. `argument` je hodnota předaná při volání.
- `rozhraní` označuje komunikaci softwarových součástí. Pro ovládání aplikace použijte `uživatelské rozhraní`.
- `komponenta` označuje část uživatelského rozhraní. `server` označuje program, nikoli fyzický počítač.
- `relace` označuje souvislost interakcí klienta a serveru, nikoli matematický nebo databázový pojem.
- `certifikát` označuje digitální certifikát. `protokol událostí` je záznam, nikoli komunikační protokol.
- Stávající `nainstalovat` má v základním slovníku mechanický význam.
  Jeho schválení se automaticky nevztahuje na instalaci softwaru. Technická slovesa posuzujte podle sekce 1 specifikace.
- Názvy příkazů, produktů a identifikátory kódu zachovejte přesně.
  Heslo `commit` je podstatné jméno, nikoli schválení slovesa „commitnout“.
- Návrat verze aplikace nemusí vrátit změny databáze. Postup obnovy dat popište samostatně.

## Další rozšíření

Slovník pokrývá základní pojmy všech čtyř oblastí. Neobsahuje úplnou terminologii konkrétních frameworků a služeb.
Nová hesla přidávejte podle skutečných textů s uživatelem podle [postupu](GUIDELINES.md) a [pravidel](RULEBOOK.md).
Sloupec `en_ste` zůstává prázdný: shoda s hesly STE nebyla ověřena.
