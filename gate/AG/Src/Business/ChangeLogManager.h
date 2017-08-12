#ifndef CHANGELOGMANAGER_H
#define CHANGELOGMANAGER_H

#include <mutex>
#include <condition_variable>
#include <list>
#include <thread>

#include "ChangeLog.h"


using namespace std;


class ChangeLogManager
{
private:
    static mutex mtx;
    static ChangeLogManager*mng;


public:
    static const int MAX_LOG_NUM=24*60*60;
private:
    int fd;
    list<ChangeLog*> listLog;
    condition_variable condLog;
    thread* trdLog;

    int lastIndex;
    int64_t lastId;


private:
    ChangeLogManager();
    void log2Server();
    bool initialized=false;

    void init();
    bool saveLog(ChangeLog *log,bool upload);
public:
    static ChangeLogManager * getInstance();


    bool addLog(ChangeLog *log);

    void start();

};

#endif // CHANGELOGMANAGER_H
