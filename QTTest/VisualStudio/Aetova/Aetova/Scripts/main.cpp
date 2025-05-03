#include "GameWindow/gamewindow.h"
#include "Common/common.h"
#include "GameProcess/GameThread/gamethread.h"

// Use QVBoxLayout for better scaling
// Qt has a file manager system call QFile

int main(int argc, char* argv[])
{
	QApplication app(argc, argv);

	GameWindow* gw = new GameWindow();
	gw->resize(SCREEN_WIDTH, SCREEN_HEIGHT);
	gw->setWindowTitle(QApplication::translate("Aetova", "Aetova"));
	gw->show();

	return app.exec();
}