# Memória por projeto

A memória existe exclusivamente em `<projeto>/.claude/`.

## Curadoria

Antes de gravar, a Skill `/memory` avalia se a informação continuará útil em
futuras implementações.

### Alta prioridade

Arquitetura, stack, integrações, regras permanentes, decisões técnicas ou de
negócio e restrições. Registrar nos documentos canônicos ou no domínio adequado.

### Média prioridade

Novos módulos, mudanças estruturais e entregas relevantes. Registrar no histórico
mensal e, quando útil, resumir no domínio.

### Baixa prioridade

Ajustes cosméticos, correções triviais, detalhes temporários ou fatos já
deriváveis do código. Não registrar como memória permanente.

## Organização por domínio

A memória pode ser dividida em:

```text
.claude/memory/
├── index.md
├── authentication.md
├── payments.md
├── users.md
└── marketing.md
```

O índice deve permanecer curto e apontar apenas para os domínios existentes.
Não criar um domínio para cada tarefa.
