#include "ChangeLogManager.h"
#include "Utils/PathUtil.h"
#include "Network/DataServer.h"

#include <iostream>

#include <unistd.h>
#include <fcntl.h>

mutex ChangeLogManager::mtx;
ChangeLogManager* ChangeLogManager::mng=nullptr;


ChangeLogManager::ChangeLogManager()
{
       lastId=-1;
       lastIndex=-1;
}

void ChangeLogManager::log2Server()
{
    unique_lock<mutex> lock(mtx,std::defer_lock);
    ChangeLog *log=nullptr;

    DataServer* server=DataServer::getInstance();



    this->init();
    while (true) {
        lock.lock();
        {
            if(listLog.size()>0)
            {
                log=listLog.front();
            }
            else
            {
                condLog.wait(lock);
            }
        }
        lock.unlock();
        if(log==nullptr) continue;

        Json::Value js;
        js["evidence_key"]=log->getEvidence();
        js["scan_time"]=log->getScannTime();

        NetRequest * req=server->createNetRequest(104,"",js);

        NetMessage *msg=req->waitFor(1000);
        if(msg!=nullptr)
        {
            int ret=msg->getRetCode();
            if(ret==0)
            {
                if(saveLog(log,true))
                {
                    lock.lock();
                    listLog.pop_front();
                    delete log;
                    lock.unlock();
                }
            }

        }
        server->deleteNetRequest(req);





    }
}

void ChangeLogManager::init()
{

    if(this->initialized) return;



    uint8_t buffer[ChangeLog::LOG_SIZE];

    string fileName= PathUtil::getFullPath("change.log");


    this->fd = open(fileName.c_str(), O_RDWR | O_CREAT | O_RSYNC, 0644);
    if (this->fd < 0)
        return;

    int index=0;
    while(true)
    {
        int length=read(fd,buffer,ChangeLog::LOG_SIZE);
        if(length<ChangeLog::LOG_SIZE) break;

        ChangeLog *log=new ChangeLog((const char*)buffer,index);

        if(log->isUploaed()==false)
        {
            listLog.push_back(log);
        }
        if(log->logId>lastId)
        {
            this->lastId=log->logId;
            this->lastIndex=index;
        }
        index++;

    }

    this->initialized=true;
}

bool ChangeLogManager::saveLog(ChangeLog *log,bool upload)
{
    if(this->fd<0) return false;

    uint8_t *buffer=log->serialize(upload);
    lseek(fd, log->index * ChangeLog::LOG_SIZE, SEEK_SET);
    int len=write(fd,buffer,ChangeLog::LOG_SIZE);
    if(len!=ChangeLog::LOG_SIZE) return false;
    fsync(fd);

    std::cout<<"logIndex:"<<log->index<<std::endl;
    return true;

}

ChangeLogManager *ChangeLogManager::getInstance()
{
    unique_lock<mutex> lock(mtx);
    if(mng==nullptr)
    {
        mng=new ChangeLogManager();
    }
    return mng;
}

bool ChangeLogManager::addLog(ChangeLog *log)
{
    bool ret=false;
    unique_lock<mutex> lock(mtx);
    if(listLog.size()<MAX_LOG_NUM)
    {
        listLog.push_back(log);
        if(this->lastId<0)
        {
            log->logId=0;
            log->index=0;
            this->lastId=0;
            this->lastIndex=0;
        }
        else
        {
            log->logId=++lastId;
            log->index=(++lastIndex)%MAX_LOG_NUM;
        }

        std::cout<<log->index<<std::endl;
        if(this->saveLog(log,false))
        {
            ret=true;
        }
        condLog.notify_all();

    }

    return ret;
}

void ChangeLogManager::start()
{
    this->trdLog=new thread(&ChangeLogManager::log2Server,this);
}
