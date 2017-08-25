var util = require('../../js/util.js').util;
var request = require('../../js/util.js').request;

var app = getApp()
Page({
  data: {
    userBalance: '--',
    dataSetDesc: ['远程', '新用户', '有余额', '有代扣'],
    hasRoute : false
  },

  onLoad: function (options) {
    // request.init(true, "hasBalance");//local test for has balance
    // request.init(true, "new");//local test for new user
  },
  onShow: function () {
    var that = this;
    this.getWallet(function (data) {
      that.setData({
        userBalance: data.balance
      });
    });
    request.get({
      url: '/sg/router/status',
      success: function (resp) {
        if (resp.code == 0) {
          switch (resp.data.status) {
            case 0: //无行程;
              that.setData({ hasRoute: false });
              break;
            case 1: //已入闸;
             that.setData({hasRoute:true});
              break;
            case 2: //隔天未出闸(异常);
              that.setData({ hasRoute: true });
              break;
            case 4: //已出闸未入闸(异常);
              that.setData({ hasRoute: true });
              break;

          }
        }
      }
    });
  },
  getWallet: function (successCB, failCB) {
    if (wx.showLoading) {
      wx.showLoading({
        title: '处理中',
      });
    }
    request.get({
      url: '/sg/wallet/info',
      success: function (p) {
        if (p.code == 0) {
          successCB(p.data);
        } else {
          util.showMsg("钱包信息获取失败");
        }
      },
      fail: function (fp) {
        console.log(fp)
        if (failCB) {
          failCB(fp);
        };
      },
      complete: function (cp) {
        if (wx.hideLoading) {
          wx.hideLoading();
        }
      }
    });
  },
  getRoute: function (successCB, failCB) {
    request.get({
      url: '/sg/router/status',
      success: function (p) {
        if (p.code == 0) {//"status": 0, // 0-无行程; 1-已入闸; 2-隔天未出闸(异常); 4-已出闸未入闸(异常)
          successCB(p.data);
        };
      },
      fail: function (fp) {
        if (failCB) {
          failCB(fp);
        };
      }
    });
  },
  gateIn: function () {
    var that = this;
    this.getWallet(function (data) {
      if (data.wxpay_quick) {
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
    var that = this;
    this.getWallet(function (data) {
      if (data.wxpay_quick) {
        util.showMsg('有代扣，可进入');
        wx.navigateTo({
          url: '../open/showCode?type=out',
        });
      } else if (data.balance > 0) {
        util.showMsg('有余额，可进入');
        wx.navigateTo({
          url: '../open/showCode?type=out',
        });
      } else {
        that.tipForRecharge('暂不出站');
      }
    });
  },
  recharge:function(){
wx.navigateTo({
  url: '../wallet/recharge',
})
  },
  tipForRecharge: function (cancelTip) {
    wx.showModal({
      title: '提醒',
      content: '没有余额且未签约代扣！',
      showCancel: true,
      cancelText: cancelTip,
      confirmText: "签约充值",
      success: function (res) {
        if (res.confirm) {
          wx.navigateTo({
            url: '../wallet/recharge',
          });
        }
        console.log(res);
      }
    })
  },
  onShareAppMessage: function (page) { },
  bindPickerChange: function (e) {
    var dataSet = ['', 'new', 'hasBalance', 'wxpay_quick'];
    request.init(e.detail.value != 0, dataSet[e.detail.value]);
  },
})
