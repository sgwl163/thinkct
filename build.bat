@echo off
setlocal

cd /d "%~dp0"

echo HotGo one-click build
powershell.exe -NoLogo -NoProfile -ExecutionPolicy Bypass -File "%~dp0build.ps1" %*

if errorlevel 1 (
    echo.
    echo Build failed. Press any key to close.
    pause >nul
    exit /b 1
)

echo.
echo Build completed successfully. Press any key to close.
pause >nul
exit /b 0
