---
name: test
description: Cria, atualiza e executa testes relevantes sem alterar requisitos de negócio.
argument-hint: <spec, plano ou módulo>
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


Use a infraestrutura existente. Não introduza framework de testes sem autorização.
Priorize testes dos critérios de aceitação e regressões prováveis. Não mude a
regra de negócio para fazer o teste passar.

Registre comandos, resultados e lacunas. Finalize com `TEST_STATUS: PASSED`,
`TEST_STATUS: FAILED` ou `TEST_STATUS: BLOCKED`.
