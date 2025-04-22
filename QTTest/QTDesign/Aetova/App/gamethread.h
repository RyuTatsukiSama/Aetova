#ifndef GAMETHREAD_H
#define GAMETHREAD_H

#include <QtCore/qthread.h>
#include "QtCore/qobject.h"
#include <windows.h>
#include <iostream>
#include <string>
using namespace std;

class GameThread : public QThread
{   
    Q_OBJECT

        void run() override;

signals :

    void threadFinish(const QString &s);

public :
    GameThread(QObject* parent = nullptr);

    QString pathToExe;
    QString exeName;
};

#endif // GAMETHREAD_H
