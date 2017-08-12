#include "DataServer.h"

#include "Storage/BasicConfig.h"
#include "Business/CryptoManager.h"

mutex DataServer::mtx;
DataServer * DataServer::server=nullptr;

DataServer::DataServer():Server()
{

}

string DataServer::getEndpoint()
{
    return BasicConfig::getInstance()->getNameNode();
}

//login
void DataServer::onConnected()
{


    std::thread tdLogin(&DataServer::login,this);
    tdLogin.detach();

}

void DataServer::login()
{

    std::this_thread::sleep_for(std::chrono::milliseconds(500));
    Json::Value js;
    int retCode=-1;


    //login
    NetRequest * req=this->createNetRequest(100,"",js);



    NetMessage *msg=req->waitFor(1000);
    if(msg!=nullptr)
    {
        retCode=msg->getRetCode();
    }
    this->deleteNetRequest(req);

    if(retCode!=0) return;

    //get key

    req=this->createNetRequest(102,"",js);



    msg=req->waitFor(1000);
    if(msg!=nullptr)
    {
        retCode=msg->getRetCode();
        if(retCode==0)
        {
            Json::Value* pJson=msg->getJson();

            string key=(*pJson)["key"].asString();

            if(key.size()>0)
            {
                CryptoManager::getInstance()->changeRSAPubKey(key);
            }
        }
    }
    this->deleteNetRequest(req);





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
