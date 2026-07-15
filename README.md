# AI Development Framework (ADF)

> Um framework open source para transformar o Claude Code em uma equipe
> completa de desenvolvimento de software.

O ADF adiciona uma camada de engenharia sobre o Claude Code através de
Skills, memória por projeto e uma CLI nativa (`adf`).

------------------------------------------------------------------------

# ✨ Recursos

-   🧠 Memória persistente por projeto
-   ⚡ Economia inteligente de contexto
-   🔍 Detecção automática da stack
-   📋 Especificação antes da implementação
-   🏗️ Planejamento técnico automático
-   🧪 Testes integrados
-   🔒 Revisão e segurança
-   📚 Documentação automática
-   🚀 Pipeline de Release
-   💻 CLI nativa para Windows, Linux e macOS

------------------------------------------------------------------------

# Instalação (30 segundos)

## Windows

``` powershell
irm https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install.ps1 | iex
```

## macOS / Linux

``` bash
curl -fsSL https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install.sh | sh
```

Depois execute:

``` text
adf doctor
```

------------------------------------------------------------------------

# Primeiro projeto

Entre na pasta do projeto:

``` text
cd MeuProjeto
```

Inicialize o workspace:

``` text
adf init
```

Abra a pasta no Claude Code e utilize:

``` text
/implement Criar sistema de login com e-mail e senha
```

Fluxo executado:

``` text
SPEC
 ↓
CONTEXT
 ↓
PLAN
 ↓
PREFLIGHT
 ↓
APROVAÇÃO DO USUÁRIO
 ↓
BUILD
 ↓
TEST
 ↓
REVIEW
 ↓
SECURITY
 ↓
REFACTOR
 ↓
DOCS
 ↓
RELEASE
```

------------------------------------------------------------------------

# Documentação

-   📘 INSTALL.md --- Guia completo de instalação
-   📗 docs/cli.md --- Comandos da CLI
-   📙 docs/skills.md --- Todas as Skills
-   📕 docs/memory.md --- Sistema de memória
-   📒 docs/architecture.md --- Arquitetura do framework
-   ❓ docs/faq.md --- Perguntas frequentes

------------------------------------------------------------------------

# Estrutura

``` text
AI-Development-Framework/
├── cmd/
├── internal/
├── skills/
├── templates/
├── docs/
├── tests/
├── install.ps1
├── install.sh
├── go.mod
└── README.md
```

------------------------------------------------------------------------

# Desenvolvimento

``` bash
git clone https://github.com/julianorubioco-cyber/ai-development-framework.git
cd ai-development-framework

git add .
git commit -m "feat: minha alteração"
git push
```

Para publicar uma nova versão:

``` bash
git tag v0.7.0
git push origin v0.7.0
```

O GitHub Actions executará testes, compilará os binários para Windows,
Linux e macOS e publicará automaticamente uma nova Release.

------------------------------------------------------------------------

# Roadmap

-   ✅ CLI nativa
-   ✅ Instalador automático
-   ✅ CI/CD
-   🔄 Auto Update (`adf self-update`)
-   🔄 Context Engine
-   🔄 Memory Engine
-   🔄 Plugin System
-   🔄 Marketplace de Skills

------------------------------------------------------------------------

# Licença

MIT

------------------------------------------------------------------------

Feito com ❤️ para tornar o desenvolvimento com IA mais organizado,
reproduzível e profissional.
