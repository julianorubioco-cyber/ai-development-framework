# Instalação

## Windows

```powershell
Set-ExecutionPolicy -Scope Process Bypass
.\scripts\install.ps1
```

## macOS/Linux

```bash
chmod +x scripts/install.sh
./scripts/install.sh
```

## Instalação manual

Copie cada diretório de `skills/` para:

```text
~/.claude/skills/
```

Depois abra uma nova sessão do Claude Code.

## Primeiro uso

1. Abra a raiz do projeto no VS Code.
2. Execute `/init-workspace`.
3. Revise os arquivos criados em `.claude/`.
4. Execute `/implement <solicitação>`.
5. Confirme somente depois de revisar o preflight.
