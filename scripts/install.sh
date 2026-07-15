#!/usr/bin/env bash
set -euo pipefail

DRY_RUN=0
[[ "${1:-}" == "--dry-run" ]] && DRY_RUN=1

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SOURCE="$REPO_ROOT/skills"
CLAUDE_ROOT="$HOME/.claude"
TARGET="$CLAUDE_ROOT/skills"
TIMESTAMP="$(date +%Y%m%d-%H%M%S)"
BACKUP="$CLAUDE_ROOT/adf-backups/$TIMESTAMP"
MANIFEST="$CLAUDE_ROOT/adf-install-manifest.json"
VERSION="$(tr -d '[:space:]' < "$REPO_ROOT/VERSION")"

mapfile -t SKILLS < <(find "$SOURCE" -mindepth 1 -maxdepth 1 -type d -printf '%f\n' | sort)

echo "ADF $VERSION"
echo "Destino: $TARGET"

if [[ "$DRY_RUN" -eq 1 ]]; then
  echo "[SIMULAÇÃO] Nenhum arquivo será alterado."
  for name in "${SKILLS[@]}"; do
    echo "[SIMULAÇÃO] Instalaria/substituiria: $TARGET/$name"
  done
  exit 0
fi

mkdir -p "$TARGET"
existing=()
for name in "${SKILLS[@]}"; do
  [[ -d "$TARGET/$name" ]] && existing+=("$name")
done

if [[ "${#existing[@]}" -gt 0 ]]; then
  mkdir -p "$BACKUP"
  for name in "${existing[@]}"; do
    cp -R "$TARGET/$name" "$BACKUP/"
  done
  echo "Backup criado em: $BACKUP"
fi

for name in "${SKILLS[@]}"; do
  rm -rf "$TARGET/$name"
  cp -R "$SOURCE/$name" "$TARGET/$name"
done

python3 - "$MANIFEST" "$VERSION" "$BACKUP" "${SKILLS[@]}" <<'PY'
import json, sys
from datetime import datetime, timezone
manifest, version, backup, *skills = sys.argv[1:]
data = {
    "framework": "AI Development Framework",
    "version": version,
    "installed_at": datetime.now(timezone.utc).isoformat(),
    "skills": skills,
    "backup_path": backup if backup else None,
}
with open(manifest, "w", encoding="utf-8") as f:
    json.dump(data, f, indent=2)
PY

echo "Skills instaladas em: $TARGET"
echo "Manifesto criado em: $MANIFEST"
