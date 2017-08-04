# SmartGate Server

## 用户协议 - HTTPS

### 1-1 小程序用户登录

说明：
1. uri: `/user/wxapp/login`
2. Method: `POST`

参数：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|code|string|是|-|登录凭证|

请求：
```json
{
  "code": "00351nLn1rO8uk0DzUMn1Jf9Ln151nLx"
}
```

返回：
```json:
{
    "code": 0,
    "msg": "success",
    "data": {
        "token": "91af8c6de82d4e17842508e3a42df412"
    }
}
```
```json:
{
    "code": -4,
    "msg": "illegal argument"
}
```


### 1-2 用户Token验证

说明：
1. uri: `/user/verifytoken`
2. Method: `POST`

参数：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|token|string|是|-|登录凭证|

请求：
```json
{
  "token": "91af8c6de82d4e17842508e3a42df412"
}
```

返回：
```json:
{
    "code": 0,
    "msg": "success",
}
```
```json:
{
    "code": -1000,
    "msg": "token was expired"
}
```


## Weapp协议 - HTTPS

### 2-1 获取用户出入闸状态

### 2-3


## Gate协议 - TCP

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

### 3-1 设置闸机编号
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


### 3-2 上报用户入闸数据
当闸机开启后上报用户入闸数据

#### 请求: MSG_USER_IN
参数说明：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|qr|string|是|-|入闸二维码数据|

请求样例：
```json
{
  "qr": "TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0aGlz",
}
```

#### 应答：MSG_USER_IN_SUCCESS
收到消息即为成功


### 3-3 上报用户出闸数据
当闸机开启后上报用户出闸数据

#### 请求: MSG_USER_OUT
参数说明：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|qr|string|是|-|出闸二维码数据|

请求样例：
```json
{
  "qr": "IHNpbmd1bGFyIHBhc3Npb24gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1c3Qgb2Yg",
}
```

#### 应答：MSG_USER_OUT_SUCCESS
收到消息即为成功

