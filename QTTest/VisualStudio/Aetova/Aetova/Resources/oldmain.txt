#include <QtWidgets/qapplication.h>
#include "gamelauncher.h"

int oldmain(int argc, char *argv[])
{
    /*QApplication app(argc, argv);
    GameLauncher gl;

    QQmlApplicationEngine engine;
    const QUrl url(mainQmlFile);
    QObject::connect(
                &engine, &QQmlApplicationEngine::objectCreated, &app,
                [url](QObject *obj, const QUrl &objUrl) {
        if (!obj && url == objUrl)
            QCoreApplication::exit(-1);
    }, Qt::QueuedConnection);

    engine.rootContext()->setContextProperty("gameLauncher",&gl);

    engine.addImportPath(QCoreApplication::applicationDirPath() + "/qml");
    engine.addImportPath(":/");
    engine.load(url);

    if (engine.rootObjects().isEmpty())
        return -1;

    return app.exec();*/

    return EXIT_SUCCESS;
}
