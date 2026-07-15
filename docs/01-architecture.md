# Arquitetura

O framework separa ferramenta e conhecimento.

## Camada global

`~/.claude/skills/` contém os comportamentos reutilizáveis. Nenhuma Skill global
deve guardar nome de cliente, decisões de projeto ou histórico operacional.

## Camada do projeto

`<projeto>/.claude/` contém somente conhecimento daquele projeto. Essa pasta
pode acompanhar o repositório, respeitando a política de privacidade da equipe.

## Contratos

Cada etapa usa artefatos curtos:

- especificação;
- plano;
- preflight;
- revisão;
- release;
- memória curada.

Os artefatos não armazenam cadeia de pensamento. Guardam conclusões, decisões,
checklists e evidências.
