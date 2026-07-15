---
name: implement
description: Orquestra o ciclo completo: entendimento, aprovação, execução, revisão, correção, release e memória.
argument-hint: <solicitação de desenvolvimento>
disable-model-invocation: true
---


## Economia de contexto

- Consulte primeiro índices e artefatos diretamente relacionados.
- Abra somente arquivos de alta relevância.
- Expanda a leitura apenas mediante dependência verificável.
- Não repita conteúdo já registrado em artefatos válidos.

## Regras obrigatórias

- Seja agnóstico à tecnologia.
- Detecte a stack antes de recomendar comandos.
- Não troque linguagem, framework, banco, ORM ou arquitetura sem autorização.
- Memória de projeto só pode ser escrita em `<raiz-do-projeto>/.claude/`.
- Nunca escreva dados de projeto em `~/.claude/skills/`.
- Não armazene cadeia de pensamento, raciocínio privado ou conversa integral.
- Use contexto progressivo: leia primeiro somente arquivos de alta relevância.
- Pare e pergunte quando houver ambiguidade material, risco destrutivo ou mudança de escopo.


# Comando principal


## Inicialização automática do workspace

Antes de iniciar o fluxo:

1. localize a raiz real do projeto;
2. verifique se `<raiz>/.claude/` existe;
3. se não existir, aplique o contrato da Skill `init-workspace`;
4. preserve qualquer conteúdo existente;
5. nunca crie memória em `~/.claude/`;
6. continue o pipeline somente após a inicialização segura.

A ausência de workspace não exige uma pergunta extra quando a solicitação já for
uma alteração explícita ou quando o usuário tiver chamado `/implement`.


Você coordena as Skills; não finja que criou processos separados se não puder
invocá-los. Aplique os contratos de cada etapa de forma explícita.

## Classificação

Classifique por tamanho e risco:

- PEQUENA: mudança localizada, baixo risco, poucos critérios.
- MÉDIA: múltiplos arquivos ou integração limitada.
- GRANDE/ALTO RISCO: novo módulo, arquitetura, dados, autenticação, pagamentos,
  segurança, migração ou impacto amplo.

Explique em uma frase a classificação.

## Fluxos mínimos

PEQUENA:
`SPEC → PREFLIGHT → aprovação → BUILD → REVIEW → MEMORY`

MÉDIA:
`SPEC → CONTEXT → PLAN → PREFLIGHT → aprovação → BUILD → TEST → REVIEW → MEMORY`

GRANDE/ALTO RISCO:
`SPEC → CONTEXT → PLAN → PREFLIGHT → aprovação → BUILD → TEST → REVIEW
→ SECURITY → REFACTOR se necessário → DOCS → RELEASE → MEMORY`

O risco prevalece sobre o tamanho.

## Gate

Não altere código antes de aprovação explícita. A aprovação da spec não equivale
à aprovação do preflight.

## Correções

Quando `REVIEW_STATUS: CHANGES_REQUIRED`, roteie:

- functional, regression, build, architecture → build;
- test → test;
- security → security;
- quality → refactor;
- documentation → docs;
- scope → parar e pedir aprovação.

Refaça uma revisão completa. Máximo de oito ciclos.

Nova aprovação é obrigatória para mudança de escopo, exclusão de dados, migração
destrutiva, custo externo, troca de tecnologia ou operação irreversível.

## Conclusão

Só finalize quando os gates aplicáveis tiverem evidência. Para fluxo grande,
exija também `READY_FOR_RELEASE`. Depois execute a curadoria de memória.

Finalize com `IMPLEMENT_STATUS: COMPLETED` ou `IMPLEMENT_STATUS: BLOCKED`.
