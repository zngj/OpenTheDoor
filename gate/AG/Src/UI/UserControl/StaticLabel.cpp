#include "StaticLabel.h"
#include <QWidget>

StaticLabel::StaticLabel(const QWidget *widget,const QRect& rect):UserControl(widget,rect)
{
}

void StaticLabel::draw(QPainter &painter)
{
    int left=this->rectV.left();
    int drawWidth=left;
    int top=this->rectV.top();
    int height=this->rectV.height();

    for(int i=0;i<(int)listString.size();i++)
    {
        int size=listSize[i];
        QString txt=listString[i];
        QColor color=listColor[i];

        if(size>0)
        {
            QFont font(".PingFang-SC",size);
            painter.setFont(font);
            painter.setPen(color);
            int txtWidth=painter.fontMetrics().width(txt);

            painter.drawText(QRect(drawWidth,top,txtWidth+5,height),Qt::AlignLeft|Qt::AlignVCenter,txt);

            drawWidth+=txtWidth;
        }
        else
        {
            QString strPath(":/Resource/Image/Icon/");
            strPath.append(txt);
            QPixmap pix(strPath);

            painter.drawPixmap(drawWidth+5,top+height/2-pix.height()/2,pix.width(),pix.height(),pix);
            drawWidth+=pix.width()+5+5;
        }

    }
}

QRect StaticLabel::rect()
{
    return this->rectV;
}

void StaticLabel::AddSegment(QString txt, int size, QColor color)
{
    this->listString.push_back(txt);
    this->listSize.push_back(size);
    this->listColor.push_back(color);
}

void StaticLabel::UpdateSegment(int index, QString txt)
{
    this->listString[index]=txt;
    this->widget->update(this->rectV);
}


