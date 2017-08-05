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
        wx.getStorage({
          key: 'token',
          success: function (res) {
            console.log(res);
            util.request({
              url: '/user/verifytoken',
              data: { access_token:res.data},
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
          fail: function (c) { that.login('StorageFail'); }
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
              wx.setStorage({
                key: 'token',
                data: loginResult.access_token,
              })
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
