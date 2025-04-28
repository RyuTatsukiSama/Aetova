#ifndef GAMEWINDOW_H
#define GAMEWINDOW_H

#include <QtWidgets/qwidget.h>
#include <QtWidgets/qapplication.h>
#include <iostream>
#include <QtWidgets/qmainwindow.h>
#include <QtWidgets/qpushbutton.h>
#include <QtWidgets/qlabel.h>
#include <QtCore/qdir.h>
#include <QtWidgets/qboxlayout.h>
#include "gamelauncher.h"
#include "common.h"

class GameWindow : public QWidget
{
	Q_OBJECT

public:
	explicit GameWindow(QWidget* parent = nullptr);

protected :

private:

};

#endif