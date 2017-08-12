#include "ListView.h"

ListView::ListView(const QWidget *widget,const QRect& rect):UserControl(widget,rect)
{
}

void ListView::draw(QPainter &painter)
{
    //drawTableHeader
    int left=rectV.left();
    int top=rectV.top();
    int width=rectV.width();
    int height=rectV.height();

    int maxRows=(height-HeaderHeight)/RowHeight;
    QFont fontNormal(".PingFang-SC",16);
    QFont fontBold(".PingFang-SC",16);
    //fontBold.setBold(true);
    painter.setPen(QColor::fromRgb(0x52,0x55,0x5d));
    painter.drawRect(QRect(left,top,width,HeaderHeight));

    painter.setPen(QColor::fromRgb(0x55,0x55,0x55));
    painter.drawText(QRect(left+16,top,200,HeaderHeight),Qt::AlignLeft|Qt::AlignVCenter,"序号");
    painter.drawText(QRect(left+75,top,200,HeaderHeight),Qt::AlignLeft|Qt::AlignVCenter,"运单号");
    painter.drawText(QRect(left+316,top,200,HeaderHeight),Qt::AlignLeft|Qt::AlignVCenter,"状态");
    painter.drawText(QRect(left+384,top,200,HeaderHeight),Qt::AlignLeft|Qt::AlignVCenter,"收件人");
    painter.drawText(QRect(left+554,top,200,HeaderHeight),Qt::AlignLeft|Qt::AlignVCenter,"投递时间");

    for(int i=0;i<maxRows;i++)
    {
        if(firstIndex+i>=(int)listContent.size()) break;
        if(firstIndex<0) break;
        painter.drawLine(left,top+HeaderHeight+i*RowHeight,left,top+HeaderHeight+(i+1)*RowHeight);
        painter.drawLine(left+width,top+HeaderHeight+i*RowHeight,left+width,top+HeaderHeight+(i+1)*RowHeight);
        painter.drawLine(left,top+HeaderHeight+(i+1)*RowHeight,left+width,top+HeaderHeight+(i+1)*RowHeight);
    }
    for(int i=0;i<maxRows;i++)
    {
        if(firstIndex+i>=(int)listContent.size()) break;
        if(firstIndex<0) break;
        std::vector<QString> *itm=listContent[firstIndex+i];

        if(firstIndex+i==selectIndex)
        {
            QRect rectBg(left-1,top+HeaderHeight+i*RowHeight-1,width+2,RowHeight+2);
            painter.fillRect(rectBg,QColor::fromRgb(0xff,0x6b,0x00));
            painter.setPen(QColor::fromRgb(0xff,0xff,0xff));
            QPixmap pix(":/Image/Icon/arrow_right_S.png");

            painter.drawPixmap(left-25,top+HeaderHeight+i*RowHeight+RowHeight/2-15,19,29,pix);
            painter.setFont(fontBold);

        }
        else
        {
            painter.setFont(fontNormal);
            painter.setPen(QColor::fromRgb(0x3d,0x3d,0x3d));
        }
        if(itm->size()>0)
        {
            QString id=itm->at(0);
            QRect rectId(left+16,top+HeaderHeight+i*RowHeight,100,RowHeight);
            painter.drawText(rectId,Qt::AlignLeft|Qt::AlignVCenter,id);
        }
        if(itm->size()>1)
        {
            QString pkgId=itm->at(1);
            QRect rectPkgId(left+75,top+HeaderHeight+i*RowHeight,250,RowHeight);
            if(pkgId.length()>17)
            {
                QString drawID=pkgId.mid(0,7);
                drawID.append("....");
                drawID.append(pkgId.mid(pkgId.length()-8));
                painter.drawText(rectPkgId,Qt::AlignLeft|Qt::AlignVCenter,drawID);
            }
            else
            {
                painter.drawText(rectPkgId,Qt::AlignLeft|Qt::AlignVCenter,pkgId);
            }


        }
        if(itm->size()>2)
        {
            QString status=itm->at(2);
            QRect rectSts(left+316,top+HeaderHeight+i*RowHeight,212,RowHeight);
            painter.drawText(rectSts,Qt::AlignLeft|Qt::AlignVCenter,status);
        }
        if(itm->size()>3)
        {
            QString persion=itm->at(3);
            QRect rectPer(left+384,top+HeaderHeight+i*RowHeight,212,RowHeight);
            painter.drawText(rectPer,Qt::AlignLeft|Qt::AlignVCenter,persion);
        }
        if(itm->size()>4)
        {
            QString tm=itm->at(4);
            QRect rectTm(left+554,top+HeaderHeight+i*RowHeight,212,RowHeight);
            painter.drawText(rectTm,Qt::AlignLeft|Qt::AlignVCenter,tm);
        }

    }
}

void ListView::setFocus(bool focus)
{
    Q_UNUSED(focus);
}

QRect ListView::rect()
{
    return this->rectV;
}

void ListView::processKey(QKeyEvent *event)
{
    int left=rectV.left();
    int top=rectV.top();
    int width=rectV.width();
    int height=rectV.height();
    int maxRows=(height-HeaderHeight)/RowHeight;

    int key=event->key();
    if(key==Qt::Key_PageDown || key==Qt::Key_Down)
    {
        if(selectIndex<(int)listContent.size()-1)
        {
            if(selectIndex-firstIndex<maxRows-1)
            {
                QRect rectOld(left-25,top+HeaderHeight+(selectIndex-firstIndex)*RowHeight-1,width+26,RowHeight+2);
                widget->update(rectOld);
                selectIndex++;
                QRect rectNew(left-25,top+HeaderHeight+(selectIndex-firstIndex)*RowHeight-1,width+26,RowHeight+2);
                widget->update(rectNew);
            }
            else
            {
                firstIndex++;
                selectIndex++;
                QRect rectUpdate(left-25,top+HeaderHeight,width+26,height-HeaderHeight);
                widget->update(rectUpdate);
            }

        }
    }
    else if(key==Qt::Key_PageUp || key==Qt::Key_Up) {

        if(selectIndex>0)
        {
            if(selectIndex==firstIndex)
            {
                selectIndex--;
                firstIndex--;
                QRect rectUpdate(left-25,top+HeaderHeight,width+26,height-HeaderHeight);
                widget->update(rectUpdate);
            }
            else
            {
                QRect rectOld(left-25,top+HeaderHeight+(selectIndex-firstIndex)*RowHeight-1,width+26,RowHeight+2);
                widget->update(rectOld);
                selectIndex--;
                QRect rectNew(left-25,top+HeaderHeight+(selectIndex-firstIndex)*RowHeight-1,width+26,RowHeight+2);
                widget->update(rectNew);
            }

        }
    }
    else if(key==Qt::Key_Return || key==Qt::Key_Enter)
    {
        if(selectIndex>=0)
        {
            this->triggerEvent(this,Confirm);
        }
    }
}

void ListView::addItem(std::vector<QString> * itm)
{
    listContent.push_back(itm);
    if(selectIndex<0) selectIndex=0;
    if(firstIndex<0) firstIndex=0;
}

void ListView::clear()
{
    listContent.clear();
    firstIndex=-1;
    selectIndex=-1;
}

void ListView::update()
{
    this->widget->update(QRect(this->rectV.left()-30,this->rect().top()-2,this->rectV.width()+32,this->rectV.height()+2));
}

int ListView::getSelectIndex()
{
    return this->selectIndex;
}

