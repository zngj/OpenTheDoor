var request = {
  dataSet: null,
  init: function (localTest, localTestDataSet) {
    this.request = localTest ? this.requestLocal : this.requestRemote;
    this.dataSet = localTestDataSet;
  },
  arround: function (delegate, newFunc) {
    return function (s) {
      newFunc(delegate, s);
    };
  },
  requestRemote: function (options) {
    var start=new Date().getTime();
    var url =options.url;
    options.url = "https://sgu.youstars.com.cn" + options.url;
    options.method = options.method || "POST";
    options.header = options.header || { 'content-type': 'application/json' };
    var originSuccess = options.success;

    options.success = this.arround(options.success, function (delegate, s) {
      if (s.statusCode == 200) {
        if (s.data.code == 1000) {//token expired
        getApp().login("TokenExpired");
          wx.reLaunch({
            url: '/pages/index/index',
          })
        } else if (delegate) {
          delegate(s.data);
        }
      } else {
        options.fail(s);
      }
    });
    options.complete = this.arround(options.complete, function (delegate, c) {
      var end = new Date().getTime();
      console.log([(end-start),url,options, c]);
      if (delegate) {
        delegate(c);
      }
    });

    var token = wx.getStorageSync('token');
    if (token) {
      options.header = { 'Access-Token': token };
    }
    wx.request(options);
  },
  get: function (options) {
    var token = wx.getStorageSync('token');
    options.method = 'GET';
    options.header = { 'Access-Token': token };
    this.request(options);
  },
  put: function (options) {
    var token = wx.getStorageSync('token');
    options.method = 'PUT';
    options.header = { 'Access-Token': token };
    this.request(options);
  },
  request: function (options) {
    this.requestRemote(options);
  },
  requestLocal: function (options) {
    var actionObj = this.actionMap[options.url];
    var result = this.getAction(actionObj, this.dataSet)(options);

    var resultData = { statusCode: 200, data: result.success || result.fail };
    if (result.success && options.success) {
      options.success(resultData);
    }
    if (result.fail && options.fail) {
      options.fail(resultData);
    }
    if (options.complete) {
      options.complete(resultData);
    }
    console.log([options, result]);
  },
  getAction: function (actionObj, dataSet) {
    if (!dataSet) {
      return actionObj['default'];
    } else if (actionObj[dataSet]) {
      return actionObj[dataSet];
    } else {
      return this.getAction(actionObj, this.dataDependentChain[dataSet]);
    }
  },

  actionMap: {
    "/user/wxapp/login": {
      default: function (options) {
        return {
          success: { 'code': 0, msg: 'success', data: { 'token': 'This is a test token' } }
        }
      },
      reject: function (options) {
        return {
          success: { 'code': -100, msg: 'Cannot get token,invalid code' }
        }
      }
    },
    "/user/check_token": {
      default: function (options) {
        if (options.data.access_token) {
          return { success: { 'code': 0, msg: 'success', data: { token: "This is a dummy user token" } } };
        } else {
          return { success: { 'code': -1, msg: 'token not exists' } };
        }
      },
      tokenExpire: function (options) {
        return { success: { 'code': -1000, msg: 'token was expired' } };

      }
    },
    "/sg/wallet/info": {
      default: function (options) {
        return { success: { code: 0, msg: 'success', data: { balance: 0, wxpay_quick: false } } }
      },
      hasBalance: function (options) {
        return { success: { code: 0, msg: 'success', data: { balance: 100, wxpay_quick: false } } }
      },
      autoPay: function (options) {
        return { success: { code: 0, msg: 'success', data: { balance: 0, wxpay_quick: true } } }
      }
    },
    "/sg/wallet/charge": {
      default: function (options) {
        return { success: { code: 0, msg: 'success', data: { balance: 0, autoPay: false } } }
      }
    },
    "/sg/router/status": {
      default: function (options) {
        return { success: { code: 0, msg: 'success', data: { status: 0 } } }
      }
    },
    "/sg/router/evidence/in": {
      default: function (options) {
        return { success: { code: 0, msg: 'success', data: { evidence_key: 'ZMh7eyMC' } } }
      }
    },
    "/sg/router/evidence/out": {
      default: function (options) {
        return { success: { code: 0, msg: 'success', data: { evidence_key: 'ZMh7e3sa' } } }
      }
    },
    "/sg/notification/router": {
      default: function (options) {
        return {
          success: {
            code: 5, msg: 'not found'
          }
        }
      },
      hadEntry: function (options) {
        return {
          success: {
            code: 0, msg: 'success', data: {
              "notification_id": 9,
              "direction": 0, //入闸
              "in_gate_id": "010100101",
              "in_station_id": "0101001",
              "in_station_name": "五一广场",
              "in_time": 1502304832
            }
          }
        }
      }
    }
  },
  dataDependentChain: {
    //default: never used
    new: 'default',
    tokenExpire: 'default',
  }

};


function redirectTo(url) {
  wx.redirectTo({
    url: url
  })
}
function isMobile(mobile) {
  if (!mobile) {
    return false;
  }
  var reg = /^(((13[0-9]{1})|(14[0-9]{1})|(15[0-9]{1})|(17[0-9]{1})|(18[0-9]{1}))+\d{8})$/;
  if (!reg.test(mobile)) {
    return false;
  }
  return true;
}
var util = {
  intToBytes: function (value) {
    var src = [];
    src[3] = ((value >> 24) & 0xFF);
    src[2] = ((value >> 16) & 0xFF);
    src[1] = ((value >> 8) & 0xFF);
    src[0] = (value & 0xFF);
    return src;
  },

  bytesToInt: function (src, offset) {
    var value = ((src[offset] & 0xFF)
      | ((src[offset + 1] & 0xFF) << 8)
      | ((src[offset + 2] & 0xFF) << 16)
      | ((src[offset + 3] & 0xFF) << 24));
    console.log(src[offset])
    return value;
  },
  mix: function (bytes) {
    return [83, 71, this.getRandom(), this.getRandom()].concat(bytes).concat(this.intToBytes(Math.ceil(new Date().getTime() / 1000)));
  },
  getRandom: function () {
    return Math.ceil(Math.random() * 255) & 0xFF;
  },
  showMsg: function (title) {
    wx.showToast({
      title: title,
      fail: function () {
        wx.showModal({
          title: '提示',
          content: title,
        })
      }
    });
  }
}
module.exports = {
  redirectTo: redirectTo,
  isMobile: isMobile,
  request: request,
  util: util
}

