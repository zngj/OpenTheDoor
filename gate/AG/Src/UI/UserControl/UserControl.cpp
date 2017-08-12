#include "UserControl.h"
#include "ControlCommon.h"

UserControl::UserControl(const QWidget *parent,const QRect &rect)
{
    this->widget=(QWidget*)parent;
    this->rectV=rect;
    this->listener=nullptr;
    this->enabled=true;
    this->visible=true;
}

void UserControl::processKey(QKeyEvent *event)
{
    Q_UNUSED(event);
}

void UserControl::addEventListener(IEventListener *listener)
{
    this->listener=listener;
}

void UserControl::triggerEvent(UserControl *control, ControlEvent event)
{
    if(this->listener!=nullptr)
    {
        this->listener->processControlEvent(control,event);
    }
}

void UserControl::setEnable(bool state)
{
    this->enabled=state;
}

bool UserControl::isEnabled()
{
    return this->enabled;
}

void UserControl::setVisibel(bool visible)
{
    this->visible=visible;
}

bool UserControl::isVisible()
{
    return this->visible;
}

void UserControl::setFocus(bool focus)
{
    Q_UNUSED(focus);
}
