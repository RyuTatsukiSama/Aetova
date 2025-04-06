@echo off
REM ------------------------------------------
REM Script pour générer un fichier MOC avec Qt
REM ------------------------------------------

REM === Nom du fichier header (à modifier si besoin) ===
set HEADER=gamethread.h

REM === Nom du fichier de sortie généré par moc ===
set OUTPUT=gamethread.cpp

REM === Chemin vers le dossier bin de Qt (à adapter à ton install) ===
set QT_BIN_PATH=C:\Qt\6.8.2\msvc2022_64\bin

REM === Ajoute le dossier bin de Qt au PATH pour cette session ===
set PATH=%QT_BIN_PATH%;%PATH%

REM === Lancer la commande moc ===
echo Génération de %OUTPUT% depuis %HEADER%
moc %HEADER% -o %OUTPUT%

REM === Vérification si la génération a fonctionné ===
if exist %OUTPUT% (
    echo ✅ Fichier %OUTPUT% généré avec succès.
) else (
    echo ❌ Erreur : le fichier %OUTPUT% n'a pas été généré.
)

pause