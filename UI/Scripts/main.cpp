#include "GameWindow/gamewindow.h"
#include "Common/common.h"
#include "GameProcess/GameThread/gamethread.h"
#include <QFontDatabase>
#include <QFont>
#include <QScreen>
#include <docLogger>
#include <QProgressBar>

// Use QVBoxLayout for better scaling
// Qt has a file manager system call QFile

void QTestWindow(QApplication& app)
{
	QWidget *testWindow = new QWidget();
	testWindow->setWindowTitle(QApplication::translate("TestWindow", "TestWindow"));
	testWindow->resize(SCREEN_WIDTH, SCREEN_HEIGHT);

	QProgressBar* pb = new QProgressBar(testWindow);
	pb->resize(500,500);
	pb->setMaximum(100);
	pb->setMinimum(0);
	pb->setValue(50);

	pb->show();

	testWindow->show();
}

int main(int argc, char *argv[])
{
	doc::LoggerOptions opts = doc::LoggerOptions::OptionsBuilder().build();
	doc::setGlobalLoggerOptions(opts);
	doc::Logger log("UI");
	log.Log(doc::LoggerSeverity::Info, "Start Aetova");

	QApplication app(argc, argv);
	GameWindow *gw = new GameWindow();
	gw->resize(SCREEN_WIDTH, 1000);
	gw->setWindowTitle(QApplication::translate("Aetova", "Aetova"));

	// put the GameWindow at the center of the screen
	QScreen *screen = QGuiApplication::primaryScreen();
	QRect screenGeometry = screen->availableGeometry();
	int x = (screenGeometry.width() - gw->width()) / 2;
	int y = (screenGeometry.height() - gw->height()) / 2;
	gw->move(x, y);

	gw->show();

	QTestWindow(app);

	int endCode = app.exec();

	if (endCode == 0)
		log.Info(std::format("Aetova end with code {}", endCode));
	else
		log.Error(std::format("Aetova end with code {}", endCode));

	return endCode;
}