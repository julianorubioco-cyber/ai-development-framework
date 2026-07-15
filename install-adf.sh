#!/usr/bin/env sh
set -eu

REPOSITORY="${ADF_REPOSITORY:-https://github.com/julianorubioco-cyber/ai-development-framework.git}"

echo "AI Development Framework - instalador macOS/Linux"

if ! command -v python3 >/dev/null 2>&1; then
  echo "Python 3.9+ não foi encontrado." >&2
  exit 1
fi

python3 -m pip install --user --upgrade "git+$REPOSITORY"
python3 -m adf_cli.cli install
python3 -m adf_cli.cli doctor

echo
echo "Instalação concluída."
echo "Se 'adf' não estiver no PATH, use: python3 -m adf_cli.cli doctor"
