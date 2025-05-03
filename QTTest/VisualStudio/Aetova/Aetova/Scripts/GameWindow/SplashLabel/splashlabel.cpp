#include "splashlabel.h"
#include "QtGui/qevent.h"
#include "QtGui/qpainter.h"

SplashLabel::SplashLabel(QWidget* parent) : QLabel(parent)
{
    setScaledContents(false); 
    
    setPixmap(LoadPixMap(":/sprite/wallpaper.png"));
    setGeometry(
        0,
        0,
        width(),
        height() / 2
    );
}

void SplashLabel::setPixmap(const QPixmap& pixmap)
{
    originalPixmap = pixmap;
    update();
}

void SplashLabel::paintEvent(QPaintEvent* event)
{
    Q_UNUSED(event);

    if (originalPixmap.isNull())
        return;

    QPainter painter(this);
    painter.setRenderHint(QPainter::SmoothPixmapTransform, true);

    QSize widgetSize = size();
    QSize pixmapSize = originalPixmap.size();
    QSize scaledSize = pixmapSize;

    scaledSize.scale(widgetSize, Qt::KeepAspectRatioByExpanding);

    QRect sourceRect = QRect(
        (scaledSize.width() - widgetSize.width()) / 2,
        (scaledSize.height() - widgetSize.height()) / 2,
        widgetSize.width(),
        widgetSize.height()
    );

    QPixmap scaled = originalPixmap.scaled(
        scaledSize,
        Qt::KeepAspectRatioByExpanding,
        Qt::SmoothTransformation
    );

    painter.drawPixmap(0, 0, scaled.copy(sourceRect));
}
