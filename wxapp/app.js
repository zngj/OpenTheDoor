var util = require('js/util.js').util;
var request = require('js/util.js').request;
App({
  data: {
  },
  onLaunch: function (options) {
    var that = this;
    request.init(false);
    wx.checkSession({
      success: function (sp) {
        console.log(sp);
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
