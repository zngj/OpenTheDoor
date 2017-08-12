#include "UI/FormMain.h"

#include <QApplication>


#include "Network/DataServer.h"

#include "Crypto/RSA1024.h"
#include "Business/ScannerCheck.h"

#include "Business/ChangeLogManager.h"

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    FormMain w;
    w.show();


    DataServer * server=DataServer::getInstance();
    server->start();

    ChangeLogManager::getInstance()->start();
    ScannerCheck::getInstance()->start();



    return a.exec();
}
