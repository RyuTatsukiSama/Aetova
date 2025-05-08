#include "buttongame.h"
#include "../../GameProcess/GameLauncher/gamelauncher.h"
#include "QtWidgets/qapplication.h"
#include "QtGui/qpainter.h"
#include "QtGui/qfontdatabase.h"

ButtonGame::ButtonGame(const QString& name, QWidget* parent) : QPushButton(name, parent)
{
	int id = QFontDatabase::addApplicationFont(":/fonts/Sansation_Regular.ttf");
	QString family = QFontDatabase::applicationFontFamilies(id).at(0);
	QFont sensation(family);
	setFont(sensation);

	launcher = new GameLauncher(this);

	setFixedSize(150, 40);

	setGeometry(
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
