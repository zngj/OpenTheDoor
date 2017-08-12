#include "FrameEdit.h"



void FrameEdit::timerEvent(QTimerEvent *event)
{
    Q_UNUSED(event);
    this->caretOn=!this->caretOn;
    if(this->focused)
    {
        this->widget->update(this->rectCaret);
    }
}

FrameEdit::FrameEdit(QWidget *widget, const QRect &rect):UserControl(widget,rect)
{
    this->timerID=this->startTimer(500);
}

FrameEdit::~FrameEdit()
{
    if(this->timerID>=0)
    this->killTimer(this->timerID);
}

void FrameEdit::setMaxChars(int chars)
{
    this->maxChars=chars;
}


void FrameEdit::draw(QPainter &painter)
{
    int left=this->rectV.left();
    int top=this->rectV.top();
    int width=this->rectV.width();
    int height=this->rectV.height();

    QPixmap pix1(":/Resource/Image/Icon/input_L1.png"); //12*12
    QRect rectLT(left,top,12,12);
    painter.drawPixmap(rectLT,pix1);

    QPixmap pix2(":/Resource/Image/Icon/input_L2.png");
    QRect rectRT(width+left-12,top,12,12);
    painter.drawPixmap(rectRT,pix2);

    QPixmap pix3(":/Resource/Image/Icon/input_L3.png");
    QRect rectLB(left,top+height-12,12,12);
    painter.drawPixmap(rectLB,pix3);

    QPixmap pix4(":/Resource/Image/Icon/input_L4.png");
    QRect rectRB(left+width-12,top+height-12,12,12);
    painter.drawPixmap(rectRB,pix4);


    QPixmap pixLine(":/Resource/Image/Icon/input_L.png");
    QRect rectLine(left+10,top,width-20,2);
    painter.drawPixmap(rectLine,pixLine);
    rectLine=QRect(left+10,top+height-2,width-20,2);
    painter.drawPixmap(rectLine,pixLine);
    rectLine=QRect(left,top+8,2,height-16);
    painter.drawPixmap(rectLine,pixLine);
    rectLine=QRect(left+width-2,top+8,2,height-16);
    painter.drawPixmap(rectLine,pixLine);

    int gridWidth=width/maxChars;
    for(int i=1;i<this->maxChars;i++)
    {
        painter.drawPixmap(left+gridWidth*i,top+13,1,height-26,pixLine);
    }

    //drawText
    if(text.length()>0)
    {
        QFont fontTxt(".PingFang-SC",45);
        painter.setFont(fontTxt);
        painter.setPen(QColor::fromRgb(0xff,0x6b,0x00));

        for(int i=0;i<text.length();i++)
        {
            painter.drawText(QRect(left+gridWidth*i,top,gridWidth,height),Qt::AlignCenter,text.mid(i,1));
        }
    }
    int num=this->text.length();
    this->rectCaret=QRect(left+gridWidth*num+15,top+12,2,height-24);
    if(this->caretOn)
    {

        if(num<this->maxChars)
        {
            painter.fillRect(this->rectCaret,QColor::fromRgb(0xff,0x6b,0x00));
        }

    }

}

void FrameEdit::processKey(QKeyEvent *event)
{
    int key=event->key();
    if(key>=Qt::Key_0 && key<=Qt::Key_9)
    {
        if(text.length()<this->maxChars)
        {
            text.append('0'+key-Qt::Key_0);
            widget->update(rectV);
            if(text.length()==this->maxChars)
            {
                triggerEvent(this,Confirm);
            }
        }
    }
    else if(key==Qt::Key_X || key==Qt::Key_Period)
    {
        if(text.length()<this->maxChars)
        {
            text.append("#");
            widget->update(rectV);
            if(text.length()==this->maxChars)
            {
                triggerEvent(this,Confirm);
            }
        }
    }
    else if(key==Qt::Key_Escape || key==Qt::Key_Backspace)
    {
        if(text.length()>0)
        {
            text.remove(text.length()-1,1);
            widget->update(rectV);
            if(text.length()==this->maxChars-1)
            {
                triggerEvent(this,TextChanged);
            }
        }
    }
}
void FrameEdit::setFocus(bool focus)
{
    this->focused=focus;
}

QString FrameEdit::getText()
{
    return text;
}

void FrameEdit::clearText()
{
    this->text="";
    widget->update(rectV);
}
QRect FrameEdit::rect()
{
    return this->rectV;
}
