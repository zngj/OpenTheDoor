#ifndef QRDISPLAY_H
#define QRDISPLAY_H

#include "UserControl.h"
#include <QWidget>
#include <QRect>

#include <list>



class QRDisplay:public UserControl
{
private:
    QString logo;
    QString content;
public:
    QRDisplay(QWidget *widget,const QRect rect);
    ~QRDisplay();
    void draw(QPainter &painter);
    void setFocus(bool focus);
    QRect rect();
    void setContent(const QString &content,const QString &logo);
    void qrChanged();
};

#endif // QRDISPLAY_H
