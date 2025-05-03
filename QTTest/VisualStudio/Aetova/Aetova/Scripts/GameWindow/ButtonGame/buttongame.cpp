#include "buttongame.h"
#include "../../GameProcess/GameLauncher/gamelauncher.h"
#include "QtWidgets/qapplication.h"
#include "QtGui/qpainter.h"

ButtonGame::ButtonGame(const QString& name, QWidget* parent) : QPushButton(name, parent)
{
	launcher = new GameLauncher(this);

	this->setFixedSize(150, 40);

	this->setGeometry(
		parent->size().width() / 2 - width() / 2,
		parent->size().height() / 2 - height() / 2,
		width(),
		height()
	);
}

void ButtonGame::connectLauncher()
{
	QApplication::connect(this, &QPushButton::released, [this]() {
		launcher->launchGame("Journeep", "Journeep");
		});
}

void ButtonGame::paintEvent(QPaintEvent* event)
{
	QPushButton::paintEvent(event);
}
