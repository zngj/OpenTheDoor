#include "BasicConfig.h"
#include <unistd.h>
#include <fstream>
#include <iostream>

#include "Utils/TimeUtil.h"
#include "Utils/PathUtil.h"

using std::string;
using std::ios;
using std::ifstream;
using std::ofstream;
using std::string;

BasicConfig * BasicConfig::config=nullptr;
std::mutex BasicConfig::mtx;

const string MiscOpt[] = { "baseSecret", "changeId","city", "lat", "lng", "locate",
        "name", "province", "road", "sn", "status", "zip", "zone" };

BasicConfig::BasicConfig()
{
    this->inialized=false;
}

BasicConfig * BasicConfig::getInstance()
{
    std::unique_lock<std::mutex> lock(mtx);
    if(config==nullptr)
    {
        config=new BasicConfig();
        config->loadConfig();
    }
    return config;
}

string BasicConfig::getKeyValue(string key)
{
    if(jsonRoot.isMember(key))
    {
        return jsonRoot[key][0].asString();
    }
    return "";
}

string BasicConfig::getNameNode()
{
    return nameNode;
}

const string BasicConfig::getAppVersion()
{
    return this->AppVersion;
}

string BasicConfig::getSN()
{
    return gateSN;
}

string BasicConfig::getMAC()
{
    return "00:E0:4C:1A:61:CE";
}

string BasicConfig::getGateName()
{
    return this->gateName;
}



void BasicConfig::updataConfig(Json::Value *json)
{
    //misc
    int mSize=sizeof(MiscOpt)/sizeof(MiscOpt[0]);
    for(int i=0;i<mSize;i++)
    {
        string key="IBM.misc."+MiscOpt[i];
        if(this->jsonRoot.isMember(key))
        {
            this->jsonRoot[key][0]=(*json)[MiscOpt[i]];
            this->jsonRoot[key][3]=TimeUtil::getInstance()->getTimeFormat();
        }
        else
        {
            Json::Value jArray;
            jArray[0]=(*json)[MiscOpt[i]];
            jArray[1]="";
            jArray[2]=TimeUtil::getInstance()->getTimeFormat();
            jArray[3]="";
            jArray[4]="y";
            this->jsonRoot[key]=jArray;
        }
    }
    //config
    Json::Value jConfig=(*json)["config"];
    int cSize=jConfig.size();
    for(int i=0;i<cSize;i++)
    {

        string key=jConfig[i]["key"].asString();
        Json::Value jVal=jConfig[i]["value"];
        if(this->jsonRoot.isMember(key))
        {
            this->jsonRoot[key][0]=jVal;
            this->jsonRoot[key][3]=TimeUtil::getInstance()->getTimeFormat();
        }
        else
        {
            Json::Value jArray;
            jArray[0]=jVal;
            jArray[1]="";
            jArray[2]=TimeUtil::getInstance()->getTimeFormat();
            jArray[3]="";
            jArray[4]="y";
            this->jsonRoot[key]=jArray;
        }
    }

    if(this->jsonRoot.isMember("IBM.misc.sn"))
    {
        this->gateSN=this->jsonRoot["IBM.misc.sn"][0].asString();
    }

    if(this->jsonRoot.isMember("IBM.misc.name"))
    {
        this->gateName=this->jsonRoot["IBM.misc.name"][0].asString();
    }

    string keyInit="IBM.app.init";

    this->jsonRoot[keyInit][0]="y";
    this->jsonRoot[keyInit][3]=TimeUtil::getInstance()->getTimeFormat();

    this->inialized=true;
}


void BasicConfig::loadConfig()
{

    string fullName=PathUtil::getFullPath("config.json");
    ifstream ifs(fullName);
    if(ifs.is_open())
    {
        Json::Reader jReader;
        jReader.parse(ifs,jsonRoot);
    }
    if(!jsonRoot.isMember("IBM.misc.sn"))
    {
        Json::Value jArray;
        jArray[0]="010100101";
        jArray[1]="闸机编号";
        jArray[2]=TimeUtil::getInstance()->getTimeFormat();
        jArray[3]="";
        jArray[4]="y";
        jsonRoot["IBM.misc.sn"]=jArray;
    }
    this->gateSN=this->jsonRoot["IBM.misc.sn"][0].asString();
    if(this->jsonRoot.isMember("IBM.misc.name"))
    {
        this->gateName=this->jsonRoot["IBM.misc.name"][0].asString();
    }
    if(!jsonRoot.isMember("IBM.network.namenode"))
    {
        Json::Value jArray;
        jArray[0]="39.108.108.102:8083";
        jArray[1]="名字节点IP:PORT";
        jArray[2]=TimeUtil::getInstance()->getTimeFormat();
        jArray[3]="";
        jArray[4]="y";
        jsonRoot["IBM.network.namenode"]=jArray;
    }
    this->nameNode=jsonRoot["IBM.network.namenode"][0].asString();
    if(!jsonRoot.isMember("IBM.app.init"))
    {
        Json::Value jArray;
        jArray[0]="n";
        jArray[1]="完成初始化";
        jArray[2]=TimeUtil::getInstance()->getTimeFormat();
        jArray[3]="";
        jArray[4]="y";
        jsonRoot["IBM.app.init"]=jArray;
    }
    else
    {
        string initV = jsonRoot["IBM.app.init"][0].asString();
        if(initV=="y")
        {
            this->inialized=true;
        }
    }


    if(ifs.is_open()==false)
    {
        saveConfig();
    }

    ifs.close();

}


void BasicConfig::saveConfig()
{
    string fullName=PathUtil::getFullPath("config.json");
     ofstream ofs(fullName);
     if(ofs.is_open())
     {
         string txt=jsonRoot.toStyledString();
         ofs.write(txt.c_str(),txt.length());
     }
     ofs.close();
}

bool BasicConfig::isInialized()
{
    return inialized;
}
