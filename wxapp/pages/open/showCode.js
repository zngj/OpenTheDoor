var wxqrcode = require('../../js/wxqrcode.js'); var util = require('../../js/util.js').util;
var request = require('../../js/util.js').request;
var Base64 = require('../../js/base64.js');
var Crypto = require('../../js/cryptojs').Crypto;

// showCode.js
Page({
  data: {
    type: 'in' //进站标识
    , evidence: ''
    , check: true
  },
  onLoad: function (options) {
    this.setData({
      type: options.type
    });
  },
  onShow: function () {
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
  nextPeople: function () {
    var page = this;
    this.getEvidence(function (data) {
      page.data.evidence = data.evidence_key;
      //data.expires_at;
    });
  },
  getEvidence: function (successCB, failCB) {
    var page = this;
    request.get({
      url: '/sg/router/evidence/' + page.data.type,
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
  makeNewQrCode: function () {
    var entryByte = this.data.type == 'in' ? 1 : (this.data.type == 'out' ? 2 : 0);
    this.setData({
      "qrImgs": [
        wxqrcode.createQrCodeImg(this.encrypt(Crypto.util.bytesToBase64(util.mix(Crypto.util.base64ToBytes(this.data.evidence).concat([entryByte]), this.data.type))), { 'size': 200 })
      ]
    });
    setTimeout(this.makeNewQrCode.bind(this), 500);
  },
  getDirectionDesc:function(direction){
    return direction==0?'入闸':'出闸';
  },
  checkNotification: function () {
    var page = this;
    request.get({
      url: '/sg/notification/router',
      success: function (p) {
        if (p.data.code == 0) {
          util.showMsg("您在 " + p.data.data.in_station_name + " " + page.getDirectionDesc(p.data.data.direction) +"成功!");
        // "notification_id": 9,
        // "direction": 0, //入闸
        // "in_gate_id": "010100101",
        // "in_station_id": "0101001",
        // "in_station_name": "五一广场",
        // "in_time": 1502304832
          request.put({
            url: '/sg/notification/consume/'+p.data.data.notification_id
        });
        }
      }, fail: function (fp) {
      }, complete: function () {
        if (page.data.check) {
          setTimeout(page.checkNotification.bind(page), 1500);
        }
      }
    });
  },
  encrypt: function (word) {
    //console.log('before:' + originWord);
    var mode = new Crypto.mode.CBC(Crypto.pad.pkcs7);
    var eb = Crypto.util.base64ToBytes(word);
    var kb = Crypto.charenc.UTF8.stringToBytes("5454395434473454");//KEY
    var vb = Crypto.charenc.UTF8.stringToBytes("6916665466156476");//IV
    var ub = Crypto.AES.encrypt(eb, kb, { iv: vb, mode: mode, asBytes: true });
    //console.log('after:' + word);
    return Crypto.util.bytesToBase64(ub);
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