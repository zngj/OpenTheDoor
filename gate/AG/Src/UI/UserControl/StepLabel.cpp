#include "StepLabel.h"

StepLabel::StepLabel(const QWidget *widget,const QRect &rect):UserControl(widget,rect)
{

}

void StepLabel::draw(QPainter &painter)
{

    int top=this->rectV.top();
    int height=this->rectV.height();

    int drawWidth=this->rectV.left();

    QFont font(".PingFang-SC",15);
    painter.setFont(font);
    painter.setPen(Qt::black);
    for(int i=0;i<(int)listSteps.size();i++)
    {
        QString stepName=listSteps[i];
        QString txtMapName(":/Resource/Image/Icon/");
        txtMapName.append(QString::number(i+1));
        txtMapName.append("_");
        if(this->currentStep>=i)
        {
            txtMapName.append("on.png");
        }
        else
        {
            txtMapName.append("off.png");
        }
        QPixmap mapNum(txtMapName);

        painter.setPen(Qt::black);
        painter.drawPixmap(drawWidth,top+height/2-12,24,24,mapNum);
        drawWidth+=24+5;

        painter.drawText(QRect(drawWidth,top,100,height),Qt::AlignLeft|Qt::AlignVCenter,stepName);
        drawWidth+=painter.fontMetrics().width(stepName);
        drawWidth+=5;
        if(i<(int)this->listSteps.size()-1)
        {
            if(i< this->currentStep)
            {
                painter.setPen(QPen(QBrush(QColor::fromRgb(0x1b,0xba,0x56)),2));
            }
            else
            {
                painter.setPen(QPen(QBrush(QColor::fromRgb(0x55,0x55,0x55)),2));
            }

            painter.drawLine(drawWidth,top+height/2,drawWidth+50,top+height/2);

            drawWidth+=50+5;

        }


    }
}

QRect StepLabel::rect()
{
    return this->rectV;
}

void StepLabel::setStepName(QString name)
{
    listSteps.push_back(name);
}

void StepLabel::setCurrentStep(int step)
{
    this->currentStep=step;
}
