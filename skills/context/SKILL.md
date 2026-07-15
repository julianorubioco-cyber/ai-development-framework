---
name: context
description: Detecta o contexto técnico mínimo do projeto e atualiza o resumo canônico.
argument-hint: [opcional: módulo ou tarefa]
---


## Economia de contexto

- Consulte primeiro índices e artefatos diretamente relacionados.
- Abra somente arquivos de alta relevância.
- Expanda a leitura apenas mediante dependência verificável.
- Não repita conteúdo já registrado em artefatos válidos.

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

Antes de criar ou atualizar contexto, detecte se o projeto já possui uma fonte
de verdade externa, como `__memoria/` ou `CLAUDE.md` na raiz.

Em modo de compatibilidade:

- leia a estrutura existente;
- não crie `.claude/context.md` duplicado;
- produza contexto técnico apenas no plano ou artefato da tarefa;
- preserve a fonte de verdade do projeto.

Não altere código. Finalize com `CONTEXT_STATUS: READY` ou `CONTEXT_STATUS: BLOCKED`.
