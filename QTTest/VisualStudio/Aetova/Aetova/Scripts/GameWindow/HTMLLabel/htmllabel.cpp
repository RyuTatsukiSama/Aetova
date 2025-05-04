#include "htmllabel.h"
#include "QtCore/qfile.h"
#include "QtCore/qtextstream.h"

HTMLLabel::HTMLLabel(const QString& pathFile, const QRect& geometry, QWidget* parent) : QLabel(parent)
{
	setFrameStyle(QFrame::Panel | QFrame::Sunken);

	QFile file(pathFile);

	if (!file.open(QIODevice::ReadOnly | QIODevice::Text))
		return;

	QTextStream in(&file);

	QString line = in.readLine();
	while (!in.atEnd()) {
		line += in.readLine();
	}

	QString textTeam = QString::fromUtf8(line.toUtf8());
	setTextFormat(Qt::RichText);
	setText(textTeam.trimmed());

	setAlignment(Qt::AlignHCenter | Qt::AlignTop);
	setWordWrap(true);

	setGeometry(geometry);

}

void HTMLLabel::paintEvent(QPaintEvent* event)
{
	QLabel::paintEvent(event);
}
