#include "NetMessage.h"

NetMessage::NetMessage(uint8_t *data,int length)
{

    this->id=data[23];
}

bool NetMessage::isUpMsg()
{
    return true;
}

uint8_t NetMessage::getID()
{
    return this->id;
}
