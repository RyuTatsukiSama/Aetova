// #include "GameWindow/gamewindow.h"
#include "Common/common.h"
// #include "GameProcess/GameThread/gamethread.h"
// #include <QFontDatabase>
// #include <QFont>
#include <QScreen>
#include <docLogger>
#include <QApplication>

#include "../../View/include/New/WebView.h"

// Use QVBoxLayout for better scaling
// Qt has a file manager system call QFile

int main(int argc, char *argv[])
{
	doc::LoggerOptions opts = doc::LoggerOptions::OptionsBuilder().build();
	doc::setGlobalLoggerOptions(opts);
	doc::Logger log("Frontend");
	log.Log(doc::LoggerSeverity::Info, "Start Aetova");

	QApplication app(argc, argv);

	WebView *AetovaWindow = new WebView();

	// put the Windows at the center of the screen
	QScreen *screen = QGuiApplication::primaryScreen();
	QRect screenGeometry = screen->availableGeometry();
	int x = (screenGeometry.width() - AetovaWindow->width()) / 2;
	int y = (screenGeometry.height() - AetovaWindow->height()) / 2;
	AetovaWindow->move(x, y);

	AetovaWindow->show();

	int endCode = app.exec();

	delete AetovaWindow;
	AetovaWindow = nullptr;

	if (endCode == 0)
		log.Info(std::format("Aetova end with code {}", endCode));
	else
		log.Error(std::format("Aetova end with code {}", endCode));

	return endCode;
}