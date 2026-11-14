# Jak přispět do ZTČ-009

Tento dokument popisuje, jak navrhnout změnu.
Dokument je napsaný v ZTČ-009.

## Základní pravidla

- Pište návrhy v ZTČ-009.
- Nekopírujte text z dokumentu ASD-STE100. Text ASD je chráněný autorským právem.
- Jeden návrh = jedno issue nebo jeden pull request.
- Každá změna pravidla musí mít minimálně jeden příklad ✗ a jeden příklad ✓.

## Typy návrhů

| Návrh | Postup |
| --- | --- |
| Nové slovo | Založte issue „Nové slovo". Potom můžete poslat pull request se změnou příslušného CSV souboru v `slovnik/`. |
| Změna pravidla | Založte issue „Změna pravidla". Počkejte na diskusi. Potom pošlete pull request. |
| Chyba v textu | Pošlete pull request přímo. |
| Otázka nebo nápad | Napište do Discussions. |

## Pull request

1. Vytvořte fork repozitáře.
2. Vytvořte novou větev.
3. Udělejte změnu.
4. Spusťte kontrolu:

   ```sh
   cd tools && go test ./... && go run ./cmd/ztcvalidate -root ..
   ```

5. Pošlete pull request. Vyplňte šablonu.

CI zkontroluje formát Markdownu, odkazy, slovník a číslování pravidel.
Pull request můžeme sloučit, jen když CI projde.

## Pravidla pro slovník

Přečtěte si [pravidla tvorby](slovnik/RULEBOOK.md) a [postup pro nový obor](slovnik/GUIDELINES.md).

- Jedno lemma má jen jeden schválený význam a jeden slovní druh.
- Neschválené slovo musí mít minimálně jednu schválenou náhradu.
- Přidejte nové schválené slovo, jen když žádné schválené slovo nemá potřebný význam.
- Formát sloupců je v souboru [slovnik/README.md](slovnik/README.md).

## Pravidla pro specifikaci

- Číslování pravidel 1.1 až 9.4 je pevné. Číslování odpovídá ASD-STE100 Issue 9.
- Nepřidávejte nová čísla pravidel. Českou úpravu napište do odstavce **[CZ]** u existujícího pravidla.
- Když změníte specifikaci, změňte také soubor `llm/ztc-009.llm.md`.
- Když změníte verzi, změňte verzi v obou souborech. CI kontroluje shodu verzí.

## Verze

Projekt používá sémantické verze (MAJOR.MINOR.PATCH).

| Změna | Verze |
| --- | --- |
| Oprava překlepu, příkladu | PATCH |
| Nové slovo, nové vysvětlení, nové doporučení | MINOR |
| Změna významu pravidla nebo limitu | MAJOR |

## Chování

Pište věcně a s respektem. Kritizujte text, ne osoby.
