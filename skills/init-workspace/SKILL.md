---
name: init-workspace
description: Inicializa a estrutura de memória exclusiva do projeto atual sem alterar código da aplicação.
argument-hint: [opcional: descrição curta da empresa ou projeto]
disable-model-invocation: true
---


## Regras obrigatórias

- Seja agnóstico à tecnologia.
- Detecte a stack antes de recomendar comandos.
- Não troque linguagem, framework, banco, ORM ou arquitetura sem autorização.
- Memória de projeto só pode ser escrita em `<raiz-do-projeto>/.claude/`.
- Nunca escreva dados de projeto em `~/.claude/skills/`.
- Não armazene cadeia de pensamento, raciocínio privado ou conversa integral.
- Use contexto progressivo: leia primeiro somente arquivos de alta relevância.
- Pare e pergunte quando houver ambiguidade material, risco destrutivo ou mudança de escopo.


# Objetivo

Criar o workspace `.claude/` na raiz real do projeto atual.

## Segurança de localização

1. Identifique a raiz do projeto pelo repositório Git ou pelos arquivos principais.
2. Mostre o caminho absoluto detectado.
3. Nunca use `~/.claude/` como destino.
4. Não prossiga se a raiz estiver ambígua.

## Estrutura

Crie, sem sobrescrever conteúdo existente:

```text
.claude/
├── CLAUDE.md
├── context.md
├── context-budget.md
├── architecture.md
├── company.md
├── decisions.md
├── memory/index.md
├── knowledge/.gitkeep
├── history/.gitkeep
├── specs/.gitkeep
├── plans/.gitkeep
├── preflights/.gitkeep
├── reviews/.gitkeep
└── releases/.gitkeep
```

Use como base os modelos deste framework quando disponíveis. Em arquivos já
existentes, preserve o conteúdo e apenas acrescente seções claramente ausentes.

Finalize com `WORKSPACE_STATUS: INITIALIZED` ou `WORKSPACE_STATUS: BLOCKED`.

## Uso automático

Esta Skill pode ser aplicada pelo `/implement` quando uma solicitação de alteração
for explícita e o projeto ainda não possuir `.claude/`.

Ela não deve ser executada automaticamente para perguntas comuns, explicações,
consultas ou análises sem alteração.
