#ifndef ISCANNERLISTENTER_H
#define ISCANNERLISTENTER_H


class IScannerListener
{
public:
    IScannerListener();
    virtual ~IScannerListener()=0 ;
    virtual void ProcessCode(const char *code)=0;

};

#endif // ISCANNERLISTENTER_H
