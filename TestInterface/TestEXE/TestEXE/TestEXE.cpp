// TestEXE.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <windows.h>

int main() {
    STARTUPINFO si = { sizeof(STARTUPINFO) };
    PROCESS_INFORMATION pi;

    system("C:\\Users\\Etudiant2\\AppData\\Local\\Discord\\Update.exe");

    // Lancer le programme
    if (CreateProcess(L"C:\\Users\\Etudiant2\\AppData\\Local\\Discord\\Update.exe",   // Chemin vers l'exécutable
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
        // Attendre la fin du processus
        WaitForSingleObject(pi.hProcess, INFINITE);
        CloseHandle(pi.hProcess);
        CloseHandle(pi.hThread);
    }
    else {
        // Gestion des erreurs
        MessageBox(NULL, L"Erreur lors du lancement", L"Erreur", MB_OK);
    }

    return 0;
}
