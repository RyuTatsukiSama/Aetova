#ifndef SPLASHLABEL_H
#define SPLASHLABEL_H

#include "../../Common/common.h"
#include <QtWidgets/qlabel.h>

class SplashLabel : public QLabel
{
	Q_OBJECT

public :
	explicit SplashLabel(QWidget* parent = nullptr);
	void Init(const QPixmap& pixmap);

	void setPixmap(const QPixmap& pixmap);

protected :
	void paintEvent(QPaintEvent* event) override;

private :
	QPixmap originalPixmap;
};

#endif // !SPLASHLABEL_H
