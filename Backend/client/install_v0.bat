@echo off

:: Delete the game if it exist
if exist ".\Shipyard\app\BuildOranys" (
    rmdir /s /q ".\Shipyard\app\BuildOranys"
    if %errorlevel% neq 0 (
        echo "Error in remove"
    )   
)

:: Replace it by an older version
"C:\Program Files\7-Zip\7z.exe" x ..\server\Shipyard\v1_0_0.zip -o.\Shipyard\app
if %errorlevel% neq 0 (
    echo "Error in compresion"
)

:: Change the version in the app manifest
set FILE=".\Shipyard\AppManifest_0.json"

powershell -Command ^
  "$j = Get-Content '%FILE%' -Raw | ConvertFrom-Json;" ^
  "$j.Version = 0;" ^
  "$j | ConvertTo-Json | Set-Content '%FILE%'"
if %errorlevel% neq 0 (
    echo "Error in manifest modif"
)

pause
exit