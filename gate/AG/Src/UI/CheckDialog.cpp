#include "CheckDialog.h"
#include "Business/ScannerCheck.h"

CheckDialog::CheckDialog(QWidget *parent):BaseDialog(parent)
{
    this->setTimeout(3);
    this->setTitle("世界之窗站");
    ScannerCheck *checker=ScannerCheck::getInstance();
    if(checker->getLastCheckResult()==0)
    {
        this->errMsg="检验通过";
        map.load(":/Resource/Image/Pic/uparrow.jpg");
    }
    else
    {
        this->errMsg=QString::fromStdString(checker->getLastErrMsg());
        map.load(":/Resource/Image/Pic/error.jpg");
    }
}

CheckDialog::~CheckDialog()
{

}

void CheckDialog::customDraw(QPainter &painter)
{
    int width=this->width();

    int bottom=this->rect().bottom();

    QFont fontTip(".PingFang-SC",18);
    painter.setPen(Qt::green);
    painter.setFont(fontTip);
    painter.drawText(QRect(0,51+52+12-25,42+200+42,50),Qt::AlignCenter,this->errMsg);


    QRect rectPic(width-350,51,350,335);
    painter.drawPixmap(rectPic,map);
}
