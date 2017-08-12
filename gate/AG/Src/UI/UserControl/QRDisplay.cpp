#include "QRDisplay.h"


#include <qrencode.h>
#include <QDebug>

QRDisplay::QRDisplay(QWidget *widget,const QRect rect):UserControl(widget,rect)
{
    this->widget=widget;
    this->rectV=rect;
}

QRDisplay::~QRDisplay()
{

}

void QRDisplay::draw(QPainter &painter)
{

    int left=this->rectV.left();
    int top=this->rectV.top();
    int width=this->rectV.width();
    int height=this->rectV.height();

    QPixmap pix1(":/Resource/Image/Icon/QR_line_1.png"); //12*12
    QRect rectLT(left,top,12,12);
    painter.drawPixmap(rectLT,pix1);

    QPixmap pix2(":/Resource/Image/Icon/QR_line_2.png");
    QRect rectRT(width+left-12,top,12,12);
    painter.drawPixmap(rectRT,pix2);

    QPixmap pix3(":/Resource/Image/Icon/QR_line_3.png");
    QRect rectLB(left,top+height-12,12,12);
    painter.drawPixmap(rectLB,pix3);

    QPixmap pix4(":/Resource/Image/Icon/QR_line_4.png");
    QRect rectRB(left+width-12,top+height-12,12,12);
    painter.drawPixmap(rectRB,pix4);

    QRect rectLine(left+10,top,width-20,1);
    QPixmap pixLine(":/Resource/Image/Icon/QR_line.png");
    painter.drawPixmap(rectLine,pixLine);
    rectLine=QRect(left+10,top+height-1,width-20,1);
    painter.drawPixmap(rectLine,pixLine);
    rectLine=QRect(left,top+10,1,height-20);
    painter.drawPixmap(rectLine,pixLine);
    rectLine=QRect(left+width-1,top+10,1,height-20);
    painter.drawPixmap(rectLine,pixLine);

    std::string qrCode = "http://tool.chinaz.com/Tools/unixtime.aspx";

    left=this->rectV.left()+13;
    top=this->rectV.top()+13;
    width=this->rectV.width()-26;
    height=this->rectV.height()-26;

    QRcode *qr = QRcode_encodeString(qrCode.c_str(), 1, QR_ECLEVEL_L, QR_MODE_8, 1);
    if(0 != qr){
        QColor fg((unsigned int)0x000000);
        QColor bg(0xffffff);
        painter.setBrush(bg);
        painter.setPen(Qt::NoPen);
        painter.drawRect(left,top,width,height);
        painter.setBrush(fg);
        const int s=qr->width>0?qr->width:1;
        const double w=width;
        const double h=height;
        const double aspect=w/h;
        const double scale=((aspect>1.0)?h:w)/s;
        for(int y=0;y<s;y++){
            const int yy=y*s;
            for(int x=0;x<s;x++){
                const int xx=yy+x;
                const unsigned char b=qr->data[xx];
                if(b &0x01){
                    const double rx1=x*scale, ry1=y*scale;
                    QRectF r(rx1+left, ry1+top, scale, scale);
                    painter.drawRects(&r,1);
                }
            }
        }
        QRcode_free(qr);
    }
    else{
        QColor error("red");
        painter.setBrush(error);
        painter.drawRect(left,top,width,height);
        qDebug()<<"QR FAIL: "<<strerror(errno);
    }
    qr=0;
    if(this->logo.isEmpty()==false)
    {
        QString logoPath(":/Resource/Image/Icon/");
        logoPath.append(this->logo);
        QPixmap pixLogo;
        pixLogo.load(logoPath);

        painter.drawPixmap(left+width/2-pixLogo.width()/2,top+height/2-pixLogo.height()/2,pixLogo.width(),pixLogo.height(),pixLogo);

    }
}

void QRDisplay::setFocus(bool focus)
{
    Q_UNUSED(focus);
}
QRect QRDisplay::rect()
{
    return this->rectV;
}

void QRDisplay::setContent(const QString &content, const QString &logo)
{
    this->content=content;
    this->logo=logo;
    this->widget->update(this->rectV);
}

void QRDisplay::qrChanged()
{
    if(this->content.length()==0)
    {
        this->widget->update(rectV);
    }
}

