//app.js
App({
  globalData: {
    userInfo: null,
    name:'DDD'
  },
  onLaunch: function (options) {

    this.getUserInfo();
  },
 
  getUserInfo:function(cb){
    var that = this
    if(this.globalData.userInfo){
      typeof cb == "function" && cb(this.globalData.userInfo)
    }else{
      //调用登录接口
      wx.login({
        success: function (loginResp) {
          //loginResp.code;
          wx.getUserInfo({
            success: function (res) {
              that.globalData.userInfo = res.userInfo
              typeof cb == "function" && cb(that.globalData.userInfo)
            }
          })
        }
      })
    }
  }
})