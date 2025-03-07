#ifndef GAMELAUNCHER_H
#define GAMELAUNCHER_H

#include <QObject>

class GameLauncher : public QObject
{
    Q_OBJECT
public:
    explicit GameLauncher(QObject *parent = nullptr);
public slots :
    void HandleFinish(const QString &s);
    void LaunchGame();
};

#endif // GAMELAUNCHER_H
