param(
    [switch]$WhatIf
)

$ErrorActionPreference = "Stop"

$RepoRoot = Split-Path -Parent $PSScriptRoot
$Source = Join-Path $RepoRoot "skills"
$ClaudeRoot = Join-Path $HOME ".claude"
$Target = Join-Path $ClaudeRoot "skills"
$Timestamp = Get-Date -Format "yyyyMMdd-HHmmss"
$Backup = Join-Path $ClaudeRoot "adf-backups\$Timestamp"
$ManifestPath = Join-Path $ClaudeRoot "adf-install-manifest.json"
$Version = (Get-Content (Join-Path $RepoRoot "VERSION") -Raw).Trim()

$skillNames = @(Get-ChildItem $Source -Directory | ForEach-Object { $_.Name })

Write-Host "ADF $Version"
Write-Host "Destino: $Target"
Write-Host "Skills: $($skillNames -join ', ')"

if ($WhatIf) {
    Write-Host "[SIMULAÇÃO] Nenhum arquivo será alterado."
    foreach ($name in $skillNames) {
        $dest = Join-Path $Target $name
        if (Test-Path $dest) {
            Write-Host "[SIMULAÇÃO] Faria backup e substituiria: $dest"
        } else {
            Write-Host "[SIMULAÇÃO] Instalaria: $dest"
        }
    }
    exit 0
}

New-Item -ItemType Directory -Force -Path $Target | Out-Null

$existing = @($skillNames | Where-Object { Test-Path (Join-Path $Target $_) })
if ($existing.Count -gt 0) {
    New-Item -ItemType Directory -Force -Path $Backup | Out-Null
    foreach ($name in $existing) {
        Copy-Item (Join-Path $Target $name) $Backup -Recurse -Force
    }
    Write-Host "Backup criado em: $Backup"
}

foreach ($name in $skillNames) {
    $sourceDir = Join-Path $Source $name
    $destDir = Join-Path $Target $name
    if (Test-Path $destDir) {
        Remove-Item $destDir -Recurse -Force
    }
    Copy-Item $sourceDir $destDir -Recurse -Force
}

$manifest = @{
    framework = "AI Development Framework"
    version = $Version
    installed_at = (Get-Date).ToString("o")
    skills = $skillNames
    backup_path = if ($existing.Count -gt 0) { $Backup } else { $null }
}
$manifest | ConvertTo-Json -Depth 4 | Set-Content $ManifestPath -Encoding UTF8

Write-Host "Skills instaladas em: $Target"
Write-Host "Manifesto criado em: $ManifestPath"
Write-Host "Abra uma nova sessão do Claude Code."
