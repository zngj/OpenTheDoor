#ifndef NETMESSAGE_H
#define NETMESSAGE_H


#ifdef __arm__
#include <json/json.h>
#else
#include <jsoncpp/json/json.h>
#endif

#include <string>
#include <mutex>

using namespace std;

class NetMessage
{
private:
     static std::mutex mtx;
     static uint32_t gID;
private:
    uint32_t id;
    string sn;
    int msgLength;
    int msgType;
    Json::Value json;
    int sndLength;
    char * sndBuffer;

    int retCode;

private:
     uint32_t getID();
public:
    NetMessage();
    NetMessage(uint8_t *data,int length);
    NetMessage(string sn,int type,Json::Value &json);

    virtual ~NetMessage();


    NetMessage * clone();

    bool isUpMsg();

    int getRetCode();

    int getMsgType();

    char *getSendFrame(int * length);

};

#endif // NETMESSAGE_H
