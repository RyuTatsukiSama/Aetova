#include "gamethread.h"


GameThread::GameThread(QObject* parent) : QThread(parent)
{

}

void GameThread::run()
{
    //std::cout << "entrer" << std::endl;
    string workingDir = "common\\" + pathToExe.toStdString();
    std::wstring widestrWD = std::wstring(workingDir.begin(), workingDir.end());
    const wchar_t* widecstrWD = widestrWD.c_str();

    string path = "common\\" + pathToExe.toStdString() + "\\" + exeName.toStdString() + ".exe";
    std::wstring widestr = std::wstring(path.begin(), path.end());
    const wchar_t* widecstr = widestr.c_str();

    STARTUPINFO si = { sizeof(STARTUPINFO) };
    PROCESS_INFORMATION pi;

    si.dwFlags = STARTF_USESHOWWINDOW;
    si.wShowWindow = SW_SHOW;

    // Launch programm
    if (CreateProcess(widecstr,   // Path to Exe
        NULL,                                // Arguments (NULL si aucun)
        NULL,                                // Attributs de sécurité du processus
        NULL,                                // Attributs de sécurité du thread
        FALSE,                               // Héritage des handles
        0,                                   // Flags
        NULL,                                // Variables d'environnement
        widecstrWD,                          // Répertoire de travail
        &si,                                 // Infos de démarrage
        &pi))                                // Infos sur le processus créé
    {
        // Wait the end of the process
        WaitForSingleObject(pi.hProcess, INFINITE);
        CloseHandle(pi.hProcess);
        CloseHandle(pi.hThread);
        emit threadFinish("End without problem");
        system("pause");
    }
    else {
        // Gestion des erreurs
        std::string errorMessage = "Error while launching the application, path : " + path;
        std::wstring widestrError = std::wstring(errorMessage.begin(), errorMessage.end());
        const wchar_t* widecstrError = widestrError.c_str();
        MessageBox(NULL, widecstrError, L"Erreur", MB_OK);
        emit threadFinish("End error");
    }
}