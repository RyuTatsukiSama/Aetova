#include "client_bridge.h"
#include <docLogger>
#include <QWebSocket>
#include <QProcess>
#include <QNetworkReply>
#include <QNetworkRequest>
#include <QNetworkAccessManager>
#include <nlohmann/json.hpp>
#include "WSReader.h"

ClientBridge::ClientBridge(QObject *parent) : QObject(parent)
{
    log = new doc::Logger();

    process = new QProcess(nullptr);
    // process->startDetached("clientGo/client.exe"); // TODO : Change this to a "CREATE_PROCESS", or a more multi platforme fonction

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

void ClientBridge::websocket()
{
    ws = new QWebSocket();
    connect(ws, &QWebSocket::connected, [this]()
            { this->wsConnected(); });
    ws->open(QUrl("ws://localhost:51419/ws"));
}

void ClientBridge::wsConnected()
{
    log->Info("WS Connected");
    connect(ws, &QWebSocket::textMessageReceived, [this](QString message)
            { this->onTextMessageReceived(message); });
    // ws->sendTextMessage(QString("FUCK UP THE WORLD!"));
}

void ClientBridge::onTextMessageReceived(QString message)
{
    auto j = json::parse(message.toStdString());
    log->Debug(message.toStdString());

    Message mess = j.get<Message>();
    if (mess.type == MONITORING)
    {
        MonitoringData md = mess.readMonitoring();
        monitoringSignal(md.dlPrc, md.dlSpeed, md.wrPrc, md.wrSpeed);
    }
    else
        mess.read();
}

void ClientBridge::wsDisconnected()
{
}

void ClientBridge::download()
{
    log->Caller();
    monitoringSignal(0, 0, 0, 0);
}

void ClientBridge::pause()
{
    log->Caller();
}

void ClientBridge::resume()
{
    log->Caller();
}