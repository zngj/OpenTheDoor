#ifndef BASEDIALOG_H
#define BASEDIALOG_H


#include "UI/UserControl/PushButton.h"
#include "UI/UserControl/LineEdit.h"
#include "UI/UserControl/QRDisplay.h"
#include "UI/UserControl/FrameEdit.h"
#include "UI/UserControl/ListView.h"
#include "UI/UserControl/StepLabel.h"
#include "UI/UserControl/StaticLabel.h"
#include "UI/UserControl/SplitterLabel.h"

#include "PopupWindow.h"

#include <QDialog>
#include <QPixmap>
#include <QRect>

#include <vector>
using std::string;

namespace Ui {
class BaseDialog;
}

class BaseDialog : public QDialog,public IEventListener
{
    Q_OBJECT

public:
    explicit BaseDialog(QWidget *parent = 0);
    ~BaseDialog();

private:
    Ui::BaseDialog *ui;
private:
    QWidget *parentWidget;
    int TitleBarHeigth=52;
    QPixmap mapLogo;
    QPixmap mapSignals[5];
    QRect rectTime;
    QRect rectSignalImg;
    QRect rectSignalDesc;
 protected:   string title;
private:
    int elapsedTime=0;
    int timerID=-1;
    QString keyUpDownTip;
    QString keyEnterTip;
    QString keyExitTip;

    std::vector<UserControl*> listControl;
    std::vector<UserControl*> listControlNoFocus;
    int focusIndex=-1;
    bool masked=false;
    bool isTout=false;
    void setMaskBg(bool mask);

protected:
    int popupWindowIndex;
    PopupWindow *popupWindow;
protected:
    void paintEvent(QPaintEvent *event);
    void timerEvent(QTimerEvent *event);
    void keyPressEvent(QKeyEvent *event);
    virtual void customDraw(QPainter & painter);
    virtual void processControlEvent(UserControl* control,ControlEvent event);
    virtual void processTimeout();
    void closeEvent(QCloseEvent *);
    void customEvent(QEvent *event);

public:
    void setTimeout(int seconds);
    void setTimeout();
    void ModemStatusChanged();
    void setKeyUpDownTip(QString tip);
    void setEnterTip(QString tip);
    void setExitTip(QString tip);
    void setTitle(string title);

    bool isTimeout();
    void moveNextControl();

    void popUpPhoneNum(QString phone);
    void popUpPhoneEnd4(QString p4n);
    void popUpTip(QString tip, QString enterTxt, QString exitTxt);
    void popUpLoading(QString tip);
    void popUpNumEdit(QString num, QString tip, bool align);
    void closePopWindow();

    virtual void processPopup(int retV,QString edit);

    PushButton * createPushButton(const QRect& rect,const QString& caption);
    LineEdit * createLineEditor(const QRect &rect,const QString& tip);
    QRDisplay * createQRDisplay(const QRect & rect);
    FrameEdit * createFrameEdit(const QRect &rect);

    ListView * createListView(const QRect &rect);

    StepLabel * createStepLabel(const QRect &rect);
    StaticLabel * createStaticLabel(const QRect &rect);

    SplitterLabel * createSplitterLabel(const QRect &rect,const QString& caption);




};

#endif // BASEDIALOG_H
