/********************************************************************************
** Form generated from reading UI file 'popupwindow.ui'
**
** Created by: Qt User Interface Compiler version 5.7.1
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_POPUPWINDOW_H
#define UI_POPUPWINDOW_H

#include <QtCore/QVariant>
#include <QtWidgets/QAction>
#include <QtWidgets/QApplication>
#include <QtWidgets/QButtonGroup>
#include <QtWidgets/QDialog>
#include <QtWidgets/QHeaderView>

QT_BEGIN_NAMESPACE

class Ui_PopupWindow
{
public:

    void setupUi(QDialog *PopupWindow)
    {
        if (PopupWindow->objectName().isEmpty())
            PopupWindow->setObjectName(QStringLiteral("PopupWindow"));
        PopupWindow->resize(400, 300);

        retranslateUi(PopupWindow);

        QMetaObject::connectSlotsByName(PopupWindow);
    } // setupUi

    void retranslateUi(QDialog *PopupWindow)
    {
        PopupWindow->setWindowTitle(QApplication::translate("PopupWindow", "Dialog", Q_NULLPTR));
    } // retranslateUi

};

namespace Ui {
    class PopupWindow: public Ui_PopupWindow {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_POPUPWINDOW_H
