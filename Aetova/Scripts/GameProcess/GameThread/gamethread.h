#ifndef GAMETHREAD_H
#define GAMETHREAD_H

#include <QtCore/qobject.h>
#include <QtCore/qthread.h>
#include <windows.h>
#include <iostream>
#include <string>
using namespace std;

class GameThread : public QThread
{
	Q_OBJECT

private:
	void run() override;

signals:
	void threadFinish(const QString& s);

public:
	QString pathToExe;
	QString exeName;

	GameThread(QObject* parent = nullptr);
};

#endif // GAMETHREAD_H
