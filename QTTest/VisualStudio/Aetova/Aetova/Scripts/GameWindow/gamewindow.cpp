#include "gamewindow.h"
#include "SplashLabel/splashlabel.h"
#include "ButtonGame/buttongame.h"
#include "HTMLLabel/htmllabel.h"

GameWindow::GameWindow(QWidget* parent) : QWidget(parent)
{
	setMinimumSize(1280, 670);

	// Splash Art 

	splashLabel = new SplashLabel(this);

	// Team

	labelTeam = new HTMLLabel(":/HTML/team.html",
		QRect(
			0,
			height() / 2,
			width() / 2,
			height() / 2
		),
		this);

	// Data Sheet

	labeldatasheet = new HTMLLabel(":/HTML/datasheet.html",
		QRect(
			width() / 2,
			height() / 2,
			width() / 2,
			height() / 2
		),
		this);

	// Button

	buttonGame = new ButtonGame(
		QApplication::translate("childwidget", "Launch Game"),
		this
	);
	buttonGame->connectLauncher();
}

void GameWindow::resizeEvent(QResizeEvent* event)
{
	QWidget::resizeEvent(event);

	splashLabel->setGeometry(
		0,
		0,
		width(),
		height() / 2
	);

	labelTeam->setGeometry(
		0,
		height() / 2,
		width() / 2,
		height() / 2
	);

	labeldatasheet->setGeometry(
		width() / 2,
		height() / 2,
		width() / 2,
		height() / 2
	);

	buttonGame->setGeometry(
		size().width() / 2 - buttonGame->width() / 2,
		size().height() / 2 - buttonGame->height() / 2,
		buttonGame->width(),
		buttonGame->height()
	);
}
