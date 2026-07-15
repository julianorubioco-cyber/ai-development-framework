param(
    [string]$Version = "latest",
    [switch]$SkipSkillInstall
)

$ErrorActionPreference = "Stop"
$Repo = "julianorubioco-cyber/ai-development-framework"
$InstallDir = Join-Path $env:LOCALAPPDATA "ADF\bin"
$Binary = Join-Path $InstallDir "adf.exe"

$arch = switch ($env:PROCESSOR_ARCHITECTURE) {
    "AMD64" { "amd64" }
    "ARM64" { "arm64" }
    default { throw "Arquitetura não suportada: $env:PROCESSOR_ARCHITECTURE" }
}

if ($Version -eq "latest") {
    $release = Invoke-RestMethod "https://api.github.com/repos/$Repo/releases/latest"
    $tag = $release.tag_name
} else {
    $tag = $Version
}

$asset = "adf_windows_$arch.exe"
$url = "https://github.com/$Repo/releases/download/$tag/$asset"

New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
$tmp = "$Binary.download"
Invoke-WebRequest -Uri $url -OutFile $tmp
Move-Item -Force $tmp $Binary

$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
$parts = @($userPath -split ";" | Where-Object { $_ })
if ($parts -notcontains $InstallDir) {
    $newPath = (($parts + $InstallDir) -join ";")
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
}
$env:Path = "$InstallDir;$env:Path"

if (-not $SkipSkillInstall) {
    & $Binary install
}
& $Binary doctor

Write-Host ""
Write-Host "ADF instalado em: $Binary"
Write-Host "Abra um novo terminal e execute: adf doctor"
