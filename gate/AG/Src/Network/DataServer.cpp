#include "DataServer.h"

#include "Storage/BasicConfig.h"

mutex DataServer::mtx;
DataServer * DataServer::server=nullptr;

DataServer::DataServer():Server()
{

}

string DataServer::getEndpoint()
{
    return BasicConfig::getInstance()->getNameNode();
}

DataServer *DataServer::getInstance()
{
    unique_lock<mutex> lock(mtx);
    if(server==nullptr)
    {
        server=new DataServer();
    }
    return server;
}
