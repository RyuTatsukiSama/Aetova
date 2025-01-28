#include <QApplication>
#include <QPushButton>
#include <windows.h>
#include <iostream>
using namespace std;

void LaunchGame()
{
    std::cout << "entrer" << std::endl;
    string path = "JourneepBuild\\Journeep.exe";
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
    }
    else {
        // Gestion des erreurs
        MessageBox(NULL, L"Erreur lors du lancement", L"Erreur", MB_OK);
    }
}

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);

    QWidget window;
    window.setFixedSize(1024, 780);

    QPushButton *button = new QPushButton("Play", &window);
    int width = 200;
    int height = 100;
    button->setGeometry(1024/2-width/2, 780/2-height/2, width, height);

    QObject::connect(button, &QPushButton::clicked, LaunchGame);

    window.show();
    return a.exec();
}
