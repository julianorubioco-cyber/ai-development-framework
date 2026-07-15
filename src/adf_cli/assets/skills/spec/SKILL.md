---
name: spec
description: Conversa com o usuário e cria uma especificação verificável sem implementar.
argument-hint: <descrição da tarefa ou caminho de artefato>
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

Entender objetivo, escopo, regras, casos extremos e critérios de aceitação.

Faça perguntas apenas sobre ambiguidades que alterem o resultado. Em tarefas
pequenas, mantenha a especificação curta. Salve em:

`.claude/specs/<slug>.md`

Inclua:

- objetivo;
- dentro e fora do escopo;
- requisitos;
- restrições;
- casos extremos;
- critérios de aceitação verificáveis;
- decisões pendentes;
- riscos conhecidos.

Não altere código. Finalize com `SPEC_STATUS: READY` ou `SPEC_STATUS: BLOCKED`.
