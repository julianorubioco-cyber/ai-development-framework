#!/usr/bin/env bash
set -euo pipefail

RESTORE=0
DRY_RUN=0
for arg in "$@"; do
  [[ "$arg" == "--restore-backup" ]] && RESTORE=1
  [[ "$arg" == "--dry-run" ]] && DRY_RUN=1
done

CLAUDE_ROOT="$HOME/.claude"
TARGET="$CLAUDE_ROOT/skills"
MANIFEST="$CLAUDE_ROOT/adf-install-manifest.json"

[[ -f "$MANIFEST" ]] || {
  echo "Manifesto não encontrado: $MANIFEST" >&2
  exit 1
}

python3 - "$MANIFEST" "$TARGET" "$RESTORE" "$DRY_RUN" <<'PY'
import json, os, shutil, sys
manifest_path, target, restore, dry_run = sys.argv[1:]
restore = restore == "1"
dry_run = dry_run == "1"

with open(manifest_path, encoding="utf-8") as f:
    data = json.load(f)

for name in data.get("skills", []):
    path = os.path.join(target, name)
    if os.path.isdir(path):
        print(("[SIMULAÇÃO] Removeria: " if dry_run else "Removendo: ") + path)
        if not dry_run:
            shutil.rmtree(path)

backup = data.get("backup_path")
if restore and backup and os.path.isdir(backup):
    for name in os.listdir(backup):
        src = os.path.join(backup, name)
        dst = os.path.join(target, name)
        if os.path.isdir(src):
            print(("[SIMULAÇÃO] Restauraria: " if dry_run else "Restaurando: ") + dst)
            if not dry_run:
                if os.path.isdir(dst):
                    shutil.rmtree(dst)
                shutil.copytree(src, dst)

if not dry_run:
    os.remove(manifest_path)
    print("Desinstalação concluída.")
PY
