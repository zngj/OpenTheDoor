# SmartGate Server

## Weapp协议

### 1-1 小程序登录


### 1-2 获取用户出入站状态


### 1-3



## Gate协议

### 消息封包格式

```code
MSG = SIZE(1) + MSGID(1) + JSON_PAYLOAD(size-2)
```

### MSGID:

|MSG|ID|Msg Reply|ID|
|---|---|---|---|
|MSG_SET_GATE_NO|100|MSG_SET_GATE_NO_SUCCESS|200|
|MSG_USER_IN|101|MSG_USER_IN_SUCCESS|201|
|MSG_USER_OUT|102|MSG_USER_OUT_SUCCESS|202|

### 2-1 设置闸机编号
当闸机成功连接到服务端后进行闸机编号设置，当闸机断线重连后也需要设置

#### 请求: MSG_SET_GATE_NO
参数说明：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|no|string|是|-|闸机编号|

请求样例：
```json
{
  "no": "010101",
}
```

#### 应答：MSG_SET_GATE_NO_SUCCESS
收到消息即为成功


### 2-2 上报用户入站数据
当闸机开启后上报用户入站数据

#### 请求: MSG_USER_IN
参数说明：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|qr|string|是|-|入站二维码数据|

请求样例：
```json
{
  "qr": "TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0aGlz",
}
```

#### 应答：MSG_USER_IN_SUCCESS
收到消息即为成功


### 2-3 上报用户出站数据
当闸机开启后上报用户出站数据

#### 请求: MSG_USER_OUT
参数说明：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|qr|string|是|-|出站二维码数据|

请求样例：
```json
{
  "qr": "IHNpbmd1bGFyIHBhc3Npb24gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1c3Qgb2Yg",
}
```

#### 应答：MSG_USER_OUT_SUCCESS
收到消息即为成功

