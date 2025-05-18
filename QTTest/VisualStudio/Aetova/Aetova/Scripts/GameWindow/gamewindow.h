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
class ButtonGame;
class HTMLLabel;
class BackgroundWidget;
class TeamLabel;
class QLabel;

class GameWindow : public QWidget
{
	Q_OBJECT

public:
	explicit GameWindow(QWidget* parent = nullptr);

protected :

	void resizeEvent(QResizeEvent* event) override;

private:
	BackgroundWidget* background;
	QLabel* labelLogo;
	ButtonGame* buttonGame;
	HTMLLabel* labelSpitch;
	TeamLabel* labelTeam;
	HTMLLabel* labeldatasheet;
};

#endif