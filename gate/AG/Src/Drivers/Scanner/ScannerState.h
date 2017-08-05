#ifndef SANNERSTATE_H
#define SANNERSTATE_H

enum ScannerState
{
    StateOK=0, //正常工作
    StateSerialPortNotSet, //串口文件未指定
    StateSerialPortNotExist,//串口文件不存在
    StateSerialPortNotOpen,//串口文件打不开(缺少权限？)

};
#endif // SANNERSTATE_H
