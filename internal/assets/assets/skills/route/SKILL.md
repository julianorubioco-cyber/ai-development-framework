---
name: route
description: Classifica uma solicitação sem alterar o projeto e informa qual fluxo seria apropriado.
argument-hint: <solicitação>
disable-model-invocation: true
---


## Regras obrigatórias

- Seja agnóstico à tecnologia.
- Use o menor contexto suficiente.
- Não altere arquivos sem intenção explícita do usuário.
- Memória de projeto só pode existir em `<raiz-do-projeto>/.claude/`.
- Não armazene cadeia de pensamento.

# Responsabilidade

Classifique a solicitação em:

- `QUESTION`
- `ANALYSIS_ONLY`
- `AMBIGUOUS_ACTION`
- `PROJECT_CHANGE`
- `EXPLICIT_IMPLEMENT`

Não abra arquivos além do estritamente necessário para classificar.

Saída:

```text
ROUTE_TYPE: <categoria>
RECOMMENDED_FLOW: <direct-answer | analysis | ask-confirmation | implement>
REASON: <uma frase curta>
```

Nunca modifique arquivos.
