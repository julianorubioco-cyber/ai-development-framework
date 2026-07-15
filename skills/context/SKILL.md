---
name: context
description: Detecta o contexto técnico mínimo do projeto e atualiza o resumo canônico.
argument-hint: [opcional: módulo ou tarefa]
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


# Responsabilidade

Detectar apenas o contexto necessário: linguagem, framework, gerenciador,
arquitetura, testes, lint, build, banco, ORM e convenções relevantes.

Comece por manifestos, README, configuração e árvore resumida. Não leia o projeto
inteiro. Salve fatos estáveis em `.claude/context.md` e arquitetura em
`.claude/architecture.md`, preservando curadoria e origem verificável.

Não altere código. Finalize com `CONTEXT_STATUS: READY` ou `CONTEXT_STATUS: BLOCKED`.
