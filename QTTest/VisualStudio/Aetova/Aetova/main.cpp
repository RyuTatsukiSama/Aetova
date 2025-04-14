#include <iostream>
#include <QtWidgets/qmainwindow.h>
#include <QtWidgets/qapplication.h>
#include <QtWidgets/qpushbutton.h>
#include "gamelauncher.h"

#define SCREEN_WIDTH 1024
#define SCREEN_HEIGHT 720

int main(int argc, char* argv[])
{
	QApplication app(argc,argv);

	QWidget window;
	window.resize(SCREEN_WIDTH, SCREEN_HEIGHT);
	window.show();
	window.setWindowTitle(QApplication::translate("Test", "Test-en"));

	QWidget window2;
	window2.resize(100, 100);

	QPushButton* button = new QPushButton(
		QApplication::translate("childwidget", "Press me"), &window);
	button->move(100, 100);
	button->show();

	QApplication::connect(button, &QPushButton::released, &window2, &QWidget::show);

	return app.exec();
}