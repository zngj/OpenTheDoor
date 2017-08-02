//index.js
//获取应用实例

var app = getApp()
Page({
  data: {
    isLogin: 0,
    userBalance: '--'
  },

  onLoad: function(options){
    var that = this;
    wx.login({
      success: function (p) {
        console.log("s:");
        console.log( p);
      },
      fail: function (p) {
        console.log("f:");
        console.log(p);
      },
      complete: function (p) {
        console.log("c:");
        console.log(p);
      }
    })
    wx.checkSession({
      success: function (p) {
        console.log("s:");
        console.log(p);
      },
      fail: function (p) {
        console.log("f:");
        console.log(p);
      },
      complete: function (p) {
        console.log("c:");
        console.log(p);
      }
    });
    /*
    wx.request({
      url: 'http://localhost/wxapp/login',
      data: {
        mobile: name,
        password: psw
      },
      method: 'POST',
      header: {
        'content-type': 'application/json'
      },
      success: function (resp) {
        if (resp.data.code == 0) { // 成功
          that.setData({
            isLogin: 1
          });
          wx.setStorageSync('loginUser', resp.data.data);
        } else {
        }
      },
      fail: function (resp) {
        if (wx.reLaunch) {
          wx.reLaunch({
            url: '/pages/propogate/propogate',
          })
        } else {
          wx.redirectTo({
            url: '/pages/propogate/propogate',
          });
        }
      }
    });
    */
  },

  onShow: function () {
    var that = this;
  },
})
