#include "../../include/New/WebView.h" // TODO : Don't forget it when include will be do in CMakeList.txt
#include <QWebEngineView>
#include <QWebChannel>
#include <QFile>
#include <docLogger>

WebView::WebView(QWidget *parent)
{
    log = new doc::Logger();
    log->Info("Init Aetova Window...");

    setWindowTitle("Aetova");
    resize(1280, 768);

    webView = new QWebEngineView(this);
    setCentralWidget(webView);

    QFile file(":/View/DefaultGamePage.html");

    if (!file.open(QIODevice::ReadOnly | QIODevice::Text))
    {
        log->Error(std::format("Opening file {} failed", file.fileName().toStdString()));
        return;
    }

    QString html = file.readAll();

    webView->setHtml(html);

#ifdef _DEBUG
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