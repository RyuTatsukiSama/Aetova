#include "backgroundwidget.h"
#include "QtGui/qpainter.h"
#include "QtGui/qpixmap.h"
#include "QtGui/qbrush.h"
#include "../SplashLabel/splashlabel.h"

BackgroundWidget::BackgroundWidget(QWidget* parent) : QWidget(parent)
{
    background = new SplashLabel(parent);
    background->lower();
}

void BackgroundWidget::paintEvent(QPaintEvent* event)
{
    QPainter painter(this);

    background->setGeometry(QRect(
        0,
        0,
        width(),
        SCREEN_HEIGHT / 2
    ));

    QLinearGradient gradient(0, SCREEN_HEIGHT / 4, 0, height());

    gradient.setColorAt(0.0, QColor(0, 0, 0, 0));
    gradient.setColorAt(SCREEN_HEIGHT / 4.f / height(), QColor(162, 144, 197, 255));
    gradient.setColorAt(1.0, QColor(241, 181, 108, 255));

    painter.fillRect(rect(), gradient);
}
