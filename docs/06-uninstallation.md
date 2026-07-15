# Desinstalação e restauração

O instalador da versão 0.2.0 cria:

```text
~/.claude/adf-install-manifest.json
```

Esse manifesto registra somente as Skills instaladas pelo ADF. O desinstalador
não deve apagar outras Skills pessoais.

## Windows

Simular:

```powershell
.\scripts\uninstall.ps1 -WhatIf
```

Remover o ADF:

```powershell
.\scripts\uninstall.ps1
```

Remover e restaurar as Skills anteriores:

```powershell
.\scripts\uninstall.ps1 -RestoreBackup
```

## macOS/Linux

Simular:

```bash
./scripts/uninstall.sh --dry-run
```

Remover:

```bash
./scripts/uninstall.sh
```

Remover e restaurar backup:

```bash
./scripts/uninstall.sh --restore-backup
```

## Remoção manual de uma instalação antiga

Instalações anteriores à versão 0.2.0 podem não ter manifesto. Nesse caso, abra:

```text
%USERPROFILE%\.claude\skills
```

no Windows, ou:

```text
~/.claude/skills
```

em macOS/Linux, e apague somente estas pastas:

```text
build
context
docs
implement
init-workspace
loop
memory
plan
preflight
refactor
release
review
security
spec
test
```

Não apague a pasta `.claude` inteira. A pasta `.claude` dentro dos projetos
também não precisa ser removida; ela contém a memória de cada projeto.
