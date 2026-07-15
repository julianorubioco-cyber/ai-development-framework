param(
    [string]$Repository = "https://github.com/julianorubioco-cyber/ai-development-framework.git"
)

$ErrorActionPreference = "Stop"

Write-Host "AI Development Framework - instalador Windows"

if (-not (Get-Command python -ErrorAction SilentlyContinue)) {
    throw "Python 3.9+ não foi encontrado. Instale pelo Microsoft Store ou python.org."
}

python -m pip install --user --upgrade "git+$Repository"

$UserBase = python -m site --user-base
$Scripts = Join-Path $UserBase "Scripts"
$env:Path = "$Scripts;$env:Path"

python -m adf_cli.cli install
python -m adf_cli.cli doctor

Write-Host ""
Write-Host "Instalação concluída."
Write-Host "Comandos disponíveis:"
Write-Host "  adf install"
Write-Host "  adf init"
Write-Host "  adf doctor"
Write-Host "  adf update"
Write-Host "  adf uninstall"
Write-Host ""
Write-Host "Se 'adf' não for reconhecido nesta janela, use:"
Write-Host "  python -m adf_cli.cli doctor"
