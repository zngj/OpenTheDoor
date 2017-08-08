# SmartGate Server

## 用户协议 - HTTPS

### 1-1 小程序用户登录

说明：
1. uri: /wxapp/login
2. Method: POST

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
        "access_token": "e062b066da834699aa10905e5d45a576",
        "expires_in": 7200
    }
}
```
```json:
{
    "code": 4,
    "msg": "illegal argument"
}
```


### 1-2 用户Token验证

说明：
1. uri: /verifytoken
2. Method: POST

Header:

|Key     |Value|描述|
|----------|----|-------|
|Access-Token|登录返回的access_token|用户访问授权|

返回：
```json:
{
    "code": 0,
    "msg": "success",
}
```
```json:
{
    "code": 1000,
    "msg": "token was expired"
}
```


## Weapp协议 - HTTPS

### 2-1 获取用户钱包信息

说明：
1. uri: /wallet/info
2. Method: GET

Header:

|Key     |Value|描述|
|----------|----|-------|
|Access-Token|登录返回的access_token|用户访问授权|

返回：
```json:
{
    "code": 0,
    "msg": "success",
    "data": {
        "balance": 100,
        "wxpay_quick": false
    }
}
```
```json:
{
    "code": 1000,
    "msg": "token was expired"
}
```


### 2-2 获取用户行程状态

说明：
1. uri: /router/status
2. Method: GET

Header:

|Key     |Value|描述|
|----------|----|-------|
|Access-Token|登录返回的access_token|用户访问授权|

返回：
```json:
{
    "code": 0,
    "msg": "success",
    "data": {
        "status": 0, // 0-无行程; 1-已入闸; 2-隔天未出闸(异常); 4-已出闸未入闸(异常)
    }
}
```
```json:
{
    "code": 1000,
    "msg": "token was expired"
}
```


### 2-3 获取用户入阐凭证

说明：
1. uri: /router/evidence/in
2. Method: GET

Header:

|Key     |Value|描述|
|----------|----|-------|
|Access-Token|登录返回的access_token|用户访问授权|

返回：
```json:
{
    "code": 0,
    "msg": "success",
    "data": {
        "evidence_key": "MjWCCOKE9yDNMarR1l/j0nVok9wxExvKPtKleA/1OiO6Cvn0BM01Fdjb9MxSF9yTYBG48Bh85ZcQdaZ97TM3o8NJ1rOoKaqD+R1LdK/c6RGxHQ6rUPdXBU7yZP2rOBeN/xhjC7ge+iHwn6/3nwURr+33V1BUb7GzJqGerU6e59Q=",
        "expires_at": 1502120737 //unxi时间戳
    }
}
```
```json:
{
    "code": 1000,
    "msg": "token was expired"
}
```


### 2-4 获取用户出阐凭证

说明：
1. uri: /router/evidence/out
2. Method: GET

Header:

|Key     |Value|描述|
|----------|----|-------|
|Access-Token|登录返回的access_token|用户访问授权|

返回：
```json:
{
    "code": 0,
    "msg": "success",
    "data": {
        "evidence_key": "MjWCCOKE9yDNMarR1l/j0nVok9wxExvKPtKleA/1OiO6Cvn0BM01Fdjb9MxSF9yTYBG48Bh85ZcQdaZ97TM3o8NJ1rOoKaqD+R1LdK/c6RGxHQ6rUPdXBU7yZP2rOBeN/xhjC7ge+iHwn6/3nwURr+33V1BUb7GzJqGerU6e59Q=",
        "expires_at": 1502120737 //unxi时间戳
    }
}
```
```json:
{
    "code": 1000,
    "msg": "token was expired"
}
```


## Gate协议 - TCP

### 消息封包格式

见《闸机后台协议.doc》

### 3-1 闸机登录
当闸机成功连接后，向服务端发送登录协议，登录成功后获取闸机的信息

#### 请求MsgID: 100 (C2S_GATE_LOGIN)

消息体：无

#### 应答MsgID：101 (S2C_GATE_LOGIN)

消息体：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|code|int8|是|-|0:登录成功;1-GateId不存在|

样例：
```json
{
    "code": 0,
    "station_name": "五一广场",
    "city_name": "长沙"
}
```
```json
{
    "code": 1
}
