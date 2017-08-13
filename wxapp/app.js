var util = require('js/util.js').util;
var request = require('js/util.js').request;
App({
  data: {
  },
  onLaunch: function (options) {
    console.log("App onLaunch");
    request.init(false);
    //this.ensureSession();
  },
  onShow: function () {
    console.log("App on show");
    this.ensureSession();
    console.log("CurrentPages:");
    console.log(getCurrentPages());
    
  },
  ensureSession:function(){
    var that = this;
    wx.checkSession({
      success: function (sp) {
        request.get({
          url: '/user/verifytoken',
          success: function (p) {
            if (p.data.code == 0) {
              // token valid
            } else {
              // token invalid
              that.login('TokenInvalid');
            }
          },
          fail: function (c) {
            that.login('verifyFail');
          }
        });
      },
      fail: function (fp) {
        that.login('SessionFail');
      }
    });
  },
  login: function (entry) {
    wx.login({
      success: function (p) {
        request.request({
          url: '/user/wxapp/login',
          data: { code: p.code },
          success: function (loginResult) {
            loginResult = loginResult.data;
            if (loginResult.code == 0) {
              wx.setStorageSync('token', loginResult.data.access_token);
            } else {
              util.showMsg(loginResult.msg);
            };
            console.log(entry);
          }, fail: function (pp) {
            console.log(entry);
          }
        });
      },
      fail: function (p) {
        console.log(entry);
      }
    });
  }
});
