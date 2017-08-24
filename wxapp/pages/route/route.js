// pages/route/route.js
var request = require('../../js/util.js').request;
Page({

  /**
   * 页面的初始数据
   */
  data: {
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
  
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {
  
  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {
    this.updateRouteList();
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {
  
  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {
  
  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {
  
  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {
  
  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {
  
  },
  updateRouteList: function (lastId) {
    var page = this;
    request.get({
      url: "/sg/router/list" + (lastId ? "?last_id=" + lastId:""),
      data: {},
      success: function (s) {
        if (lastId){
          page.setData({ "currentRoutes": page.data.currentRoutes.concat(s.data) });
        } else {
          page.setData({ "currentRoutes": s.data });
        
        }
      },
      fail: function () { }
    });

  },
  onReachBottom:function(){
    this.updateRouteList(this.data.currentRoutes[this.data.currentRoutes.length - 1].id);
  }
})