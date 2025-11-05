#ifndef COMMON_H
#define COMMON_H

constexpr int SCREEN_WIDTH = 1315;
constexpr int SCREEN_HEIGHT = 568;

#include <QtCore/qstring.h>
#include <QtWidgets/qwidget.h>

QPixmap LoadPixMap(QString path);

#endif
