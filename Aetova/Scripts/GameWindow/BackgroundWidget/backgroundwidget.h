#ifndef BACKGROUNDWIDGET_H
#define BACKGROUNDWIDGET_H

#include "../../Common/common.h"

class SplashLabel;

class BackgroundWidget : public QWidget
{
	Q_OBJECT

public:
	explicit BackgroundWidget(QWidget* parent = nullptr);

protected:
	void paintEvent(QPaintEvent* event) override;

private:
	SplashLabel* background;
};

#endif // !BACKGROUNDWIDGET_H