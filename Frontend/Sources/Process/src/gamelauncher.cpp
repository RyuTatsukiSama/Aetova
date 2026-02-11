#include "gamelauncher.h"
#include <iostream>
#include "../GameThread/gamethread.h"

GameLauncher::GameLauncher(QObject *parent)
	: QObject{parent}
{
}

void GameLauncher::handleFinish(const QString &s, const doc::LoggerSeverity &_severity)
{
	log.Log(_severity, s.toStdString());
	delete gt;
	gt = nullptr;
}

void GameLauncher::launchGame(const QString pathToExe, const QString exeName)
{
	log.Info("Launch " + exeName.toStdString());
	gt = new GameThread(exeName.toStdString(),this);
	gt->pathToExe = pathToExe;
	gt->exeName = exeName;
	connect(gt, &GameThread::threadFinish, this, &GameLauncher::handleFinish);
	connect(gt, &GameThread::finished, gt, &QObject::deleteLater);
	gt->start();
}

void GameLauncher::stopGame()
{
	if (gt)
	{
		gt->kill();
		gt->terminate();
	}
}