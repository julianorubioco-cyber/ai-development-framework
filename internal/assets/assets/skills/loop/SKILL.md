---
name: loop
description: Alias compatível do comando implement, preservando o mesmo gate e pipeline adaptativo.
argument-hint: <solicitação de desenvolvimento>
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


Recomende brevemente `/implement` para novos fluxos e siga integralmente as
instruções da Skill implement. Não pule preflight, aprovação, revisão ou memória.
Finalize com o mesmo estado do implement.
