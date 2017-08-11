var request = {
  dataSet: null,
  init: function (localTest, localTestDataSet) {
    this.request = localTest ? this.requestLocal : this.requestRemote;
    this.dataSet = localTestDataSet;
  },
  requestRemote: function (options) {
    options.url = "https://sgu.youstars.com.cn" + options.url;
    options.method = options.method || "POST";
    options.header = options.header || { 'content-type': 'application/json' };
    var originComplete = options.complete;
    options.complete = function (c) {
      if (originComplete) {
        originComplete(c);
      }
      console.log(options);
      console.log(c);
    };
    wx.request(options);
  },
  request: function (options) {
    this.requestRemote(options);
  },
  requestLocal: function (options) {
    var actionObj = this.actionMap[options.url];
    var result = this.getAction(actionObj, this.dataSet)(options);

    var resultData = { data: result.success || result.fail`` };
    if (result.success && options.success) {
      options.success(resultData);
    }
    if (result.fail && options.fail) {
      options.fail(resultData);
    }
    if (options.complete) {
      options.complete(resultData);
    }
    console.log(options);
    console.log(result);
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
    "/user/verifytoken": {
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
    "/wallet/info": {
      default: function (options) {
        return { success: { code: 0, msg: 'success', data: { balance: 0, autoPay: false } } }
      },
      hasBalance: function (options) {
        return { success: { code: 0, msg: 'success', data: { balance: 100, autoPay: false } } }
      },
      autoPay: function (options) {
        return { success: { code: 0, msg: 'success', data: { balance: 0, autoPay: true } } }
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
function initRequest() {
  //request.init(arguments[0]);
  request.init.apply(request, arguments);
}
function theRequest(options) {
  request.request(options);
}
function showMsg(title) {
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
    var newBytes = [this.getRandom(), this.getRandom()];
    newBytes = newBytes.concat(bytes).concat(this.intToBytes(new Date().getTime()));
    return newBytes;
  },
  getRandom: function () {
    return Math.ceil(Math.random() * 255);
  },
}
module.exports = {
  redirectTo: redirectTo,
  isMobile: isMobile,
  initRequest: initRequest,
  request: theRequest,
  showMsg: showMsg,
  util: util
}

