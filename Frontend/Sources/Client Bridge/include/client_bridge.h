#pragma once

#include <QObject>
#include <QString>
#include "monitoring_data.h"

namespace doc
{
    class Logger;
}

class ClientBridge : public QObject
{
    Q_OBJECT

private:
    class QProcess *process;
    class QWebSocket *ws;
    class QNetworkAccessManager *manager;
    class GameLauncher *launcher;
    doc::Logger *log;

    QString baseUrl = "://localhost:";
    QString port = "51419";

public:
    explicit ClientBridge(QObject *parent = nullptr);

public slots:

    void CallFuncByName(const QString &name);
    void StartBinding();

private:
    void ConfPort();

    void download();
    void pause();
    void resume();
    void launch();

    void websocket();
    void wsConnected();
    void wsDisconnected();
    void onTextMessageReceived(QString message);

signals:
    void bindFunctionToButton(const QString &text, const QString &name);
    void monitoringSignal(float dlPrc, float dlSpeed, float wrPrc, float wrSpeed);
};