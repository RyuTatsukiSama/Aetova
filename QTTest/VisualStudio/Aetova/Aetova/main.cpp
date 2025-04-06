#include <iostream>
#include <QtWidgets/qmainwindow.h>
#include <QtWidgets/qapplication.h>
#include "gamelauncher.h"

int main(int argc, char* argv[])
{
	QApplication app(argc,argv);

	QMainWindow window;

	window.show();

	GameLauncher launch;
	launch.launchGame("Journeep", "Journeep");

	return app.exec();
}