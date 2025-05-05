#ifndef COMMON_H
#define COMMON_H

constexpr int SCREEN_WIDTH = 1280;
constexpr int SCREEN_HEIGHT = 670;

#include <QtCore/qstring.h>
#include <QtWidgets/qwidget.h>

QPixmap LoadPixMap(QString path);

#endif
