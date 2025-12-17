#include "buttongame.h"
#include "../../GameProcess/GameLauncher/gamelauncher.h"
#include <QApplication>
#include <QPainter>
#include <QFontDatabase>
#include <QPalette>
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
	background-color: #D4AF37;
	color: white;
	border: none;
	border-radius: 12px;
	padding: 10px 20px;
	font-weight: bold;
	}
	QPushButton:hover {
	   background-color: #F0C758; /* l�g�rement plus sombre */
	}
	QPushButton:pressed {
		background-color: #bd9d46; /* encore plus sombre */
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
	setStyleSheet(R"(
	QPushButton {
    background-color: #366CD4;
    color: white;
    border: none;
    border-radius: 12px;
    padding: 10px 20px;
    font-weight: bold;
	}
	QPushButton:hover {
    	background-color: #5086E0; /* légèrement plus clair */
	}
	QPushButton:pressed {
    	background-color: #2556B0; /* plus foncé */
	}
	)");
	disconnect(this, &QPushButton::released, nullptr, nullptr);
	QApplication::connect(this, &QPushButton::released, [this]()
						  {
							  stopConnect();
							  launcher->launchGame("BuildOranys", "Oranys"); // TODO : need modulable
						  });
}

void ButtonGame::stopConnect()
{
	setText("Stop");

	disconnect(this, &QPushButton::released, nullptr, nullptr);
	QApplication::connect(this, &QPushButton::released, [this]()
						  {
							  launchConnect();
							  launcher->launchGame("BuildOranys", "Oranys"); // TODO : need modulable
						  });
}

void ButtonGame::paintEvent(QPaintEvent *event)
{
	QPushButton::paintEvent(event);
}
