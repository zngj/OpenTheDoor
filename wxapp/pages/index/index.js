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
    var token = wx.getStorageSync('token');
    console.log(token);
    wx.getStorage({
      key: 'token',
      success: function (res) {
        if (wx.showLoading) {
          wx.showLoading({
            title: '处理中',
          });
        }
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
          },
          complete:function(){
            if(wx.hideLoading){
              wx.hideLoading();
            }
          }
        });
      },
    });
  },
  gateIn: function () {
    var that = this;
    this.getWallet(function (data) {
      if (data.autoPay) {
        util.showMsg('已签约代扣，可进入');
        wx.navigateTo({
          url: '../open/showCode?type=in',
        });
      } else if (data.balance > 0) {
        util.showMsg('有余额，可进入');
        wx.navigateTo({
          url: '../open/showCode?type=in',
        });
      } else {
        that.tipForRecharge('暂不进站');
      }
    });
  },
  gateOut: function () {
    var that = this ;
    this.getWallet(function (data) {
      if (data.autoPay) {
        util.showMsg('有代扣，可进入');
        wx.navigateTo({
          url: '../open/showCode?type=out',
        });
      } else if (data.balance > 0) {
        if(!route || route.fee<data.balance){
        util.showMsg('有余额，可进入');
        wx.navigateTo({
          url: '../open/showCode?type=out',
        });
        } 
      } else {
        that.tipForRecharge('暂不出站');
      }
    });
  },
  tipForRecharge: function (cancelTip){
    wx.showModal({
      title: '提醒',
      content: '没有余额且未签约代扣！',
      showCancel: true,
      cancelText: cancelTip,
      confirmText: "签约充值",
      success: function (res) {
        if (res.confirm) {
          wx.navigateTo({
            url: '../recharge/recharge',
          });
        }
        console.log(res);
      }
    })
  },
  onShareAppMessage: function (page) { },
  bindPickerChange: function (e) {
    var dataSet = ['', 'new', 'hasBalance', 'autoPay'];
    util.initRequest(e.detail.value != 0, dataSet[e.detail.value]); 
  },
})
