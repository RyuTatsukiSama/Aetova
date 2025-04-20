#include <iostream>
#include <QtWidgets/qmainwindow.h>
#include <QtWidgets/qapplication.h>
#include <QtWidgets/qpushbutton.h>
#include <QtWidgets/qlabel.h>
#include <QtCore/qdir.h>
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
	window->show();
	window->setWindowTitle(QApplication::translate("Aetova", "Aetova"));

	// Splash Art

	QLabel* imageLabel = new QLabel(window);

	imageLabel->setGeometry(
		0,
		0,
		window->width(),
		window->height() / 2.f
	);


	QPixmap pixmap(":/sprite/wallpaper.png");
	if (pixmap.isNull())
	{
		std::cout << "Error : Qt didn't load the image at path :/sprite/wallpaper.png" << std::endl;
	}
	imageLabel->setPixmap(pixmap);
	imageLabel->setScaledContents(true);
	imageLabel->show();

	// Team

	QLabel* labelTeam = new QLabel(window);

	labelTeam->setFrameStyle(QFrame::Panel | QFrame::Sunken);

	labelTeam->setGeometry(
		0,
		window->height() / 2.f,
		window->width() / 2.f,
		window->height() / 2.f
	);

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
	labelTeam->show();

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
	labelDataSheet->show();

	// Button

	GameLauncher launcher;

	QPushButton* button = new QPushButton(
		QApplication::translate("childwidget", "Launch Game"),
		window
	);

	button->setGeometry(
		window->width() / 2 - button->width() / 2,
		window->height() / 2 - button->height() / 2,
		100,
		50
	);

	button->show();

	QApplication::connect(button, &QPushButton::released, [&launcher]() {
		launcher.launchGame("Journeep", "Journeep");
		});

	return app.exec();
}