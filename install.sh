#!/usr/bin/env sh
set -eu

REPO="julianorubioco-cyber/ai-development-framework"
VERSION="${ADF_VERSION:-latest}"
INSTALL_DIR="${ADF_INSTALL_DIR:-$HOME/.local/bin}"

os="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$os" in
  darwin) os="darwin" ;;
  linux) os="linux" ;;
  *) echo "Sistema não suportado: $os" >&2; exit 1 ;;
esac

arch="$(uname -m)"
case "$arch" in
  x86_64|amd64) arch="amd64" ;;
  arm64|aarch64) arch="arm64" ;;
  *) echo "Arquitetura não suportada: $arch" >&2; exit 1 ;;
esac

if [ "$VERSION" = "latest" ]; then
  VERSION="$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" |
    awk -F'"' '/"tag_name":/{print $4; exit}')"
fi

asset="adf_${os}_${arch}"
url="https://github.com/$REPO/releases/download/$VERSION/$asset"

mkdir -p "$INSTALL_DIR"
tmp="$INSTALL_DIR/.adf.download"
curl -fsSL "$url" -o "$tmp"
chmod +x "$tmp"
mv "$tmp" "$INSTALL_DIR/adf"

"$INSTALL_DIR/adf" install
"$INSTALL_DIR/adf" doctor

echo
echo "ADF instalado em: $INSTALL_DIR/adf"
case ":$PATH:" in
  *":$INSTALL_DIR:"*) ;;
  *) echo "Adicione ao PATH: export PATH=\"$INSTALL_DIR:\$PATH\"" ;;
esac
