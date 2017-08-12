#include "ScannerCheck.h"
#include <iostream>
#include <string.h>

#include "Lib/Base64/Base64.h"
#include "CryptoManager.h"
#include "Utils/TimeUtil.h"

#include "Network/DataServer.h"

#include <math.h>



ScannerCheck * ScannerCheck::checker=nullptr;
mutex ScannerCheck::mtx;
ScannerCheck::ScannerCheck()
{

}

void ScannerCheck::checkCode()
{
    uint8_t aesDec[256];
    uint8_t rsaDec[256];
    unique_lock<mutex> lock(mtx,std::defer_lock);

    char * qrCode;
    while (true) {

        lock.lock();
        {
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
        }
        lock.unlock();

        if(qrCode==nullptr) continue;



        CryptoManager *crypto=CryptoManager::getInstance();


        int len=crypto->aesDecrypt(string(qrCode),aesDec);

        delete qrCode;

        if(aesDec[0]!='S' || aesDec[1]!='G') continue;
        int dir=aesDec[132];

        int unixTime=(int)(aesDec[133]|(aesDec[134]<<8)|(aesDec[135]<<16)|(aesDec[136]<<24));


        int timeNow=(int)(TimeUtil::getInstance()->getUnixTime(""));

        if(abs(timeNow-unixTime)>5*60) continue;


        string evidence_key=base64_encode(aesDec+4,128);
        //rsa decode

        len=crypto->rsaDecrypt(aesDec+4,rsaDec);

        if(len<42) continue;



        sscanf((char*)rsaDec+32,"%d",&unixTime);

        if(abs(timeNow-unixTime)>24*60*60) continue;




        Json::Value js;

        js["evidence_key"]=evidence_key;


        DataServer *server=DataServer::getInstance();

        NetRequest *req= server->createNetRequest(103,"010100101",js);

        NetMessage *msg=req->waitFor(1000);

        if(msg!=nullptr)
        {
            if(msg->getRetCode()==0)
            {
                std::cout<<"OK"<<std::endl;
            }

        }

        server->deleteNetRequest(req);




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


