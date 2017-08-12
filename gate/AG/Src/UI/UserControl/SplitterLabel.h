#ifndef SPLITTERLABEL_H
#define SPLITTERLABEL_H


#include "UserControl.h"

#include <QWidget>
#include <QRect>

class SplitterLabel:public UserControl
{
private:
    QString caption;
protected:
    void draw(QPainter &painter);
    QRect rect();
public:
    SplitterLabel(const QWidget *widget,const QRect rect,const QString caption);
};

#endif // SPLITTERLABEL_H
