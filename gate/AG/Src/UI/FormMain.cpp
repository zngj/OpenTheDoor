#include "FormMain.h"
#include "UI/UIListener.h"
#include "CheckDialog.h"

FormMain::FormMain(QWidget*parent):BaseDialog(parent)
{
    this->setTitle("世界之窗站");
    qrDisplay=this->createQRDisplay(QRect(42,95+52,200,200));
}

FormMain::~FormMain()
{

}


void FormMain::customDraw(QPainter &painter)
{
    int width=this->width();

    int bottom=this->rect().bottom();

    QFont fontTip(".PingFang-SC",18);
    painter.setPen(QColor::fromRgb(0x55,0x55,0x55));
    painter.setFont(fontTip);
    painter.drawText(QRect(0,51+52+12-25,42+200+42,50),Qt::AlignCenter,"扫码开闸门");


    QRect rectPic(width-350,51,350,335);
    QPixmap pixLogo(":/Resource/Image/Pic/logo.jpg");
    painter.drawPixmap(rectPic,pixLogo);

    painter.fillRect(QRect(0,bottom-94,width,95),QColor::fromRgb(0xff,0x6b,0x00));

    painter.setPen(Qt::white);
    QFont fontBmMsg(".PingFang-SC",22);
    painter.setFont(fontBmMsg);
    painter.drawText(QRect(267,bottom-95,300,95),Qt::AlignLeft|Qt::AlignVCenter,"欢迎使用云闸机系统");

}

void FormMain::keyPressEvent(QKeyEvent *event)
{
    BaseDialog::keyPressEvent(event);
    int key=event->key();

}

void FormMain::processControlEvent(UserControl *control, ControlEvent event)
{
    Q_UNUSED(event);

}

void FormMain::customEvent(QEvent *event)
{
    BaseDialog::customEvent(event);
    int type=event->type();
    if(type==WM_CHECK_OK)
    {
        CheckDialog * dialog=new CheckDialog(this);
        dialog->exec();
        delete(dialog);
    }

}
