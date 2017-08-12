#ifndef FORMMAIN_H
#define FORMMAIN_H

#include <QWidget>

#include "BaseDialog.h"

class FormMain:public BaseDialog
{
public:
    explicit FormMain(QWidget *parent = 0);
    ~FormMain();
protected:
    void customDraw(QPainter &painter);
    void keyPressEvent(QKeyEvent *event);
    void processControlEvent(UserControl *control, ControlEvent event);
    void customEvent(QEvent *event);
private:
   FrameEdit * frameEdit;
   QRDisplay *qrDisplay;
};

#endif // FORMMAIN_H
