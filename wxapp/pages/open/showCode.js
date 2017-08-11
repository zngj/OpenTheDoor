var wxqrcode = require('../../js/wxqrcode.js');
var aes = require('../../js/aes.js');
var util = require('../../js/util.js');
var Base64 = require('../../js/base64.js'); 
var Crypto = require('../../js/cryptojs').Crypto;

// showCode.js
Page({
  data: {
    type: 'in' //进站标识
    
  },

  onLoad: function (options) {
    this.setData({
      type: options.type
    });
    var page = this;
    this.getEvidence(function (data) {
      page.makeCode(data.evidence_key);
      //page.makeCode("12345678901234567890123456789012");
      //data.expires_at;
    });
  },

  getEvidence: function (successCB, failCB) {
    var page=this;
    var token = wx.getStorageSync('token');
    util.request({
      url: '/sg/router/evidence/'+page.data.type,
      method: 'GET',
      header: { 'Access-Token': token },
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
  mix:function(rawData){
    var bytes = Crypto.util.base64ToBytes(rawData);
    var newBytes= [getRandom(), getRandom()];
    newBytes=newBytes.concat(bytes);

  },
  makeCode: function (rawData) {
    this.setData({
      "qrImgs": [
        wxqrcode.createQrCodeImg(this.encrypt(Crypto.util.bytesToBase64(util.util.mix(Crypto.util.base64ToBytes(rawData)))), { 'size': 200 }),
        wxqrcode.createQrCodeImg(this.encrypt(Crypto.util.bytesToBase64(util.util.mix(Crypto.util.base64ToBytes(rawData)))), { 'size': 200 }),
        wxqrcode.createQrCodeImg(this.encrypt(Crypto.util.bytesToBase64(util.util.mix(Crypto.util.base64ToBytes(rawData)))), { 'size': 200 })
    ]});
  },
  nextPeople: function () {
    var page = this;
    this.getEvidence(function (data) {
      page.makeCode(data.evidence_key);
      //data.expires_at;
    });
  },
  encrypt: function (originWord) {
    console.log('before:' + originWord);
    var key = aes.CryptoJS.enc.Utf8.parse("5454395434473454");   //十六位十六进制数作为秘钥
    var iv = aes.CryptoJS.enc.Utf8.parse('6916665466156476');  //十六位十六进制数作为秘钥偏移量
    var srcs = aes.CryptoJS.enc.Base64.parse(originWord);
    var encrypted = aes.CryptoJS.AES.encrypt(srcs, key, { iv: iv, mode: aes.CryptoJS.mode.CBC, padding: aes.CryptoJS.pad.Pkcs7 });
    var word = encrypted.ciphertext.toString(aes.CryptoJS.enc.Base64);
    console.log('after:' + word);
    return word;
  },
  decrypt: function (word) {
    var key = aes.CryptoJS.enc.Utf8.parse("5454395434473454");   //十六位十六进制数作为秘钥
    var iv = aes.CryptoJS.enc.Utf8.parse('6916665466156476');  //十六位十六进制数作为秘钥偏移量
    var encryptedHexStr = aes.CryptoJS.enc.Base64.parse(word);
    var srcs = aes.CryptoJS.enc.Base64.stringify(encryptedHexStr);
    console.log(srcs);
    var decrypt = aes.CryptoJS.AES.decrypt(srcs, key, { iv: iv, mode: aes.CryptoJS.mode.CBC, padding: aes.CryptoJS.pad.Pkcs7 });
    var decryptedStr = decrypt.toString(aes.CryptoJS.enc.Base64);
    return decryptedStr.toString();
  },
  Encrypt : function (word) {
    var mode = new Crypto.mode.CBC(Crypto.pad.pkcs7);
    var eb = Crypto.charenc.UTF8.stringToBytes(word);
    var kb = Crypto.charenc.UTF8.stringToBytes("5454395434473454");//KEY
    var vb = Crypto.charenc.UTF8.stringToBytes("6916665466156476");//IV
    var ub = Crypto.AES.encrypt(eb, kb, { iv: vb, mode: mode, asBpytes: true });
    return ub;
  },
  Decrypt: function (word) {
    var mode = new Crypto.mode.CBC(Crypto.pad.pkcs7);
    var eb = Crypto.util.base64ToBytes(word);
    var kb = Crypto.charenc.UTF8.stringToBytes("5454395434473454");//KEY
    var vb = Crypto.charenc.UTF8.stringToBytes("6916665466156476");//IV
    var ub = Crypto.AES.decrypt(eb, kb, { asBpytes: true, mode: mode, iv: vb });
    return ub;
  }
})