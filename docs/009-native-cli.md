# CLI nativa

A partir da versão 0.6.0, a distribuição principal do ADF usa uma CLI compilada
em Go. O usuário final não precisa instalar Python, Go ou Git.

## Plataformas

- Windows x64 e ARM64;
- Linux x64 e ARM64;
- macOS Intel e Apple Silicon.

## Instalação

### Windows

```powershell
irm https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install.ps1 | iex
```

### macOS/Linux

```bash
curl -fsSL https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install.sh | sh
```

O instalador consulta a última GitHub Release, identifica sistema e arquitetura,
baixa somente o executável correto, instala as Skills embutidas e executa
`adf doctor`.

## Comandos

```text
adf install
adf init
adf doctor
adf uninstall
adf version
```

## Distribuição

Uma tag como `v0.6.0` inicia o workflow de release. O GitHub Actions compila seis
binários, gera checksums e publica todos na GitHub Release.

## Segurança operacional

- Skills anteriores recebem backup;
- o manifesto controla a desinstalação;
- `adf init` preserva arquivos existentes;
- memória continua exclusivamente no projeto.
