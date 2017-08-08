#ifndef PATHUTIL_H
#define PATHUTIL_H

#include <string>

using std::string;

class PathUtil
{
public:
    PathUtil();
    static string getFullPath(string file);
};

#endif // PATHUTIL_H
