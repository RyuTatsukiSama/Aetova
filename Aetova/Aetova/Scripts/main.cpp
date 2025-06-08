#include "GameWindow/gamewindow.h"
#include "Common/common.h"
#include "GameProcess/GameThread/gamethread.h"
#include "QtGui/qfontdatabase.h"
#include "QtGui/qfont.h"
#include "QtGui/qscreen.h"

// Use QVBoxLayout for better scaling
// Qt has a file manager system call QFile

int main(int argc, char* argv[])
{
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