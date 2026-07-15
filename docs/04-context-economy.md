# Economia de contexto

O agente deve buscar progressivamente:

1. Ler os índices e artefatos diretamente relacionados.
2. Mapear arquivos por alta, média e baixa relevância.
3. Abrir primeiro os de alta relevância.
4. Expandir somente mediante evidência de dependência.
5. Evitar reler artefatos não alterados.
6. Preferir resumos canônicos a históricos extensos.
7. Ler histórico apenas por período ou domínio relevante.
8. Nunca percorrer dependências geradas, binários, caches ou diretórios de build.

Economia de contexto não autoriza ignorar arquivos necessários para segurança,
arquitetura, testes ou critérios de aceitação.
