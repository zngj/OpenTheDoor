#ifndef FRAMEEDIT_H
#define FRAMEEDIT_H

#include "UserControl.h"
#include <QWidget>
#include <QRect>
#include <QPainter>
#include <QKeyEvent>
#include <QString>

class FrameEdit:public UserControl
{
private:
    int maxChars=6;
    QString text;
    bool focused=false;
    int timerID=-1;
    bool caretOn=false;
    QRect rectCaret;
protected:
    void timerEvent(QTimerEvent *event);
public:
    FrameEdit(QWidget *widget,const QRect& rect);
    ~FrameEdit();
public:
    void draw(QPainter &painter);
    void processKey(QKeyEvent *event);
    QRect rect();
    void setMaxChars(int chars);
    void setFocus(bool focus);
    QString getText();
    void clearText();

};

#endif // FRAMEEDIT_H
