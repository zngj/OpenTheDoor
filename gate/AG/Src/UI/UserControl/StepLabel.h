#ifndef STEPLABEL_H
#define STEPLABEL_H

#include "UserControl.h"
#include <QWidget>
#include <QString>
#include <vector>

class StepLabel :public UserControl
{
private:
    std::vector<QString> listSteps;
    int currentStep=-1;
public:
    StepLabel(const QWidget *widget,const QRect &rect);
    void draw(QPainter &painter);
    QRect rect();
    void setFocus();
    void setStepName(QString name);
    void setCurrentStep(int step);

};

#endif // STEPLABEL_H
