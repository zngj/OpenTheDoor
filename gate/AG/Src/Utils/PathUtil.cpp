#include "PathUtil.h"
#include <unistd.h>

PathUtil::PathUtil()
{

}

string PathUtil::getFullPath(string file)
{
    char exeName[256];
    readlink("/proc/self/exe", exeName, 256);

    string exePath=string(exeName);
    int index= exePath.find_last_of("/");

    string path=exePath.substr(0,index+1);

    path.append(file);

    return path;
}
