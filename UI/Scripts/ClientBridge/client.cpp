#include "client.h"
#include <QProcess>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QNetworkRequest>
#include <docLogger>
#include <QWebSocket>
#include "../GameWindow/ButtonGame/buttongame.h"

Client::Client(QObject* parent) : QObject(parent)
{
    log = new doc::Logger();

    process = new QProcess(nullptr);
    process->startDetached("clientGo/client.exe"); // TODO : Change this to a "CREATE_PROCESS", or a more multi platforme fonction

    manager = new QNetworkAccessManager(nullptr);

    QNetworkReply *reply = manager->get(QNetworkRequest(QUrl("http://localhost:51419/health")));
    QObject::connect(reply, &QNetworkReply::finished,
                     this, [reply, this]()
                     {
    if (reply->error() == QNetworkReply::NoError) 
    {
        QByteArray data = reply->readAll();
        log->Info(data.toStdString());
        websocket();
    } 
    else 
    {
        log->Error(reply->errorString().toStdString() + " " + reply->readAll().toStdString());
    }
    reply->deleteLater(); });
}

void Client::websocket()
{
    ws = new QWebSocket();
    connect(ws, &QWebSocket::connected, [this](){
        this->wsConnected();
    });
    ws->open(QUrl("ws://localhost:51419/ws"));
}

void Client::wsConnected()
{
    log->Info("WS Connected");
    connect(ws, &QWebSocket::textMessageReceived, [this](QString message){
        this->onTextMessageReceived(message);
    });
    //ws->sendTextMessage(QString("FUCK UP THE WORLD!"));
}

void Client::onTextMessageReceived(QString message)
{
    log->Debug(message.toStdString());
}

void Client::wsDisconnected()
{
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
        button->resumeConnect();
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
        button->resumeConnect();
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
        button->resumeConnect();
    } });
}

Client::~Client()
{
    delete process;
    process = nullptr;

    manager->get(QNetworkRequest(QUrl("http://localhost:51419/kill")));
    delete manager;
    manager = nullptr;
}