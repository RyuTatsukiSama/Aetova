#pragma once

#include <QMainWindow>

namespace doc 
{
    class Logger;
}

class AetovaWindow
{
private:
    doc::Logger* log;

public:
    AetovaWindow(QWidget *parent = nullptr);
    ~AetovaWindow();
};
