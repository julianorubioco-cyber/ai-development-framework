# Modo de compatibilidade

O ADF pode coexistir com frameworks e Skills que já organizam memória, briefing,
identidade, marketing ou conhecimento do projeto.

## Problema evitado

Sem compatibilidade, o projeto poderia ter duas fontes para o mesmo fato:

```text
__memoria/empresa.md
.claude/company.md
```

Quando uma fosse atualizada e a outra não, o agente receberia informações
contraditórias.

## Detecção

O comando:

```text
adf detect
```

procura:

- `__memoria/`;
- `_memoria/`;
- `memoria/`;
- `memory/`;
- `knowledge/`;
- `CLAUDE.md` na raiz;
- `identidade/`;
- `marketing/`;
- `dados/`;
- `saidas/`.

Uma estrutura com `__memoria/`, `CLAUDE.md` e pastas de negócio é reconhecida
como compatível com o padrão observado no MazyOS.

## Comportamento

### Projeto sem outra estrutura

O ADF cria seu workspace completo em `.claude/`.

### Projeto com memória existente

O ADF:

- reutiliza a memória existente;
- mantém `CLAUDE.md` da raiz como instrução canônica;
- não cria `company.md`, `context.md`, `decisions.md` ou `memory/` paralelos;
- cria somente artefatos técnicos:

```text
.claude/
├── compatibility.json
├── specs/
├── plans/
├── preflights/
├── reviews/
├── releases/
└── history/
```

## Fonte de verdade

Em modo de compatibilidade:

| Tipo | Fonte |
|---|---|
| Empresa e negócio | memória existente |
| Preferências | memória existente |
| Estratégia | memória existente |
| Instruções gerais | `CLAUDE.md` da raiz |
| Especificações técnicas | `.claude/specs/` |
| Planos | `.claude/plans/` |
| Revisões | `.claude/reviews/` |
| Releases | `.claude/releases/` |

## Uso

Primeiro diagnostique:

```text
adf detect
```

Depois inicialize:

```text
adf init
```

O resultado mostra os arquivos criados, preservados e ignorados para evitar
duplicação.
