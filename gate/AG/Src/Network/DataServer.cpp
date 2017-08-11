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
    NetRequest * req=this->createNetRequest(100,"010100101",js);



    NetMessage *msg=req->waitFor(1000);
    if(msg!=nullptr)
    {
        retCode=msg->getRetCode();
    }
    this->deleteNetRequest(req);

    if(retCode!=0) return;

    //get key

    req=this->createNetRequest(102,"010100101",js);



    msg=req->waitFor(1000);
    if(msg!=nullptr)
    {
        retCode=msg->getRetCode();
    }
    this->deleteNetRequest(req);


    js["evidence_key"]="pQjfmNL7SK3lV3CKJLAhOfB27VirJAhSLT59HC1uenn2DUGfxV4bLs0Xni4uf7JEjQiob1940NzBUQ9E7XTA5KnII0/L1qrjpyaa/qQi9cF6hvOqNwCmoRd5vO3l28pd0mfIZB9hO6jHuXSAdQI1KsnXBXz05ZJeMBvG2r19K9k=";



    req=this->createNetRequest(103,"010100101",js);



    msg=req->waitFor(1000);
    if(msg!=nullptr)
    {
        retCode=msg->getRetCode();
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
