#include "gamelauncher.h"
#include "gamethread.h"
#include <qdebug.h>

GameLauncher::GameLauncher(QObject *parent)
    : QObject{parent}
{}

void GameLauncher::launchGame(const QString name)
{
    std::cout << "LaunchGame" << std::endl;
    GameThread* gt = new GameThread();
    gt->gName = name;
    connect(gt, &GameThread::threadFinish, this, &GameLauncher::handleFinish);
    connect(gt, &GameThread::finished, gt, &QObject::deleteLater);
    gt->start();
}

void GameLauncher::handleFinish(const QString &s)
{
    qDebug() << s;
}
