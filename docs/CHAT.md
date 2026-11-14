# ZTČ v Claude a ChatGPT

Základní způsob použití je textová instrukce v běžném chatu.
Plugin, terminál, soubory na disku ani GitHub nejsou podmínkou.
Název „ZTČ-009“ sám o sobě nezaručuje, že model zná pravidla tohoto projektu.

## Rychlé použití

Vložte do chatu tento samostatný text:

```text
Odpovídej mi v TZČ-009 podle následujícího stručného profilu ZTČ-009.
TZČ-009 je zde přijímaná varianta označení ZTČ-009 (Zjednodušená technická čeština).
Nejprve se jednou zeptej na profesi a témata, pro která chci tento styl používat, pokud je ještě neznáš.
Moje volby si ponech v kontextu tohoto chatu. Pokud řeknu „všechna témata“, respektuj to.
Piš krátké, konkrétní věty. Jedna věta má jedno hlavní sdělení.
V postupu používej vykání, rozkazovací způsob a nejvýše 20 slov ve větě. Jeden krok obsahuje jednu instrukci.
V popisu používej nejvýše 25 slov ve větě a nejvýše 6 vět v odstavci.
Používej činný rod. Podmínku napiš před instrukcí.
Pro jednu věc používej stejný název. Technickou zkratku poprvé vysvětli: CI (průběžná integrace).
Vynech slang, zbytečné výplně a zkratky jako „tj.“ nebo „např.“. Nepoužívej středník.
Zachovej význam, míru jistoty, rozdíl mezi doporučením a povinností, čísla, kód a doslovné citace.
U triviálního dotazu, který lze úplně zodpovědět jednou přirozenou větou, styl nevynucuj.
Složitý dotaz kvůli této výjimce nestlačuj do jedné věty. Mimo zvolená témata odpovídej běžně.
Výslovný požadavek na převod konkrétního textu do ZTČ platí i pro jednu větu.
Pokud chybí slovník mého oboru, nabídni postupnou tvorbu hesel s významem a příkladem přímo v chatu.
Neoznačuj vygenerované významy za odborně ověřené. Nikam je nenahrávej bez mého konkrétního souhlasu.
Na požádání změň témata nebo tento styl vypni. Neslibuj zapamatování v jiných chatech.
```

Tento text je stručný profil, nikoli náhrada všech pravidel specifikace.
Je použitelný bez dalších nástrojů a bez stahování souborů.

## Úplná pravidla

Pro plný režim vložte do chatu obsah [LLM profilu](../llm/ztc-009.llm.md).
Přidejte obsah [základního slovníku](../slovnik/ztc-009.csv) a vhodného oborového slovníku, pokud jej chcete používat.
Pro software použijte [softwarový slovník](../slovnik/software-development.csv).
Soubory lze také přiložit, pokud je dané prostředí umí číst. Samotný odkaz není důkaz načtení souboru.

Potom stačí například:

> Odpovídej mi v TZČ-009 pro vývoj softwaru. Triviální jednověté odpovědi nech přirozené.

Bez známých témat se model nejprve zeptá. Při uvedení témat v zadání otázku neopakuje.
Žádost „přepiš tuto větu do ZTČ“ platí jen pro daný text, nikoli automaticky pro další chat.

## Slovník bez souborového systému

Nový slovník lze připravit jako tabulku nebo CSV přímo v odpovědi.
Potvrzené významy a témata zůstávají v kontextu chatu. Model nesmí tvrdit, že je uložil na disk.
Při dlouhém nebo novém chatu předejte pravidla a potvrzená hesla znovu, pokud již nejsou dostupná.

Publikace je volitelná. Pokud ji uživatel chce, postupujte podle [návodu k příspěvku](INSTALL.md#dobrovolný-příspěvek-přes-pr).
Ověřte skutečné možnosti connectoru nebo jiného nástroje. Bez nich předejte obsah a návod k ručnímu PR.
Souhlas s používáním slovníku není souhlas s jeho zveřejněním.

## Trvalé použití a plugin

Pokud hostitel nabízí vlastní nebo projektové instrukce, lze do nich vložit stejný text.
Dostupnost a rozsah těchto funkcí závisí na prostředí. Tento návod nevyžaduje jejich existenci.
[Instalace pluginu](HARNESSES.md) je volitelné rozšíření podle podpory konkrétního prostředí.
Claude chat a Cowork mají vlastní upload cestu. Dostupnost v účtu a automatické hooky ověřte odděleně.
Claude Code není totožný s běžným chatem Claude. Codex není totožný s běžným ChatGPT.
