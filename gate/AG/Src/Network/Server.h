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
public:
    Server();
    void start();
};

#endif // SERVER_H
