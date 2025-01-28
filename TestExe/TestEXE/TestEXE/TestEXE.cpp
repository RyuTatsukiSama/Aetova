// TestEXE.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <windows.h>
#include <fstream>
#include <iostream>
using namespace std;

int main()
{
	string str;
	try
	{
		std::ifstream in("List.txt");
		in.exceptions(std::ifstream::failbit); // may throw
		in >> str; // may throw
	}
	catch (const std::ios_base::failure& fail)
	{
		// handle exception here
		std::cout << fail.what() << '\n';
	}

	string path = "Build\\Journeep.exe";
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

	return 0;
}
