#include "New/WebView.h"
#include <QWebEngineView>
#include <QWebChannel>
#include <QFile>
#include <QPainter>
#include <QPainterPath>
#include <docLogger>

WebView::WebView(QWidget *parent)
{
    log = new doc::Logger();
    log->Info("Init Aetova Window...");

    setWindowTitle("Aetova");

    // Frameless Window
    // setAttribute(Qt::WA_TranslucentBackground);
    // setWindowFlags(Qt::FramelessWindowHint | Qt::Window);

    webView = new QWebEngineView(this);
    setCentralWidget(webView);

    webView->setUrl(QUrl("qrc:/DGP/View/DefaultGamePage.html"));

#ifdef _DEBUG
    // Web dev tools
    devTools = new QWebEngineView();
    webView->page()->setDevToolsPage(devTools->page());
    devTools->show();
#endif

    log->Info("Aetova Window Init");
}

WebView::~WebView()
{
    delete devTools;
    devTools = nullptr;

    delete webView;
    webView = nullptr;

    delete log;
    log = nullptr;
}