#ifndef IEVENTLISTENER_H
#define IEVENTLISTENER_H

#include "ControlCommon.h"

class UserControl;

class IEventListener
{
public:
    IEventListener();

    virtual void processControlEvent(UserControl*control,ControlEvent event)=0;
};

#endif // IEVENTLISTENER_H
