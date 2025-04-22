#include "gamewindow.h"

GameWindow::GameWindow(QWidget* parent) : QWidget(parent)
{
	// root layout
	QVBoxLayout* mainLayout = new QVBoxLayout(this);

	// Splash Art 
	// -- TO DO --
	// - Make it look more clean. make a custom label and override resizeEvent

	QLabel* imageLabel = new QLabel();

	QPixmap pixmap(":/sprite/wallpaper.png");
	if (pixmap.isNull()) // Do a function like LoadSprite of Pierre to don't have to write the test everytime
	{
		std::cout << "Error : Qt didn't load the image at path :/sprite/wallpaper.png" << std::endl;
	}

	imageLabel->setPixmap(pixmap);
	imageLabel->setScaledContents(true);
	imageLabel->setMinimumSize(QSize(1294 / 2, 733 / 2));
	imageLabel->setSizePolicy(QSizePolicy::Expanding, QSizePolicy::Expanding);
	mainLayout->addWidget(imageLabel);

	// Button
	GameLauncher* launcher = new GameLauncher(this);

	QPushButton* button = new QPushButton(
		QApplication::translate("childwidget", "Launch Game"),
		this
	);

	button->setFixedSize(150, 40);
	mainLayout->addWidget(button, 0, Qt::AlignHCenter);

	QApplication::connect(button, &QPushButton::released, [&launcher]() {
		launcher->launchGame("Journeep", "Journeep");
		});

	// Bottom Layout for Team & DataSheet
	QHBoxLayout* bottomLayout = new QHBoxLayout();

	// Team

	QLabel* labelTeam = new QLabel(this);

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
		)"); // Voir pour de esoace
	labelDataSheet->setText(textDS.trimmed());

	labelDataSheet->setAlignment(Qt::AlignHCenter | Qt::AlignTop);
	labelDataSheet->setWordWrap(true);
	bottomLayout->addWidget(labelDataSheet);

	mainLayout->addLayout(bottomLayout);


	this->setLayout(mainLayout);
}