---
name: memory
description: Faz curadoria da memória exclusivamente dentro do projeto, por prioridade e domínio.
argument-hint: <resultado, decisão ou módulo>
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


# Destino obrigatório

Escreva somente em `<raiz-do-projeto>/.claude/`. Nunca grave conhecimento do
projeto em `~/.claude/`.

# Curadoria

Antes de guardar, avalie:

- será útil daqui a meses;
- evita redescoberta ou nova pergunta;
- representa decisão, regra, integração ou arquitetura;
- não é facilmente derivável do código;
- pertence a este projeto.

## Prioridade alta

Atualize documentos canônicos ou `.claude/memory/<dominio>.md`:
arquitetura, stack, regras permanentes, integrações, restrições e decisões.

## Prioridade média

Registre entrega ou mudança estrutural em
`.claude/history/AAAA-MM.md`; atualize domínio apenas quando necessário.

## Prioridade baixa

Não memorize: ajustes cosméticos, correções triviais, fatos temporários, logs,
conversa completa ou detalhes já visíveis no código.

# Domínios

Use nomes estáveis, como `authentication.md`, `payments.md`, `users.md`.
Não crie um arquivo por tarefa. Atualize `.claude/memory/index.md` com descrições
curtas dos domínios.

Remova ou substitua informação obsoleta quando houver evidência. Não apague
decisões históricas relevantes; marque-as como substituídas.

Finalize com `MEMORY_STATUS: UPDATED`, `MEMORY_STATUS: NO_CHANGES` ou
`MEMORY_STATUS: BLOCKED`.
