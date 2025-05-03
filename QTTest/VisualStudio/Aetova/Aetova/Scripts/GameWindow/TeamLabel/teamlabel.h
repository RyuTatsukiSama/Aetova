#ifndef TEAMLABEL_H
#define TEAMLABEL_H

#include "../../Common/common.h"
#include "QtWidgets/qlabel.h"

class TeamLabel : public QLabel
{
	Q_OBJECT

public:
	explicit TeamLabel(QWidget* parent = nullptr);

protected:
	void paintEvent(QPaintEvent* event) override;

private:
};

#endif // !TEAMLABEL_H
