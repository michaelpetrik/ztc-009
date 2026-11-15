#!/bin/sh
guard="${ZTC_GUARD_BIN:-$HOME/.local/bin/ztcguard}"
if [ ! -x "$guard" ]; then
  printf '%s\n' '{"continue":false,"stopReason":"ZTČ-009: chybí spustitelný ztcguard. Dokončete instalaci podle docs/ENFORCEMENT.md.","systemMessage":"ZTČ-009: kontrola odpovědí není dostupná. Tah je zastaven."}'
  exit 0
fi
if output=$("$guard" --action "$1") && verified=$(printf '%s' "$output" | "$guard" --action validate-output); then
  case "$verified" in
    'ZTC-GUARD/1 '{*'}') printf '%s\n' "${verified#ZTC-GUARD/1 }"; exit 0 ;;
  esac
fi
printf '%s\n' '{"continue":false,"stopReason":"ZTČ-009: kontrolor selhal nebo vrátil neplatný protokol.","systemMessage":"ZTČ-009: kontrola selhala. Tah je zastaven."}'
