#include "SplitterLabel.h"



SplitterLabel::SplitterLabel(const QWidget *widget,const QRect rect,const QString caption):UserControl(widget,rect)
{
    this->caption=caption;
}

void SplitterLabel::draw(QPainter &painter)
{
    int left=this->rectV.left();
    int top=this->rectV.top();
    int width=this->rectV.width();
    int height=this->rectV.height();

    painter.setFont(QFont(".PingFang-SC",15));
    QPen pen=QPen(QBrush(QColor::fromRgb(0x99,0x99,0x99)),2);
    painter.setPen(pen);
    if(this->caption.isEmpty())
    {
        painter.drawLine(left+width/2,top,left+width/2,top+height);
    }
    else
    {
        int txtNum=this->caption.length();
        int ftHeight=painter.fontMetrics().height();
        int txtStart=top+height/2-txtNum*ftHeight/2;
        painter.drawLine(left+width/2,top,left+width/2,txtStart-4);
        for(int i=0;i<txtNum;i++)
        {
            painter.drawText(QRect(left-20,txtStart,width+40,ftHeight),Qt::AlignCenter,caption.mid(i,1));
            txtStart+=ftHeight;
        }
        painter.drawLine(left+width/2,txtStart+4,left+width/2,top+height);
    }
}

QRect SplitterLabel::rect()
{
    return this->rectV;
}
