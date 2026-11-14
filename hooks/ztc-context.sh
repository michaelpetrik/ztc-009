#!/bin/sh
if [ -z "${CLAUDE_PLUGIN_ROOT:-}" ] || [ ! -f "$CLAUDE_PLUGIN_ROOT/skills/ztc-response/SKILL.md" ]; then
  exit 0
fi
printf '%s\n' 'ZTČ: respektujte potvrzený profil a zvolená témata. Výchozí profil je .ztc/profile.md v projektu; jinou zvolenou cestu určují instrukce hostitele.'
printf '%s\n' 'Bez potvrzeného rozsahu odpovídejte běžně a instalaci zahajte jen na žádost. Mimo rozsah, při vypnutí nebo u triviálního dotazu zodpověditelného jednou větou odpovídejte běžně. Výslovný požadavek na ZTČ pro danou odpověď má přednost před těmito výjimkami.'
printf '%s\n' 'Pouze při aktivním ZTČ načtěte skill z této cesty:'
printf '"%s"\n' "$CLAUDE_PLUGIN_ROOT/skills/ztc-response/SKILL.md"
