#pragma once

#include <QMainWindow>

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
    doc::Logger *log;
};