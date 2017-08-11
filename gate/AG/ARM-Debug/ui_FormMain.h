/********************************************************************************
** Form generated from reading UI file 'FormMain.ui'
**
** Created by: Qt User Interface Compiler version 5.6.2
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_FORMMAIN_H
#define UI_FORMMAIN_H

#include <QtCore/QVariant>
#include <QtWidgets/QAction>
#include <QtWidgets/QApplication>
#include <QtWidgets/QButtonGroup>
#include <QtWidgets/QDialog>
#include <QtWidgets/QHeaderView>

QT_BEGIN_NAMESPACE

class Ui_FormMain
{
public:

    void setupUi(QDialog *FormMain)
    {
        if (FormMain->objectName().isEmpty())
            FormMain->setObjectName(QStringLiteral("FormMain"));
        FormMain->resize(400, 300);

        retranslateUi(FormMain);

        QMetaObject::connectSlotsByName(FormMain);
    } // setupUi

    void retranslateUi(QDialog *FormMain)
    {
        FormMain->setWindowTitle(QApplication::translate("FormMain", "Dialog", 0));
    } // retranslateUi

};

namespace Ui {
    class FormMain: public Ui_FormMain {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_FORMMAIN_H
