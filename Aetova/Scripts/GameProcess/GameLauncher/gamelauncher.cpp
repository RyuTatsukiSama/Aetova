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
}

void GameLauncher::launchGame(const QString pathToExe, const QString exeName)
{
	log.Info("Launch " + exeName.toStdString());
	GameThread *gt = new GameThread(exeName.toStdString(),this);
	gt->pathToExe = pathToExe;
	gt->exeName = exeName;
	connect(gt, &GameThread::threadFinish, this, &GameLauncher::handleFinish);
	connect(gt, &GameThread::finished, gt, &QObject::deleteLater);
	gt->start();
}