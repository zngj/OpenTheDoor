#ifndef DATASERVER_H
#define DATASERVER_H

#include "Server.h"
#include <mutex>

class DataServer : public Server
{
private:
    static DataServer *server;
    static mutex mtx;
    DataServer();
protected:
     string getEndpoint();

public:
     static DataServer * getInstance();
};

#endif // DATASERVER_H
