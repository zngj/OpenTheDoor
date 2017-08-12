#ifndef LINEEDIT_H
#define LINEEDIT_H

#include "UserControl.h"

#include <QWidget>
#include <QString>
#include <QRect>
#include <QPainter>

class LineEdit :public UserControl
{
private:
    QString tip;
    QString text;
    QString textPasswd;
    QString textPhone;

    bool focus;
    bool indicatorAllow;
    bool caretOn;
    int timerID;
    QRect rectCaret;
    int maxChars=-1;
    bool pswdInput=false;
    bool onlyDigit=false;
    bool phoneNumMode=false;
    bool allowSlash=false;
    bool alignCenter=false;
public:
    LineEdit(QWidget *widget,const QRect& rect,const QString & tip);
    ~LineEdit();
    void draw(QPainter& painter);
    void setFocus(bool focus);
    QRect rect();
    void allowIndicator(bool allow);

    void setMaxChars(int max);
    void setPasswdInput();
    void setOnlyDigit();
    void setPhoneNumMode();
    void setAllowSlash();
    QString getText();
    QString getDisplayText();

    void setText(QString txt);
    void setAlignCenter(bool align);
    void processKey(QKeyEvent *event);
protected:

    void timerEvent(QTimerEvent *event);
};

#endif // LINEEDIT_H
