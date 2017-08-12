#include "CheckDialog.h"

CheckDialog::CheckDialog(QWidget *parent):BaseDialog(parent)
{
    this->setTimeout(3);
    this->setTitle("世界之窗站");
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
    painter.drawText(QRect(0,51+52+12-25,42+200+42,50),Qt::AlignCenter,"检验通过");


    QRect rectPic(width-350,51,350,335);
    QPixmap pixLogo(":/Resource/Image/Pic/uparrow.jpg");
    painter.drawPixmap(rectPic,pixLogo);
}
