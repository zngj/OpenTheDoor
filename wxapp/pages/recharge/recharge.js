// recharge.js
var util = require('../../js/util.js').util;
var request = require('../../js/util.js').request;
var app = getApp()
Page({

  /**
   * 页面的初始数据
   */
  data: {
    confirmRecharge:0,
    rechargeMoney: 0
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onShow: function () {
    var that = this;
    this.setData({
      recharge:0,
      confirmRecharge: 0,
      userBalance: '--',
      rechargeSolution: [
        { "sol1": 10,"sol2":50,"sol3":100 },
      ]
    });
  },
  rechargeSelected:function(e){
    this.setData({
      recharge:e.target.dataset.money
    });
    this.data.rechargeMoney = e.target.dataset.money;
  },
  confirmRecharge:function(e){
    var that = this;
    this.setData({
      confirmRecharge: 1
    });
  },
  payRecharge:function(e){
    
    var rechargeMoney = this.data.rechargeMoney;
    wx.login({
      success: function(res) {
        request.request({
          url: '/user/wxapp/rechargeMakeOrder',
          data: {
            idaccess_token: wx.getStorageInfoSync('token'),
            rechargeMoney: rechargeMoney,
            jsCode: res.code
          },
          success: function(res) {
            var data = res.data.data;
            console.log(res.data.code);
            if(res.data.code ==-1 || data==null || data.paySign==null || !data || !data.paySign){
              wx.showToast({
                title: res.data.msg,
              });
              return;
            }
            wx.requestPayment({
              timeStamp: data.timeStamp,
              nonceStr: data.nonceStr,
              package: data.package,
              signType: data.signType,
              paySign: data.paySign,
              success: function (res) {
                wx.showToast({
                  title: '支付成功',
                });
                wx.navigateBack({
                  url: '/pages/index/index',
                });
              },
              fail: function (res) {
                wx.showToast({
                  title: '支付失败',
                });
              }
            });
          },
          fail: function(res) {
            wx.showToast({
              title: '请登录重试',
            })
          },
          complete: function(res) {},
        })
      },
      fail: function(res) {},
      complete: function(res) {},
    })
  }, //payRecharge

});
