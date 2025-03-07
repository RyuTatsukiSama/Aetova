#include "gamelauncher.h"
#include "gamethread.h"
#include <qdebug.h>

GameLauncher::GameLauncher(QObject *parent)
    : QObject{parent}
{}

void GameLauncher::LaunchGame()
{
    std::cout << "LaunchGame" << std::endl;
    GameThread* gt = new GameThread();
    connect(gt, &GameThread::threadFinish, this, &GameLauncher::HandleFinish);
    connect(gt, &GameThread::finished, gt, &QObject::deleteLater);
    gt->start();
}

void GameLauncher::HandleFinish(const QString &s)
{
    qDebug() << s;
}
