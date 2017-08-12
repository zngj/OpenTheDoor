#include "PopupWindow.h"
#include "ui_popupwindow.h"

#include <QPainter>
#include <QKeyEvent>
#include "BaseDialog.h"


PopupWindow::PopupWindow(QWidget *parent) :
    QDialog(parent),
    ui(new Ui::PopupWindow)
{
    this->lineEdit=nullptr;
    this->phoneEdit=nullptr;
    setWindowFlags(Qt::FramelessWindowHint);
    this->setAttribute(Qt::WA_TranslucentBackground);//设置背景透明
    pixmapCorner[0].load(":/Resource/Image/Icon/up_icon_1.png");
    pixmapCorner[1].load(":/Resource/Image/Icon/up_icon_2.png");
    pixmapCorner[2].load(":/Resource/Image/Icon/up_icon_3.png");
    pixmapCorner[3].load(":/Resource/Image/Icon/up_icon_4.png");

    pixLoading[0].load(":/Resource/Image/Icon/loading1.png");
    pixLoading[1].load(":/Resource/Image/Icon/loading2.png");
    pixLoading[2].load(":/Resource/Image/Icon/loading3.png");
    pixLoading[3].load(":/Resource/Image/Icon/loading4.png");
    pixLoading[4].load(":/Resource/Image/Icon/loading5.png");
    pixLoading[5].load(":/Resource/Image/Icon/loading6.png");

    pixEnter.load(":/Resource/Image/Icon/confirm_icon.png");
    pixExit.load(":/Resource/Image/Icon/return_icon.png");
    this->phoneEdit=nullptr;
    this->loadingRect=QRect(162,94,36,36);
    ui->setupUi(this);
}

PopupWindow::~PopupWindow()
{
    if(this->phoneEdit!=nullptr)
    {
        delete this->phoneEdit;
    }
    if(this->timerId>=0)
    {
        this->killTimer(this->timerId);
    }
    if(this->lineEdit!=nullptr)
    {
        delete this->lineEdit;
    }
    delete ui;
}

void PopupWindow::setEnterTip(QString tip)
{
    this->keyEnterTip=tip;
}

void PopupWindow::setExitTip(QString tip)
{
    this->keyExitTip=tip;
}

void PopupWindow::setPhoneNum(QString txt)
{
    this->phoneTxt=txt;
}

void PopupWindow::setPhoneEdit(QString p4n)
{
    this->phone4Num=p4n;
    this->phoneEdit=new FrameEdit(this,QRect(70,94,300,80));
    this->phoneEdit->setMaxChars(4);
    this->phoneEdit->setFocus(true);
    this->phoneEdit->addEventListener(this);
}

void PopupWindow::setTipMsg(QString tip)
{
    this->tipMsg=tip;
}

void PopupWindow::setLoadingTip(QString tip)
{
    this->loadingMsg=tip;
    this->loadingIndex=0;
    this->startTimer(100);
}

void PopupWindow::setLineEdit(QString txt, QString tip,bool align)
{
    this->lineEditTip=tip;
    this->lineEdit=new LineEdit(this,QRect(70+50,94,200,80),"");
    this->lineEdit->setOnlyDigit();
    this->lineEdit->setMaxChars(2);
    this->lineEdit->setFocus(true);
    this->lineEdit->setText(txt);
    this->lineEdit->setAlignCenter(align);
}

void PopupWindow::paintEvent(QPaintEvent *event)
{
    Q_UNUSED(event);

    QPainter painter(this);
    int width=this->width();
    int height=this->height();

    painter.drawPixmap(0,0,8,8,pixmapCorner[0]);
    painter.drawPixmap(width-8,0,8,8,pixmapCorner[1]);
    painter.drawPixmap(0,height-8,8,8,pixmapCorner[2]);
    painter.drawPixmap(width-8,height-8,8,8,pixmapCorner[3]);

    painter.fillRect(QRect(0,8,8,height-16),Qt::white);
    painter.fillRect(QRect(width-8,8,8,height-16),Qt::white);
    painter.fillRect(QRect(8,0,width-16,height),Qt::white);

    if(!this->tipMsg.isEmpty())
    {
        painter.setPen(Qt::black);
        painter.setFont(QFont(".PingFang-SC",18));
        painter.drawText(QRect(0,100,width,32),Qt::AlignCenter,this->tipMsg);
    }

    if(!this->phoneTxt.isEmpty())
    {
        painter.setPen(Qt::black);
        painter.setFont(QFont(".PingFang-SC",18));
        painter.drawText(QRect(0,60,width,32),Qt::AlignCenter,"确认手机号码");
        painter.setPen(QColor::fromRgb(0xff,0x6b,0x00));
        painter.setFont(QFont(".PingFang-SC",28));
        painter.drawText(QRect(0,106,width,32),Qt::AlignCenter,this->phoneTxt);

    }
    if(!this->lineEditTip.isEmpty())
    {
        painter.setPen(Qt::black);
        painter.setFont(QFont(".PingFang-SC",18));
        painter.drawText(QRect(0,50,width,30),Qt::AlignCenter,this->lineEditTip);
    }
    if(this->lineEdit!=nullptr)
    {
        this->lineEdit->draw(painter);
    }
    if(this->phoneEdit!=nullptr)
    {
        QString p4n=this->phoneEdit->getText();
        if(p4n.length()==4 && p4n!=this->phone4Num)
        {
            painter.setPen(QColor::fromRgb(0xff,0x00,0x00));
            painter.setFont(QFont(".PingFang-SC",18));
            painter.drawText(QRect(0,50,width,30),Qt::AlignCenter,"输入错误，请重试");
        }
        else
        {
            painter.setPen(Qt::black);
            painter.setFont(QFont(".PingFang-SC",18));
            painter.drawText(QRect(0,50,width,30),Qt::AlignCenter,"请输入登录手机号后4位");
        }

        this->phoneEdit->draw(painter);
    }

    if(!this->loadingMsg.isEmpty())
    {
        painter.setPen(Qt::black);
        painter.setFont(QFont(".PingFang-SC",18));
        painter.drawText(QRect(0,52,width,32),Qt::AlignCenter,this->loadingMsg);
    }
    if(this->loadingIndex>=0)
    {
        painter.drawPixmap(this->loadingRect,this->pixLoading[this->loadingIndex]);
    }

    painter.setPen(QColor::fromRgb(0x3d,0x3d,0x3d));
    QFont ftTip(".PingFang-SC",16);
    painter.setFont(ftTip);
    if(!keyEnterTip.isEmpty())
    {
        int drawWidth=width-40-painter.fontMetrics().width(keyEnterTip);
        painter.drawText(QRect(drawWidth,height-25-30,300,32),Qt::AlignLeft|Qt::AlignBottom,keyEnterTip);
        drawWidth-=(5+pixEnter.width());
        painter.drawPixmap(drawWidth,height-20-pixEnter.height(),pixEnter.width(),pixEnter.height(),pixEnter);
        drawWidth-=(5+painter.fontMetrics().width("按"));
        painter.drawText(QRect(drawWidth,height-25-30,100,32),Qt::AlignLeft|Qt::AlignBottom,"按");

    }
    if(!keyExitTip.isEmpty())
    {
        int drawWidth=40;
        painter.drawText(QRect(drawWidth,height-25-30,300,32),Qt::AlignLeft|Qt::AlignBottom,"按");
        drawWidth+=(5+painter.fontMetrics().width("按"));
        painter.drawPixmap(drawWidth,height-20-pixExit.height(),pixExit.width(),pixExit.height(),pixExit);
        drawWidth+=(5+pixExit.width());
        painter.drawText(QRect(drawWidth,height-25-30,300,32),Qt::AlignLeft|Qt::AlignBottom,keyExitTip);
    }


}

void PopupWindow::keyPressEvent(QKeyEvent *event)
{
    int key=event->key();
    if(key==Qt::Key_Enter || key==Qt::Key_Return)
    {
        if(this->phoneEdit!=nullptr)
        {
            if(this->phone4Num==this->phoneEdit->getText())
            {
                BaseDialog *dialogParent=(BaseDialog*)this->parent();
                dialogParent->processPopup(0,"");
                this->close();
            }
        }
        else
        {
            if(!this->keyEnterTip.isEmpty())
            {
                BaseDialog *dialogParent=(BaseDialog*)this->parent();
                this->returnV=0;
                QString txtRet;
                if(this->lineEdit!=nullptr)
                {
                    txtRet=this->lineEdit->getText();
                }
                dialogParent->processPopup(returnV,txtRet);
                this->close();
            }
        }

    }
    else if(key==Qt::Key_Escape) {

        if(!this->keyExitTip.isEmpty())
        {
            BaseDialog *dialogParent=(BaseDialog*)this->parent();
            this->returnV=1;
            dialogParent->processPopup(returnV,"");
            this->close();
        }
    }
    if(this->phoneEdit!=nullptr)
    {
        this->phoneEdit->processKey(event);
    }
    if(this->lineEdit!=nullptr)
    {
        this->lineEdit->processKey(event);
    }
}

void PopupWindow::timerEvent(QTimerEvent *event)
{
    this->loadingIndex++;
    if(this->loadingIndex>5)
    {
        this->loadingIndex=0;
    }
    this->update(this->loadingRect);
}

void PopupWindow::processControlEvent(UserControl *control, ControlEvent event)
{
    int width=this->width();
    QRect rect=QRect(0,50,width,30);
    if(this->phoneEdit==control)
    {
        this->update(rect);
    }
}
