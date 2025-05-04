#ifndef TEAMLABEL_H
#define TEAMLABEL_H

#include "../../Common/common.h"
#include "QtWidgets/qlabel.h"

class HTMLLabel : public QLabel
{
	Q_OBJECT

public:
	explicit HTMLLabel(const QString& pathFile, const QRect& geometry, QWidget* parent = nullptr);

protected:
	void paintEvent(QPaintEvent* event) override;

private:
};

#endif // !TEAMLABEL_H
