//var qr = require('../../js/qr.js')
var wxqrcode = require('../../js/wxqrcode.js')
var aes = require('../../js/aes.js')

// showCode.js
Page({
  data: {
    type: 'in' //进站标识
    
  },

  onLoad: function (options) {
    this.setData({
      type: options.type
    });

    this.makeCode(this.getRandom(12));
  },

  makeCode: function (rawData) {
    this.setData({
      "qrImgs": [
        wxqrcode.createQrCodeImg(this.encrypt(this.getRandom(2) + rawData + ":" + new Date().getTime() + this.getRandom(2)), { 'size': 200 })
        ,wxqrcode.createQrCodeImg(this.encrypt(this.getRandom(2) + rawData + ":" + new Date().getTime() + this.getRandom(2)), { 'size': 200 })
        ,wxqrcode.createQrCodeImg(this.encrypt(this.getRandom(2) + rawData + ":" + new Date().getTime() + this.getRandom(2)), { 'size': 200 })
    ]});
  },
  getRandom(len){
    return (Math.random() + '').replace("0.", "").substr(0, len);
  },
 
  encrypt: function (word) {
    var key = aes.CryptoJS.enc.Utf8.parse("3454345434543454");   //十六位十六进制数作为秘钥
    var iv = aes.CryptoJS.enc.Utf8.parse('6666666666666666');  //十六位十六进制数作为秘钥偏移量
    var srcs = aes.CryptoJS.enc.Utf8.parse(word);
    var encrypted = aes.CryptoJS.AES.encrypt(srcs, key, { iv: iv, mode: aes.CryptoJS.mode.CBC, padding: aes.CryptoJS.pad.Pkcs7 });
    return encrypted.ciphertext.toString().toUpperCase();
  },
  decrypt: function (word) {
    var key = aes.CryptoJS.enc.Utf8.parse("3454345434543454");   //十六位十六进制数作为秘钥
    var iv = aes.CryptoJS.enc.Utf8.parse('6666666666666666');  //十六位十六进制数作为秘钥偏移量
    var encryptedHexStr = aes.CryptoJS.enc.Hex.parse(word);
    var srcs = aes.CryptoJS.enc.Base64.stringify(encryptedHexStr);
    var decrypt = aes.CryptoJS.AES.decrypt(srcs, key, { iv: iv, mode: aes.CryptoJS.mode.CBC, padding: aes.CryptoJS.pad.Pkcs7 });
    var decryptedStr = decrypt.toString(aes.CryptoJS.enc.Utf8);
    return decryptedStr.toString();
  }
})