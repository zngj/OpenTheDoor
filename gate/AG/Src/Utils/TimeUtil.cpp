#include "TimeUtil.h"

TimeUtil * TimeUtil::mng=nullptr;
std::mutex TimeUtil::mtx;

TimeUtil::TimeUtil()
{

}

TimeUtil *TimeUtil::getInstance()
{
    std::unique_lock<std::mutex> lock(mtx);
    if(mng==nullptr)
    {
        mng=new TimeUtil();
    }
    return mng;
}

string TimeUtil::getTimeTxt()
{
    struct timeval tv;
    struct timezone tz;
    struct tm *now;

    gettimeofday(&tv, &tz);
    now = localtime(&tv.tv_sec);

    char tmTxt[18];
    sprintf(tmTxt, "%4d%02d%02d%02d%02d%02d%03d", 1900 + now->tm_year,
            1 + now->tm_mon, now->tm_mday, now->tm_hour, now->tm_min,
            now->tm_sec, (int) (tv.tv_usec / 1000));

    return string(tmTxt);
}

string TimeUtil::getTimeTxt(uint32_t unixTm)
{
    struct tm *now;
    time_t tm_t=(time_t)unixTm;
    now = localtime(&tm_t);
    char tmTxt[18];
    sprintf(tmTxt, "%02d/%02d/%02d %02d:%02d", 1900 + now->tm_year-2000,
            1 + now->tm_mon, now->tm_mday, now->tm_hour, now->tm_min);

    return string(tmTxt);
}

string TimeUtil::getTimeFormat()
{
    struct timeval tv;
    struct timezone tz;
    struct tm *now;

    gettimeofday(&tv, &tz);
    now = localtime(&tv.tv_sec);

    char tmTxt[24];
    sprintf(tmTxt, "%4d-%02d-%02d %02d:%02d:%02d", 1900 + now->tm_year,
            1 + now->tm_mon, now->tm_mday, now->tm_hour, now->tm_min,
            now->tm_sec);

    return string(tmTxt);
}

//字符串格式为:2016-09-03 15:47:12 或者 20160903154712000
uint32_t TimeUtil::getUnixTime(string tmTxt)
{
    struct tm time;
    if(tmTxt.length()==0)
    {
        struct timeval tv;
        struct timezone tz;

        gettimeofday(&tv, &tz);
        time = *localtime(&tv.tv_sec);
    }
    else
    {
        const char *tmStr=tmTxt.c_str();
        if (tmStr[4] == '-') {
            time.tm_year = atoi(tmStr) - 1900;
            time.tm_mon = atoi(tmStr + 5) - 1;
            time.tm_mday = atoi(tmStr + 8);
            time.tm_hour = atoi(tmStr + 11);
            time.tm_min = atoi(tmStr + 14);
            time.tm_sec = atoi(tmStr + 17);
            time.tm_isdst = 0; //不使用夏令时
        } else {
            int year, month, day, hour, min, sec, milis;
            sscanf(tmStr, "%4d%02d%02d%02d%02d%02d%03d", &year, &month, &day, &hour,
                    &min, &sec, &milis);

            time.tm_year = year - 1900;
            time.tm_mon = month - 1;
            time.tm_mday = day;
            time.tm_hour = hour;
            time.tm_min = min;
            time.tm_sec = sec;
            time.tm_isdst = 0; //不使用夏令时

        }
    }


    time_t tm = mktime(&time);

    return (uint32_t) tm;
}


