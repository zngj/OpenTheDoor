var util = require('../../js/util.js');

var app = getApp()
Page({
  data: {
    isLogin: 0,
    userBalance: '--',
    dataSetDesc: ['远程', '新用户', '有余额', '有代扣']
  },

  onLoad: function (options) {
    var that = this;
    this.getWallet(function (data) {
      that.setData({
        userBalance: data.balance
      });
    });
  },
  getWallet: function (successCB, failCB) {
    wx.getStorage({
      key: 'token',
      success: function (res) {
        util.request({
          url: '/user/wallet/info',
          data: { token: res.data },
          success: function (p) {
            if (p.data.code == 0) {
              successCB(p.data.data);
            };
          },
          fail: function (fp) {
            if (failCB) {
              failCB(fp);
            };
          }
        });
      },
    });
  },
  gateIn: function () {
    //../open/showCode?type=out
    this.getWallet(function (data) {
      if (data.autoPay) {
        util.showMsg('有代扣，可进入');
        wx.navigateTo({
          url: '../open/showCode?type=in',
        });
      } else if (data.balance > 0) {
        util.showMsg('有余额，可进入');
        wx.navigateTo({
          url: '../open/showCode?type=in',
        });
      } else {
        util.showMsg('需要充值、或者签约代扣');
        return;
      }
    });
  },
  gateOut: function () {
    wx.navigateTo({
      url: '../open/showCode?type=out',
    });

  },
  onShareAppMessage: function (page) { },
  bindPickerChange: function (e) {
    var dataSet = ['', 'new', 'hasBalance', 'autoPay'];
    util.initRequest(e.detail.value != 0, dataSet[e.detail.value]); 
  },
})
