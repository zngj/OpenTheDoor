#ifndef CHECKDIALOG_H
#define CHECKDIALOG_H


#include "BaseDialog.h"

#include <QPainter>

class CheckDialog:public BaseDialog
{
public:
    CheckDialog(QWidget *parent);
    virtual ~CheckDialog();

protected:
    void customDraw(QPainter &painter);
};

#endif // CHECKDIALOG_H
