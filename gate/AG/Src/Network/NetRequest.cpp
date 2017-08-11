#include "NetRequest.h"


NetRequest::NetRequest()
{
    this->frameBuffer=nullptr;
    this->frameLength=0;
    this->met10=false;
    this->frameOK=false;

    this->msgRcv=nullptr;
}

NetRequest::NetRequest(string sn,int type,Json::Value &json):NetRequest()
{
    this->msgRcv=nullptr;
    this->msgSnd=new NetMessage(sn,type,json);
}

NetRequest::~NetRequest()
{
    if(this->frameBuffer!=nullptr)
    {
        delete this->frameBuffer;
    }
}

bool NetRequest::putData(uint8_t data)
{
    if(frameBuffer==nullptr)
    {
        frameBuffer=new uint8_t[MAX_FRAME_SIZE];
    }
    if(this->frameOK || frameLength>=MAX_FRAME_SIZE)
    {
        this->frameOK=false;
        this->frameLength=0;
        this->met10=false;
    }
    switch (data)
    {
    case 0x10:
    {
        if (frameLength == 0)
        {
            frameBuffer[frameLength++] = 0x10;

        }
        else if (frameLength > 1)
        {
            if(!met10)
            {
                frameBuffer[frameLength++] = 0x10;
            }
        }

        met10 = !met10;
    }
        break;
    case 0x02:
    {
        if (met10) //met a frame header
        {
            frameBuffer[0] = 0x10;
            frameBuffer[1] = 0x02;
            frameLength = 2;

        }
        else if (frameLength > 1)
        {
            frameBuffer[frameLength++] = 0x02;
        }

        met10 = false;

    }
        break;
    case 0x03:
    {
        if (met10) //met a frame tailer
        {
            frameBuffer[frameLength++] = 0x03;

            //做长度和校验和验证

            if (frameLength >=8) //2(帧头)+2(长度)+N(数据)+2(校验和)+2(帧尾)
            {
                uint16_t fLen = frameBuffer[3]|(frameBuffer[2]<<8);

                uint16_t sum = 0;
                if (fLen == (frameLength - 2 - 2 - 2))//做长度的校验
                {
                    sum = 0;
                    for (int j = 0; j < fLen - 2; j++)
                    {
                        sum += frameBuffer[4 + j];
                    }

                    uint8_t checkLow = frameBuffer[2 + fLen+1];
                    uint8_t checkHigh = frameBuffer[2 + fLen ];

                    if (((sum & 0xff) == checkLow) && ((sum >> 8) & 0xff) == checkHigh) //校验和通过
                    {
                        //这个时候，数据帧完全OK
                        this->frameOK=true;
                        if(this->msgRcv!=nullptr)
                        {
                            delete this->msgRcv;
                        }

                        this->msgRcv=new NetMessage(frameBuffer+4,frameLength-8);
                    }

                }
            }
            if(!frameOK)
            {
                this->frameLength=0;
            }

        }
        else if (frameLength > 1)
        {
            frameBuffer[frameLength++] = 0x03;
        }
        met10 = false;
    }
        break;
    default:
    {
        if (frameLength > 1)
        {
            frameBuffer[frameLength++] = data;
        }
    }
        break;
    }
    return this->frameOK;
}

NetMessage *NetRequest::getRcvMsg()
{
    return this->msgRcv;
}

bool NetRequest::matchMessage(NetRequest *req)
{
    uint8_t idSnd=this->msgSnd->getMsgType();
    uint8_t idRcv=req->msgRcv->getMsgType();
    return (idSnd==idRcv)?true:false;
    return false;
}

void NetRequest::assignMessage(NetRequest *req)
{
    this->msgRcv=req->msgRcv->clone();
    std::unique_lock<std::mutex> lock(this->mtx);
    this->cond.notify_all();
}

NetMessage *NetRequest::getSendMsg()
{
    return this->msgSnd;
}

NetMessage *NetRequest::waitFor(int mills)
{
    std::unique_lock<std::mutex> lock(this->mtx);
    if(mills>0)
    {
        cond.wait_for(lock,std::chrono::milliseconds(mills));
    }

    return this->msgRcv;
}


