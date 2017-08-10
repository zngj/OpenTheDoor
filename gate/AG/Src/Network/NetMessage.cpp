#include "NetMessage.h"

#include <string.h>

#include <iostream>

std::mutex NetMessage::mtx;

uint32_t NetMessage::gID=0;

NetMessage::NetMessage()
{
    this->sndBuffer=nullptr;
    this->sndLength=0;
    this->retCode=-1;
}

NetMessage::NetMessage(uint8_t *data,int length):NetMessage()
{
    if(data[0]!='S') return;
    if(data[1]!='G') return;
    this->msgLength=(data[21]<<8)|data[22];
    this->msgType=data[23];
    this->id=(data[16]<<24)|(data[17]<<16)|(data[18]<<8)|(data[19]<<0);


    data[msgLength]=0;


    string jTxt=string((char*)(data+32));


    std::cout<<jTxt<<std::endl;
    Json::Reader reader;

    reader.parse(jTxt, this->json);

    this->retCode=0;
    if(this->json.isMember("errcode"))
    {
        this->retCode=json["errcode"].asInt();
    }




}

NetMessage::NetMessage(string sn, int type, Json::Value &json):NetMessage()
{
    this->sn=sn;
    this->msgType=type;
    if(json.isNull())
    {
        Json::Reader reader;
        reader.parse("{}",this->json);
    }
    else
    {
        this->json=json;
    }
    this->id=this->getID();
}
NetMessage::~NetMessage()
{
    if(this->sndBuffer!=nullptr)
    {
        delete this->sndBuffer;
    }
}

NetMessage *NetMessage::clone()
{
    NetMessage *msg=new NetMessage();
    if(msg!=nullptr)
    {
        msg->id=this->id;
        msg->json=this->json;
        msg->msgType=this->msgType;
        msg->sn=this->sn;
        msg->retCode=this->retCode;

    }
    return msg;
}

bool NetMessage::isUpMsg()
{
    return true;
}

int NetMessage::getRetCode()
{
    return this->retCode;
}

int NetMessage::getMsgType()
{
    return this->msgType;
}

uint32_t NetMessage::getID()
{
    uint32_t id=0;
    std::unique_lock<std::mutex> lock(NetMessage::mtx);

    this->id=NetMessage::gID;
    NetMessage::gID++;

    return id;
}

char *NetMessage::getSendFrame(int *length)
{
    std::unique_lock<std::mutex> lock(NetMessage::mtx);
    if(this->sndLength>0)
    {
        *length=this->sndLength;
        return this->sndBuffer;
    }
    else
    {
       const int HeaderSize=32;
       Json::FastWriter jw;
       string jsTxt= jw.write(this->json);
       uint32_t jsSize=jsTxt.length();
       uint32_t msgSize=HeaderSize+jsSize;
       uint8_t header[HeaderSize];

       memset(header,0,HeaderSize);
       header[0] = 'S';
       header[1] = 'G';
       header[2] = 1; //后台与闸机消息
       header[3]=1;//version=1
       header[4] = 0;
       int len = sn.size();
       if (len > 12)
           len = 12; //SN最大12字节
       for (int i = 0; i < len; i++) {
           header[4 + i] = sn[i];
       }
       //消息流水号
       header[16] = ((this->id) >> 24) & 0xff;
       header[17] = ((this->id) >> 16) & 0xff;
       header[18] = ((this->id) >> 8) & 0xff;
       header[19] = ((this->id) >> 0) & 0xff;
       //message attribute
       header[20]=0;
       // frame Length
       header[21] = ((msgSize >> 8) & 0xff); //21
       header[22] = ((msgSize >> 0) & 0xff); //22

       header[23]=this->msgType;
       //frame start
       int lenSnd = 2; //10 03

       uint16_t frameSize=2+msgSize;
       lenSnd+=2; //frame length
       uint8_t lowFrame=(uint8_t)(frameSize&0xff);
       uint8_t highFrame=(uint8_t)((frameSize>>8)&0xff);
       if(lowFrame==0x10)
       {
            lenSnd++;
       }
       if(highFrame==0x10)
       {
           lenSnd++;
       }


       uint16_t checkSum = 0;
       for (int i = 0; i < HeaderSize; i++) {
           checkSum += header[i];
           if(header[i]==0x10)
           {
               lenSnd++;
           }
           lenSnd++;
       }
       for (int i = 0; i < jsSize; i++) {
           uint8_t c=(uint8_t)jsTxt[i];
           checkSum += c;
           if(c==0x10)
           {
               lenSnd++;
           }
           lenSnd++;
       }

       //check sum
       lenSnd+=2;

       uint8_t lowSum=(uint8_t)(checkSum&0xff);
       if(lowSum==0x10)
       {
            lenSnd++;
       }
       uint8_t highSum=(uint8_t)((checkSum>>8)&0xff);
       if(highSum==0x10)
       {
           lenSnd++;
       }

       //frame end
       lenSnd+=2;




       this->sndLength=lenSnd;
       this->sndBuffer=new char[this->sndLength];

       this->sndBuffer[0]=0x10;
       this->sndBuffer[1]=0x02;

       int index = 2;

        //frame length

       sndBuffer[index++]= highFrame;
       if(highFrame==0x10)
       {
             sndBuffer[index++] = 0x10;
       }
       sndBuffer[index++]= lowFrame;
       if(lowFrame==0x10)
       {
          sndBuffer[index++] = 0x10;
       }

       //header

       for (int i = 0; i < HeaderSize; i++) {
           sndBuffer[index++] = header[i];
           if (header[i] ==  0x10)
           {
                sndBuffer[index++]=0x10;
           }
       }
       //body
       for(int i=0;i<jsSize;i++)
       {
           uint8_t c=jsTxt[i];
           sndBuffer[index++] = c;
           if (c ==  0x10) {
               sndBuffer[index++] = c;
           }
       }

       //check sum
       sndBuffer[index++]= highSum;
       if(highSum==0x10)
       {
             sndBuffer[index++] = 0x10;
       }
       sndBuffer[index++]= lowSum;
       if(lowSum==0x10)
       {
          sndBuffer[index++] = 0x10;
       }


       //frame end
       sndBuffer[index++] = 0x10;
       sndBuffer[index++]=0x03;

       *length=this->sndLength;

       return this->sndBuffer;



    }
}
