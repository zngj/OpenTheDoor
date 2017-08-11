#include "ScannerCheck.h"
#include <iostream>
#include <string.h>

#include "Lib/Base64/Base64.h"
#include "Crypto/AES128.h"

ScannerCheck * ScannerCheck::checker=nullptr;
mutex ScannerCheck::mtx;
ScannerCheck::ScannerCheck()
{

}

void ScannerCheck::checkCode()
{
    uint8_t aesDec[256];
    AES128 aes("5454395434473454","6916665466156476"); //now this is fixed key
    unique_lock<mutex> lock(mtx,std::defer_lock);

    char * qrCode;
    while (true) {

        lock.lock();
        qrCode=nullptr;
        if(lstQrCode.size()>0)
        {
            qrCode=lstQrCode.front();
            lstQrCode.pop_front();
        }
        if(qrCode==nullptr)
        {
            condQr.wait(lock);
        }

        lock.unlock();

        if(qrCode==nullptr) continue;




        //(1)aes128 decode
        int length=aes.decrypto(string(qrCode+2),aesDec);


        string ss= base64_encode(aesDec,length);


        std::cout<<qrCode<<std::endl;

        delete qrCode;


    }
}

void ScannerCheck::ProcessCode(const char *code)
{
    int size=strlen(code);
    if(size>0)
    {
        char* qrCode=new char[size+1];
        if(qrCode!=nullptr)
        {
            memcpy(qrCode,code,size);
            qrCode[size]=0;


            unique_lock<mutex> lock(mtx);
            lstQrCode.push_back(qrCode);
            condQr.notify_one();

        }
    }
}

ScannerCheck *ScannerCheck::getInstance()
{
    unique_lock<mutex> lock(mtx);

    if(checker==nullptr)
    {
        checker=new ScannerCheck();

        ScannerManager* mng=ScannerManager::GetInstance();
        mng->AddListenter(checker);

    }
    return checker;

}

int ScannerCheck::start()
{
    ScannerManager* mng=ScannerManager::GetInstance();
    mng->Open("/dev/ttyACM0");

    thrdCheck =new thread(&ScannerCheck::checkCode,this);

}


