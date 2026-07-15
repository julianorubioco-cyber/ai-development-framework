# Roteamento inteligente de solicitações

## Objetivo

Permitir conversa normal com o Claude Code sem iniciar o pipeline completo para
perguntas simples, mas ativar o framework quando houver intenção real de alterar
o projeto.

## Regra principal

O framework só deve ser ativado automaticamente quando a solicitação implicar
mudança, criação, correção, migração, configuração, implantação ou outra ação
sobre o projeto.

Perguntas explicativas, consultas e pedidos de leitura não devem iniciar o
pipeline de implementação.

## Categorias

### Pergunta comum

Exemplos:

- "O que é JWT?"
- "Como este módulo funciona?"
- "Qual arquivo controla o login?"
- "Explique este erro."
- "Qual banco este projeto usa?"

Comportamento:

- responder diretamente;
- usar contexto mínimo;
- não criar workspace automaticamente;
- não criar spec;
- não executar preflight;
- não alterar arquivos.

### Consulta ou análise sem alteração

Exemplos:

- "Revise este arquivo e me diga o que acha."
- "Encontre onde a autenticação é configurada."
- "Explique a arquitetura atual."
- "Quais riscos existem neste código?"

Comportamento:

- analisar somente os arquivos necessários;
- não alterar código;
- não iniciar `/implement`;
- recomendar fluxo apropriado apenas quando houver necessidade concreta.

### Alteração real de projeto

Exemplos:

- "Crie uma tela de login."
- "Corrija o checkout."
- "Adicione pagamento com Stripe."
- "Troque o texto do botão."
- "Atualize as dependências."
- "Faça deploy."
- "Migre o banco."

Comportamento:

- inicializar o workspace se necessário;
- classificar tamanho e risco;
- seguir o fluxo de `/implement`;
- exigir preflight e aprovação explícita antes de modificar arquivos.

### Comando explícito

Quando o usuário executar:

```text
/implement <solicitação>
```

o pipeline deve ser iniciado mesmo que a tarefa pareça pequena.

## Casos ambíguos

Exemplo:

- "Você consegue colocar login aqui?"
- "Dá para melhorar essa tela?"
- "Esse erro tem como arrumar?"

Nesses casos:

1. identifique se o usuário está apenas perguntando sobre possibilidade;
2. não altere arquivos imediatamente;
3. responda brevemente e pergunte se ele deseja iniciar a implementação;
4. só ative o pipeline após confirmação ou uso explícito de `/implement`.

## Economia de contexto

O roteador deve decidir primeiro a categoria da solicitação antes de ler arquivos.
Isso evita carregar contexto de implementação para perguntas que não exigem ação.
