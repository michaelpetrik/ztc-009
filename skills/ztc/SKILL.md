---
name: ztc
description: Přepněte ZTČ-009 pro aktuální relaci příkazem /ztc. Zapnutý stav vypne, vypnutý zapne. Nemění uložený profil ani jiné relace.
---

# Přepínač ZTČ-009 pro relaci

Přepněte stav pouze při přímém vyvolání tohoto skillu uživatelem nebo jeho samostatném příkazu `/ztc`.
V Claude pluginu má stejné chování `/ztc-009:ztc`, v Codexu explicitní vyvolání `$ztc`.
Citovaný příkaz, popis funkce ani pouhé načtení tohoto souboru nejsou pokynem k přepnutí.

1. Zjistěte poslední stav ZTČ v této relaci. Relační změna má přednost před uloženým profilem.
   Bez relační změny použijte potvrzené nastavení z chatu nebo dostupného profilu.
   Bez aktivního nastavení je výchozí stav vypnutý.
   Triviální nebo vyloučený dotaz nemění stav přepínače.
2. Přepněte zapnuto na vypnuto nebo vypnuto na zapnuto právě jednou pro tento uživatelský příkaz.
   Při opětovném načtení skillu v témže tahu nepřepínejte znovu.
3. Nový stav uchovejte jen v kontextu této relace jako `ztc_session_enabled: true` nebo `false`.
   Neupravujte `.ztc/profile.md`, globální instrukce, jiné relace ani soubory pluginu.
   Při shrnutí kontextu zachovejte poslední stav i témata; nová prázdná relace tuto výjimku nedědí.
4. Při vypnutí odpovězte: „ZTČ-009 je pro tuto relaci vypnuté.“
   Další běžné odpovědi nepřebírají zapnutí z uloženého profilu.
5. Při zapnutí zachovejte známá potvrzená témata, výjimky a slovníky.
   Pokud témata chybí, ponechte přepínač zapnutý a jednou se zeptejte na profesi a témata.
   Do potvrzení témat styl neaplikujte. Další `/ztc` tento čekající stav vypne.
   Pokud témata znáte, načtěte [odpovědní skill](../ztc-response/SKILL.md) a potřebný LLM profil.
   Po úspěšném načtení potvrďte: „ZTČ-009 je pro tuto relaci zapnuté.“
   Chybějící pravidla oznamte; dostupnost plného režimu nepředstírejte.

Zapnutí zachovává výjimku pro triviální jednověté odpovědi a výběr témat.
Výslovný převod konkrétního textu může styl použít i při vypnutém přepínači, ale stav přepínače nemění.
Přepnutí nezahajuje instalaci, tvorbu slovníku, kontrolu GitHubu ani nabídku PR.
Pokud předchozí stav není po ztrátě kontextu dostupný, zjistěte jej od uživatele; netvrďte, že jej znáte.
