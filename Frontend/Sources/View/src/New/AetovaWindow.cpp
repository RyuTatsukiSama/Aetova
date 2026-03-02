#include "New/AetovaWindow.h" // TODO : Don't forget it when include will be do in CMakeList.txt
#include <docLogger>

AetovaWindow::AetovaWindow(QWidget *parent)
{
    log = new doc::Logger();
}

AetovaWindow::~AetovaWindow()
{
    delete log;
    log = nullptr;
}