# AI Development Framework (ADF)

Framework de desenvolvimento orientado por Skills para Claude Code.

O ADF organiza o trabalho do agente em etapas independentes: compreensão, contexto,
planejamento, validação antes de agir, implementação, testes, revisão por evidências,
segurança, documentação, release e memória persistente por projeto.

## Princípios

- Skills reutilizáveis ficam globalmente em `~/.claude/skills/`.
- Conhecimento de clientes nunca é salvo nas Skills globais.
- A memória operacional fica somente em `<projeto>/.claude/`.
- Nenhuma implementação começa sem aprovação explícita no preflight.
- A revisão posterior exige evidências verificáveis.
- O pipeline se adapta ao tamanho e ao risco da tarefa.
- O agente lê o menor contexto necessário.
- A memória é curada por prioridade e organizada por domínio.
- Cadeia de pensamento e raciocínio privado nunca são armazenados.

## Comando principal

Depois de instalar, abra um projeto no VS Code e execute:

```text
/implement criar um sistema de login com e-mail e senha
```

Fluxo principal:

```text
SPEC → CONTEXT → PLAN → PREFLIGHT → APROVAÇÃO DO USUÁRIO
→ BUILD → TEST → REVIEW → etapas necessárias → RELEASE → MEMORY
```

O comando `/loop` é mantido como alias compatível.

## Instalação rápida no Windows

No PowerShell, dentro deste repositório:

```powershell
Set-ExecutionPolicy -Scope Process Bypass
.\scripts\install.ps1
```

## Instalação rápida no macOS/Linux

```bash
chmod +x scripts/install.sh
./scripts/install.sh
```

O instalador copia as Skills para `~/.claude/skills/`, preservando backups quando
encontra arquivos existentes.

## Inicializar o workspace de um projeto

Abra o projeto no terminal e execute:

```text
/init-workspace
```

Isso cria somente dentro do projeto:

```text
.claude/
├── CLAUDE.md
├── context.md
├── architecture.md
├── company.md
├── decisions.md
├── memory/
│   └── index.md
├── knowledge/
├── history/
├── specs/
├── plans/
├── preflights/
├── reviews/
└── releases/
```

## Pipeline adaptativo

### Pequeno
`SPEC → PREFLIGHT → BUILD → REVIEW → MEMORY`

### Médio
`SPEC → CONTEXT → PLAN → PREFLIGHT → BUILD → TEST → REVIEW → MEMORY`

### Grande ou de alto risco
`SPEC → CONTEXT → PLAN → PREFLIGHT → BUILD → TEST → REVIEW → SECURITY
→ REFACTOR quando necessário → DOCS → RELEASE → MEMORY`

A classificação não deve ser usada para pular controles de segurança. Uma alteração
pequena, mas destrutiva ou sensível, recebe fluxo de alto risco.

## Documentação

- [Arquitetura](docs/01-architecture.md)
- [Orquestração](docs/02-orchestration.md)
- [Memória por projeto](docs/03-project-memory.md)
- [Economia de contexto](docs/04-context-economy.md)
- [Instalação e uso](docs/05-installation.md)
- [Desinstalação e restauração](docs/06-uninstallation.md)
- [Desenvolvimento e contribuições](CONTRIBUTING.md)

## Estado do projeto

Versão atual: `v0.2.0`.

Esta versão é intencionalmente conservadora. O framework fornece instruções e
contratos operacionais para o Claude Code; ele não substitui testes, permissões,
revisão humana ou controles de implantação.
