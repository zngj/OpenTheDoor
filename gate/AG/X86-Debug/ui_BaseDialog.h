/********************************************************************************
** Form generated from reading UI file 'BaseDialog.ui'
**
** Created by: Qt User Interface Compiler version 5.7.1
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_BASEDIALOG_H
#define UI_BASEDIALOG_H

#include <QtCore/QVariant>
#include <QtWidgets/QAction>
#include <QtWidgets/QApplication>
#include <QtWidgets/QButtonGroup>
#include <QtWidgets/QDialog>
#include <QtWidgets/QHeaderView>

QT_BEGIN_NAMESPACE

class Ui_BaseDialog
{
public:

    void setupUi(QDialog *BaseDialog)
    {
        if (BaseDialog->objectName().isEmpty())
            BaseDialog->setObjectName(QStringLiteral("BaseDialog"));
        BaseDialog->resize(800, 480);

        retranslateUi(BaseDialog);

        QMetaObject::connectSlotsByName(BaseDialog);
    } // setupUi

    void retranslateUi(QDialog *BaseDialog)
    {
        BaseDialog->setWindowTitle(QApplication::translate("BaseDialog", "Dialog", Q_NULLPTR));
    } // retranslateUi

};

namespace Ui {
    class BaseDialog: public Ui_BaseDialog {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_BASEDIALOG_H
