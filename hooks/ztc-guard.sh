#!/bin/sh
guard="${ZTC_GUARD_BIN:-$HOME/.local/bin/ztcguard}"
if [ ! -x "$guard" ]; then
  printf '%s\n' '{"continue":false,"stopReason":"ZTČ-009: chybí spustitelný ztcguard. Dokončete instalaci podle docs/ENFORCEMENT.md.","systemMessage":"ZTČ-009: kontrola odpovědí není dostupná. Tah je zastaven."}'
  exit 0
fi
if output=$("$guard" --action "$1"); then
  printf '%s\n' "$output"
else
  printf '%s\n' '{"continue":false,"stopReason":"ZTČ-009: kontrolor skončil chybou.","systemMessage":"ZTČ-009: kontrola selhala. Tah je zastaven."}'
fi
