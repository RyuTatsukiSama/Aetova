@echo off
REM ------------------------------------------
REM Script for generate a resources cpp file with Qt
REM ------------------------------------------

REM === Input for the name of the header ===
set /p BASENAME=Enter the name of the header : 

REM === File name ===
set RESOURCE=%BASENAME%.qrc
set OUTPUT=%BASENAME%.cpp

REM === path to the bin folder of Qt ===
set QT_BIN_PATH=C:\Qt\6.8.2\msvc2022_64\bin

REM === add the bin folder of Qt to PATH ===
set PATH=%QT_BIN_PATH%;%PATH%

REM === Execute rcc ===
echo Generation of %OUTPUT% from %HEADER%
rcc %RESOURCE% -o %OUTPUT%

REM === Verification of success ===
if exist %OUTPUT% (
    echo File %OUTPUT% generation succeed.
) else (
    echo Error : the file %OUTPUT% was'nt generated.
)

pause