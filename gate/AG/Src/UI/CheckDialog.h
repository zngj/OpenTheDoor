#ifndef CHECKDIALOG_H
#define CHECKDIALOG_H


#include "BaseDialog.h"

#include <QPainter>

class CheckDialog:public BaseDialog
{
public:
    CheckDialog(QWidget *parent);
    virtual ~CheckDialog();
private:
    QString errMsg;
    QPixmap map;
protected:
    void customDraw(QPainter &painter);
};

#endif // CHECKDIALOG_H
