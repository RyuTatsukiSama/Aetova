#ifndef GAMELAUNCHER_H
#define GAMELAUNCHER_H

#include <QObject>
#include <docLogger>

class GameLauncher : public QObject
{
    Q_OBJECT
public:
    explicit GameLauncher(QObject *parent = nullptr);
public slots:
    void handleFinish(const QString &s, const doc::LoggerSeverity &_severity);
    void launchGame(const QString pathToExe, const QString exeName);
    void stopGame();
private:
    doc::Logger log;
    class GameThread* gt;
};

#endif // GAMELAUNCHER_H
