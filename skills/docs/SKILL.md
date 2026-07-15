---
name: docs
description: Atualiza somente a documentação afetada pela mudança aprovada.
argument-hint: <spec, release ou módulo>
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


Atualize README, contratos, endpoints, configuração ou documentação de operação
apenas quando necessário. Não gere documentação volumosa sem uso. Não exponha
segredos ou detalhes internos inadequados.

Finalize com `DOCS_STATUS: UPDATED`, `DOCS_STATUS: NOT_NEEDED` ou
`DOCS_STATUS: BLOCKED`.
