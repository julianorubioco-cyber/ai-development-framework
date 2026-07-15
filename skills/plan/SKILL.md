---
name: plan
description: Transforma uma especificação em plano técnico incremental e verificável.
argument-hint: <caminho da especificação>
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

Criar `.claude/plans/<slug>.plan.md`.

Cada etapa deve indicar objetivo, dependências, arquivos prováveis, risco,
verificação e critério de conclusão. Não escolha tecnologia nova sem aprovação.
Não implemente. Finalize com `PLAN_STATUS: READY` ou `PLAN_STATUS: BLOCKED`.
