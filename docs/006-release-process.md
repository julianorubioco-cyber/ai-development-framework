# Processo de versões

O projeto usa versionamento semântico:

```text
MAJOR.MINOR.PATCH
```

## PATCH

Correções compatíveis e documentação.

## MINOR

Novas Skills, políticas ou recursos compatíveis.

## MAJOR

Mudanças incompatíveis em contratos, estrutura ou instalação.

## Fluxo de lançamento

1. definir escopo;
2. alterar em branch;
3. validar;
4. executar testes;
5. atualizar changelog;
6. mesclar em `main`;
7. criar tag;
8. publicar GitHub Release.

## Comandos

```powershell
git add .
git commit -m "docs: establish project foundations"
git push
git tag v0.3.0
git push origin v0.3.0
```
