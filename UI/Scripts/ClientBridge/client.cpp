#include "client.h"
#include <QProcess>

Client::Client(/* args */)
{
    QProcess *process = new QProcess(nullptr);

    process->startDetached("clientGo/client.exe");
}

Client::~Client()
{
}