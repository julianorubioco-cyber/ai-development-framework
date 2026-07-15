# 🚀 Instalação do AI Development Framework (ADF)

O ADF instala as Skills do Claude Code e a CLI (`adf`) para Windows,
macOS e Linux.

------------------------------------------------------------------------

# Requisitos

-   Claude Code instalado
-   Git instalado
-   Conexão com a Internet

> Não é necessário instalar Python ou Go.

------------------------------------------------------------------------

# Windows

## 1. Abra o PowerShell

Pressione **Win**, digite **PowerShell** e abra-o.

## 2. Execute o instalador

``` powershell
irm https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install.ps1 | iex
```

O instalador irá:

-   detectar o Windows;
-   baixar o executável correto;
-   instalar a CLI `adf`;
-   instalar as Skills do Claude Code;
-   criar backup das Skills existentes;
-   adicionar o `adf` ao PATH (quando necessário).

## 3. Feche e abra o PowerShell novamente

## 4. Verifique a instalação

``` powershell
adf doctor
```

Se tudo estiver correto, o diagnóstico indicará que o ADF está saudável.

------------------------------------------------------------------------

# macOS

Abra o Terminal e execute:

``` bash
curl -fsSL https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install.sh | sh
```

Depois:

``` bash
adf doctor
```

------------------------------------------------------------------------

# Linux

Abra o Terminal e execute:

``` bash
curl -fsSL https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install.sh | sh
```

Depois:

``` bash
adf doctor
```

------------------------------------------------------------------------

# Criando um projeto

Entre na pasta do projeto:

``` text
cd MeuProjeto
```

Inicialize o workspace:

``` text
adf init
```

Será criada uma pasta:

``` text
.claude/
```

A memória e o contexto ficam **somente dentro desse projeto**.

------------------------------------------------------------------------

# Atualizando o ADF

Quando houver uma nova versão:

``` bash
git pull
```

Depois:

``` text
adf update
```

> (Caso o comando ainda não exista na versão instalada, siga as
> instruções da release mais recente.)

------------------------------------------------------------------------

# Desinstalar

``` text
adf uninstall
```

Para restaurar o backup das Skills:

``` text
adf uninstall --restore-backup
```

------------------------------------------------------------------------

# Solução de problemas

## O comando `adf` não foi encontrado

Feche e abra o terminal novamente.

Depois execute:

``` text
adf doctor
```

## O Claude Code não encontrou as Skills

Execute:

``` text
adf install
```

------------------------------------------------------------------------

# Verificar a versão

``` text
adf version
```

------------------------------------------------------------------------

# Fluxo recomendado

Instalar:

``` text
Windows:
irm https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install.ps1 | iex

macOS/Linux:
curl -fsSL https://raw.githubusercontent.com/julianorubioco-cyber/ai-development-framework/main/install.sh | sh
```

Usar:

``` text
cd MeuProjeto
adf init
```

Abrir a pasta no Claude Code e começar a trabalhar.
