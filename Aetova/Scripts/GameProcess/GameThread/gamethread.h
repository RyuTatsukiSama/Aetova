#ifndef GAMETHREAD_H
#define GAMETHREAD_H

#include <QThread>
#include <QObject>
#include <windows.h>
#include <iostream>
#include <string>
#include <Logger.h>
using namespace std;

class GameThread : public QThread
{
	Q_OBJECT

private:
	void run() override;

signals:
	void threadFinish(const QString &s, const doc::LoggerSeverity &_severity);

public:
	QString pathToExe;
	QString exeName;

	GameThread(std::string _threadName, QObject *parent = nullptr);

private:
	doc::Logger log;
};

#endif // GAMETHREAD_H
