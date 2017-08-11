#include "UI/FormMain.h"

#include <QApplication>


#include "Network/DataServer.h"

#include "Crypto/RSA1024.h"

#include "Utils/PathUtil.h"
#include "Lib/Base64/Base64.h"

#include "Business/ScannerCheck.h"
#include "Crypto/AES128.h"
int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    FormMain w;
    w.show();


    AES128 aes("5454395434473454","6916665466156476");

        char raw[]="D/g8+5RYMv2edAjMyLqmIL0V6guYzrUPcPQDcPUyZp1J9hSn3CmduX/Cxw+W9dhdEhsukyJcEyYc8aNQxF2LPbul4b/lPUJssQh5t+0utJXCKjJY4aqweX75/YUk9kk8FApVQTffyxEF5s0OCWZbtYh2o26nILudu+WcsNEjRt+NxjmLW4LrC06pjCtH1QO1";

        char test[256];

        char * testData="ssdsdd";
        char enTest[64];
        char deTest[64];
        int lenEn=aes.encrypto((uint8_t*)testData,strlen(testData),(uint8_t*)enTest);

        aes.decrypto((uint8_t*)enTest,lenEn,(uint8_t*)deTest);

        int num=aes.decrypto(string(raw),(uint8_t*)test);

        string base64=base64_encode((const unsigned char*)test,num);

        std::cout<<base64<<std::endl;

    DataServer * server=DataServer::getInstance();
    server->start();

    ScannerCheck::getInstance()->start();



    return a.exec();
}
