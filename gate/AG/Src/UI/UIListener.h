#ifndef UILISTENER_H
#define UILISTENER_H
#include <mutex>
#include <QWidget>
#include <list>

enum CustomEvent
{
    WM_CHECK_OK=0x0400,
    WM_CHECK_FAIL,
    WM_SYNC_BEGIN,
    WM_SYNC_END,
    WM_BOX_OPEN_BEGIN,
    WM_BOX_OPEN_END,
    WM_BOX_CHANGE_BEGIN,
    WM_BOX_CHANGE_END,
    WM_BOX_GETINFO_BEGIN,
    WM_BOX_GETINFO_END,
    WM_CREATE_PARCEL_BEGIN,
    WM_CREATE_PARCEL_END,
    WM_ADMIN_OPEN_BEGIN,
    WM_ADMIN_OPEN_END,
};

class UIListener
{

private:
    UIListener();
    static UIListener * listner;
    static std::mutex mtx;

    std::list<QWidget *> listMsgListener;
public:

    static UIListener * getInstance();


public:
    void notifyCheckOK();
    void notifyCheckFail();
    void notifyOpenBoxBegin();
    void notifyOpenBoxEnd();
    void notifyChangeBoxBegin();
    void notifyChangeBoxEnd();
    void notifyBoxInfoBegin();
    void notifyBoxInfoEnd();

    void notifyParcelCreateBegin();
    void notifyParcelCreateEnd();

    void notifySyncBegin();
    void notifySyncEnd();

    void addListener(QWidget *w);
    void removeListener(QWidget *w);

};

#endif // UILISTENER_H
