#include "gamewindow.h"
#include "SplashLabel/splashlabel.h"
#include "ButtonGame/buttongame.h"
#include "QtCore/qfile.h"
#include "QtCore/qtextstream.h"

GameWindow::GameWindow(QWidget* parent) : QWidget(parent)
{
	setMinimumSize(1280, 670);

	// Splash Art 
	// -- TO DO --
	// - Make it look more clean. make a custom label and override resizeEvent

	splashLabel = new SplashLabel(this);

	// Button

	buttonGame = new ButtonGame(
		QApplication::translate("childwidget", "Launch Game"),
		this
	);
	buttonGame->connectLauncher();

	// Team

	QLabel* labelTeam = new QLabel();

	labelTeam->setFrameStyle(QFrame::Panel | QFrame::Sunken);

	QFile file(":/HTML/team.html");

	if (!file.open(QIODevice::ReadOnly | QIODevice::Text))
		return;

	QTextStream in(&file);

	QString line = in.readLine();
	while (!in.atEnd()) {
		line += in.readLine();
	}

	QString textTeam = QString::fromUtf8(line.toUtf8());
	labelTeam->setTextFormat(Qt::RichText);
	labelTeam->setText(textTeam.trimmed());

	labelTeam->setAlignment(Qt::AlignHCenter | Qt::AlignTop);
	labelTeam->setWordWrap(true);

	// Data Sheet

	QLabel* labelDataSheet = new QLabel(this);

	labelDataSheet->setFrameStyle(QFrame::Panel | QFrame::Sunken);

	labelDataSheet->setGeometry(
		this->width() / 2.f,
		this->height() / 2.f,
		this->width() / 2.f,
		this->height() / 2.f
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

		<b>Editor :</b>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>Developer :</b><br>
				<img src=":/sprite/creajeux.png" width="98" height="70">&nbsp;&nbsp;&nbsp;&nbsp;
				<img src=":/sprite/g3d.png" width="71" height="70"><br>
		)"); // Voir pour de espace
	labelDataSheet->setText(textDS.trimmed());

	labelDataSheet->setAlignment(Qt::AlignHCenter | Qt::AlignTop);
	labelDataSheet->setWordWrap(true);
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

	buttonGame->setGeometry(
		size().width() / 2 - buttonGame->width() / 2,
		size().height() / 2 - buttonGame->height() / 2,
		buttonGame->width(),
		buttonGame->height()
	);
}
