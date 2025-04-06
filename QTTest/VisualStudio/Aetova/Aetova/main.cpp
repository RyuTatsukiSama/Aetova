#include <iostream>
#include <QtWidgets/qmainwindow.h>
#include <QtWidgets/qapplication.h>

int main(int argc, char* argv[])
{
	QApplication app(argc,argv);

	QMainWindow window;

	window.show();

	return app.exec();
}