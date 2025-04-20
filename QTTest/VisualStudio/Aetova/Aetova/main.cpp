#include <iostream>
#include <QtWidgets/qmainwindow.h>
#include <QtWidgets/qapplication.h>
#include <QtWidgets/qpushbutton.h>
#include <QtWidgets/qlabel.h>
#include <QtCore/qdir.h>
#include <QtWidgets/qboxlayout.h>
#include "gamelauncher.h"
#include "gamewindow.h"
#include "common.h"

// Use QVBoxLayout for better scaling
// Qt has a file manager system call QFile

int main(int argc, char* argv[])
{
	QApplication app(argc, argv);

	QWidget* window = new QWidget();
	window->resize(SCREEN_WIDTH, SCREEN_HEIGHT);
	window->setWindowTitle(QApplication::translate("Aetova", "Aetova"));

	// root layout
	QVBoxLayout* mainLayout = new QVBoxLayout(window);

	// Splash Art 
	// -- TO DO --
	// - Make it look more clean

	QLabel* imageLabel = new QLabel();

	QPixmap pixmap(":/sprite/wallpaper.png");
	if (pixmap.isNull()) // Do a function like LoadSprite of Pierre to don't have to write the test everytime
	{
		std::cout << "Error : Qt didn't load the image at path :/sprite/wallpaper.png" << std::endl;
	}

	imageLabel->setPixmap(pixmap);
	imageLabel->setScaledContents(true);
	imageLabel->setSizePolicy(QSizePolicy::Expanding, QSizePolicy::Expanding);
	imageLabel->setMaximumHeight(window->height() / 2);
	mainLayout->addWidget(imageLabel);

	// Button
	GameLauncher launcher;

	QPushButton* button = new QPushButton(
		QApplication::translate("childwidget", "Launch Game"),
		window
	);

	button->setFixedSize(150, 40);
	mainLayout->addWidget(button, 0, Qt::AlignHCenter);

	QApplication::connect(button, &QPushButton::released, [&launcher]() {
		launcher.launchGame("Journeep", "Journeep");
		});

	// Bottom Layout for Team & DataSheet
	QHBoxLayout* bottomLayout = new QHBoxLayout();

	// Team

	QLabel* labelTeam = new QLabel(window);

	labelTeam->setFrameStyle(QFrame::Panel | QFrame::Sunken);

	QString textTeam = QString::fromUtf8(u8R"(
	<b><u>Team :</u></b><br><br>

	<b>Programmers :</b><br>
	Bréand Amaryne<br>
	<b>Daniel Colin</b><br>
	Paris-Chevalier Thomas<br>
	Rebattet Mickaël<br>
	Ribault Dorian<br>
	Thomas Sylvain<br><br>

	<b>Artists :</b><br>
	Cart Lou Anne<br>
	Langlois Maxime<br>
	Raynal Lucas<br>
	<b>Sarton Kathleen</b>
	)");
	labelTeam->setText(textTeam.trimmed());

	labelTeam->setAlignment(Qt::AlignHCenter | Qt::AlignTop);
	labelTeam->setWordWrap(true);
	bottomLayout->addWidget(labelTeam);

	// Data Sheet

	QLabel* labelDataSheet = new QLabel(window);

	labelDataSheet->setFrameStyle(QFrame::Panel | QFrame::Sunken);

	labelDataSheet->setGeometry(
		window->width() / 2.f,
		window->height() / 2.f,
		window->width() / 2.f,
		window->height() / 2.f
	);

	QString textDS = QString::fromUtf8(u8R"(
	<b><u>Data Sheet :</u></b><br><br>
	<b>Type:</b> Delivery - Sport game / 3D / 3rd person view<br>
 
	<b>Platform :</b> PC<br>
 
	<b>Controls :</b> Controller / Keyboard - Mouse<br>
 
	<b>Language :</b> English<br>
 
	<b>Targeted audience :</b> Intermediate - Experienced<br>
			Explorers - Achievers<br><br>

	<b>Software used :</b><br>
            <img src=":/sprite/unity.png" width="70" height="70">
            <img src=":/sprite/fmod.png" width="189" height="67">
            <img src=":/sprite/maya.png" width="70" height="70">
            <img src=":/sprite/blender.png" width="85" height="70">
            <img src=":/sprite/ps.png" width="70" height="70"><br>
            <img src=":/sprite/pt.png" width="70" height="70">
            <img src=":/sprite/ds.png" width="70" height="70">
            <img src=":/sprite/Notion.png" width="70" height="70"><br><br>

	<b>Equivalent to : </b><br>
			<img src=":/sprite/PEGI_7.png"width="70" height="86"><br>
	)");
	labelDataSheet->setText(textDS.trimmed());

	labelDataSheet->setAlignment(Qt::AlignHCenter | Qt::AlignTop);
	labelDataSheet->setWordWrap(true);
	bottomLayout->addWidget(labelDataSheet);

	mainLayout->addLayout(bottomLayout);


	window->setLayout(mainLayout);
	window->show();

	return app.exec();
}