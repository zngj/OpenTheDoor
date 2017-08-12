#include "ChangeLogManager.h"
#include "Utils/PathUtil.h"
#include "Network/DataServer.h"

#include <unistd.h>
#include <fcntl.h>

mutex ChangeLogManager::mtx;
ChangeLogManager* ChangeLogManager::mng=nullptr;


ChangeLogManager::ChangeLogManager()
{

}

void ChangeLogManager::log2Server()
{
    unique_lock<mutex> lock(mtx,std::defer_lock);
    ChangeLog *log=nullptr;

    DataServer* server=DataServer::getInstance();

    Json::Value js;

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

        if(lastLog==nullptr)
        {
            this->lastLog=log;
        }
        else if(lastLog->logId<log->logId)
        {
            delete lastLog;
            this->lastLog=log;
        }

        index++;

    }

    this->initialized=true;
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
        if(this->lastLog==nullptr)
        {
            log->logId=0;
        }
        else
        {
            log->logId=lastLog->logId+1;
        }
        this->lastLog=log;
        condLog.notify_all();
        ret=true;
    }

    return ret;
}

void ChangeLogManager::start()
{
    this->trdLog=new thread(&ChangeLogManager::log2Server,this);
}
