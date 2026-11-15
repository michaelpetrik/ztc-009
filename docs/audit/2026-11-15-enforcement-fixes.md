# Opravy po auditu kontroloru

Navazuje na [reprodukce obejití](2026-11-15-enforcement-bypass.md).

| Nález | Výsledek |
| --- | --- |
| Model si udělí `outside` nebo `off` | Změnové CLI akce jsou odmítnuté. Stav a výjimky řídí přesný uživatelský vstup v hooku. |
| Próza ukrytá v citaci nebo kódu | Samotný obal kontrolu nevypne. Chráněné úseky musí mít shodný otisk s aktuálním zadáním. Nový kód povoluje uživatel. |
| Falešný nadpis a dlouhá závorka | Obsah se počítá do limitu slov. |
| Zkratka s neviditelným znakem | Kontrolní kopie odstraní znaky kategorie Cf a dekóduje HTML entity. |
| Prázdný úspěch spouštěče | Validuje se JSON, schéma a výstupní protokol. Prázdný program nebo chybný protokol zastaví tah. |
| Konstrukce „Ventil musí být uzavřen.“ | Přidána cílená kontrola rozpoznaného trpného příčestí po „musí být“. Nejde o úplný jazykový analyzátor. |
| Zpožděný Stop | Pokud hostitel poskytne ID tahu, musí odpovídat aktuálnímu stavu. |

Regresní testy obsahují původní obcházející vstupy, odmítnuté změnové akce a skutečný shellový spouštěč.
Testuje se i zachování doloženého kódu a citace; změněný citační obsah ochranu nezíská.

## Zbývající hranice hostitele

Podvržení celé hook události nebo úpravu souborů agentem s neomezeným shellem plugin neizoluje.
Původní reprodukce falešného nového promptu není vydávána za bezpečnostně vyřešenou.
Odstranění změnového API z běžného modelového workflow tuto hranici nenahrazuje.
Vyžaduje hostitelem chráněné úložiště, spouštěč a důvěryhodný kanál událostí mimo oprávnění agenta.

Stop také neodstraňuje již streamovaný návrh a testy necertifikují všech 53 jazykových pravidel.
Tyto meze jsou výslovně uvedené v [návodu kontroly](../ENFORCEMENT.md#hranice-důvěry).
Živé ověření v přihlášených klientských účtech nebylo součástí tohoto uzavření.
