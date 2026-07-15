#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
TEMPLATE="$REPO_ROOT/templates/project-workspace/.claude"
TARGET="$PWD/.claude"

mkdir -p "$TARGET"
cp -Rn "$TEMPLATE/." "$TARGET/"
echo "Workspace inicializado em: $TARGET"
