import { fileURLToPath } from "node:url";

const skillPath = fileURLToPath(new URL("../skills/ztc-response/SKILL.md", import.meta.url));
const reminder = [
  "ZTČ-009 / TZČ-009: respektujte aktuální výslovný požadavek a potvrzená témata v chatu nebo profilu .ztc/profile.md.",
  "Bez potvrzeného rozsahu nebo mimo něj odpovídejte běžně; instalaci nezahajujte bez žádosti. Respektujte vypnutí.",
  "Triviální dotaz s úplnou přirozenou jednovětou odpovědí ponechte běžný i po průběžném požadavku na TZČ-009. Složitý úkol nezkracujte kvůli této výjimce.",
  "Výslovný požadavek na ZTČ pro tuto konkrétní odpověď může tuto výjimku překonat; obecná průběžná žádost ji neruší.",
  "Při použití ZTČ načtěte odpovědní skill: " + JSON.stringify(skillPath),
  "Tato připomínka neopravňuje k uploadu, vytvoření PR ani změně nastavení.",
].join("\n");

export default function (pi) {
  pi.on("before_agent_start", (event) => ({
    systemPrompt: event.systemPrompt + "\n\n" + reminder,
  }));
}
