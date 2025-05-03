#ifndef GAMEWINDOW_H
#define GAMEWINDOW_H

#include <QtWidgets/qwidget.h>
#include <QtWidgets/qapplication.h>
#include <iostream>
#include <QtWidgets/qmainwindow.h>
#include <QtWidgets/qpushbutton.h>
#include <QtWidgets/qlabel.h>
#include <QtCore/qdir.h>
#include <QtWidgets/qboxlayout.h>
#include "../GameProcess/GameLauncher/gamelauncher.h"
#include "../Common/common.h"

class SplashLabel;

class GameWindow : public QWidget
{
	Q_OBJECT

public:
	explicit GameWindow(QWidget* parent = nullptr);

protected :

	void resizeEvent(QResizeEvent* event) override;

private:
	SplashLabel* splashLabel;
};

#endif