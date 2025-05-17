#include "htmllabel.h"
#include "QtCore/qfile.h"
#include "QtCore/qtextstream.h"
#include "QtGui/qfontdatabase.h"

HTMLLabel::HTMLLabel(const QString& pathFile, const QRect& geometry, QWidget* parent) : QLabel(parent)
{
	int id = QFontDatabase::addApplicationFont(":/fonts/Sansation_Regular.ttf");
	QString family = QFontDatabase::applicationFontFamilies(id).at(0);
	QFont sensation(family);
	setFont(sensation);

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

	setStyleSheet(R"(
    background: qlineargradient(
            x1: 0, y1: 0,
            x2: 0, y2: 1,
            stop: 0 #a290c5,
            stop: 1 #f1b56c
        );
    padding: 8px;
	)");

	setGeometry(geometry);

}

void HTMLLabel::paintEvent(QPaintEvent* event)
{
	QLabel::paintEvent(event);
}
