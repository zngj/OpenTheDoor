#ifndef SCANNERMANAGER_H
#define SCANNERMANAGER_H

#include <pthread.h>
#include "ScannerState.h"
#include "IScannerListener.h"

#define CODE_SIZE   (512)

class ScannerManager
{
private:
    ScannerManager();

    pthread_t thread_rx; //串口接收数据
    int spFd; //串口描述符
    bool isRunning;

    enum ScannerState state;

    IScannerListener * listenter;



    char    code[CODE_SIZE];
    int codeLength;

public:
    static ScannerManager * GetInstance();

    int Open(const char * portName);
    int Close();

    enum  ScannerState GetState();

    void AddListenter(IScannerListener* listenter);
    void RemoveListenter(IScannerListener* listenter);

private:
    static void * ProcCodeRx(void *);
    static ScannerManager * manager;
    static   pthread_mutex_t mutexThis;

    void PortSetting(unsigned int flags);

    void ProcessCode();

};

#endif // SCANNERMANAGER_H
