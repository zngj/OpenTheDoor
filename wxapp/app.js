var util = require('js/util.js').util;
var request = require('js/util.js').request;
App({
  data: {
  },
  onLaunch: function (options) {
    //console.log("App onLaunch");
    request.init(false);
    this.ensureSession();
  },
  onShow: function () {
    //console.log("App on show");
    this.ensureSession();
    //console.log("CurrentPages:");
    //console.log(getCurrentPages());
    this.initSocket();
  },
  ensureSession: function () {
    var that = this;
    wx.checkSession({
      success: function (sp) {
        request.get({
          url: '/user/check_token',
          success: function (p) {
            if (p.code != 0) { // token invalid
              that.login('TokenInvalid');
            }
          },
          fail: function (c) {
            that.login('verifyFail');
          }
        });
      },
      fail: function (fp) {
        that.login('SessionFail');
      }
    });
  },
  login: function (entry) {
    var app = this;
    if (app.loging) {
      return;
    }
    app.loging = true;
    console.log("start login");
    wx.login({
      success: function (p) {
        request.request({
          url: '/user/wxapp/login',
          data: { code: p.code },
          success: function (loginResult) {
            if (loginResult.code == 0) {
              wx.setStorageSync('token', loginResult.data.access_token);
              wx.reLaunch({
                url: '/pages/index/index',
              });
            } else {
              util.showMsg(loginResult.msg);
            };
          }, complete: function (pp) {
            app.loging = false;
            console.log(entry);
          }
        });
      },
      complete: function (p) {
        console.log(entry);
      }
    });
  },
  connect: function () {
    wx.connectSocket({
      url: 'wss://sgu.youstars.com.cn/ws',
      protocols: [wx.getStorageSync('token')],
      method: "GET",
      complete: function (c) {
        console.log(c);
      }
    });
  },
  initSocket: function () {
    var app = this;
    app.connect();
    app.heartBeat();
    wx.onSocketOpen(function (res) {
      console.log('WebSocketOpen！');
    });
    wx.onSocketError(function (res) {
      console.log('WebSocket连接打开失败，请检查！')
    });
    wx.onSocketMessage(function (res) {
      console.log('收到服务器内容：' + res.data);
      try {
        if (res.data !== "Pong") {
          var json = JSON.parse(res.data);
            app.handleMessage(json);
        }
      } catch (e) {
        console.log(e);
      }
    });
    wx.onSocketClose(function (res) {
      console.log('WebSocket Close！' )
      console.log(res)
      setTimeout(app.connect,1000);
    })
  },
  heartBeat:function(){
    console.log("Ping")
    wx.sendSocketMessage({
      data: "Ping"
    });
    setTimeout(this.heartBeat,45000);
  },
  handleMessage:function(msg){
    if(msg.cmd===101&& msg.body.code===1000){//登录失效
      this.ensureSession();
    } else{
      wx.sendSocketMessage({
        data: JSON.stringify({ cmd: json.cmd, id: json.body.data.id })
      });
    }

    var pages = getCurrentPages();
    if (pages.length > 0 ) {
      var currPage = pages[pages.length - 1];
      if (currPage && currPage.handleMessage){
        currPage.handleMessage(msg);
      }
    }
  }

});
