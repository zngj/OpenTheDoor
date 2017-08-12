#ifndef POPUPWINDOW_H
#define POPUPWINDOW_H

#include <QDialog>
#include <QPixmap>

#include "UI/UserControl/FrameEdit.h"
#include "UI/UserControl/LineEdit.h"

namespace Ui {
class PopupWindow;
}

class PopupWindow : public QDialog ,public IEventListener
{
    Q_OBJECT

public:
    explicit PopupWindow(QWidget *parent = 0);
    ~PopupWindow();

    void setEnterTip(QString tip);
    void setExitTip(QString tip);

    void setPhoneNum(QString txt);
    void setPhoneEdit(QString p4n);
    void setTipMsg(QString tip);
    void setLoadingTip(QString tip);

    void setLineEdit(QString txt,QString tip,bool align);
protected:
    void paintEvent(QPaintEvent *event);
    void keyPressEvent(QKeyEvent *event);
    void timerEvent(QTimerEvent *event);
    void processControlEvent(UserControl *control, ControlEvent event);
private:
    Ui::PopupWindow *ui;

    QPixmap pixmapCorner[4];
    QPixmap pixEnter;
    QPixmap pixExit;

    QPixmap pixLoading[6];
    QString keyExitTip;
    QString keyEnterTip;
    int returnV=-1;

    //PhoneNum Tip
    QString phoneTxt;

    FrameEdit * phoneEdit;
    QString phone4Num;
    //LineEdit
    QString lineEditTip;
    LineEdit *lineEdit;

    QString tipMsg;

    QString loadingMsg;
    int loadingIndex=-1;
    QRect loadingRect;
    int timerId=-1;
};

#endif // POPUPWINDOW_H
