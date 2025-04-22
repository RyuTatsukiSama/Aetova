#include "gamewindow.h"
#include "common.h"
#include "gamethread.h"

// Use QVBoxLayout for better scaling
// Qt has a file manager system call QFile

int main(int argc, char* argv[])
{
	QApplication app(argc, argv);

	GameWindow* test = new GameWindow();
	test->resize(SCREEN_WIDTH, SCREEN_HEIGHT);
	test->setWindowTitle(QApplication::translate("Aetova", "Aetova"));
	test->show();

	return app.exec();
}