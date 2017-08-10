#ifndef NETREQUEST_H
#define NETREQUEST_H

#include <stdint.h>
#include "NetMessage.h"
#include <mutex>
#include <condition_variable>
#include <string>
using namespace std;
#define MAX_FRAME_SIZE  (64*1024)

class NetRequest
{

private:
    NetMessage *msgRcv;
    NetMessage *msgSnd;
    std::mutex mtx;
    std::condition_variable cond;
private:
    uint8_t * frameBuffer;
    int frameLength;
    bool frameOK;
    bool met10;

public:
    NetRequest();
    NetRequest(string sn,int type,Json::Value &json);
    virtual ~NetRequest();

    bool putData(uint8_t data);

    NetMessage * getRcvMsg();

    bool matchMessage(NetRequest *req);

    void assignMessage(NetRequest *req);

    NetMessage * getSendMsg();
    NetMessage *waitFor(int mills);
};

#endif // NETREQUEST_H
