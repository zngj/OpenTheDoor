#ifndef LISTVIEW_H
#define LISTVIEW_H

#include "UserControl.h"
#include <QWidget>
#include <QRect>
#include <QKeyEvent>
#include <vector>

class ListView :public UserControl
{
private:
    std::vector<std::vector<QString> *> listContent;

    int firstIndex=-1;
    int selectIndex=-1;
    const int  HeaderHeight=40;
    const int RowHeight=54;
protected:
    void processKey(QKeyEvent *event);
public:
    ListView(const QWidget *widget,const QRect& rect);
    void draw(QPainter &painter);
    void setFocus(bool focus);
    QRect rect();
    void addItem(std::vector<QString> *itm);
    void clear();
    void update();
    int getSelectIndex();
};

#endif // LISTVIEW_H
