#include "labelteam.h"
#include "QtGui/qpainter.h"
#include "QtWidgets/qstyle.h"
#include "QtWidgets/qstyleoption.h"

TeamLabel::TeamLabel(const QString& pathFile, const QRect& geometry, QWidget* parent) : HTMLLabel(pathFile, geometry, parent)
{
    setStyleSheet(R"(
    QLabel {
        padding: 10px;
        border-radius: 8px;
    }
	)");
}

void TeamLabel::paintEvent(QPaintEvent* event)
{
    QStyleOption opt;
    opt.initFrom(this);
    QPainter painter(this);
    QColor overlay(255, 255, 255, 100); // Noir avec alpha (sombre partiellement)
    painter.fillRect(rect(), overlay);
    style()->drawPrimitive(QStyle::PE_Widget, &opt, &painter, this);
    QLabel::paintEvent(event);
}
