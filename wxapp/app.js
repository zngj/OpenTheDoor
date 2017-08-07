//app.js
var util = require('js/util.js');
App({
  data: {
  },
  onLaunch: function (options) {
    var that = this;
    util.initRequest(false);
    wx.checkSession({
      success: function (sp) {
        var token = wx.getStorageSync("token");
        console.log(token);
        util.request({
          url: '/user/verifytoken',
          header: { 'Access-Token': token },
          method: "GET",
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
        util.request({
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
