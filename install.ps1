param(
    [string]$Target = "$env:LOCALAPPDATA\Microsoft\WindowsApps"
)

$Binary = "fm.exe"

function Build {
    Write-Host "Building $Binary..."
    go build -o $Binary .
    if ($LASTEXITCODE -ne 0) { Write-Error "Build failed."; exit 1 }
    Write-Host "Done."
}

function Install {
    Build
    Write-Host "Installing to $Target..."
    Copy-Item -Force $Binary "$Target\$Binary"
    Write-Host "Installed. You can now run 'fm' from any terminal."
}

function Uninstall {
    $path = "$Target\$Binary"
    if (Test-Path $path) {
        Remove-Item $path
        Write-Host "Removed $path"
    } else {
        Write-Host "$path not found, nothing to uninstall."
    }
}

function Clean {
    if (Test-Path $Binary) {
        Remove-Item $Binary
        Write-Host "Removed local $Binary"
    } else {
        Write-Host "Nothing to clean."
    }
}

switch ($args[0]) {
    "build"     { Build }
    "uninstall" { Uninstall }
    "clean"     { Clean }
    default     { Install }
}
