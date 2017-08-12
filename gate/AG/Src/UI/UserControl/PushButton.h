#ifndef PUSHBUTTON_H
#define PUSHBUTTON_H

#include "UserControl.h"
#include <QRect>
#include <QString>
#include <QPainter>
#include <QWidget>

class PushButton :public UserControl
{
private:
    QString captionV;
    bool focus;
    int buttonStyle=0;
    QPixmap mapOn;
    QPixmap mapOff;
    QPixmap mapArrow;
    QPixmap mapRight;
    QPixmap mapCornerOn[4];
    QPixmap mapCornerOff[4];

    QPixmap mapBoxOn[4];
    QPixmap mapBoxOff[4];
    int boxType=-1;
    QString oldPrice;
    QString newPrice;
    int boxNum=0;
public:
    PushButton(QWidget* widget,const QRect& rect,const QString& caption);
    void setFocus(bool focus);
    bool focused();
    QRect rect();
    QString caption();
    void setImageName(QString imgName);
    void setBoxType(int type,int boxNum,QString old,QString now);
    void draw(QPainter& painter);
protected:
    void processKey(QKeyEvent *event);
};

#endif // PUSHBUTTON_H
