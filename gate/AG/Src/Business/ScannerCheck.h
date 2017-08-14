#ifndef SCANNERCHECK_H
#define SCANNERCHECK_H

#include <mutex>
#include <condition_variable>
#include <thread>
#include <list>

#include "Drivers/Scanner/ScannerManager.h"
#include "Drivers/Scanner/IScannerListener.h"



using namespace std;

class ScannerCheck:IScannerListener
{
private:
    static ScannerCheck* checker;
    static mutex mtx;
    const static int MAX_SEG_NUM=4;
private:
    condition_variable condQr;
    ScannerCheck();
    list<char*> lstQrCode;
    thread *thrdCheck;

    int lastCheckResult;
    string lastErrMsg;
private:
    void checkCode();
protected:
    void ProcessCode(const char *code);
public:
    static ScannerCheck * getInstance();

    int start();

    int getLastCheckResult();
    string getLastErrMsg();


};

#endif // SCANNERCHECK_H
