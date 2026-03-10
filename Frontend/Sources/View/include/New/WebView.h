#pragma once

#include <QMainWindow>
#include <QObject>

namespace doc
{
    class Logger;
}

class WebView : public QMainWindow
{
    Q_OBJECT

public:
    WebView(QWidget *parent = nullptr);
    ~WebView();

private:
    class QWebEngineView *webView;
    class QWebEngineView *devTools;
    class ClientBridge *bridge;
    doc::Logger *log;
};