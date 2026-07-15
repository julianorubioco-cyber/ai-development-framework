# Política de economia de tokens e contexto

## Objetivo

Reduzir leituras, repetições e saídas desnecessárias sem comprometer segurança,
correção ou verificabilidade.

## Regra central

> Use o menor contexto suficiente para concluir a etapa atual com segurança.

## Protocolo de descoberta

Antes de abrir arquivos:

1. leia a solicitação e os artefatos diretamente relacionados;
2. consulte `.claude/memory/index.md`;
3. identifique os domínios relacionados;
4. monte uma lista curta de arquivos candidatos;
5. classifique-os em alta, média e baixa relevância;
6. abra somente os de alta relevância;
7. expanda a busca apenas quando surgir uma dependência verificável.

## Ordem preferencial

```text
1. especificação ou relatório diretamente citado
2. índice de memória
3. memória do domínio atual
4. manifesto e configuração
5. arquivos de entrada
6. dependências diretas
7. testes relacionados
8. arquivos de relevância média
```

## O que não deve ser lido por padrão

- repositório inteiro;
- histórico completo;
- todos os domínios de memória;
- dependências instaladas;
- diretórios de build;
- binários;
- caches;
- arquivos gerados;
- logs extensos sem relação com o erro.

## Reutilização

Antes de reconstruir conhecimento, procure:

```text
.claude/context.md
.claude/architecture.md
.claude/memory/index.md
.claude/memory/<dominio>.md
.claude/specs/
.claude/plans/
.claude/reviews/
```

Reutilize artefatos válidos e atualize somente o que mudou.

## Saídas econômicas

- não repetir a especificação inteira em cada etapa;
- referenciar caminhos de artefatos;
- apresentar resumos curtos;
- registrar detalhes extensos no arquivo correspondente;
- evitar narrar operações triviais;
- não armazenar cadeia de pensamento.

## Limites de exploração

Use limites orientativos, não absolutos:

### Tarefa pequena

- começar com até 5 arquivos candidatos;
- ampliar apenas se houver dependência.

### Tarefa média

- começar com até 12 arquivos candidatos;
- priorizar arquivos de entrada, contratos e testes.

### Tarefa grande

- mapear módulos primeiro;
- ler amostras e interfaces antes de implementações completas;
- dividir por domínio e etapa.

## Quando economizar menos

Expanda o contexto quando necessário para:

- segurança;
- autenticação e autorização;
- migrações;
- consistência de dados;
- mudanças arquiteturais;
- regressões;
- critérios de aceitação distribuídos.

Economia nunca justifica uma aprovação sem evidência.
