#pragma once

#include <QObject>
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
    doc::Logger *log;

public:
    explicit ClientBridge(QObject *parent = nullptr);

    Q_INVOKABLE void download();
    Q_INVOKABLE void pause();
    Q_INVOKABLE void resume();

private:
    void websocket();
    void wsConnected();
    void wsDisconnected();
    void onTextMessageReceived(QString message);

signals:
    void monitoringSignal(float dlPrc, float dlSpeed, float wrPrc, float wrSpeed);
};