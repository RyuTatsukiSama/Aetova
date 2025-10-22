#include "GameWindow/gamewindow.h"
#include "Common/common.h"
#include "GameProcess/GameThread/gamethread.h"
#include <QFontDatabase>
#include <QFont>
#include <QScreen>
#include <Logger.h>

// Use QVBoxLayout for better scaling
// Qt has a file manager system call QFile

int main(int argc, char* argv[])
{
	doc::LoggerOptions opts = doc::LoggerOptions::OptionsBuilder().build();
	doc::setGlobalLoggerOptions(opts);
	doc::Logger log;

	std::string message = "Start of Aetova";
	log.Caller();
	log.Log(doc::LoggerSeverity::Info, message);
	message = "Start of Aetova";

	QApplication app(argc, argv);
	GameWindow* gw = new GameWindow();
	gw->resize(SCREEN_WIDTH, 1000);
	gw->setWindowTitle(QApplication::translate("Aetova", "Aetova"));

	QScreen* screen = QGuiApplication::primaryScreen();
	QRect screenGeometry = screen->availableGeometry();
	int x = (screenGeometry.width() - gw->width()) / 2;
	int y = (screenGeometry.height() - gw->height()) / 2;
	gw->move(x, y);

	gw->show();

	return app.exec();
}