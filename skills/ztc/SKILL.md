---
name: ztc
description: Přepněte ZTČ-009 pro aktuální relaci příkazem /ztc. Zapnutý stav vypne, vypnutý zapne. Nemění uložený profil ani jiné relace.
---

# Přepínač ZTČ-009 pro relaci

Přepněte stav pouze při přímém vyvolání tohoto skillu uživatelem nebo jeho samostatném příkazu `/ztc`.
V Claude pluginu má stejné chování `/ztc-009:ztc`, v Codexu explicitní vyvolání `$ztc`.
Citovaný příkaz, popis funkce ani pouhé načtení tohoto souboru nejsou pokynem k přepnutí.

## Hostitel s kontrolorem

Hook zpracuje přesný uživatelský příkaz `/ztc` a předá nový stav v `ZTČ-GUARD`.
Stav je zdroj pravdy. Sám nevolejte shell pro přepnutí ani nesimulujte událost hooku.
Pokud klient příkaz rozbalil před hookem a přepnutí nepotvrdil, požádejte o běžnou zprávu `ztc toggle`.
Pouhé vyvolání tohoto skillu není důkaz, že se stav změnil.

- Při potvrzeném vypnutí odpovězte: „ZTČ-009 je pro tuto relaci vypnuté.“
- Při potvrzeném zapnutí použijte známá témata a [odpovědní skill](../ztc-response/SKILL.md).
  Chybějící témata zjistěte jednou krátkou otázkou. Kontrolu kvůli otázce nevypínejte.
- Uložený profil, globální instrukce ani stav jiné relace nepřepisujte.
- Výjimky pro tento tah schvaluje uživatel příkazy popsanými v [návodu](../../docs/ENFORCEMENT.md).

## Textový režim bez kontroloru

Bez kontextu `ZTČ-GUARD` je dostupné pouze instrukční přepínání, nikoli vynucení hookem.
Poslední relační volba má přednost před profilem. Bez aktivní volby začíná stav vypnutý.
Přepněte jej jednou a udržujte `ztc_session_enabled` pouze v kontextu chatu.
Triviální odpověď, citovaný příkaz ani opětovné načtení skillu stav nepřepínají.
Při zapnutí bez témat se na ně jednou zeptejte; do potvrzení styl neaplikujte.
Ztrátu kontextu nenahrazujte odhadem stavu. Uživatel musí potvrdit nové nastavení.

Zapnutí zachovává výjimku pro triviální jednověté odpovědi a výběr témat.
Jednorázový převod může uživatel v kontrolovaném režimu zapnout příkazem `/ztc once`; stav celé relace se nemění.
Přepnutí nezahajuje tvorbu slovníku, kontrolu GitHubu ani nabídku PR.
