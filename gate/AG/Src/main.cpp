#include "UI/FormMain.h"

#include <QApplication>


#include "Network/DataServer.h"


int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    FormMain w;
    w.show();


    DataServer * server=DataServer::getInstance();
    server->start();

    return a.exec();
}
