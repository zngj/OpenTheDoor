#include "Server.h"
#include "NetRequest.h"
#include "NetMessage.h"
#include "Storage/BasicConfig.h"

#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <poll.h>
#include <unistd.h>
#include <string.h>

#include <iostream>

Server::Server()
{
    this->isRunning=false;
    this->threadPerRequest=false;
}

void Server::start()
{
    this->isRunning=true;
    if(this->threadPerRequest==false)
    {
       threadRx=new thread(&Server::procRx,this);
    }
}


void Server::procRx()
{

    uint8_t buffer[BUFFER_SIZE];
    NetRequest *request=new NetRequest();
    while (true)
    {
        this->isRunning=false;
        if(this->isInitialized()==false)
        {
            std::this_thread::sleep_for(std::chrono::seconds(1));
            continue;
        }
        string ipEndPoint=this->getEndpoint();
        if(ipEndPoint.length()==0)
        {
            std::this_thread::sleep_for(std::chrono::seconds(3));
            continue;
        }
        int index=ipEndPoint.find(':');

        string ipAddr=ipEndPoint.substr(0,index);
        int ipPort=atoi(ipEndPoint.substr(index+1).c_str());
        struct sockaddr_in serverAddr;
        serverAddr.sin_family = AF_INET;
        serverAddr.sin_port = htons(ipPort);
        serverAddr.sin_addr.s_addr = inet_addr(ipAddr.c_str());

        sockfd = socket(AF_INET, SOCK_STREAM, 0);
        int conn = connect(sockfd, (struct sockaddr *) &serverAddr,sizeof(serverAddr));
        if (conn < 0) {
            close(sockfd);
            return;
        }

        this->isRunning=true;

        this->onConnected();
        while (true)
        {
            int length = recv(sockfd, buffer, BUFFER_SIZE, 0);
            if (length <= 0)
                break;
            for(int i=0;i<length;i++)
            {
                if(request->putData(buffer[i]))
                {
                    NetMessage *msg=request->getRcvMsg();
                    if(msg->isUpMsg())
                    {
                        std::unique_lock<std::mutex> lock(mtxReq,std::defer_lock);
                        lock.lock();

                        std::cout<<lstReq.size()<<std::endl;
                        for(NetRequest* req : lstReq)
                        {
                            if(req->matchMessage(request))
                            {
                                req->assignMessage(request);
                            }
                        }
                        lock.unlock();
                    }
                    else
                    {
                        //processDownMsg(msg);
                    }
                }

            }

        }
    }
}

bool Server::isInitialized()
{
    return true;
}


NetRequest *Server::createNetRequest(int type, string sn, Json::Value &json)
{
    string snReal=sn;
    if(sn.length()==0)
    {
        snReal=BasicConfig::getInstance()->getSN();
    }
    NetRequest *request=new NetRequest(snReal,type,json);
    if(this->threadPerRequest)
    {

    }
    else
    {
         std::unique_lock<std::mutex> lock(this->mtxReq,std::defer_lock);

         lock.lock();

         this->lstReq.push_back(request);

         lock.unlock();
         int sndLength=0;
         NetMessage *msg=request->getSendMsg();
         char *data=(char*)msg->getSendFrame(&sndLength);
         if(this->isRunning)
         {
             send(sockfd,data,sndLength,MSG_NOSIGNAL);
         }

    }
    return request;
}

void Server::deleteNetRequest(NetRequest *request)
{
    if(!this->threadPerRequest)
    {
        std::unique_lock<std::mutex> lock(this->mtxReq,std::defer_lock);
        lock.lock();

        this->lstReq.remove(request);

        lock.unlock();
    }
    delete request;
}

