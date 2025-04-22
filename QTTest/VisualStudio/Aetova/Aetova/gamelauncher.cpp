#include "gamelauncher.h"
#include <iostream>
#include "gamethread.h"

GameLauncher::GameLauncher(QObject* parent)
	: QObject{ parent }
{
}

void GameLauncher::handleFinish(const QString& s)
{
	std::cout << s.toStdString() << std::endl;;
}

void GameLauncher::launchGame(const QString pathToExe, const QString exeName)
{
	std::cout << "LaunchGame" << std::endl;
	GameThread* gt = new GameThread(this);
	gt->pathToExe = pathToExe;
	gt->exeName = exeName;
	connect(gt, &GameThread::threadFinish, this, &GameLauncher::handleFinish);
	connect(gt, &GameThread::finished, gt, &QObject::deleteLater);
	gt->start();
}