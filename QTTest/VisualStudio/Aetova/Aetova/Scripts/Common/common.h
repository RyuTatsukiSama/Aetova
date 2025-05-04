#ifndef COMMON_H
#define COMMON_H

constexpr int SCREEN_WIDTH = 1024;
constexpr int SCREEN_HEIGHT = 768;

#include <QtCore/qstring.h>
#include <QtWidgets/qwidget.h>

QPixmap LoadPixMap(QString path);

#endif
