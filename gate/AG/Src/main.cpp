#include "UI/FormMain.h"

#include <QApplication>


#include "Network/DataServer.h"

#include "Crypto/RSA1024.h"

#include "Utils/PathUtil.h"
#include "Lib/Base64/Base64.h"

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    FormMain w;
    w.show();



    DataServer * server=DataServer::getInstance();
    server->start();

    return a.exec();
}
