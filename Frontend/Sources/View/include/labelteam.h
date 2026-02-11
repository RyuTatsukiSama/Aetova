#ifndef TEAMLABEL_H
#define TEAMLABEL_H

#include "../../../Common/common.h"
#include "../htmllabel.h"

class TeamLabel : public HTMLLabel
{
	Q_OBJECT

public:
	explicit TeamLabel(const QString& pathFile, const QRect& geometry, QWidget* parent = nullptr);

protected:
	void paintEvent(QPaintEvent* event) override;

private:
};

#endif // !TEAMLABEL_H
