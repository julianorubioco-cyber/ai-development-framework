# CLI e instalação automática

## Opção recomendada: um comando

### Windows PowerShell

```powershell
irm https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install-adf.ps1 | iex
```

### macOS/Linux

```bash
curl -fsSL https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install-adf.sh | sh
```

O instalador:

1. detecta o sistema operacional pelo script usado;
2. verifica Python;
3. instala ou atualiza a CLI;
4. copia as Skills para `~/.claude/skills/`;
5. cria backup das Skills existentes;
6. grava o manifesto da instalação;
7. executa diagnóstico.

## Opção com Git

```bash
git clone https://github.com/julianorubioco-cyber/ai-development-framework.git
cd ai-development-framework
python -m pip install --user .
adf install
```

No Windows, também é possível usar:

```powershell
python -m adf_cli.cli install
```

## Comandos

### Instalar Skills

```text
adf install
```

### Criar workspace no projeto atual

```text
adf init
```

### Diagnosticar

```text
adf doctor
```

### Atualizar clone Git e reinstalar

```text
adf update
```

### Desinstalar

```text
adf uninstall
```

### Restaurar Skills anteriores

```text
adf uninstall --restore-backup
```

### Simular sem alterar

```text
adf install --dry-run
adf init --dry-run
adf uninstall --dry-run
```

## Dependência

A CLI v0.5.0 usa somente a biblioteca padrão do Python e requer Python 3.9+.

Uma versão futura poderá oferecer binários nativos para dispensar o Python.
