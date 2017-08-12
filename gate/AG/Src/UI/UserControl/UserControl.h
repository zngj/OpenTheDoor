#ifndef USERCONTROL_H
#define USERCONTROL_H

#include <QObject>
#include <QPainter>
#include <QKeyEvent>

#include "IEventListener.h"
#include "ControlCommon.h"

class UserControl:public QObject
{
private:
    IEventListener *listener;

protected:
    QWidget *widget;
    QRect rectV;
    bool enabled;
    bool visible;
public:
    UserControl(const QWidget *parent,const QRect &rect);
    virtual QRect rect()=0;
    virtual void draw(QPainter &painter)=0;
    virtual void setFocus(bool focus);
    virtual void processKey(QKeyEvent *event);
    virtual void addEventListener(IEventListener* listener);
    virtual void triggerEvent(UserControl * control,ControlEvent event);
    void setEnable(bool state);
    bool isEnabled();

    void setVisibel(bool visible);
    bool isVisible();
};

#endif // USERCONTROL_H
