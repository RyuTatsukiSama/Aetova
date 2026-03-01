#include "common.h"
#include <iostream>

QPixmap LoadPixMap(QString path)
{
	QPixmap pixmap(path);
	if (pixmap.isNull()) // Do a function like LoadSprite of Pierre to don't have to write the test everytime
	{
		std::cout << "Error : Qt didn't load the image at path :/sprite/wallpaper.png" << std::endl;
	}

	return pixmap;
}
