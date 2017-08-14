#include "ScannerManager.h"

#include <stdio.h>
#include <unistd.h>
#include <fcntl.h>
#include <pthread.h>
#include <termios.h>

ScannerManager *ScannerManager::manager = NULL;

pthread_mutex_t ScannerManager::mutexThis = PTHREAD_MUTEX_INITIALIZER;

ScannerManager::ScannerManager()
{
    this->isRunning=false;
    this->state=StateSerialPortNotSet;
    this->listenter=NULL;
    this->codeLength=-1;
}


ScannerManager * ScannerManager::GetInstance()
{
    pthread_mutex_lock(&ScannerManager::mutexThis);
    if(manager==NULL)
    {
        manager=new ScannerManager();
    }
    pthread_mutex_unlock(&ScannerManager::mutexThis);

    return manager;
}


enum ScannerState ScannerManager::GetState()
{
    return this->state;
}

int ScannerManager::Open(const char *portName)
{
    int ret = 0;
    pthread_mutex_lock(&mutexThis);
    if (this->isRunning == false) {
        spFd = open(portName, O_RDWR | O_NOCTTY);
        if (spFd >= 0) {
            // set port baudrate
            this->PortSetting(B115200 | CS8 | CLOCAL | CREAD);
            tcflush(spFd, TCIFLUSH);

            this->isRunning = true;
            //以分离状态启动线程
            pthread_attr_t attr;
            pthread_attr_init(&attr);
            pthread_attr_setdetachstate(&attr, PTHREAD_CREATE_DETACHED);
            pthread_create(&thread_rx, &attr, ScannerManager::ProcCodeRx, this);
            pthread_attr_destroy(&attr);

            this->state=StateOK;

            ret = 0;
        } else {
            this->state=StateSerialPortNotOpen;
            ret = -1;
        }
    }

    pthread_mutex_unlock(&mutexThis);

    return ret;

}

void ScannerManager::PortSetting(unsigned int flags) {
    if (spFd >= 0) {
        // set port baudrate
        struct termios serial_struct;
        serial_struct.c_cflag = flags;
        serial_struct.c_iflag = IGNPAR;
        serial_struct.c_oflag = 0;
        serial_struct.c_lflag = 0;
        serial_struct.c_cc[VTIME] = 2; // timeout after 0.2s that isn't working
        serial_struct.c_cc[VMIN] = 0;

        tcsetattr(spFd, TCSANOW, &serial_struct);
    }
}


int ScannerManager::Close()
{
    pthread_mutex_lock(&ScannerManager::mutexThis);
    if (this->isRunning) {
        this->isRunning = false;
        close(spFd);
        spFd = -1;
    }
    pthread_mutex_unlock(&ScannerManager::mutexThis);
    return 0;
}



void ScannerManager::ProcessCode()
{

    pthread_mutex_lock(&ScannerManager::mutexThis);

    this->code[this->codeLength]=0;
    if(this->codeLength>0)
    {
        if(this->listenter!=NULL)
        {
            this->listenter->ProcessCode(code);
        }
    }

    this->codeLength=0;
    pthread_mutex_unlock(&ScannerManager::mutexThis);
}

void * ScannerManager::ProcCodeRx(void * args)
{
    ScannerManager * manager = (ScannerManager*) args;

    char *code=manager->code;

    char  buffer[CODE_SIZE];
    while (manager->isRunning) {
        int length = read(manager->spFd, buffer, 128);

        if(length>0)
        {
            for(int i=0;i<length;i++)
            {
                if(buffer[i]=='^')
                {
                    manager->codeLength=0;
                }
                else if(buffer[i]=='$')
                {

                    manager->ProcessCode();
                    manager->codeLength=-1;
                }
                else if(manager->codeLength>=0)
                {
                    code[manager->codeLength++]=buffer[i];
                }
                if(manager->codeLength>=CODE_SIZE)
                {
                    manager->codeLength=-1;
                }
            }
        }

    }


    return NULL;
}

void ScannerManager::AddListenter(IScannerListener *listenter)
{

    pthread_mutex_lock(&mutexThis);
    if(this->listenter==NULL)
    {
        this->listenter=listenter;
    }
    pthread_mutex_unlock(&mutexThis);
}


void ScannerManager::RemoveListenter(IScannerListener *listenter)
{
    pthread_mutex_lock(&mutexThis);
    if(this->listenter==listenter)
    this->listenter=NULL;
    pthread_mutex_unlock(&mutexThis);
}
