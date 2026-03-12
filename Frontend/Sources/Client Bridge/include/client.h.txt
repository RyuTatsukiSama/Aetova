#ifndef CLIENT_H
#define CLIENT_H

#include <QObject>

class ButtonGame;
namespace doc
{
	class Logger;
}

class Client : QObject
{
	Q_OBJECT

private:
	class QNetworkAccessManager *manager;
	class QProcess *process;
	class QWebSocket *ws;
	doc::Logger *log;

public:
	Client(QObject *parent);
	~Client();

	void websocket();

	void download(ButtonGame *button);
	void pause(ButtonGame *button);
	void resume(ButtonGame *button);

private:
	void wsConnected();
	void wsDisconnected();
	void onTextMessageReceived(QString message);
};

#endif