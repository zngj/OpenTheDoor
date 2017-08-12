#include "BaseDialog.h"
#include "ui_BaseDialog.h"
#include <QPainter>
#include <QColor>
#include <QKeyEvent>
#include "UIListener.h"


BaseDialog::BaseDialog(QWidget *parent) :
    QDialog(parent),
    ui(new Ui::BaseDialog)
{
    ui->setupUi(this);
    this->parentWidget=parent;

    mapLogo.load(":/Resource/Image/Icon/logo_icon.png");
    mapSignals[0].load(":/Resource/Image/Icon/network-breaks_icon.png");
    mapSignals[1].load(":/Resource/Image/Icon/network_icon1.png");
    mapSignals[2].load(":/Resource/Image/Icon/network_icon2.png");
    mapSignals[3].load(":/Resource/Image/Icon/network_icon3.png");
    mapSignals[4].load(":/Resource/Image/Icon/network_icon4.png");

    this->masked=false;
    this->popupWindow=nullptr;

    UIListener::getInstance()->addListener(this);

    if(parent!=nullptr)
    {
        QString parentName=parent->objectName();
        if(parentName=="BaseDialog")
        {
            BaseDialog *pDialog=(BaseDialog*)parent;
            pDialog->setTimeout(0);
        }
    }

}

BaseDialog::~BaseDialog()
{
    for(UserControl* control : listControl)
    {
            delete control;
    }
    for(UserControl* control : listControlNoFocus)
    {
            delete control;
    }
    delete ui;
}

void BaseDialog::paintEvent(QPaintEvent *event)
{
    Q_UNUSED(event);

    int width=this->width();
    int height=this->height();
    QColor colorTitle=QColor::fromRgb(0x18,0x18,0x18);
    QColor colorTime=QColor::fromRgb(0xff,0x6b,0x00);
    QFont fontTimeTxt(".PingFang-SC",22);
    QFont fontSignal(".PingFang-SC",15);
    QFont fontTitle(".PingFang-SC",16);
    QFont fontTip(".PingFang-SC",15);
    QPainter painter(this);

    //draw titlebar
    painter.fillRect(0,0,width,TitleBarHeigth,colorTitle);

    painter.drawPixmap(20,TitleBarHeigth/2-mapLogo.height()/2,mapLogo);
    painter.setPen(Qt::white);
    painter.setFont(fontTitle);
    if(!this->title.empty())
    {
        QRect rectTitle(0,0,width,TitleBarHeigth);
        painter.drawText(rectTitle,Qt::AlignCenter,title.c_str());
    }
    //drawtime
    QString txtTm;
    if(this->elapsedTime>0)
    {
       rectTime= QRect(width-92,0,92-10,TitleBarHeigth);

        painter.setPen(colorTime);
        painter.setFont(fontTimeTxt);
        txtTm.append(QString::number(this->elapsedTime));
        txtTm.append("s");
        painter.drawText(rectTime,Qt::AlignVCenter|Qt::AlignRight,txtTm);
    }
    //draw signal
    int drawWidth=width-painter.fontMetrics().width(txtTm)-10-mapSignals[0].width()-20;
    QRect rectSignal(drawWidth,17,22,18);

    int signalVal=3;
    painter.drawPixmap(rectSignal,mapSignals[signalVal]);
    //draw signal type
    painter.setPen(Qt::white);
    painter.setFont(fontSignal);
    rectSignalDesc=QRect(drawWidth-50-10,0,50,TitleBarHeigth);

    painter.drawText(rectSignalDesc,Qt::AlignCenter|Qt::AlignRight,"LTE");


    //draw body
    painter.fillRect(0,TitleBarHeigth,width,height-2*TitleBarHeigth,Qt::white);

    //draw bottom
    painter.fillRect(0,height-TitleBarHeigth,width,TitleBarHeigth,colorTitle);
    painter.setFont(fontTip);
    if(this->keyUpDownTip.length()>0)
    {
        QRect rectPress(20,height-TitleBarHeigth,100,TitleBarHeigth);
        painter.drawText(rectPress,Qt::AlignVCenter|Qt::AlignLeft,"按");

        QRect rectUp(46,height-TitleBarHeigth+14,64,28);
        QPixmap pixUp(":/Resource/Image/Icon/last_s_icon.png");
        painter.drawPixmap(rectUp,pixUp);

        QRect rectSlash(116,height-TitleBarHeigth,120,TitleBarHeigth);
        painter.drawText(rectSlash,Qt::AlignVCenter|Qt::AlignLeft,"/");

        QRect rectDown(129,height-TitleBarHeigth+14,64,28);
        QPixmap pixDown(":/Resource/Image/Icon/next_s_icon.png");
        painter.drawPixmap(rectDown,pixDown);

        QRect rectDesc(197,height-TitleBarHeigth,120,TitleBarHeigth);
        painter.drawText(rectDesc,Qt::AlignVCenter|Qt::AlignLeft,this->keyUpDownTip);


    }
    if(this->keyEnterTip.length()>0)
    {
        QRect rectPress(307,height-TitleBarHeigth,100,TitleBarHeigth);
        painter.drawText(rectPress,Qt::AlignVCenter|Qt::AlignLeft,"按");

        QRect rectEnter(333,height-TitleBarHeigth+14,46,28);
        QPixmap pixEnter(":/Resource/Image/Icon/confirm_s_icon.png");
        painter.drawPixmap(rectEnter,pixEnter);

        QRect rectDesc(383,height-TitleBarHeigth,120,TitleBarHeigth);
        painter.drawText(rectDesc,Qt::AlignVCenter|Qt::AlignLeft,this->keyEnterTip);
    }
    if(this->keyExitTip.length()>0)
    {
        int offset=painter.fontMetrics().width(this->keyExitTip)+20;
        QRect rectTip(width-offset,height-TitleBarHeigth,200,TitleBarHeigth);
        painter.drawText(rectTip,Qt::AlignVCenter|Qt::AlignLeft,keyExitTip);

        QPixmap pix(":/Resource/Image/Icon/return_s_icon.png");
        offset+=pix.width()+5;
        QRect rectExit(width-offset,height-TitleBarHeigth+14,46,28);

        painter.drawPixmap(rectExit,pix);

        offset+=painter.fontMetrics().width("按")+5;
        QRect rectPress(width-offset,height-TitleBarHeigth,100,TitleBarHeigth);
        painter.drawText(rectPress,Qt::AlignVCenter|Qt::AlignLeft,"按");


    }


    //draw control
    for(UserControl* control : listControl)
    {
        if(control->isVisible())
        {
            control->draw(painter);
        }
    }
    for(UserControl* control : listControlNoFocus)
    {
        if(control->isVisible())
        {
            control->draw(painter);
        }
    }

    this->customDraw(painter);

    if(this->masked)
    {
        QRect rectMask(0,this->TitleBarHeigth,width,height-TitleBarHeigth);

        painter.fillRect(rectMask,QBrush(QColor::fromRgbF(0,0,0,0.6f)));
    }
}

void BaseDialog::setTimeout(int seconds)
{
    this->elapsedTime=seconds;
    this->update(rectTime);
    if(seconds>0)
    {
        if(this->timerID<0)
        {
            this->timerID=this->startTimer(1000);
        }

    }
    else
    {
        if(this->timerID>=0)
        {
            this->killTimer(this->timerID);
            this->timerID=-1;
        }
    }
    this->update(QRect(0,0,this->width(),TitleBarHeigth));
}

void BaseDialog::setTimeout()
{
    this->isTout=true;
}

void BaseDialog::timerEvent(QTimerEvent *event)
{
    Q_UNUSED(event);
    if(this->elapsedTime>0)
    {
        this->elapsedTime--;
        this->update(QRect(this->width()-200,0,200,TitleBarHeigth));
    }
    if(this->elapsedTime<=0)
    {
        this->isTout=true;
        processTimeout();
        if(this->popupWindow!=nullptr)
        {
            this->popupWindow->close();
        }
        this->close();
    }
}

void BaseDialog::ModemStatusChanged()
{
   this->update(rectSignalDesc);
   this->update(rectSignalImg);
}

void BaseDialog::setKeyUpDownTip(QString txt)
{
    this->keyUpDownTip=txt;
}
void BaseDialog::setEnterTip(QString tip)
{
    this->keyEnterTip=tip;
}
void BaseDialog::setExitTip(QString tip)
{
    this->keyExitTip=tip;
}
void BaseDialog::setTitle(string title)
{
    this->title=title;
}

bool BaseDialog::isTimeout()
{
    return this->isTout;
}

void BaseDialog::moveNextControl()
{
    UserControl * control=  listControl[focusIndex];
    control->setFocus(false);
    this->update(control->rect());
    focusIndex++;
    if(focusIndex>=(int)listControl.size()) focusIndex=0;
    control=  listControl[focusIndex];
    control->setFocus(true);
    this->update(control->rect());
}

void BaseDialog::popUpPhoneNum(QString phone)
{
    this->setMaskBg(true);
    this->popupWindow=new PopupWindow(this);
    this->popupWindow->move(180,110);
    this->popupWindow->resize(440,260);
    this->popupWindow->setPhoneNum(phone);
    this->popupWindow->setEnterTip("开");
    this->popupWindow->setExitTip("返回");
    this->popupWindow->show();
    this->popupWindow->setFocus();
}

void BaseDialog::popUpPhoneEnd4(QString p4n)
{
    this->setMaskBg(true);
    this->popupWindow=new PopupWindow(this);
    this->popupWindow->move(180,105);
    this->popupWindow->resize(440,260);

    this->popupWindow->setPhoneEdit(p4n);
    this->popupWindow->setEnterTip("开");
    this->popupWindow->setExitTip("返回");
    this->popupWindow->show();
    this->popupWindow->setFocus();
}

void BaseDialog::popUpTip(QString tip,QString enterTxt,QString exitTxt)
{
    this->setMaskBg(true);
    this->popupWindow=new PopupWindow(this);
    this->popupWindow->move(180,105);
    this->popupWindow->resize(440,260);
    this->popupWindow->setTipMsg(tip);
    if(enterTxt.isEmpty()==false)
    {
        this->popupWindow->setEnterTip(enterTxt);
    }
    if(exitTxt.isEmpty()==false)
    {
        this->popupWindow->setExitTip(exitTxt);
    }

    this->popupWindow->show();
    this->popupWindow->setFocus();
}

void BaseDialog::popUpLoading(QString tip)
{
    this->setMaskBg(true);
    this->popupWindow=new PopupWindow(this);
    this->popupWindow->move(228,160);
    this->popupWindow->resize(360,160);
    this->popupWindow->setLoadingTip(tip);

    this->popupWindow->show();
    this->popupWindow->setFocus();

}

void BaseDialog::popUpNumEdit(QString num, QString tip,bool align)
{
    this->setMaskBg(true);
    this->popupWindow=new PopupWindow(this);
    this->popupWindow->move(180,105);
    this->popupWindow->resize(440,260);

    this->popupWindow->setLineEdit(num,tip,align);
    this->popupWindow->setEnterTip("开");
    this->popupWindow->setExitTip("返回");
    this->popupWindow->show();
    this->popupWindow->setFocus();
}

void BaseDialog::closePopWindow()
{
    if(this->popupWindow!=nullptr)
    {
        this->setMaskBg(false);
        this->popupWindow->close();
        this->popupWindow=nullptr;
        this->setFocus();
    }
}

void BaseDialog::setMaskBg(bool mask)
{
    masked=mask;
    this->update(QRect(0,TitleBarHeigth,this->width(),this->height()-TitleBarHeigth+1));
}

void BaseDialog::closeEvent(QCloseEvent *)
{
    UIListener::getInstance()->addListener(this->parentWidget);
    //UIListener::getInstance()->removeListener(this);
}

void BaseDialog::customEvent(QEvent *event)
{
    int type=(int)event->type();
    if(type==WM_LOGIN_BEGIN)
    {
        this->popUpLoading("正在登录，请稍后...");
    }
    else if(type==WM_LOGIN_END)
    {
        this->closePopWindow();
    }
    else if(type==WM_SYNC_BEGIN)
    {
        this->popUpLoading("同步中，请勿走开...");
    }
    else if(type==WM_SYNC_END)
    {
        this->closePopWindow();
    }
    else if(type==WM_BOX_OPEN_BEGIN)
    {
        this->popUpLoading("正在为您开箱，请勿走开...");
    }
    else if(type==WM_BOX_OPEN_END)
    {
        this->closePopWindow();
    }
    else if(type==WM_BOX_GETINFO_BEGIN)
    {
        this->popUpLoading("正在获取信息...");
    }
    else if(type==WM_BOX_GETINFO_END)
    {
        this->closePopWindow();
    }
    else if(type==WM_BOX_CHANGE_BEGIN)
    {
        this->popUpLoading("正在为您，请勿走开...");
    }
    else if(type==WM_BOX_CHANGE_END)
    {
        this->closePopWindow();
    }
    else if(type==WM_CREATE_PARCEL_BEGIN)
    {
        this->popUpLoading("正在为您，请勿走开...");
    }
    else if(type==WM_CREATE_PARCEL_END)
    {
        this->closePopWindow();
    }

}

PushButton *BaseDialog::createPushButton(const QRect &rect, const QString &caption)
{
    PushButton * button=new PushButton(this,rect,caption);
    listControl.push_back(button);
    if(focusIndex<0){
        focusIndex=0;
        button->setFocus(true);
    }
    return button;
}

LineEdit *BaseDialog::createLineEditor(const QRect &rect, const QString &tip)
{
    LineEdit * editor=new LineEdit(this,rect,tip);
    listControl.push_back(editor);
    if(focusIndex<0){
        focusIndex=0;
        editor->setFocus(true);
    }
    return editor;
}

QRDisplay *BaseDialog::createQRDisplay(const QRect &rect)
{
    QRDisplay * display=new QRDisplay(this,rect);

    listControlNoFocus.push_back(display);
    return display;
}

FrameEdit *BaseDialog::createFrameEdit(const QRect &rect)
{
    FrameEdit * edit=new FrameEdit(this,rect);
    listControl.push_back(edit);
    if(focusIndex<0){
        focusIndex=0;
        edit->setFocus(true);
    }
    return edit;
}

ListView *BaseDialog::createListView(const QRect &rect)
{
    ListView * listView=new ListView(this,rect);
    listControl.push_back(listView);
    if(focusIndex<0){
        focusIndex=0;
        listView->setFocus(true);
    }
    return listView;
}

StepLabel *BaseDialog::createStepLabel(const QRect &rect)
{
    StepLabel * label=new StepLabel(this,rect);

    listControlNoFocus.push_back(label);
    return label;
}

StaticLabel *BaseDialog::createStaticLabel(const QRect &rect)
{
    StaticLabel * label=new StaticLabel(this,rect);

    listControlNoFocus.push_back(label);
    return label;
}

SplitterLabel *BaseDialog::createSplitterLabel(const QRect &rect, const QString &caption)
{
    SplitterLabel * label=new SplitterLabel(this,rect,caption);

    listControlNoFocus.push_back(label);
    return label;
}

void BaseDialog::keyPressEvent(QKeyEvent *event)
{
    int key=event->key();
    if(key==Qt::Key_Escape)
    {
        if(this->keyExitTip.length()>0)
        {
            this->close();
        }
    }
    if((key==Qt::Key_PageDown || key==Qt::Key_Down) && listControl.size()>1)
    {
        UserControl * control=  listControl[focusIndex];
        control->setFocus(false);
        this->update(control->rect());
        int selIndex=focusIndex+1;

        for(int i=0;i<listControl.size();i++)
        {
            if(selIndex>=(int)listControl.size())
            {
                selIndex=0;
            }
            control=  listControl[selIndex];
            if(control->isEnabled())
            {
                focusIndex=selIndex;
                break;
            }
            selIndex++;
        }

        control->setFocus(true);
        this->update(control->rect());

    }
    else if((key==Qt::Key_PageUp || key==Qt::Key_Up) && listControl.size()>1)
    {
        UserControl * control=  listControl[focusIndex];
        control->setFocus(false);
        this->update(control->rect());
        int selIndex=focusIndex-1;

        for(unsigned int i=0;i<listControl.size();i++)
        {
            if(selIndex<0) selIndex=(int)listControl.size()-1;
            control=  listControl[selIndex];
            if(control->isEnabled())
            {
                focusIndex=selIndex;
                break;
            }
            selIndex--;
        }
        control->setFocus(true);
        this->update(control->rect());
    }
    else
    {
        if(focusIndex>=0 && focusIndex<(int)listControl.size())
        {
            UserControl * control=listControl[focusIndex];
            control->processKey(event);
        }
    }
}

void BaseDialog::customDraw(QPainter & painter)
{
    Q_UNUSED(painter);
}

void BaseDialog::processControlEvent(UserControl *control, ControlEvent event)
{
    Q_UNUSED(control);
    Q_UNUSED(event);
}

void BaseDialog::processTimeout()
{

}


void BaseDialog::processPopup(int retV, QString edit)
{
    Q_UNUSED(retV);
    Q_UNUSED(edit);
    this->setMaskBg(false);
    this->setFocus();
    this->popupWindow=nullptr;
}


