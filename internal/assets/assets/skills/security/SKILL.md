---
name: security
description: Audita riscos de segurança relevantes à stack e aplica correções somente quando autorizadas.
argument-hint: <spec, módulo ou alterações>
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


Avalie autenticação, autorização, validação, injeções, XSS, CSRF, segredos,
uploads, dependências e exposição de dados conforme aplicável.

Não invente ameaças irrelevantes à stack. Separe achados comprovados de riscos
hipotéticos. Correções que alterem arquitetura, dados, custo ou UX exigem aprovação.

Finalize com `SECURITY_STATUS: PASSED`, `SECURITY_STATUS: CHANGES_REQUIRED`
ou `SECURITY_STATUS: BLOCKED`.
