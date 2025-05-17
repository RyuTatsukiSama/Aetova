#include "gamewindow.h"
#include "SplashLabel/splashlabel.h"
#include "ButtonGame/buttongame.h"
#include "HTMLLabel/htmllabel.h"

GameWindow::GameWindow(QWidget* parent) : QWidget(parent)
{
	setMinimumSize(SCREEN_WIDTH, SCREEN_HEIGHT);

	// Splash Art 

	splashLabel = new SplashLabel(this);

	// Team

	labelTeam = new HTMLLabel(":/HTML/team.html",
		QRect(
			0,
			SCREEN_HEIGHT / 2,
			width() / 2,
			height() - SCREEN_HEIGHT / 2
		),
		this);

	// Data Sheet

	labeldatasheet = new HTMLLabel(":/HTML/datasheet.html",
		QRect(
			width() / 2,
			SCREEN_HEIGHT / 2,
			width() / 2,
			height() - SCREEN_HEIGHT / 2
		),
		this);

	// Button

	buttonGame = new ButtonGame(
		QApplication::translate("childwidget", "Launch Game"),
		this
	);
	buttonGame->connectLauncher();

	// Border window

	//setWindowFlag(Qt::FramelessWindowHint);
}

void GameWindow::resizeEvent(QResizeEvent* event)
{
	QWidget::resizeEvent(event);

	splashLabel->setGeometry(
		0,
		0,
		width(),
		SCREEN_HEIGHT / 2
	);

	labelTeam->setGeometry(
		0,
		SCREEN_HEIGHT / 2,
		width() / 2,
		height() - SCREEN_HEIGHT / 2
	);

	labeldatasheet->setGeometry(
		width() / 2,
		SCREEN_HEIGHT / 2,
		width() / 2,
		height() - SCREEN_HEIGHT / 2
	);

	buttonGame->setGeometry(
		width() / 2 - buttonGame->width() / 2,
		SCREEN_HEIGHT / 2 - buttonGame->height(),
		buttonGame->width(),
		buttonGame->height()
	);
}
