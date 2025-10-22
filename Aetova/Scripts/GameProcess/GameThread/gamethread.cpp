#include "gamethread.h"
#include <format>

GameThread::GameThread(std::string _threadName, QObject *parent) : QThread(parent)
{
    log = doc::Logger(_threadName);
}

void GameThread::run()
{
    // Construct the path to the working directory ( directory of the exe )
    string workingDir = std::format("apps\\{}", pathToExe.toStdString());
    std::wstring widestrWD = std::wstring(workingDir.begin(), workingDir.end());
    const wchar_t *widecstrWD = widestrWD.c_str();

    // Construct the path to the exe
    string path = std::format("{}\\{}.exe", workingDir, exeName.toStdString());
    std::wstring widestr = std::wstring(path.begin(), path.end());
    const wchar_t *widecstr = widestr.c_str();

    STARTUPINFO si = {sizeof(STARTUPINFO)};
    PROCESS_INFORMATION pi;

    si.dwFlags = STARTF_USESHOWWINDOW;
    si.wShowWindow = SW_SHOW;

    // Launch programm
    if (CreateProcess(widecstr,   // Path to Exe
                      NULL,       // Arguments
                      NULL,       // Security attributs of the process
                      NULL,       // Security attributs of the thread
                      FALSE,      // Handle heritage
                      0,          // Flags
                      NULL,       // Environement variable
                      widecstrWD, // Working directory
                      &si,        // Startup info
                      &pi))       // Info on the process created
    {
        // Wait the end of the process
        WaitForSingleObject(pi.hProcess, INFINITE);
        CloseHandle(pi.hProcess);
        CloseHandle(pi.hThread);
        emit threadFinish(exeName + " end without problem", doc::LoggerSeverity::Info);
    }
    else
    {
        // Gestion des erreurs
        std::string errorMessage = "Error while launching the application, path : " + path;
        std::wstring widestrError = std::wstring(errorMessage.begin(), errorMessage.end());
        const wchar_t *widecstrError = widestrError.c_str();
        // Warn the user about the problem
        emit threadFinish(exeName + " end with error " + widecstrError, doc::LoggerSeverity::Error);
    }
}