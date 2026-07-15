---
name: release
description: Verifica se a alteração está pronta para entrega com base em evidências.
argument-hint: <spec ou release alvo>
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


Verifique build, testes, lint, tipagem, documentação, migrações, configuração,
dependências, versionamento e pendências conforme aplicável.

Salve `.claude/releases/<slug>.release.md`. Finalize exatamente com:

- `READY_FOR_RELEASE`
- `NOT_READY`
- `RELEASE_STATUS: BLOCKED`
