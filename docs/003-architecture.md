# Arquitetura oficial

## Visão em camadas

```text
Usuário
  ↓
Orquestrador (/implement)
  ↓
Skills especializadas
  ↓
Arquivos do projeto
  ↓
Workspace local (.claude/)
```

## Camada global

Local:

```text
~/.claude/skills/
```

Contém somente Skills reutilizáveis.

## Camada do projeto

Local:

```text
<projeto>/.claude/
```

Contém:

- contexto técnico;
- arquitetura;
- decisões;
- memória por domínio;
- especificações;
- planos;
- preflights;
- revisões;
- releases;
- histórico.

## Componentes

### Orquestrador

`/implement` classifica tamanho e risco e escolhe o fluxo mínimo seguro.

### Gate anterior

`/preflight` impede alterações antes de o entendimento ser revisado e aprovado.

### Execução

`/build` altera o projeto somente dentro do escopo aprovado.

### Gate posterior

`/review` verifica critérios por evidência e encaminha problemas.

### Curadoria

`/memory` registra somente conhecimento reutilizável, com prioridade e domínio.

## Estados principais

```text
PREFLIGHT_STATUS: APPROVED
BUILD_STATUS: COMPLETED
REVIEW_STATUS: APPROVED
READY_FOR_RELEASE
IMPLEMENT_STATUS: COMPLETED
```

Estados ausentes ou conflitantes não devem ser tratados como aprovação.
