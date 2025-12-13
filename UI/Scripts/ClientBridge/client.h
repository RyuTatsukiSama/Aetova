#include "../Common/common.h"
#include <QObject>

class ButtonGame;
namespace doc {
	class Logger;
}

class Client : QObject
{
	Q_OBJECT

private:
	class QNetworkAccessManager* manager;
	class QProcess* process;
	doc::Logger* log;

public:
	Client();
	~Client();

	void download(ButtonGame* button);
	void pause(ButtonGame* button);
	void resume(ButtonGame* button);
};