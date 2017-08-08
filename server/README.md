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

### 闸机编号
编号是由0-9组成的字符串
```code
国家(2位) + 城市(2位) + 站点(3位) + 闸机(2位)
```
示例
```code
五一广场的01号闸机：010100101
黄兴广场的02号闸机：010100202
```

### 3-1 闸机登录
当闸机成功连接后，向后台发送登录协议，登录成功后获取闸机的信息

#### MsgID: 100
协议：闸机登录<br/>
别名：C2S_GATE_LOGIN<br/>
发送：闸机 -> 后台<br/>
数据：无

#### MsgID：101
协议：登录结果<br/>
别名：S2C_GATE_LOGIN<br/>
发送：后台 -> 闸机<br/>
数据：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|code|int|是|-|0-登录成功;3100-GateId不存在|
|gate_direction|string|否|-|0-入;1-出|
|station_name|string|否|-|站点名称|
|city_name|string|否|-|城市名称|

示例：
```json
{
    "code": 0,
    "gate_direction": 0,
    "station_name": "五一广场",
    "city_name": "长沙"
}
```
```json
{
    "code": 1
}
```

#### MsgID: 102
协议：闸机未登录<br/>
别名：S2C_NOT_LOGIN<br/>
发送：后台 -> 闸机<br/>
描述：后台收到闸机发送的非登录协议，检查到闸机未登录，会向闸机发送此协议<br/>
数据：无

### 3-2 获取私钥
闸机获取私钥用于解密出入闸凭证

#### MsgID: 103
协议：请求私钥<br/>
别名：C2S_RSA_KEY<br/>
发送：闸机 -> 后台<br/>
数据：无

#### MsgID：104
协议：私钥结果<br/>
别名：S2C_RSA_KEY<br/>
发送：后台 -> 闸机<br/>
数据：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|key|string|是|-|私钥|

示例：
```json
{
    "key": "\n-----BEGIN RSA PRIVATE KEY-----\nMIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y\n7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7\nAte8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB\nAoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM\nZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1\nXKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB\n/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40\nIRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG\n4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9\nDAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8\n9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw\nDPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO\nAQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O\n-----END RSA PRIVATE KEY-----\n"
}
```

### 3-3 验证凭证
闸机能过扫描二维码，获取用户的出入闸凭证并解密，发送到后台验证

#### MsgID: 200
协议：请求验证凭证<br/>
别名：C2S_VERIFY_EVIDENCE<br/>
发送：闸机 -> 后台<br/>
数据：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|evidence_key|string|是|-|凭证|

示例：
```json
{
    "evidence_key": "0cf4524163c34d13916e2e649d889367"
}
```

#### MsgID: 201
协议：验证凭证结果<br/>
别名：S2C_VERIFY_EVIDENCE<br/>
发送：后台 -> 闸机<br/>
数据：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|code|int|是|-|0-通过;<br/>3201-凭证不存在;<br/>3202-凭证已过期;<br/>3203-凭证与机闸不匹配;<br/>3204-用户不符合付费标准|

示例：
```json
{
    "code": 0
}
```

### 3-4 用户出入闸数据
开闸后，闸机上报用户入阐数据

#### MsgID: 202
协议：上报出入闸数据<br/>
别名：C2S_USER_EVIDENCE<br/>
发送：闸机 -> 后台<br/>
数据：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|evidence_key|string|是|-|出入闸凭证|
|scan_time|int64|是|-|扫码时间|

示例：
```json
{
    "evidence_key": "0cf4524163c34d13916e2e649d889367",
    "scan_time": 1502120737000  //Millisecond
}
```

#### MsgID: 203
协议：上报结果<br/>
别名：S2C_USER_EVIDENCE<br/>
发送：后台 -> 闸机<br/>
数据：

|参数名     |类型|是否必须|默认值  |说明    |
|----------|----|-------|-------|--------|
|success|bool|是|-|是否成功|

示例：
```json
{
    "success": true
}
```

