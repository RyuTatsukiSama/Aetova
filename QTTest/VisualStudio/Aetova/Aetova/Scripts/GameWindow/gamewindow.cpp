#include "gamewindow.h"
#include "SplashLabel/splashlabel.h"
#include "ButtonGame/buttongame.h"
#include "HTMLLabel/htmllabel.h"
#include "BackgroundWidget/backgroundwidget.h"
#include "HTMLLabel/LabelTeam/labelteam.h"

GameWindow::GameWindow(QWidget* parent) : QWidget(parent)
{
	setMinimumSize(SCREEN_WIDTH, SCREEN_HEIGHT);

	// Background

	background = new BackgroundWidget(this);

	// Spitch

	labelSpitch = new HTMLLabel(":/HTML/spitch.html",
		QRect(
			0,
			SCREEN_HEIGHT / 2,
			width(),
			100
		),
		this);

	// Data Sheet

	labeldatasheet = new HTMLLabel(":/HTML/datasheet.html",
		QRect(
			40,
			SCREEN_HEIGHT / 2 + labelSpitch->height(),
			width() / 2,
			height() - SCREEN_HEIGHT / 2 + labelSpitch->height()
		),
		this);

	// Team

	labelTeam = new TeamLabel(":/HTML/team.html",
		QRect(
			width() / 2,
			SCREEN_HEIGHT / 2 + labelSpitch->height(),
			width() / 2,
			height() - SCREEN_HEIGHT / 2 + labelSpitch->height()
		),
		this);


	/*// Button

	buttonGame = new ButtonGame(
		QApplication::translate("childwidget", "Launch Game"),
		this
	);
	buttonGame->connectLauncher();

	// Border window

	//setWindowFlag(Qt::FramelessWindowHint);*/
}

void GameWindow::resizeEvent(QResizeEvent* event)
{
	QWidget::resizeEvent(event);

	background->setGeometry(
		0,
		0,
		width(),
		height()
	);

	labelSpitch->setGeometry(
		0,
		SCREEN_HEIGHT / 2,
		width(),
		100
	);

	labeldatasheet->setGeometry(
		40,
		SCREEN_HEIGHT / 2 + labelSpitch->height(),
		width() / 2,
		height() - SCREEN_HEIGHT / 2 + labelSpitch->height()
	);

	labelTeam->setGeometry(
		width() / 2,
		SCREEN_HEIGHT / 2 + labelSpitch->height(),
		width() / 2,
		height() - SCREEN_HEIGHT / 2 + labelSpitch->height()
	);



	/*buttonGame->setGeometry(
		width() / 2 - buttonGame->width() / 2,
		SCREEN_HEIGHT / 4 - buttonGame->height() / 2,
		buttonGame->width(),
		buttonGame->height()
	);*/
}
