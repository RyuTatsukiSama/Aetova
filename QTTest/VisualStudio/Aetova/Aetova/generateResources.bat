@echo off
REM ------------------------------------------
REM Script pour générer un fichier resources avec Qt
REM ------------------------------------------

REM === Nom du fichier qrc ===
set RESOURCE=resources.qrc

REM === Nom du fichier de sortie généré par rcc ===
set OUTPUT=resources.cpp

REM === Chemin vers le dossier bin de Qt ===
set QT_BIN_PATH=C:\Qt\6.8.2\msvc2022_64\bin

REM === Ajoute le dossier bin de Qt au PATH pour cette session ===
set PATH=%QT_BIN_PATH%;%PATH%

REM === Lancer la commande rcc ===
echo Generation of %OUTPUT% from %HEADER%
rcc %RESOURCE% -o %OUTPUT%

REM === Vérification si la génération a fonctionné ===
if exist %OUTPUT% (
    echo File %OUTPUT% generation succeed.
) else (
    echo Error : the file %OUTPUT% was'nt generated.
)

pause