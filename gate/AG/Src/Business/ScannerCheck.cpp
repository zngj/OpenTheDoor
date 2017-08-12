#include "ScannerCheck.h"
#include <iostream>
#include <string.h>

#include "Lib/Base64/Base64.h"
#include "CryptoManager.h"
#include "Utils/TimeUtil.h"

#include "Network/DataServer.h"
#include "Business/ChangeLogManager.h"
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

    uint8_t frame[256];
    int frameLen=1;
    int segIndex=0;


    unique_lock<mutex> lock(mtx,std::defer_lock);

    bool checkPass;
    uint32_t scanTime;
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


        char t= frame[frameLen-1];
        int segLen=base64_decode(string(qrCode),(char*)(frame+frameLen-1));


        int sNum=(frame[frameLen-1]>>4);
        int sIndex=(frame[frameLen-1]&0x0f);

        frame[frameLen-1]=t;
         std::cout<<sNum<<":"<<sIndex<<std::endl;

        if(sNum>MAX_SEG_NUM){segIndex=0 ;continue;}

        if(sIndex==segIndex)
        {
            frameLen+=(segLen-1);
            segIndex++;
        }
        else if(sIndex>segIndex || (sIndex< segIndex-1) )
        {
            segIndex=0;
            frameLen=1;
            continue;
        }
        if(segIndex<sNum)
        {
            continue;
        }
        segIndex=0;








        int len=crypto->aesDecrypt(frame+1, frameLen-1,aesDec);
        frameLen=1;
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


        //
        checkPass=true;
        scanTime=TimeUtil::getInstance()->getUnixTime("");
        Json::Value js;

        js["evidence_key"]=evidence_key;


        DataServer *server=DataServer::getInstance();

        NetRequest *req= server->createNetRequest(103,"010100101",js);

        NetMessage *msg=req->waitFor(1000);

        if(msg!=nullptr)
        {
            if(msg->getRetCode()!=0)
            {
                checkPass=false;
            }
            else
            {
                std::cout<<"OK"<<std::endl;
            }
        }

        server->deleteNetRequest(req);


        if(checkPass==false) continue;

        ChangeLog * log=new ChangeLog(scanTime,(const char *)(aesDec+4));

        if(ChangeLogManager::getInstance()->addLog(log)==false)
        {
            delete log;
            continue;
        }
        //let it go


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


