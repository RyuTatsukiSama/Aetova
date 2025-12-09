#include "buttongame.h"
#include "../../GameProcess/GameLauncher/gamelauncher.h"
#include "QtWidgets/qapplication.h"
#include "QtGui/qpainter.h"
#include "QtGui/qfontdatabase.h"
#include "../../ClientBridge/client.h"
#include <filesystem>
#include <iostream>

namespace fs = std::filesystem;

ButtonGame::ButtonGame(const QString &name, QWidget *parent) : QPushButton(name, parent)
{
	client = new Client();

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
	   background-color: #8d7bb0; /* l�g�rement plus sombre */
	}
	QPushButton:pressed {
		background-color: #756498; /* encore plus sombre */
	}	
	)");

	setGeometry(
		parent->size().width() / 2 - width() / 2,
		parent->size().height() / 2 - height() / 2,
		width(),
		height());
}

void ButtonGame::startConnect()
{
	std::string path = "wd/downloads/BuildOranys";
	if (fs::exists(path) && fs::is_directory(path)) // the download already started or is done
	{
		if (fs::exists("wd/manifest.json")) // the download isn't finish
		{
			resumeConnect();
		}
		else
		{
			launchConnect();
		}
	}
	else
	{
		downloadConnect();
	}
}

void ButtonGame::downloadConnect()
{
	setText("Download");
	disconnect(this, &QPushButton::released, nullptr, nullptr);
	QApplication::connect(this, &QPushButton::released, [this]()
						  { client->download(this); });
}

void ButtonGame::pauseConnect()
{
	setText("Pause");
	disconnect(this, &QPushButton::released, nullptr, nullptr);
	QApplication::connect(this, &QPushButton::released, [this]()
						  { client->pause(this); });
}

void ButtonGame::resumeConnect()
{
	setText("Resume");
	disconnect(this, &QPushButton::released, nullptr, nullptr);
	QApplication::connect(this, &QPushButton::released, [this]()
						  { client->resume(this); });
}

void ButtonGame::launchConnect()
{
	setText("Launch");
	disconnect(this, &QPushButton::released, nullptr, nullptr);
	QApplication::connect(this, &QPushButton::released, [this]()
						  {
							  launcher->launchGame("BuildOranys", "Oranys"); // TODO : need modulable
						  });
}

void ButtonGame::paintEvent(QPaintEvent *event)
{
	QPushButton::paintEvent(event);
}
