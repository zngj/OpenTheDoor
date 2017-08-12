#ifndef STATICLABEL_H
#define STATICLABEL_H

#include "UserControl.h"
#include <vector>
#include <QString>
#include <QColor>
#include <QRect>

class StaticLabel:public UserControl
{
private:
    std::vector<QString> listString;
    std::vector<int> listSize;
    std::vector<QColor> listColor;
public:
    StaticLabel(const QWidget *widget,const QRect& rect);
    void draw(QPainter &painter);
    QRect rect();
    void AddSegment(QString txt,int size,QColor color);
    void UpdateSegment(int index,QString txt);
};

#endif // STATICLABEL_H
