#ifndef NETMESSAGE_H
#define NETMESSAGE_H


#ifdef __arm__
#include <json/json.h>
#else
#include <jsoncpp/json/json.h>
#endif

class NetMessage
{
private:
    uint8_t id;
    Json::Value json;
public:
    NetMessage(uint8_t *data,int length);
    bool isUpMsg();
    uint8_t getID();
};

#endif // NETMESSAGE_H
