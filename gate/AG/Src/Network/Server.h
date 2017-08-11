#ifndef SERVER_H
#define SERVER_H


#include <thread>
#include <mutex>
#include <list>

#include "NetRequest.h"

using namespace std;

#define BUFFER_SIZE (256)

class Server
{
private:
    thread *threadRx;
    bool isRunning;

    int sockfd;

    mutex mtxReq;

    list<NetRequest *> lstReq;
     void procRx();

protected:

    bool threadPerRequest;
    virtual string getEndpoint()=0;
    virtual bool isInitialized();
    virtual void onConnected()=0;
public:
    Server();
    void start();
    NetRequest *createNetRequest(int type, string sn, Json::Value &json);
    void deleteNetRequest(NetRequest *request);
};

#endif // SERVER_H
