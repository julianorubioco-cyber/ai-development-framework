# Changelog

## 0.5.0 — 2026-07-15

- Nova CLI multiplataforma `adf`.
- Comandos install, init, doctor, update e uninstall.
- Detecção automática de sistema operacional.
- Instaladores remotos de um comando para Windows, macOS e Linux.
- Empacotamento Python sem dependências externas.
- Assets de Skills e templates incluídos no pacote.
- Testes da CLI e isolamento da memória.

## 0.4.0 — 2026-07-15

- Roteamento inteligente entre pergunta, análise, ambiguidade e alteração.
- Perguntas comuns deixam de iniciar o pipeline.
- Solicitações reais podem iniciar `/implement` automaticamente.
- `/implement` inicializa o workspace quando necessário.
- Nova Skill manual `/route` para diagnosticar a classificação.
- Política de economia aplicada antes da leitura de arquivos.
- Testes para garantir que o roteador preserve perguntas simples.

## 0.3.0 — 2026-07-15

- Definição oficial de visão, princípios e arquitetura.
- Política formal de economia de tokens e contexto.
- Terminologia padronizada.
- Processo de releases documentado.
- Novo `context-budget.md` no workspace.
- Regras econômicas adicionadas ao contexto e ao orquestrador.
- Testes para garantir a presença dos fundamentos.

## 0.2.0 — 2026-07-15

- Instalador passa a registrar exatamente quais Skills foram instaladas.
- Novo desinstalador para Windows, macOS e Linux.
- Restauração opcional do backup anterior.
- Modo de simulação (`-WhatIf` / `--dry-run`) para instalação e remoção.
- Validação mais rigorosa dos frontmatters.
- Testes automatizados para estrutura, isolamento de memória e contratos.
- Workflow de CI para GitHub Actions.
- Guia de desinstalação e recuperação.
- Manifesto de instalação para evitar apagar Skills que não pertencem ao framework.

## 0.1.0 — 2026-07-15

- Primeira versão funcional.
- Pipeline adaptativo.
- Gate de preflight com aprovação explícita.
- Revisão posterior baseada em evidências.
- Memória exclusiva por projeto.
- Curadoria por prioridade e organização por domínio.
- Política de economia de contexto.
- Instaladores para Windows e sistemas Unix.
