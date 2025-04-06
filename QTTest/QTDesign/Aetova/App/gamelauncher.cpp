#include "gamelauncher.h"
#include "gamethread.h"
#include <QtCore/qdebug.h>

GameLauncher::GameLauncher(QObject *parent)
    : QObject{parent}
{}

void GameLauncher::launchGame(const QString pathToExe, const QString exeName)
{
    std::cout << "LaunchGame" << std::endl;
    GameThread* gt = new GameThread();
    gt->pathToExe = pathToExe;
    gt->exeName = exeName;
    connect(gt, &GameThread::threadFinish, this, &GameLauncher::handleFinish);
    connect(gt, &GameThread::finished, gt, &QObject::deleteLater);
    gt->start();
}

void GameLauncher::handleFinish(const QString &s)
{
    qDebug() << s;

}
