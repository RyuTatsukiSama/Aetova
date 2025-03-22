#ifndef GAMETHREAD_H
#define GAMETHREAD_H

#include <QObject>
#include <QThread>
#include <windows.h>
#include <iostream>
#include <string>
using namespace std;

class GameThread : public QThread
{   
    Q_OBJECT

    void run() override
    {
        std::cout << "entrer" << std::endl;
        string path = "common\\" + pathToExe.toStdString() + "\\" + exeName.toStdString() + ".exe";
        std::wstring widestr = std::wstring(path.begin(), path.end());
        const wchar_t* widecstr = widestr.c_str();

        STARTUPINFO si = { sizeof(STARTUPINFO) };
        PROCESS_INFORMATION pi;

        // Launch programm
        if (CreateProcess(widecstr,   // Path to Exe
                          NULL,                                // Arguments (NULL si aucun)
                          NULL,                                // Attributs de sécurité du processus
                          NULL,                                // Attributs de sécurité du thread
                          FALSE,                               // Héritage des handles
                          0,                                   // Flags
                          NULL,                                // Variables d'environnement
                          NULL,                                // Répertoire de travail
                          &si,                                 // Infos de démarrage
                          &pi))                                // Infos sur le processus créé
        {
            // Wait the end of the process
            WaitForSingleObject(pi.hProcess, INFINITE);
            CloseHandle(pi.hProcess);
            CloseHandle(pi.hThread);
            emit threadFinish("End without problem");
        }
        else {
            // Gestion des erreurs
            std::string errorMessage = "Error lors du lancement de l'application " + path;
            std::wstring widestrError = std::wstring(errorMessage.begin(), errorMessage.end());
            const wchar_t* widecstrError = widestrError.c_str();
            MessageBox(NULL, widecstrError, L"Erreur", MB_OK);
            emit threadFinish("End error");
        }
    }

signals :
    void threadFinish(const QString &s);
public :
    QString pathToExe;
    QString exeName;
};

#endif // GAMETHREAD_H
