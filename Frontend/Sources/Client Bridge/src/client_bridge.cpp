#include "client_bridge.h"
#include <docLogger>
#include <QWebSocket>
#include <QProcess>
#include <QNetworkReply>
#include <QNetworkRequest>
#include <QNetworkAccessManager>
#include <nlohmann/json.hpp>
#include "WSReader.h"
#include <gamelauncher.h>
#include <filesystem>
#include <fstream>
#define fs std::filesystem

ClientBridge::ClientBridge(QObject *parent) : QObject(parent)
{
    log = new doc::Logger();

    process->startDetached("clientGo/client.exe"); // TODO : Change this to a "CREATE_PROCESS", or a more multi platforme fonction

    manager = new QNetworkAccessManager(nullptr);

    launcher = new GameLauncher(this);

    QNetworkReply *reply = manager->get(QNetworkRequest(QUrl("http" + baseUrl + port + "/health")));
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
        ConfPort();
    }
    reply->deleteLater(); });
}

ClientBridge::~ClientBridge()
{
    exit();

    delete process;
    process = nullptr;

    delete ws;
    ws = nullptr;

    delete manager;
    manager = nullptr;

    delete launcher;
    launcher = nullptr;

    delete log;
    log = nullptr;
}

void ClientBridge::ConfPort()
{
    std::fstream confFile("Shipyard/.conf", std::ios::in);

    std::string stdPort;
    confFile >> stdPort;

    port = QString::fromStdString(stdPort);

    QNetworkReply *reply = manager->get(QNetworkRequest(QUrl("http" + baseUrl + port + "/health")));
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
    ws->open(QUrl("ws" + baseUrl + port + "/ws"));
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
    // log->Debug(message.toStdString());

    Message mess = j.get<Message>();
    if (mess.type == MONITORING)
    {
        MonitoringData md = mess.readMonitoring();
        monitoringSignal(md.dlPrc, md.dlSpeed, md.wrPrc, md.wrSpeed);
    }
    else if (mess.type == DOWNLOAD_DONE)
    {
        bindFunctionToButton("Launch", "launch");
    }
    else if (mess.type == NEED_UPDATE)
    {
        if (!isResume)
        {
            log->Info("Need update");
            needUpdate = true;
            bindFunctionToButton("Update", "update");
            monitoringSignal(0.f, 0.f, 0.f, 0.f);
        }
    }
    else
        mess.read();
}

void ClientBridge::wsDisconnected()
{
}

void ClientBridge::CallFuncByName(const QString &name)
{
    log->Info(name.toStdString() + "has been called");
    if (name == "download")
    {
        download();
    }
    else if (name == "update")
    {
        update();
    }
    else if (name == "pause")
    {
        pause();
    }
    else if (name == "resume")
    {
        resume();
    }
    else if (name == "resume_upt")
    {
        resumeUpt();
    }
    else if (name == "launch")
    {
        launch();
    }
    else
    {
        log->Error("Function named " + name.toStdString() + " doesn't exist");
    }
}

void ClientBridge::StartBinding()
{
    log->Info("Start Biding");
    std::string path = "Shipyard/app/BuildOranys";

    if (fs::exists(path) && fs::is_directory(path)) // the download already started or is done
    {
        // TODO : Update it when PostreSQL will be here
        if (fs::exists(std::format("Shipyard/downloads/Mfs_{}.json", 0))) // the download isn't finish
        {
            bindFunctionToButton("Resume", "resume");
            isResume = true;
        }
        else if (fs::exists(std::format("Shipyard/downloads/UMfs_{}.json", 0))) // the download isn't finish
        {
            bindFunctionToButton("Resume", "resume_upt");
            isResume = true;
        }
        else
        {
            if (needUpdate)
            {
                bindFunctionToButton("Update", "update");
            }
            else
            {
                bindFunctionToButton("Launch", "launch");
                monitoringSignal(100.f, 0.f, 100.f, 0.f);
            }
        }
    }
    else
    {
        bindFunctionToButton("Download", "download");
    }
}

void ClientBridge::download()
{
    log->Info("Download Start");
    Message dlMessage;
    dlMessage.type = TEXT;
    std::string dlstr = "download";
    dlMessage.data = dlstr;
    json j = dlMessage;
    ws->sendTextMessage(QString::fromStdString(j.dump()));
    bindFunctionToButton("Pause", "pause");
}

void ClientBridge::update()
{
    log->Info("Update Start");
    Message dlMessage;
    dlMessage.type = TEXT;
    std::string dlstr = "update";
    dlMessage.data = dlstr;
    json j = dlMessage;
    ws->sendTextMessage(QString::fromStdString(j.dump()));
    bindFunctionToButton("Pause", "pause");
}

void ClientBridge::pause()
{
    log->Info("Pause called");
    Message dlMessage;
    dlMessage.type = TEXT;
    std::string dlstr = "pause";
    dlMessage.data = dlstr;
    json j = dlMessage;
    ws->sendTextMessage(QString::fromStdString(j.dump()));
    bindFunctionToButton("Resume", "resume");
}

void ClientBridge::resume()
{
    log->Info("Resume called");
    Message dlMessage;
    dlMessage.type = TEXT;
    std::string dlstr = "resume";
    dlMessage.data = dlstr;
    json j = dlMessage;
    ws->sendTextMessage(QString::fromStdString(j.dump()));
    bindFunctionToButton("Pause", "pause");
}

void ClientBridge::resumeUpt()
{
    log->Info("Resume update called");
    Message dlMessage;
    dlMessage.type = TEXT;
    std::string dlstr = "resume_upt";
    dlMessage.data = dlstr;
    json j = dlMessage;
    ws->sendTextMessage(QString::fromStdString(j.dump()));
    bindFunctionToButton("Pause", "pause");
}

void ClientBridge::exit()
{
    log->Info("Resume update called");
    Message dlMessage;
    dlMessage.type = EXIT;
    std::string dlstr = "";
    dlMessage.data = dlstr;
    json j = dlMessage;
    ws->sendTextMessage(QString::fromStdString(j.dump()));
}

void ClientBridge::launch()
{
    log->Info("Launch called");
    launcher->launchGame("BuildOranys", "Oranys");
}