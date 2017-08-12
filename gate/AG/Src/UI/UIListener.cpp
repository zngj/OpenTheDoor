#include "UIListener.h"
#include <QApplication>
#include <QEvent>
#include <qevent.h>

std::mutex UIListener::mtx;
UIListener * UIListener::listner=nullptr;



UIListener::UIListener()
{

}

UIListener *UIListener::getInstance()
{
    std::unique_lock<std::mutex> lock(mtx);
    if(listner==nullptr)
    {
        listner=new UIListener();
    }
    return listner;
}

void UIListener::notifyLogInBegin()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_LOGIN_BEGIN));
    }
}

void UIListener::notifyLogInEnd()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_LOGIN_END));
    }
}

void UIListener::notifyOpenBoxBegin()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_BOX_OPEN_BEGIN));
    }
}

void UIListener::notifyOpenBoxEnd()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_BOX_OPEN_END));
    }
}

void UIListener::notifyChangeBoxBegin()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_BOX_CHANGE_BEGIN));
    }
}

void UIListener::notifyChangeBoxEnd()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_BOX_CHANGE_END));
    }
}

void UIListener::notifyBoxInfoBegin()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_BOX_GETINFO_BEGIN));
    }
}

void UIListener::notifyBoxInfoEnd()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_BOX_GETINFO_END));
    }
}

void UIListener::notifyParcelCreateBegin()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_CREATE_PARCEL_BEGIN));
    }
}

void UIListener::notifyParcelCreateEnd()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_CREATE_PARCEL_END));
    }
}

void UIListener::notifySyncBegin()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_SYNC_BEGIN));
    }
}

void UIListener::notifySyncEnd()
{
    std::unique_lock<std::mutex> lock(mtx);
    for (auto val : listMsgListener)
    {
        QWidget *widget=(QWidget*)val;
        QApplication::postEvent(val,new QEvent((QEvent::Type)WM_SYNC_END));
    }
}

void UIListener::addListener(QWidget *w)
{
    std::unique_lock<std::mutex> lock(mtx);
    listMsgListener.clear();
    listMsgListener.push_back(w);

}

void UIListener::removeListener(QWidget *w)
{
    std::unique_lock<std::mutex> lock(mtx);
    listMsgListener.remove(w);
}




