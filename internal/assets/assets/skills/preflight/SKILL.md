---
name: preflight
description: Valida entendimento, plano, riscos e testabilidade antes de qualquer alteração.
argument-hint: <spec, plano ou descrição>
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


# Responsabilidade

Pensar antes de agir sem registrar raciocínio privado.

Execute até três rodadas de verificação sem novas informações:

- objetivo e escopo estão claros;
- decisões do usuário estão resolvidas;
- stack e arquitetura não estão sendo presumidas;
- mudanças destrutivas e custos estão identificados;
- arquivos prováveis e estratégia de testes estão definidos;
- cada critério de aceitação é verificável.

Se houver dúvida material, faça perguntas e encerre com
`PREFLIGHT_STATUS: BLOCKED`.

Quando estiver pronto, crie `.claude/preflights/<slug>.preflight.md` com
entendimento, escopo, riscos, estratégia e critérios. Mostre o resumo e pergunte:

`Posso iniciar a implementação conforme este entendimento?`

Encerre com `PREFLIGHT_STATUS: AWAITING_APPROVAL`. Após resposta explícita,
registre a aprovação no arquivo e produza `PREFLIGHT_STATUS: APPROVED`.

Nunca implemente.
