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

rem === Active the two following command if qt is not in the PATH environnement ===

REM === Path to the Qt bin folder ===
REM QT_BIN_PATH=C:\Qt\6.10.10\msvc2022_64\bin

REM === Add Qt bin to PATH ===
REM PATH=%QT_BIN_PATH%;%PATH%

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