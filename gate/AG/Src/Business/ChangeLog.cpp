#include "ChangeLog.h"
#include <string.h>
#include "Lib/Base64/Base64.h"

ChangeLog::ChangeLog(const char *mem,int index)
{
    memcpy(this->buffer,mem,LOG_SIZE);
    this->index=index;
    this->logId=0;

    for(int i=0;i<8;i++)
    {
        this->logId|=(((int64_t)(this->buffer[8]))<<(i*8));
    }
}

ChangeLog::ChangeLog(const char *key, uint32_t time)
{
    memset(this->buffer,0,LOG_SIZE);
    buffer[0]='L';
    buffer[1]='O';
    buffer[2]='G';
    buffer[4]=(uint8_t)(time&0xff);
    buffer[5]=(uint8_t)((time>>8)&0xff);
    buffer[6]=(uint8_t)((time>>16)&0xff);
    buffer[7]=(uint8_t)((time>>24)&0xff);
    memcpy(buffer+128,key,128);
}

string ChangeLog::getEvidence()
{
    return base64_encode((const unsigned char*)(this->buffer+128),128);
}

int ChangeLog::getScannTime()
{
    return (this->buffer[4])|(this->buffer[5]<<8)|(this->buffer[6]<<16)|(this->buffer[7]<<24);
}

uint8_t *ChangeLog::serialize()
{

    for(int i=0;i<8;i++)
    {
        this->buffer[i+8]=(uint8_t)((this->logId>>(i*8))&0xff);
    }
    return this->buffer;
}

bool ChangeLog::isUploaed()
{
    return (this->buffer[3]==1)?true:false;
}


