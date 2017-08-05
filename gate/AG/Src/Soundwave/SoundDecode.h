#ifndef SOUNDDECODE_H
#define SOUNDDECODE_H


#include <thread>
#include <iostream>
#include <mutex>

using namespace std;
class SoundDecode
{
public:
    SoundDecode();
    void start();
    void stop();
private:

    uint frameLength=0;

    const static int TxtSize=256;
    char frameTxt[TxtSize];

    thread * threadDecode;
    bool isRunning;
    void decode();

    static mutex mtx;
};

#endif // SOUNDDECODE_H
