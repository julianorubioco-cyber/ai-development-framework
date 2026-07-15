# Instruções do projeto

## Roteamento de intenção

Antes de agir, classifique a solicitação em uma destas categorias:

### 1. Pergunta comum

Exemplos: explicação, conceito, dúvida, localização de arquivo, entendimento de erro.

Comportamento:

- responda diretamente;
- use contexto mínimo;
- não inicialize o workspace;
- não crie spec, plano ou preflight;
- não altere arquivos.

### 2. Análise sem alteração

Exemplos: revisar, explicar arquitetura, localizar riscos, analisar um módulo.

Comportamento:

- leia apenas os arquivos necessários;
- não altere o projeto;
- não inicie automaticamente o pipeline de implementação.

### 3. Solicitação ambígua

Exemplos: "dá para melhorar?", "consegue corrigir?", "tem como adicionar?".

Comportamento:

- não altere arquivos;
- explique brevemente o que seria necessário;
- pergunte se o usuário deseja iniciar a implementação.

### 4. Alteração real de projeto

Exemplos: criar, adicionar, corrigir, remover, migrar, configurar, publicar,
atualizar, refatorar ou implementar.

Comportamento:

- se `.claude/` ainda não estiver inicializada, inicialize o workspace na raiz;
- classifique tamanho e risco;
- aplique o fluxo do `/implement`;
- não altere arquivos antes do preflight aprovado.

### 5. `/implement`

Quando o usuário chamar `/implement`, inicie o pipeline explicitamente,
independentemente da classificação automática.

## Ordem de leitura

Somente depois de classificar a intenção:

1. artefato citado pela tarefa;
2. `.claude/memory/index.md`;
3. apenas domínios relacionados;
4. manifestos e arquivos de alta relevância;
5. dependências diretas quando necessárias.

## Regras

- Use o menor contexto suficiente.
- Não percorra o repositório inteiro por padrão.
- Não leia todos os históricos ou domínios.
- Preserve decisões e arquitetura documentadas.
- Não grave cadeia de pensamento ou conversas completas.
- Atualize memória somente por curadoria.
- Memória deste projeto não pode sair desta pasta `.claude/`.
