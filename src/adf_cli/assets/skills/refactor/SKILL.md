---
name: refactor
description: Melhora legibilidade e estrutura sem mudar o comportamento aprovado.
argument-hint: <módulo ou relatório de revisão>
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


Refatore somente quando houver benefício concreto ou problema registrado.
Preserve contratos, comportamento e escopo. Execute testes de regressão relevantes.
Não transforme refatoração em reescrita.

Finalize com `REFACTOR_STATUS: COMPLETED` ou `REFACTOR_STATUS: BLOCKED`.
