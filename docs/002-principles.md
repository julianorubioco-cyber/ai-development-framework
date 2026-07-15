# Princípios

## 1. Entender antes de alterar

Nenhum fluxo orquestrado modifica código antes de um preflight aprovado.

## 2. Evidência antes de conclusão

A revisão não aprova por impressão. Cada critério deve possuir evidência.

## 3. Menor contexto suficiente

O agente deve usar o menor conjunto de arquivos capaz de executar a tarefa com
segurança.

## 4. Busca progressiva

Primeiro índices e arquivos de alta relevância. O escopo só aumenta quando houver
uma dependência comprovada.

## 5. Memória pertence ao projeto

Skills globais são ferramentas. Contexto, decisões e histórico ficam somente em
`<projeto>/.claude/`.

## 6. Memória é curada

Nem toda informação deve ser salva. Fatos temporários, conversas completas e
detalhes facilmente derivados do código não viram memória permanente.

## 7. Risco prevalece sobre tamanho

Uma alteração de uma linha pode exigir fluxo completo quando envolve
autenticação, pagamentos, dados, permissões ou produção.

## 8. Sem substituição silenciosa de tecnologia

Linguagem, framework, banco, ORM, arquitetura e infraestrutura só mudam com
autorização explícita.

## 9. Responsabilidade única

Cada Skill possui um papel definido. Orquestração não deve virar implementação
oculta.

## 10. Degradação segura

Quando algo importante não puder ser verificado, o estado correto é `BLOCKED`,
não `APPROVED`.
