#include "common.h"

int main(int argc, char *argv[]) 
{
    QApplication app(argc, argv);
    
    QMainWindow window;
    window.setWindowTitle("Hello world!");

    window.show();
    return app.exec();
}