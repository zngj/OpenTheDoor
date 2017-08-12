#ifndef CHANGELOG_H
#define CHANGELOG_H


#include <string>

using namespace std;

class ChangeLog
{
   friend class ChangeLogManager;

public:
       static const int LOG_SIZE   =256;
private:
    uint8_t buffer[LOG_SIZE];
    int index;
    int64_t logId;
public:
    ChangeLog(const char * mem,int index);
    ChangeLog(const char *key, uint32_t time);
    string getEvidence();
    int getScannTime();

    uint8_t *serialize();

    bool isUploaed();
};

#endif // CHANGELOG_H
