var wxqrcode = require('../../js/wxqrcode.js'); 
var util = require('../../js/util.js').util;
var request = require('../../js/util.js').request;
var Crypto = require('../../js/cryptojs').Crypto;

Page({
  data: {
    type: 'in' //进站标识
    , typeDesc: ''
    , evidence: ''
    , check: true
    , qrInterval: 200
    , currSeg: -1
    , segCount: 2
    , qrImg: ''
    , qrTime: ''
    , qrImgs: []
    , qrBuf: {}
    , shouldLeave: false
    , leaveTimeout: 15 * 1000
    , autoRefreshQr: false
  },
  onLoad: function (options) {
    this.setData({
      type: options.type,
      typeDesc: 'in' == options.type ? '进站' : '出站',
    });
    wx.setNavigationBarTitle({
      title: "e畅行"+('in' == this.data.type ? '进站二维码' : '出站二维码'),
    });
    if (wx.onUserCaptureScreen) {
      wx.onUserCaptureScreen(function (res) {
        console.log('用户截屏了')
      })
    }
  },
  onShow: function () {
    if (wx.setKeepScreenOn) {
      wx.setKeepScreenOn({ keepScreenOn: true });
    }
    var page = this;
    this.data.check = true;
    this.getEvidence(function (data) {
      page.data.evidence = data.evidence_key;
      //data.expires_at;
      page.makeNewQrCode();
    })
    setTimeout(this.checkNotification.bind(this), 1500);
  },
  onHide: function () {
    this.data.check = false;
  },
  onUnload: function () {
    this.data.check = false;
  },
  nextPeople: function () {
    var page = this;
    this.getEvidence(function (data) {
      page.data.shouldLeave = false;
      page.data.evidence = data.evidence_key;
      page.data.currSeg = -1;
      page.data.qrImgs = []
      //data.expires_at;
    });
  },
  getEvidence: function (successCB, failCB) {
    var page = this;
    request.get({
      url: '/sg/router/evidence/' + page.data.type,
      success: function (p) {
        if (p.code == 0) {
          successCB(p.data);
        }
      },
      fail: function (fp) {
        if (failCB) {
          failCB(fp);
        };
      }
    });
  },
  split: function (wholeBytes, segCount) {
    var byteArr = [];
    var newLen = wholeBytes.length / segCount;
    for (var i = 0; i < segCount; i++) {
      byteArr[i] = [(segCount << 4) + i].concat(wholeBytes.slice(i * newLen, (i == segCount - 1) ? wholeBytes.length : newLen * (i + 1)));
    }
    return byteArr;
  },
  makeNewQrCode: function () {
    var page = this;
    if (this.data.currSeg >= 0 && this.data.currSeg < this.data.segCount - 1) {
      this.setData({
        "qrImg": this.data.qrImgs[++this.data.currSeg]
      });
    } else {
      if (!this.data.autoRefreshQr && this.data.qrImgs.length > 0) {
        page.setData({
          "qrImg": this.data.qrImgs[0]
          , "currSeg": 0
        });
      } else {
        this.getNextQrImgs(this.data.evidence, function (qrImgs) {
          page.setData({
            "qrImgs": qrImgs
            , "qrImg": qrImgs[0]
            , "currSeg": 0
            , "qrTime": new Date().getTime() % 100000
          });
        });
      }
    }
    if (this.data.check) {
      setTimeout(this.makeNewQrCode.bind(this), this.data.qrInterval);
    }
  },
  getNextQrImgs(evidence, callback) {
    var start = new Date().getTime();
    if (this.data.evidence != this.data.qrBuf.evidence) {
      this.data.qrBuf.evidence = this.data.evidence;
      this.data.qrBuf.imgArrays = [];
    }
    if (this.data.qrBuf.imgArrays.length > 0) {
    } else {
      this.data.qrBuf.imgArrays.push(this.generateNewQrCode(this.data.evidence));
    }
    if (callback) {
      callback(this.data.qrBuf.imgArrays.pop());
    }
    if (this.data.autoRefreshQr) {
      setTimeout(this.generateNewQrBuffer.bind(this), 1);
    }
    //console.log("newCode:" + (new Date().getTime()-start))
  },
  generateNewQrBuffer: function () {
    var evidence = this.data.evidence;
    var newCode = this.generateNewQrCode(evidence);
    if (this.data.qrBuf.evidence == evidence) {
      this.data.qrBuf.imgArrays.push(newCode);
    }
  },
  generateNewQrCode: function (evidence) {
    var entryByte = this.data.type == 'in' ? 1 : (this.data.type == 'out' ? 2 : 0);
    var wholeBytes = util.mix(Crypto.util.base64ToBytes(evidence).concat([entryByte]));
    var encrytBytes = this.encrypt(Crypto.util.bytesToBase64(wholeBytes));
    var byteArr = this.split(encrytBytes, this.data.segCount);
    var qrImgs = [];
    for (var i = 0; i < byteArr.length; i++) {
      var s = new Date().getTime();
      qrImgs[i] = wxqrcode.createQrCodeImg("^" + Crypto.util.bytesToBase64(byteArr[i]) + "$", { 'size': 200 });
      var e = new Date().getTime();
      console.log("create" + (e - s));
    }
    return qrImgs;
  },
  getStationDesc: function (data) {
    if (data.direction == 0) {
      return data.in_station_name + " " + '入闸';
    } else {
      return data.out_station_name + " " + '出闸';
    }
  },
  checkNotification: function () {
    var page = this;
    request.get({
      url: '/sg/notification/router',
      success: function (p) {
        if (p.code == 0) {
          util.showMsg("您在 " + page.getStationDesc(p.data) + "成功!");
          // "notification_id": 9,
          // "direction": 0, //入闸
          // "in_gate_id": "010100101",
          // "in_station_id": "0101001",
          // "in_station_name": "五一广场",
          // "in_time": 1502304832
          request.put({
            url: '/sg/notification/consume/' + p.data.notification_id
          });
          page.data.shouldLeave = true;
          setTimeout(page.leavePage.bind(page), page.data.leaveTimeout);
        }
      }, fail: function (fp) {
      }, complete: function () {
        if (page.data.check) {
          setTimeout(page.checkNotification.bind(page), 1500);
        }
      }
    });
  },
  leavePage: function () {
    if (this.data.shouldLeave) {
      if (wx.reLaunch) {
        wx.reLaunch({
          url: '/pages/index/index'
        });
      }
    }
  },
  encrypt: function (word) {
    //console.log('before:' + originWord);
    var mode = new Crypto.mode.CBC(Crypto.pad.pkcs7);
    var eb = Crypto.util.base64ToBytes(word);
    var kb = Crypto.charenc.UTF8.stringToBytes("5454395434473454");//KEY
    var vb = Crypto.charenc.UTF8.stringToBytes("6916665466156476");//IV
    var ub = Crypto.AES.encrypt(eb, kb, { iv: vb, mode: mode, asBytes: true });
    //console.log('after:' + word);
    //return Crypto.util.bytesToBase64(ub);
    return ub;//return byte
  },
  decrypt: function (word) {
    var mode = new Crypto.mode.CBC(Crypto.pad.pkcs7);
    var eb = Crypto.util.base64ToBytes(word);
    var kb = Crypto.charenc.UTF8.stringToBytes("5454395434473454");//KEY
    var vb = Crypto.charenc.UTF8.stringToBytes("6916665466156476");//IV
    var ub = Crypto.AES.decrypt(eb, kb, { asBytes: true, mode: mode, iv: vb });
    return Crypto.util.bytesToBase64(ub);
  }
})