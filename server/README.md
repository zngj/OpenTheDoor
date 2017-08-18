# SmartGate Server

## 用户中心协议 - HTTPS

### 检查手机号码

检查用户输入的手机号码是否已被注册

请求地址：https://sgu.youstars.com.cn/user/check_phone_number/:phone_number

请求方式：GET

请求参数：

|参数|类型|必填|说明|
|---|----|---|--- |
|:phone_number|string|是|手机号码|

请求说明：
```json
GET https://sgu.youstars.com.cn/user/check_phone_number/13800138000
```


返回说明：
```json
//正常返回的JSON数据包
{
    "code": 0,
    "msg": "成功"
}
//手机号已被注册
{
    "code": 1001,
    "msg": "手机号已被注册"
}
```

### 用户手机注册

注册用户手机号和密码

请求地址：https://sgu.youstars.com.cn/user/signup

请求方式：POST

请求参数：

|参数|类型|必填|说明|
|---|----|---|--- |
|phone_number|string|是|手机号码|
|password|string|是|密码|

请求说明：
```json
{
	"phone_number": "13800138000",
	"password": "123456"
}
```

返回说明：
```json
//正常返回的JSON数据包
{
    "code": 0,
    "msg": "成功"
}
//手机号已被注册
{
    "code": 1001,
    "msg": "手机号已被注册"
}
```


### 用户手机号码登录

用户使用手机号码和密码登录

请求地址：https://sgu.youstars.com.cn/user/login

请求方式：POST

请求参数：

|参数|类型|必填|说明|
|---|----|---|--- |
|phone_number|string|是|手机号码|
|password|string|是|密码|

请求说明：
```json
{
	"phone_number": "13800138000",
	"password": "123456"
}
```

返回说明：
```json
//登录成功
{
    "code": 0,
    "msg": "成功",
    "data": {
        "access_token": "e7b154f559544ae78f8da201546c1a37",
        "expires_in": 7200
    }
}
//登录失败：手机号码不存在
{
    "code": 1002,
    "msg": "手机号码不存在"
}
//登录失败：密码不正确
{
    "code": 1003,
    "msg": "密码不正确"
}
```


### 应用APP微信登录

应用APP使用微信登录

请求地址：https://sgu.youstars.com.cn/user/login_wxapi

请求方式：POST

请求参数：

|参数|类型|必填|说明|
|---|----|---|--- |
|code|string|是|微信授权临时票据|

请求说明：
```json
{
	"code": "013i8LU00jGznD1WP4S00OCIU00i8LUR"
}
```

返回说明：
```json
//登录成功
{
    "code": 0,
    "msg": "成功",
    "data": {
        "access_token": "e7b154f559544ae78f8da201546c1a37",
        "expires_in": 7200
    }
}
//无效code
{
    "code": 40029,
    "msg": "无效的微信授权临时票据(code)"
}
```


### 应用APP微信登录

应用APP使用微信登录

请求地址：https://sgu.youstars.com.cn/user/login_wxapi

请求方式：POST

请求参数：

|参数|类型|必填|说明|
|---|----|---|--- |
|code|string|是|微信授权临时票据|

请求说明：
```json
{
	"code": "013i8LU00jGznD1WP4S00OCIU00i8LUR"
}
```

返回说明：
```json
//登录成功
{
    "code": 0,
    "msg": "成功",
    "data": {
        "access_token": "e7b154f559544ae78f8da201546c1a37",
        "expires_in": 7200
    }
}
//无效code
{
    "code": 40029,
    "msg": "无效的微信授权临时票据(code)"
}
```


### 小程序微信登录

用户打开小程序，自动使用微信登录

请求地址：https://sgu.youstars.com.cn/user/login_weapp

请求方式：POST

请求参数：

|参数|类型|必填|说明|
|---|----|---|--- |
|code|string|是|微信授权临时票据|

请求说明：
```json
{
	"code": "00351nLn1rO8uk0DzUMn1Jf9Ln151nLx"
}
```

返回说明：
```json
//登录成功
{
    "code": 0,
    "msg": "成功",
    "data": {
        "access_token": "e7b154f559544ae78f8da201546c1a37",
        "expires_in": 7200
    }
}
//无效code
{
    "code": 40029,
    "msg": "无效的微信授权临时票据(code)"
}
```


### Token检查

检查客户端Token是否到期

请求地址：https://sgu.youstars.com.cn/user/check_token

请求方式：GET

请求HEADER：

|Key|Value|描述|
|----------|----|-------|
|Access-Token|登录返回的access_token|用户访问授权|

请求说明：
```json
GET /check_token HTTP/1.1
Host: localhost:8081
Access-Token: 9276a1988a70477eb2a6700bebbe99b7
Cache-Control: no-cache
```

返回说明：
```json
//正常返回的JSON数据
{
    "code": 0,
    "msg": "成功"
}
//登录失效
{
    "code": 1000,
    "msg": "登录失效"
}
```


## 后台协议 - HTTPS

### 获取用户行程状态

获取用户当前行程的状态，用于判断用户打开客户是准备进站还是出站

请求地址：https://sgu.youstars.com.cn/sg/router/status

请求方式：GET

请求HEADER：

|Key|Value|描述|
|---|-----|---|
|Access-Token|登录返回的access_token|用户访问授权|


返回说明：
//正常返回的JSON数据
```json:
{
    "code": 0,
    "msg": "成功",
    "data": {
        "status": 0, // 0-无行程; 1-已入闸;
    }
}
//Token失效
{
    "code": 1000,
    "msg": "登录失效"
}
```


### 获取用户钱包信息

获取用户的钱包信息

请求地址：https://sgu.youstars.com.cn/sg/wallet/info

请求方式：GET

请求HEADER：

|Key|Value|描述|
|---|-----|---|
|Access-Token|登录返回的access_token|用户访问授权|

返回参数：

|参数|说明|
|---|---|
|balance|用户钱包余额|
|wxpay_quick|是否开通小额免密|

返回说明：
//正常返回的JSON数据
```json:
{
    "code": 0,
    "msg": "成功",
    "data": {
        "balance": 100,
        "wxpay_quick": false
    }
}
//Token失效
{
    "code": 1000,
    "msg": "登录失效"
}
```


### 获取用户进站凭证

获取用户进站凭证

请求地址：https://sgu.youstars.com.cn/sg/evidence/in

请求方式：GET

请求HEADER：

|Key|Value|描述|
|---|-----|---|
|Access-Token|登录返回的access_token|用户访问授权|

返回参数：

|参数|说明|
|---|---|
|evidence_key|加密码后的用户凭证，用于显示进站二维码|
|expires_at|失效时间，unix时间戳|

返回说明：
//正常返回的JSON数据
```json:
{
    "code": 0,
    "msg": "成功",
    "data": {
        "evidence_key": "MjWCCOKE9yDNMarR1l/j0nVok9wxExvKPtKleA/1OiO6Cvn0BM01Fdjb9MxSF9yTYBG48Bh85ZcQdaZ97TM3o8NJ1rOoKaqD+R1LdK/c6RGxHQ6rUPdXBU7yZP2rOBeN/xhjC7ge+iHwn6/3nwURr+33V1BUb7GzJqGerU6e59Q=",
        "expires_at": 1502120737
    }
}
//Token失效
{
    "code": 1000,
    "msg": "登录失效"
}
```


### 获取用户出站凭证

获取用户出站凭证

请求地址：https://sgu.youstars.com.cn/sg/evidence/out

请求方式：GET

请求HEADER：

|Key|Value|描述|
|---|-----|---|
|Access-Token|登录返回的access_token|用户访问授权|

返回参数：

|参数|说明|
|---|---|
|evidence_key|加密码后的用户凭证，用于显示出站二维码|
|expires_at|失效时间，unix时间戳|

返回说明：
//正常返回的JSON数据
```json:
{
    "code": 0,
    "msg": "成功",
    "data": {
        "evidence_key": "MjWCCOKE9yDNMarR1l/j0nVok9wxExvKPtKleA/1OiO6Cvn0BM01Fdjb9MxSF9yTYBG48Bh85ZcQdaZ97TM3o8NJ1rOoKaqD+R1LdK/c6RGxHQ6rUPdXBU7yZP2rOBeN/xhjC7ge+iHwn6/3nwURr+33V1BUb7GzJqGerU6e59Q=",
        "expires_at": 1502120737
    }
}
//Token失效
{
    "code": 1000,
    "msg": "登录失效"
}
```


### 获取用户的通知

获取系统发下给用户的通知，客户端根据通知类型，请求相应的接口拉取最新的信息，实现消息推送

请求地址：https://sgu.youstars.com.cn/sg/notification

请求方式：GET

请求HEADER：

|Key|Value|描述|
|---|-----|---|
|Access-Token|登录返回的access_token|用户访问授权|

返回参数：

|参数|说明|
|---|---|
|notification_id|通知ID|
|notification_type|通知类型：1-用户入站；2-用户出站；|

返回说明：
//正常返回的JSON数据
```json:
{
    "code": 0,
    "msg": "成功",
    "data": {
        "notification_id": 9,
        "notification_type": 1
    }
}
//Token失效
{
    "code": 1000,
    "msg": "登录失效"
}
```

回调说明：

|消息类型|回调说明|回调地址|
|---|---|---|
|1|获取用户当天入站信息|https://sgu.youstars.com.cn/sg/router/in/list|
|2|获取用户当天出站信息|https://sgu.youstars.com.cn/sg/router/out/list|


### 推送消息消费

客户端成功获取推送的信息后，请求后台进行通知消息消费


请求地址：https://sgu.youstars.com.cn/sg/notification/consume/:notification_id

请求方式：PUT

请求HEADER：

|Key|Value|描述|
|---|-----|---|
|Access-Token|登录返回的access_token|用户访问授权|

请求参数：

|参数|类型|必填|说明|
|---|----|---|--- |
|:notification_id|string|是|通知ID|

请求说明：
```json
PUT https://sgu.youstars.com.cn/sg/notification/consume/9
```

返回说明：
//正常返回的JSON数据
```json:
{
    "code": 0,
    "msg": "成功"
}
//Token失效
{
    "code": 1000,
    "msg": "登录失效"
}
```


### 获取用户当天入站信息

获取用户当天入站信息列表，用户可实时查看进站的信息


请求地址：https://sgu.youstars.com.cn/sg/router/in/list

请求方式：GET

请求HEADER：

|Key|Value|描述|
|---|-----|---|
|Access-Token|登录返回的access_token|用户访问授权|

返回说明：
//正常返回的JSON数据
```json:
{
    "code": 0,
    "msg": "成功",
    "data": [
        {
            "id": 8,
            "in_station_name": "五一广场",
            "in_time": 1503044457,
            "status": 1,
            "statusName": "已入站",
            "pay": false
        },
        {
            "id": 9,
            "in_station_name": "五一广场",
            "in_time": 1503045250,
            "status": 1,
            "statusName": "已入站",
            "pay": false
        }
    ]
}
//Token失效
{
    "code": 1000,
    "msg": "登录失效"
}
```


### 获取用户当天出站信息

获取用户当天出站信息列表，用户可实时查看出站的信息

请求地址：https://sgu.youstars.com.cn/sg/router/out/list

请求方式：GET

请求HEADER：

|Key|Value|描述|
|---|-----|---|
|Access-Token|登录返回的access_token|用户访问授权|

返回说明：
//正常返回的JSON数据
```json:
{
    "code": 0,
    "msg": "成功",
    "data": [
        {
            "id": 7,
            "in_station_name": "五一广场",
            "in_time": 1503044456,
            "out_station_name": "黄兴广场",
            "out_time": 1503044463,
            "status": 2,
            "statusName": "已出站",
            "money": 2,
            "pay": false
        },
        {
            "id": 8,
            "in_station_name": "五一广场",
            "in_time": 1503044457,
            "status": 1,
            "statusName": "已入站",
            "pay": false
        }
    ]
}
//Token失效
{
    "code": 1000,
    "msg": "登录失效"
}
```




## Gate协议 - TCP

### 出入闸凭证格式

#### Demo阶段
32位uuid+10位时间戳
如：0cf4524163c34d13916e2e649d8893671502120737

#### 正式阶段
将10位时间戳插入到32位uuid中，加入验证位，再加密混淆


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
别名：GATE_LOGIN<br/>
发送：闸机 <-> 后台<br/>

请求数据：无

返回数据：

|参数名     |类型|默认值  |说明    |
|----------|----|-------|--------|
|gate_id|string|-|闸机ID|
|gate_direction|string|-|0-入;1-出|
|station_name|string|-|站点名称|
|city_name|string|-|城市名称|
|errcode|int|-|1至999-通用错误<br/>3100-无效的闸机ID|
|errmsg|string|-|错误内容|

返回示例：
```json
{
    "gate_id": "010100101",
    "gate_direction": 0,
    "station_name": "五一广场",
    "city_name": "长沙"
}
```
```json
{
    "errcode": 3100,
    "errmsg": "invalid gate id"
}
```

#### MsgID: 101
协议：闸机未登录<br/>
别名：NOT_LOGIN<br/>
发送：后台 -> 闸机<br/>
描述：后台收到闸机发送的非登录协议，检查到闸机未登录，会向闸机发送此协议<br/>
数据：无

### 3-2 获取私钥
闸机获取私钥用于解密出入闸凭证

#### MsgID: 102
协议：获取公钥<br/>
别名：RSA_KEY<br/>
发送：闸机 <-> 后台<br/>

请求数据：无

返回数据：

|参数名     |类型|默认值  |说明    |
|----------|----|-------|--------|
|key|string|-|公钥|
|errcode|int|-|1至999-通用错误|
|errmsg|string|-|错误内容|

返回示例：
```json
{
    "key": "\n-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv\nppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd\nwYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL\nAUeJ6PeW+DAkmJWF6QIDAQAB\n-----END PUBLIC KEY-----\n"
}
```
```json
{
    "errcode": 1,
    "errmsg": "not found"
}
```


### 3-3 验证凭证
闸机能过扫描二维码，获取用户的出入闸凭证并解密，发送到后台验证

#### MsgID: 103
协议：请求验证凭证<br/>
别名：VERIFY_EVIDENCE<br/>
发送：闸机 <-> 后台<br/>

请求数据：

|参数名     |类型|默认值  |说明    |
|----------|----|-------|--------|
|evidence_key|string|-|凭证|

请求示例：
```json
{
    "evidence_key": "MjWCCOKE9yDNMarR1l/j0nVok9wxExvKPtKleA/1OiO6Cvn0BM01Fdjb9MxSF9yTYBG48Bh85ZcQdaZ97TM3o8NJ1rOoKaqD+R1LdK/c6RGxHQ6rUPdXBU7yZP2rOBeN/xhjC7ge+iHwn6/3nwURr+33V1BUb7GzJqGerU6e59Q="
}
```

返回数据：

|参数名     |类型|默认值  |说明    |
|----------|----|-------|--------|
|errcode|int|-|1到999-通用错误<br/>3201-无效凭证<br/>3202-凭证已过期<br/>3203-凭证与机闸不匹配<br/>3204-用户不符合付费标准|
|errmsg|string|-|错误内容|

返回示例：
```json
{
   //成功
}
```
```json
{
    "errcode": 3201,
    "errmsg": "invalid envidence"
}
```
```json
{
    "errcode": 3202,
    "errmsg": "expired envidence"
}
```

### 3-4 提交出入凭证
开闸后，闸机提交用户出入凭证

#### MsgID: 104
协议：提交凭证<br/>
别名：SUBMIT_EVIDENCE<br/>
发送：闸机 <-> 后台<br/>

请求数据：

|参数名     |类型|默认值  |说明    |
|----------|----|-------|--------|
|evidence_key|string|-|出入闸凭证|
|scan_time|int64|-|扫码unix时间戳|

请求示例：
```json
{
    "evidence_key": "MjWCCOKE9yDNMarR1l/j0nVok9wxExvKPtKleA/1OiO6Cvn0BM01Fdjb9MxSF9yTYBG48Bh85ZcQdaZ97TM3o8NJ1rOoKaqD+R1LdK/c6RGxHQ6rUPdXBU7yZP2rOBeN/xhjC7ge+iHwn6/3nwURr+33V1BUb7GzJqGerU6e59Q=",
    "scan_time": 1502120737
}
```

返回数据：

|参数名     |类型|默认值  |说明    |
|----------|----|-------|--------|
|errcode|int|-|1到999-通用错误|
|errmsg|string|-|错误内容|

返回示例：
```json
{
   //成功
}
```
```json
{
    "errcode": 3,
    "errmsg": "wrong argument"
}
```

