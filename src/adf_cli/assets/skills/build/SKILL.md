---
name: build
description: Implementa somente o escopo aprovado, respeitando o projeto e registrando resultados.
argument-hint: <spec ou plano aprovado>
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


# Gate

Em fluxo orquestrado, exija um preflight com `PREFLIGHT_STATUS: APPROVED`.
Sem isso, não altere arquivos e retorne `BUILD_STATUS: BLOCKED_BY_PREFLIGHT`.

# Implementação

- siga especificação e plano;
- faça mudanças mínimas e coerentes;
- não amplie escopo;
- não execute operações irreversíveis sem confirmação;
- registre arquivos criados, alterados e removidos;
- execute verificações apropriadas à stack.

Finalize com `BUILD_STATUS: COMPLETED` ou `BUILD_STATUS: BLOCKED`.
