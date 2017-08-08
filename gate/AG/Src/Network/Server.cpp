#include "Server.h"
#include "NetRequest.h"
#include "NetMessage.h"

#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <poll.h>
#include <unistd.h>

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

