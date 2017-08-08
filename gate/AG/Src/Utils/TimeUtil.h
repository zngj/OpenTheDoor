#ifndef TIMEUTIL_H
#define TIMEUTIL_H


#include <mutex>
#include <chrono>
#include <string>

#include <time.h>
#include <sys/time.h>

using std::string;


class TimeUtil
{
private:
    TimeUtil();

    static std::mutex mtx;
    static TimeUtil *mng;

public:
    static TimeUtil *getInstance();
    string getTimeTxt();
    string getTimeTxt(uint32_t unixTm);
    string getTimeFormat();


    uint32_t getUnixTime(string tmTxt);
};

#endif // TIMEUTIL_H
