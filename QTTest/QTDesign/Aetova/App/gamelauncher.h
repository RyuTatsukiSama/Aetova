#ifndef GAMELAUNCHER_H
#define GAMELAUNCHER_H

#include <QObject>

class GameLauncher : public QObject
{
    Q_OBJECT
public:
    explicit GameLauncher(QObject *parent = nullptr);
public slots :
    void handleFinish(const QString &s);
    void launchGame(const QString pathToExe, const QString exeName);
};

#endif // GAMELAUNCHER_H
