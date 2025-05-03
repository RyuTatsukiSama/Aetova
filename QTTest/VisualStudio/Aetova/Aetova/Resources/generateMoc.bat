@echo off
REM ------------------------------------------
REM Script to generate a MOC file with Qt
REM ------------------------------------------

REM === Ask for the path to the header folder ===
set /p HEADER_PATH=Enter the relative or absolute path to the header (e.g., ../include or C:\Project\headers): 

REM === Ask for the base name of the header file ===
set /p BASENAME=Enter the name of the header file (without .h): 

REM === Compose full paths ===
set HEADER="%HEADER_PATH%\%BASENAME%.h"
set OUTPUT="%HEADER_PATH%\moc_%BASENAME%.cpp"

REM === Path to the Qt bin folder ===
set QT_BIN_PATH=C:\Qt\6.8.2\msvc2022_64\bin

REM === Add Qt bin to PATH ===
set PATH=%QT_BIN_PATH%;%PATH%

REM === Run moc ===
echo Generating %OUTPUT% from %HEADER%
moc %HEADER% -o %OUTPUT%

REM === Check if the output file was generated ===
if exist %OUTPUT% (
    echo File %OUTPUT% generated successfully.
) else (
    echo Error: file %OUTPUT% was not generated.
)

pause