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


## Instalação em um comando

A CLI é um executável nativo. O usuário não precisa instalar Python, Go ou Git.

### Windows PowerShell

```powershell
irm https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install.ps1 | iex
```

### macOS/Linux

```bash
curl -fsSL https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install.sh | sh
```

Depois:

```text
adf doctor
adf init
```


## Instalação manual alternativa no Windows

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




## Atualização automática

```text
adf update
```

ou:

```text
adf self-update
```

A CLI baixa a última Release, valida o checksum e substitui o executável correto
para o sistema operacional.


## Compatibilidade com estruturas existentes

O ADF não cria uma segunda memória quando o projeto já usa outro sistema.

```text
adf detect
adf init
```

Se encontrar `__memoria/`, `CLAUDE.md`, `identidade/`, `marketing/` ou outra
estrutura reconhecida, ele ativa o modo de compatibilidade e cria somente os
artefatos técnicos que faltarem.


## Conversa normal ou `/implement`

Você não precisa usar `/implement` para perguntas comuns.

Exemplos que **não** iniciam o pipeline:

```text
O que é JWT?
Como funciona este módulo?
Qual arquivo controla o login?
Explique este erro.
```

Exemplos que iniciam o fluxo de implementação:

```text
Crie uma tela de login.
Corrija o erro do checkout.
Adicione pagamentos com Stripe.
```

Pedidos ambíguos, como “dá para melhorar essa tela?”, devem gerar uma explicação
curta e um pedido de confirmação antes de qualquer alteração.

O comando `/implement` sempre força o pipeline disciplinado.


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

- [Visão](docs/001-vision.md)
- [Princípios](docs/002-principles.md)
- [Arquitetura oficial](docs/003-architecture.md)
- [Economia de tokens](docs/004-token-economy.md)
- [Terminologia](docs/005-terminology.md)
- [Processo de versões](docs/006-release-process.md)
- [Roteamento inteligente](docs/007-intent-routing.md)
- [CLI e instalação automática](docs/008-cli-installation.md)
- [CLI nativa](docs/009-native-cli.md)
- [Modo de compatibilidade](docs/010-compatibility-mode.md)
- [Autoatualização](docs/011-self-update.md)

- [Arquitetura](docs/01-architecture.md)
- [Orquestração](docs/02-orchestration.md)
- [Memória por projeto](docs/03-project-memory.md)
- [Economia de contexto](docs/04-context-economy.md)
- [Instalação e uso](docs/05-installation.md)
- [Desinstalação e restauração](docs/06-uninstallation.md)
- [Desenvolvimento e contribuições](CONTRIBUTING.md)

## Estado do projeto

Versão atual: `v0.8.0`.

Esta versão é intencionalmente conservadora. O framework fornece instruções e
contratos operacionais para o Claude Code; ele não substitui testes, permissões,
revisão humana ou controles de implantação.
