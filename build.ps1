[CmdletBinding()]
param(
    [switch]$SkipInstall
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$rootDir = $PSScriptRoot
$webDir = Join-Path $rootDir 'web'
$serverDir = Join-Path $rootDir 'server'
$webDistDir = Join-Path $webDir 'dist'
$adminResourceDir = Join-Path $serverDir 'resource\public\admin'

function Require-Command {
    param([Parameter(Mandatory = $true)][string]$Name)

    if (-not (Get-Command $Name -ErrorAction SilentlyContinue)) {
        throw "Required command '$Name' was not found in PATH."
    }
}

function Invoke-ExternalCommand {
    param(
        [Parameter(Mandatory = $true)][string]$FilePath,
        [Parameter()][string[]]$Arguments = @(),
        [Parameter(Mandatory = $true)][string]$Description
    )

    Write-Host "==> $Description"
    & $FilePath @Arguments
    if ($LASTEXITCODE -ne 0) {
        throw "$Description failed with exit code $LASTEXITCODE."
    }
}

if (-not (Test-Path (Join-Path $webDir 'package.json'))) {
    throw "Web project was not found at '$webDir'."
}

if (-not (Test-Path (Join-Path $serverDir 'go.mod'))) {
    throw "Server project was not found at '$serverDir'."
}

Require-Command 'pnpm'
Require-Command 'gf'

Push-Location $webDir
try {
    if (-not $SkipInstall -and -not (Test-Path (Join-Path $webDir 'node_modules'))) {
        Invoke-ExternalCommand -FilePath 'pnpm' -Arguments @('install', '--frozen-lockfile') -Description 'Installing web dependencies'
    }

    Invoke-ExternalCommand -FilePath 'pnpm' -Arguments @('run', 'build') -Description 'Building web application'
}
finally {
    Pop-Location
}

if (-not (Test-Path $webDistDir)) {
    throw "Web build output was not found at '$webDistDir'."
}

Write-Host '==> Updating server admin static resources'
if (Test-Path $adminResourceDir) {
    Remove-Item -Path $adminResourceDir -Recurse -Force
}
New-Item -Path $adminResourceDir -ItemType Directory -Force | Out-Null
Copy-Item -Path (Join-Path $webDistDir '*') -Destination $adminResourceDir -Recurse -Force

Push-Location $serverDir
try {
    Write-Host '==> Building server application'
    'y' | & gf build
    if ($LASTEXITCODE -ne 0) {
        throw "Building server application failed with exit code $LASTEXITCODE."
    }
}
finally {
    Pop-Location
}

Write-Host ''
Write-Host 'Build completed successfully.'
Write-Host "Web resources: $adminResourceDir"
Write-Host "Server output: $(Join-Path $serverDir 'temp')"
