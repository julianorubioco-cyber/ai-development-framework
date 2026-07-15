param(
    [switch]$RestoreBackup,
    [switch]$WhatIf
)

$ErrorActionPreference = "Stop"

$ClaudeRoot = Join-Path $HOME ".claude"
$Target = Join-Path $ClaudeRoot "skills"
$ManifestPath = Join-Path $ClaudeRoot "adf-install-manifest.json"

if (-not (Test-Path $ManifestPath)) {
    throw "Manifesto não encontrado em $ManifestPath. Para remoção manual, apague somente as pastas das Skills do ADF."
}

$manifest = Get-Content $ManifestPath -Raw | ConvertFrom-Json
$skills = @($manifest.skills)

Write-Host "Skills registradas: $($skills -join ', ')"

foreach ($name in $skills) {
    $dest = Join-Path $Target $name
    if (Test-Path $dest) {
        if ($WhatIf) {
            Write-Host "[SIMULAÇÃO] Removeria: $dest"
        } else {
            Remove-Item $dest -Recurse -Force
            Write-Host "Removida: $dest"
        }
    }
}

if ($RestoreBackup -and $manifest.backup_path) {
    $backup = [string]$manifest.backup_path
    if (Test-Path $backup) {
        Get-ChildItem $backup -Directory | ForEach-Object {
            $dest = Join-Path $Target $_.Name
            if ($WhatIf) {
                Write-Host "[SIMULAÇÃO] Restauraria: $($_.FullName) -> $dest"
            } else {
                if (Test-Path $dest) { Remove-Item $dest -Recurse -Force }
                Copy-Item $_.FullName $dest -Recurse -Force
                Write-Host "Restaurada: $dest"
            }
        }
    } else {
        Write-Warning "Backup registrado não foi encontrado: $backup"
    }
}

if (-not $WhatIf) {
    Remove-Item $ManifestPath -Force
    Write-Host "Manifesto removido."
    Write-Host "Desinstalação concluída. Abra uma nova sessão do Claude Code."
}
