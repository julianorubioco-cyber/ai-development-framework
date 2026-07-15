---
name: review
description: Revisa a implementação de forma independente e exige evidência por critério.
argument-hint: <spec e plano>
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


# Revisão posterior

Não confie no relatório do build. Inspecione diretamente alterações, comportamento,
testes, lint, tipagem e build relevantes.

Para cada critério, registre:

- critério;
- evidência;
- arquivo ou teste;
- resultado;
- status: `MET`, `NOT_MET` ou `NOT_VERIFIED`.

Problemas devem ir para `.claude/reviews/<slug>.review.md` com ID, gravidade,
categoria, localização, esperado, observado, evidência, correção e validação.

Categorias: functional, test, security, quality, documentation, architecture,
regression, scope, build e dependency.

Finalize exatamente com um estado:

- `REVIEW_STATUS: APPROVED`
- `REVIEW_STATUS: CHANGES_REQUIRED`
- `REVIEW_STATUS: BLOCKED`

Não aprove com critério importante não verificado.
