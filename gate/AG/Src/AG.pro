#-------------------------------------------------
#
# Project created by QtCreator 2017-08-05T16:31:25
#
#-------------------------------------------------

QT       += core gui

greaterThan(QT_MAJOR_VERSION, 4): QT += widgets

TARGET = AG
TEMPLATE = app

LIBS += -lfftw3 -lasound -ljsoncpp -lssl -lcrypto -lqrencode

# The following define makes your compiler emit warnings if you use
# any feature of Qt which as been marked as deprecated (the exact warnings
# depend on your compiler). Please consult the documentation of the
# deprecated API in order to know how to port your code away from it.
DEFINES += QT_DEPRECATED_WARNINGS

# You can also make your code fail to compile if you use deprecated APIs.
# In order to do so, uncomment the following line.
# You can also select to disable deprecated APIs only up to a certain version of Qt.
#DEFINES += QT_DISABLE_DEPRECATED_BEFORE=0x060000    # disables all the APIs deprecated before Qt 6.0.0



SOURCES += main.cpp\
    Drivers/Scanner/IScannerListener.cpp \
    Drivers/Scanner/ScannerManager.cpp \
    Soundwave/SoundDecode.cpp \
    Network/NetMessage.cpp \
    Network/NetRequest.cpp \
    Network/Server.cpp \
    Network/DataServer.cpp \
    Storage/BasicConfig.cpp \
    Utils/PathUtil.cpp \
    Utils/TimeUtil.cpp \
    Lib/Base64/Base64.cpp \
    Crypto/AES128.cpp \
    Crypto/RSA1024.cpp \
    Business/ScannerCheck.cpp \
    Business/CryptoManager.cpp \
    Business/ChangeLogManager.cpp \
    Business/ChangeLog.cpp \
    UI/BaseDialog.cpp \
    UI/PopupWindow.cpp \
    UI/UIListener.cpp \
    UI/UserControl/FrameEdit.cpp \
    UI/UserControl/IEventListener.cpp \
    UI/UserControl/LineEdit.cpp \
    UI/UserControl/ListView.cpp \
    UI/UserControl/PushButton.cpp \
    UI/UserControl/QRDisplay.cpp \
    UI/UserControl/SplitterLabel.cpp \
    UI/UserControl/StaticLabel.cpp \
    UI/UserControl/StepLabel.cpp \
    UI/UserControl/UserControl.cpp \
    UI/FormMain.cpp \
    UI/CheckDialog.cpp

HEADERS  += \
    Drivers/Scanner/IScannerListener.h \
    Drivers/Scanner/ScannerManager.h \
    Drivers/Scanner/ScannerState.h \
    Soundwave/SoundDecode.h \
    Network/NetMessage.h \
    Network/NetRequest.h \
    Network/Server.h \
    Network/DataServer.h \
    Storage/BasicConfig.h \
    Utils/PathUtil.h \
    Utils/TimeUtil.h \
    Lib/Base64/Base64.h \
    Crypto/AES128.h \
    Crypto/RSA1024.h \
    Business/ScannerCheck.h \
    Business/CryptoManager.h \
    Business/ChangeLogManager.h \
    Business/ChangeLog.h \
    UI/BaseDialog.h \
    UI/PopupWindow.h \
    UI/UIListener.h \
    UI/UserControl/ControlCommon.h \
    UI/UserControl/FrameEdit.h \
    UI/UserControl/IEventListener.h \
    UI/UserControl/LineEdit.h \
    UI/UserControl/ListView.h \
    UI/UserControl/PushButton.h \
    UI/UserControl/QRDisplay.h \
    UI/UserControl/SplitterLabel.h \
    UI/UserControl/StaticLabel.h \
    UI/UserControl/StepLabel.h \
    UI/UserControl/UserControl.h \
    UI/FormMain.h \
    UI/CheckDialog.h

FORMS    += \
    UI/BaseDialog.ui \
    UI/popupwindow.ui

RESOURCES += \
    resource.qrc

