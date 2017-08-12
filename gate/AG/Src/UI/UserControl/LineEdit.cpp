#include "LineEdit.h"
#include <QLabel>
LineEdit::LineEdit(QWidget* widget,const QRect& rect,const QString & tip):UserControl(widget,rect)
{
    this->tip=tip;
    this->focus=false;
    this->indicatorAllow=false;
    this->caretOn=false;
    timerID=this->startTimer(500);
}

LineEdit::~LineEdit()
{
    this->killTimer(timerID);
}

QRect LineEdit::rect()
{
    if(this->indicatorAllow)
    {
        return QRect(this->rectV.left()-25,this->rectV.top(),this->rectV.width()+25,this->rectV.height());
    }
    return this->rectV;

}
void LineEdit::draw(QPainter& painter)
{
    int left=this->rectV.left();
    int top=this->rectV.top();
    int width=this->rectV.width();
    int height=this->rectV.height();
    QFont fontTip(".PingFang-SC",15);
    QFont fontTxt(".PingFang-SC",height>70?35:22);
    int leftMargin=height>70?20:10;

    if(indicatorAllow && focus)
    {
        QPixmap pix1(":/Resource/Image/Icon/o-1.png"); //10*10
        QRect rectLT(left,top,10,10);
        painter.drawPixmap(rectLT,pix1);

        QPixmap pix2(":/Resource/Image/Icon/o-2.png");
        QRect rectRT(width+left-10,top,10,10);
        painter.drawPixmap(rectRT,pix2);

        QPixmap pix3(":/Resource/Image/Icon/o-3.png");
        QRect rectLB(left,top+height-10,10,10);
        painter.drawPixmap(rectLB,pix3);

        QPixmap pix4(":/Resource/Image/Icon/o-4.png");
        QRect rectRB(left+width-10,top+height-10,10,10);
        painter.drawPixmap(rectRB,pix4);

        QRect rectLine(left+5,top,width-10,4);
        QPixmap pixLine(":/Resource/Image/Icon/o-5.png");
        painter.drawPixmap(rectLine,pixLine);

        rectLine=QRect(left+5,top+height-4,width-10,4);
        painter.drawPixmap(rectLine,pixLine);
        rectLine=QRect(left,top+5,4,height-10);
        painter.drawPixmap(rectLine,pixLine);
        rectLine=QRect(left+width-4,top+5,4,height-10);
        painter.drawPixmap(rectLine,pixLine);

        //draw arrow
        QPixmap pixArrow(":/Resource/Image/Icon/arrow_right_S.png");

        painter.drawPixmap(this->rectV.left()-25,this->rectV.top()+this->rectV.height()/2-15,19,29,pixArrow);
    }
    else
    {
        QPixmap pix1(":/Resource/Image/Icon/g-1.png"); //10*10
        QRect rectLT(left,top,10,10);
        painter.drawPixmap(rectLT,pix1);

        QPixmap pix2(":/Resource/Image/Icon/g-2.png");
        QRect rectRT(width+left-10,top,10,10);
        painter.drawPixmap(rectRT,pix2);

        QPixmap pix3(":/Resource/Image/Icon/g-3.png");
        QRect rectLB(left,top+height-10,10,10);
        painter.drawPixmap(rectLB,pix3);

        QPixmap pix4(":/Resource/Image/Icon/g-4.png");
        QRect rectRB(left+width-10,top+height-10,10,10);
        painter.drawPixmap(rectRB,pix4);

        QRect rectLine(left+5,top,width-10,2);
        QPixmap pixLine(":/Resource/Image/Icon/g-5.png");
        painter.drawPixmap(rectLine,pixLine);

        rectLine=QRect(left+5,top+height-2,width-10,2);
        painter.drawPixmap(rectLine,pixLine);
        rectLine=QRect(left,top+5,2,height-10);
        painter.drawPixmap(rectLine,pixLine);
        rectLine=QRect(left+width-2,top+5,2,height-10);
        painter.drawPixmap(rectLine,pixLine);
    }
    QString drawText=text;
    if(text.isEmpty())
    {
        if(tip.isEmpty()==false)
        {
            drawText=tip;
            painter.setFont(fontTip);
            painter.setPen(QColor::fromRgb(0x99,0x99,0x99));
            QRect rectTxt(this->rectV.left()+leftMargin,this->rectV.top(),this->rectV.width()-12,this->rectV.height());
            painter.drawText(rectTxt,Qt::AlignLeft|Qt::AlignVCenter,this->tip);
        }

    }
    else
    {

        painter.setFont(fontTxt);
        painter.setPen(Qt::black);
        QRect rectTxt(this->rectV.left()+leftMargin,this->rectV.top(),this->rectV.width()-leftMargin,this->rectV.height());

        if(this->pswdInput)
        {
            drawText=textPasswd;
        }
        else if(this->phoneNumMode)
        {
            drawText=textPhone;
            painter.drawText(rectTxt,Qt::AlignLeft|Qt::AlignVCenter,textPhone);
        }
        if(this->alignCenter)
        {
            painter.drawText(this->rectV,Qt::AlignCenter,drawText);
        }
        else
        {
            painter.drawText(rectTxt,Qt::AlignLeft|Qt::AlignVCenter,drawText);
        }


    }
    //drawCaret
    int txtWidth=0;
    if(text.isEmpty()==false)
    {
        txtWidth=painter.fontMetrics().width(drawText);
    }
    if(this->alignCenter==false)
    {
        this->rectCaret=QRect(this->rectV.left()+leftMargin+2+txtWidth-1,this->rectV.top()+7,2,this->rectV.height()-15);
        if(this->focus && this->caretOn)
        {

            painter.setPen(QPen(QBrush(QColor::fromRgb(0xff,0x6b,0x00)),2) );
            painter.drawLine(this->rectV.left()+leftMargin+2+txtWidth,this->rectV.top()+8,this->rectV.left()+leftMargin+2+txtWidth,this->rectV.bottom()-8);
        }
    }

}

void LineEdit::setFocus(bool focus)
{
    this->focus=focus;
}

void LineEdit::allowIndicator(bool allow)
{
    this->indicatorAllow=allow;
}

void LineEdit::setMaxChars(int max)
{
    this->maxChars=max;
}

void LineEdit::setPasswdInput()
{
    this->pswdInput=true;
}

void LineEdit::setOnlyDigit()
{
    this->onlyDigit=true;
}

void LineEdit::setPhoneNumMode()
{
    this->onlyDigit=true;
    this->maxChars=11;
    this->phoneNumMode=true;
}

void LineEdit::setAllowSlash()
{
    this->allowSlash=true;
}

QString LineEdit::getText()
{
    return this->text;
}

QString LineEdit::getDisplayText()
{
    if(this->phoneNumMode)
    {
        return this->textPhone;
    }
    else if(this->pswdInput)
    {
        return this->textPasswd;
    }
    return text;
}

void LineEdit::setText(QString txt)
{
    this->text=txt;
    if(this->pswdInput)
    {

    }
    else
    {

    }
    widget->update(this->rect());
}

void LineEdit::setAlignCenter(bool align)
{
    this->alignCenter=align;
}

void LineEdit::processKey(QKeyEvent *event)
{
    int key=event->key();
    if(key>=Qt::Key_0 && key<=Qt::Key_9)
    {
        if(this->maxChars>0 && text.length()>=this->maxChars) return;
        text.append('0'+key-Qt::Key_0);
        textPasswd.append("●");
        textPhone.append('0'+key-Qt::Key_0);
        if(text.length()==3 || text.length()==7)
        {
            textPhone.append(" ");
        }
        widget->update(this->rect());
        triggerEvent(this,TextChanged);
    }
    else if(key>=Qt::Key_A && key<=Qt::Key_Z)
    {
        if(this->onlyDigit) return;
        if(this->maxChars>0 && text.length()>=this->maxChars) return;
        if(key==Qt::Key_X && this->allowSlash)
        {
            text.append('\\');
            textPasswd.append("●");
            widget->update(this->rect());
        }
        else
        {
            if(event->modifiers()&Qt::ShiftModifier)
            {
                text.append('A'+key-Qt::Key_A);
                textPasswd.append("●");
                widget->update(this->rect());
            }
            else
            {
                text.append('a'+key-Qt::Key_A);
                textPasswd.append("●");
                widget->update(this->rect());
            }
        }

        triggerEvent(this,TextChanged);

    }
    else if(key==Qt::Key_Period)
    {
        if(this->onlyDigit) return;
        if(this->maxChars>0 && text.length()>=this->maxChars) return;
        if(this->allowSlash)
        {
            text.append('\\');
            textPasswd.append("●");
            widget->update(this->rect());
        }
    }
    else if(key==Qt::Key_Backspace) {
        if(!text.isEmpty())
        {

            text.remove(text.length()-1,1);
            textPhone.remove(textPhone.length()-1,1);
            if(text.length()==2 || text.length()==6)
            {
                textPhone.remove(textPhone.length()-1,1);
            }
            textPasswd.remove(textPasswd.length()-1,1);
            widget->update(this->rect());
            triggerEvent(this,TextChanged);
        }

    }
    else if(key==Qt::Key_Return || key==Qt::Key_Enter)
    {
        triggerEvent(this,Confirm);
    }

}

void LineEdit::timerEvent(QTimerEvent *event)
{
    Q_UNUSED(event);
    this->caretOn=!this->caretOn;
    if(this->focus)
    {
        this->widget->update(this->rectCaret);
    }
}


