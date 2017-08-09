var wxqrcode = require('../../js/wxqrcode.js');
var aes = require('../../js/aes.js');
var util = require('../../js/util.js');
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
  makeCode: function (rawData) {
    console.log(rawData.length);
    console.log(rawData);
    this.setData({
      "qrImgs": [
         wxqrcode.createQrCodeImg(this.encrypt(this.getRandom(2) + rawData + "=" + new Date().getTime() ), { 'size': 200 })
        ,wxqrcode.createQrCodeImg(this.encrypt(this.getRandom(2) + rawData + "=" + new Date().getTime() ), { 'size': 200 })
        ,wxqrcode.createQrCodeImg(this.encrypt(this.getRandom(2) + rawData + "=" + new Date().getTime() ), { 'size': 200 })
    ]});
  },
  nextPeople: function () {
    var page = this;
    this.getEvidence(function (data) {
      page.makeCode(data.evidence_key);
      //data.expires_at;
    });
  },
  getRandom(len){
    var s="";
    for(var i=0;i<len;i++){
      s+=String.fromCharCode("A".charCodeAt(0) + Math.ceil(Math.random() * 25))
    }
    return s;
  },
 
  encrypt: function (word) {
    console.log(word);
    var key = aes.CryptoJS.enc.Utf8.parse("5454395434473454");   //十六位十六进制数作为秘钥
    var iv = aes.CryptoJS.enc.Utf8.parse('6916665466156476');  //十六位十六进制数作为秘钥偏移量
    var srcs = aes.CryptoJS.enc.Base64.parse(word);
    //var srcs = aes.CryptoJS.enc.Utf8.parse(word);
    var encrypted = aes.CryptoJS.AES.encrypt(srcs, key, { iv: iv, mode: aes.CryptoJS.mode.CBC, padding: aes.CryptoJS.pad.Pkcs7 });
    var word = encrypted.ciphertext.toString(aes.CryptoJS.enc.Base64);
    console.log(word.length);
    console.log(word);
    return word;
  },
  decrypt: function (word) {
    var key = aes.CryptoJS.enc.Utf8.parse("5454395434473454");   //十六位十六进制数作为秘钥
    var iv = aes.CryptoJS.enc.Utf8.parse('6916665466156476');  //十六位十六进制数作为秘钥偏移量
    //var encryptedHexStr = aes.CryptoJS.enc.Hex.parse(word);
    var encryptedHexStr = aes.CryptoJS.enc.Base64.parse(word);
    var srcs = aes.CryptoJS.enc.Base64.stringify(encryptedHexStr);
    console.log(srcs);
    var decrypt = aes.CryptoJS.AES.decrypt(srcs, key, { iv: iv, mode: aes.CryptoJS.mode.CBC, padding: aes.CryptoJS.pad.Pkcs7 });
    //var decryptedStr = decrypt.toString(aes.CryptoJS.enc.Utf8);
    var decryptedStr = decrypt.toString(aes.CryptoJS.enc.Base64);
    return decryptedStr.toString();
  }
})