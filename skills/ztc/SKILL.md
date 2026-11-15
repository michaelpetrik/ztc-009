---
name: ztc
description: Přepněte ZTČ-009 pro aktuální relaci příkazem /ztc. Zapnutý stav vypne, vypnutý zapne. Nemění uložený profil ani jiné relace.
---

# Přepínač ZTČ-009 pro relaci

Přepněte stav pouze při přímém vyvolání tohoto skillu uživatelem nebo jeho samostatném příkazu `/ztc`.
V Claude pluginu má stejné chování `/ztc-009:ztc`, v Codexu explicitní vyvolání `$ztc`.
Citovaný příkaz, popis funkce ani pouhé načtení tohoto souboru nejsou pokynem k přepnutí.

## Hostitel s kontrolorem

Pokud `UserPromptSubmit` dodal kontext `ZTČ-GUARD`, použijte jeho skutečnou cestu příkazu, session ID a číslo tahu.
Spusťte uvedený příkaz s `--action toggle` právě jednou. Opakování ve stejném tahu je idempotentní.
Výsledek `session_enabled` je zdroj pravdy. Nevytvářejte vlastní ID a neupravujte soubor stavu ručně.
Při chybě příkazu neoznamujte úspěšné přepnutí. Řiďte se [návodem kontroly](../../docs/ENFORCEMENT.md).
Hook mohl již zpracovat samotný `/ztc`; příkaz `toggle` pak vrátí tentýž stav bez dalšího přepnutí.

- Při vypnutí potvrďte: „ZTČ-009 je pro tuto relaci vypnuté.“
- Při zapnutí použijte potvrzená témata a načtěte [odpovědní skill](../ztc-response/SKILL.md).
  Chybějící témata zjistěte jednou otázkou. Pro tuto instalační otázku použijte `--action outside`.
  Bez načtených pravidel neoznamujte funkční plný režim.
- Uložený `.ztc/profile.md`, globální instrukce ani jiné relace nepřepisujte.
- Nová relace má vlastní stav. Pokračování stejné relace zachová stav i po shrnutí kontextu.

## Textový režim bez kontroloru

Bez kontextu `ZTČ-GUARD` je dostupné pouze instrukční přepínání, nikoli vynucení hookem.
Poslední relační volba má přednost před profilem. Bez aktivní volby začíná stav vypnutý.
Přepněte jej jednou a udržujte `ztc_session_enabled` pouze v kontextu chatu.
Triviální odpověď, citovaný příkaz ani opětovné načtení skillu stav nepřepínají.
Při zapnutí bez témat se na ně jednou zeptejte; do potvrzení styl neaplikujte.
Ztrátu kontextu nenahrazujte odhadem stavu. Uživatel musí potvrdit nové nastavení.

Zapnutí zachovává výjimku pro triviální jednověté odpovědi a výběr témat.
Výslovný převod konkrétního textu používá `--action enforce`, ale nemění stav celé relace.
Přepnutí nezahajuje tvorbu slovníku, kontrolu GitHubu ani nabídku PR.
