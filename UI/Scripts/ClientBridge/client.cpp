#include "client.h"
#include <QProcess>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QNetworkRequest>
#include "../GameWindow/ButtonGame/buttongame.h"

Client::Client()
{
    log = new doc::Logger();

    process = new QProcess(nullptr);
    process->startDetached("clientGo/client.exe");

    manager = new QNetworkAccessManager(nullptr);

    QNetworkReply *reply = manager->get(QNetworkRequest(QUrl("http://localhost:51419/health")));
    QObject::connect(reply, &QNetworkReply::finished,
                     this, [reply, this]()
                     {
    if (reply->error() == QNetworkReply::NoError) 
    {
        QByteArray data = reply->readAll();
        log->Info(data.toStdString());
    } 
    else 
    {
        log->Error(reply->errorString().toStdString() + " " + reply->readAll().toStdString());
    }
    reply->deleteLater(); });
}

void Client::download(ButtonGame *button)
{
    log->Info("Download Start");
    button->pauseConnect();
    QNetworkReply *reply = manager->post(QNetworkRequest(QUrl("http://localhost:51419/download")), QByteArray());
    connect(reply, &QNetworkReply::finished, [reply, button, this]()
            { 
    if (reply->error() == QNetworkReply::NoError) 
    {
        log->Info("Download done");
        button->launchConnect();
    } 
    else 
    {
        log->Error(reply->errorString().toStdString() + " " + reply->readAll().toStdString());
    } });
}

void Client::pause(ButtonGame *button)
{
    log->Info("Pause Start");
    QNetworkReply *reply = manager->post(QNetworkRequest(QUrl("http://localhost:51419/pause")), QByteArray());
    connect(reply, &QNetworkReply::finished, [reply, button, this]()
            { 
                if (reply->error() == QNetworkReply::NoError) 
    {
        log->Info("Pause done");
        button->resumeConnect();
    } 
    else 
    {
        log->Error(reply->errorString().toStdString() + " " + reply->readAll().toStdString());
    } });
}

void Client::resume(ButtonGame *button)
{
    log->Info("Resume Start");
    button->pauseConnect();
    QNetworkReply *reply = manager->post(QNetworkRequest(QUrl("http://localhost:51419/resume")), QByteArray());
    connect(reply, &QNetworkReply::finished, [reply, button, this]()
            {
    if (reply->error() == QNetworkReply::NoError) 
    {
        log->Info("Resume done");
        button->launchConnect();
    } 
    else 
    {
        log->Error(reply->errorString().toStdString() + " " + reply->readAll().toStdString());
    } });
}

Client::~Client()
{
    delete process;
    process = nullptr;

    delete manager;
    manager = nullptr;
}