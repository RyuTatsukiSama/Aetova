// TestEXE.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <windows.h>

int main() {
    STARTUPINFO si = { sizeof(STARTUPINFO) };
    PROCESS_INFORMATION pi;

    system("C:\\Users\\Etudiant2\\AppData\\Local\\Discord\\Update.exe");

    // Lancer le programme
    if (CreateProcess(L"C:\\Users\\Etudiant2\\AppData\\Local\\Discord\\Update.exe",   // Chemin vers l'ex�cutable
        NULL,                                // Arguments (NULL si aucun)
        NULL,                                // Attributs de s�curit� du processus
        NULL,                                // Attributs de s�curit� du thread
        FALSE,                               // H�ritage des handles
        0,                                   // Flags
        NULL,                                // Variables d'environnement
        NULL,                                // R�pertoire de travail
        &si,                                 // Infos de d�marrage
        &pi))                                // Infos sur le processus cr��
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
