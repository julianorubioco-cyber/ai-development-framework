$ErrorActionPreference = "Stop"
$RepoRoot = Split-Path -Parent $PSScriptRoot
$Template = Join-Path $RepoRoot "templates\project-workspace\.claude"
$Target = Join-Path (Get-Location) ".claude"

if (-not (Test-Path $Template)) { throw "Template não encontrado: $Template" }
New-Item -ItemType Directory -Force -Path $Target | Out-Null

Get-ChildItem $Template -Recurse -File | ForEach-Object {
    $relative = $_.FullName.Substring($Template.Length).TrimStart('\','/')
    $dest = Join-Path $Target $relative
    $destDir = Split-Path -Parent $dest
    New-Item -ItemType Directory -Force -Path $destDir | Out-Null
    if (-not (Test-Path $dest)) {
        Copy-Item $_.FullName $dest
    }
}
Write-Host "Workspace inicializado em: $Target"
