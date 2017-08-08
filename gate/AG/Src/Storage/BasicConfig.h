#ifndef BASICCONFIG_H
#define BASICCONFIG_H

#include <mutex>
#ifdef __arm__
#include <json/json.h>
#else
#include <jsoncpp/json/json.h>
#endif


using std::string;

class BasicConfig
{
private:
    BasicConfig();
    static BasicConfig *config;
    static std::mutex mtx;
private:
    const string AppVersion="3.1.1.1";
    Json::Value jsonRoot;
    bool inialized;
    void loadConfig();
    string nameNode;
    string gateSN;
    string gateName;
public:
    bool isInialized();
    void saveConfig();
    static BasicConfig *getInstance();
    string getKeyValue(string key);

    string getNameNode();
    const string getAppVersion();

    string getSN();
    string getMAC();
    string getGateName();



    void updataConfig(Json::Value *json);

};

#endif // BASICCONFIG_H
