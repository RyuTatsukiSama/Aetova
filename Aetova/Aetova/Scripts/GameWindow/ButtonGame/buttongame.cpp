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
	sensation.setPointSizeF(14);
	setFont(sensation);

	launcher = new GameLauncher(this);

	setFixedSize(250, 55);

	setStyleSheet(R"(
    QPushButton {
    background-color: #a290c5;
    color: white;
    border: none;
    border-radius: 12px;
    padding: 10px 20px;
    font-weight: bold;
	}
	QPushButton:hover {
	   background-color: #8d7bb0; /* légèrement plus sombre */
	}
	QPushButton:pressed {
		background-color: #756498; /* encore plus sombre */
	}	
	)");

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
		launcher->launchGame("Oranys", "Oranys");
		});
}

void ButtonGame::paintEvent(QPaintEvent* event)
{
	QPushButton::paintEvent(event);
}
