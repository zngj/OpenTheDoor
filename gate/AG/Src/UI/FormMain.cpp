#include "FormMain.h"
#include "ui_FormMain.h"

FormMain::FormMain(QWidget *parent) :
    QDialog(parent),
    ui(new Ui::FormMain)
{
    ui->setupUi(this);
}

FormMain::~FormMain()
{
    delete ui;
}
