#include "PushButton.h"

PushButton::PushButton(QWidget * widget,const QRect& rect,const QString& caption):UserControl(widget,rect)
{
    this->widget=(QWidget*)widget;
    this->rectV=rect;
    this->captionV=caption;
    this->focus=false;
    mapRight.load(":/Resource/Image/Icon/arrow_right_S.png");
    mapArrow.load(":/Resource/Image/Icon/arrow_up.png");
    mapCornerOn[0].load(":/Resource/Image/Icon/o_icon_1.png");
    mapCornerOn[1].load(":/Resource/Image/Icon/o_icon_2.png");
    mapCornerOn[2].load(":/Resource/Image/Icon/o_icon_3.png");
    mapCornerOn[3].load(":/Resource/Image/Icon/o_icon_4.png");

    mapCornerOff[0].load(":/Resource/Image/Icon/g_icon_1.png");
    mapCornerOff[1].load(":/Resource/Image/Icon/g_icon_2.png");
    mapCornerOff[2].load(":/Resource/Image/Icon/g_icon_3.png");
    mapCornerOff[3].load(":/Resource/Image/Icon/g_icon_4.png");

    mapBoxOn[0].load(":/Resource/Image/Icon/Sbox_on_icon.png");
    mapBoxOn[1].load(":/Resource/Image/Icon/Mbox_on_icon.png");
    mapBoxOn[2].load(":/Resource/Image/Icon/Lbox_on_icon.png");
    mapBoxOn[3].load(":/Resource/Image/Icon/XLbox_on_icon.png");

    mapBoxOff[0].load(":/Resource/Image/Icon/Sbox_off_icon.png");
    mapBoxOff[1].load(":/Resource/Image/Icon/Mbox_off_icon.png");
    mapBoxOff[2].load(":/Resource/Image/Icon/Lbox_off_icon.png");
    mapBoxOff[3].load(":/Resource/Image/Icon/XLbox_off_icon.png");
}

void PushButton::setFocus(bool focus)
{
    this->focus=focus;
}

bool PushButton::focused()
{
    return this->focus;
}

QRect PushButton::rect()
{
    int left=this->rectV.left();
    int top=this->rectV.top();
    int width=this->rectV.width();
    int height=this->rectV.height();
    if(this->buttonStyle==0)
    {
        QRect rect(left-30,top-2,width+30,height+4);
        return rect;
    }
    else //bigIconButton
    {
        QRect rect(left-2,top-2,width+4,height+50);
        return rect;

    }

}

QString PushButton::caption()
{
    return this->captionV;
}

void PushButton::setImageName(QString imgName)
{
    this->buttonStyle=1;
    QString strOn(":/Resource/Image/Icon/");
    strOn.append(imgName);
    strOn.append("_on_icon.png");
    this->mapOn.load(strOn);

    QString strOff(":/Resource/Image/Icon/");
    strOff.append(imgName);
    strOff.append("_off_icon.png");
    this->mapOff.load(strOff);
}

void PushButton::setBoxType(int type, int boxNum, QString old, QString now)
{
    this->buttonStyle=2;
    this->boxType=type;
    this->boxNum=boxNum;
    this->oldPrice=old;
    this->newPrice=now;
}

void PushButton::draw(QPainter& painter)
{
    QFont fontTxt(".PingFang-SC",this->captionV.length()>10?18:22);
    painter.setFont(fontTxt);

    int left=this->rectV.left();
    int right=this->rectV.right();
    int bottom=this->rectV.bottom();
    int width=this->rectV.width();
    int height=this->rectV.height();
    int top=this->rectV.top();

    if(this->buttonStyle==0)
    {

        QRect rectLeftTop(this->focus?left:left+13,this->focus?top:top+2,8,8);
        painter.drawPixmap(rectLeftTop,this->focus?mapCornerOn[0]:mapCornerOff[0]);

        QRect rectRightTop(this->focus?right-8:right-8-13,this->focus?top:top+2,8,8);
        painter.drawPixmap(rectRightTop,this->focus?mapCornerOn[1]:mapCornerOff[1]);

        QRect rectLeftBottom(this->focus?left:left+13,this->focus?bottom-8:bottom-8-2,8,8);
        painter.drawPixmap(rectLeftBottom,this->focus?mapCornerOn[2]:mapCornerOff[2]);

        QRect rectRightBottom(this->focus?right-8:right-8-13,this->focus?bottom-8:bottom-8-2,8,8);
        painter.drawPixmap(rectRightBottom,this->focus?mapCornerOn[3]:mapCornerOff[3]);

        painter.fillRect(QRect(this->focus?left:left+13,this->focus?top+7:top+2+7,8,this->focus?height-14:height-14-4),this->focus?QColor::fromRgb(0xff,0x6b,0x00):QColor::fromRgb(0x55,0x55,0x55));
        painter.fillRect(QRect(this->focus?right-8:right-8-13,this->focus?top+7:top+2+7,8,this->focus?height-14:height-14-4),this->focus?QColor::fromRgb(0xff,0x6b,0x00):QColor::fromRgb(0x55,0x55,0x55));

        painter.fillRect(QRect(this->focus?left+8:left+8+13,this->focus?top:top+2,this->focus?width-16:width-16-13*2,this->focus?height-1:height-2*2-1),this->focus?QColor::fromRgb(0xff,0x6b,0x00):QColor::fromRgb(0x55,0x55,0x55));

        if(this->focus)
        {
            painter.drawPixmap(left-19-4,top+height/2-15,19,29,mapRight);
        }

        if(this->focus==false)
        {
            QFont fontTxtS(".PingFang-SC",this->captionV.length()>10?16:18);
            painter.setFont(fontTxtS);
        }
        painter.drawText(QRect(left,top,width,height),Qt::AlignCenter,this->captionV);
    }
    else if(this->buttonStyle==1)
    {
        if(this->focus)
        {
            QPen pen(QBrush(QColor::fromRgb(0xff,0x6b,0x00)),4);
            painter.setPen(pen);
            painter.drawRect(this->rectV);
            painter.setPen(QColor::fromRgb(0xff,0x6b,0x00));

            int pixWidth=mapOn.width();
            int pixHeight=mapOn.height();

            painter.drawPixmap(left+width/2-pixWidth/2,top+20,pixWidth,pixHeight,mapOn);

            pixWidth=mapArrow.width();
            pixHeight=mapArrow.height();
            painter.drawPixmap(left+width/2-pixWidth/2,bottom+10,pixWidth,pixHeight,mapArrow);
        }
        else
        {
            QPen pen(QBrush(QColor::fromRgb(0x52,0x55,0x5d)),2);
            painter.setPen(pen);
            painter.drawRect(this->rectV);
            painter.setPen(Qt::black);

            int pixWidth=mapOff.width();
            int pixHeight=mapOff.height();

            painter.drawPixmap(left+width/2-pixWidth/2,top+20,pixWidth,pixHeight,mapOff);
        }
        QRect rectTxt(left,bottom-30-40,width,30+40);

        painter.drawText(rectTxt,Qt::AlignCenter,this->captionV);

    }
    else if(buttonStyle==2)
    {
        QString names[4]={QString("小箱("),QString("中箱("),QString("大箱("),QString("超大箱(")};
        QFont font(".PingFang-SC",(this->boxType==3 || this->boxNum>=100)?(this->rectV.width()>210?25:18):(this->rectV.width()>210?25:20));
        painter.setFont(font);
        if(this->focus)
        {
            QPen pen(QBrush(QColor::fromRgb(0xff,0x6b,0x00)),4);
            painter.setPen(pen);
            painter.drawRect(this->rectV);
            int iconTop=top+30;
            int iconHeight=mapBoxOn[boxType].height();
            int iconWidth=mapBoxOn[boxType].width();
            if(oldPrice.isEmpty() && newPrice.isEmpty())
            {
                iconTop=top+height/2-iconHeight/2;
            }
            painter.drawPixmap(left+10,iconTop,iconWidth,iconHeight,this->enabled?mapBoxOn[boxType]:mapBoxOff[boxType]);

            painter.setPen(this->enabled?Qt::black:QColor::fromRgb(0x55,0x55,0x55));

            int drawStart=left+iconWidth+15;
            painter.drawText(QRect(drawStart,iconTop,200,iconHeight),Qt::AlignLeft|Qt::AlignVCenter,names[boxType]);
            drawStart+=painter.fontMetrics().width(names[boxType]);
            painter.setPen(this->enabled?QColor::fromRgb(0xff,0x6b,0x00):QColor::fromRgb(0x55,0x55,0x55));

            painter.drawText(QRect(drawStart,iconTop,200,iconHeight),Qt::AlignLeft|Qt::AlignVCenter,QString::number(boxNum));
            drawStart+=painter.fontMetrics().width(QString::number(boxNum));
            painter.setPen(this->enabled?Qt::black:QColor::fromRgb(0x55,0x55,0x55));
            painter.drawText(QRect(drawStart,iconTop,200,iconHeight),Qt::AlignLeft|Qt::AlignVCenter,")");


            //drawPrice
            painter.setFont(QFont(".PingFang-SC",18));
            painter.setPen(QColor::fromRgb(0x55,0x55,0x55));
            if(!oldPrice.isEmpty())
            {
                painter.drawText(QRect(left+width/2-10-100,bottom-25-50,100,50),Qt::AlignRight|Qt::AlignBottom,oldPrice);
                painter.drawLine(left+width/2-10,bottom-25-15,left+width/2-10-painter.fontMetrics().width(oldPrice),bottom-25-15);
            }
            if(!newPrice.isEmpty())
            {
                painter.setPen(QColor::fromRgb(0x1b,0xba,0x56));
                painter.drawText(QRect(left+width/2+10,bottom-25-50,100,50),Qt::AlignLeft|Qt::AlignBottom,newPrice);
            }

            int pixWidth=mapArrow.width();
            int pixHeight=mapArrow.height();
            painter.drawPixmap(left+width/2-pixWidth/2,bottom+10,pixWidth,pixHeight,mapArrow);
        }
        else
        {
            int iconTop=top+30;
            int iconHeight=mapBoxOn[boxType].height();
            int iconWidth=mapBoxOn[boxType].width();
            if(oldPrice.isEmpty() && newPrice.isEmpty())
            {
                iconTop=top+height/2-iconHeight/2;
            }

            if(this->enabled)
            {
                QPen pen(QBrush(QColor::fromRgb(0x52,0x55,0x5d)),2);
                painter.setPen(pen);
                painter.drawRect(this->rectV);

                painter.drawPixmap(left+10,iconTop,iconWidth,iconHeight,mapBoxOn[boxType]);
            }
            else
            {
                QPen pen(QBrush(QColor::fromRgb(0xaa,0xaa,0xaa)),2);
                painter.setPen(pen);
                painter.drawRect(this->rectV);

                painter.drawPixmap(left+10,iconTop,iconWidth,iconHeight,mapBoxOff[boxType]);
            }



            painter.setPen(this->enabled? Qt::black:QColor::fromRgb(0x55,0x55,0x55));
            int drawStart=left+mapBoxOn[boxType].width()+15;
            painter.drawText(QRect(drawStart,iconTop,200,iconHeight),Qt::AlignLeft|Qt::AlignVCenter,names[boxType]);
            drawStart+=painter.fontMetrics().width(names[boxType]);
            painter.setPen(this->enabled?QColor::fromRgb(0xff,0x6b,0x00):QColor::fromRgb(0x55,0x55,0x55));
            painter.drawText(QRect(drawStart,iconTop,200,iconHeight),Qt::AlignLeft|Qt::AlignVCenter,QString::number(boxNum));
            drawStart+=painter.fontMetrics().width(QString::number(boxNum));
            painter.setPen(this->enabled?Qt::black:QColor::fromRgb(0x55,0x55,0x55));
            painter.drawText(QRect(drawStart,iconTop,200,iconHeight),Qt::AlignLeft|Qt::AlignVCenter,")");


            //drawPrice
            painter.setFont(QFont(".PingFang-SC",18));
            painter.setPen(QColor::fromRgb(0x55,0x55,0x55));
            painter.drawText(QRect(left+width/2-10-100,bottom-25-50,100,50),Qt::AlignRight|Qt::AlignBottom,oldPrice);
            painter.drawLine(left+width/2-10,bottom-25-15,left+width/2-10-painter.fontMetrics().width(oldPrice),bottom-25-15);
            if(this->enabled)
            {
                painter.setPen(QColor::fromRgb(0x1b,0xba,0x56));
            }
            painter.drawText(QRect(left+width/2+10,bottom-25-50,100,50),Qt::AlignLeft|Qt::AlignBottom,newPrice);

        }

    }
}

void PushButton::processKey(QKeyEvent *event)
{
    int key=event->key();
    if((key==Qt::Key_Enter || key==Qt::Key_Return)&&this->focus)
    {
        if(this->enabled)
        {
            triggerEvent(this,Confirm);
        }

    }
}
