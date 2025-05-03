#ifndef BUTTONGAME_H
#define BUTTONGAME_H

#include "../../Common/common.h"
#include "QtWidgets/qpushbutton.h"

class GameLauncher;


class ButtonGame : public QPushButton
{
	Q_OBJECT

public:
	explicit ButtonGame(const QString& name, QWidget* parent = nullptr);
	void connectLauncher();

protected:

	void paintEvent(QPaintEvent* event) override;

private:
	GameLauncher* launcher = nullptr;
};

#endif // !BUTTONGAME_H