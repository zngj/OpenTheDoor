#weapp

二维码加密算法
	aes(2位随机字母+服务端返回的凭证+"="+二维码生成时间戳, "5454395434473454" ,"6916665466156476");

说明:
	rawData: 服务端返回的凭证
	var word = this.getRandom(2) + rawData + ":" + new Date().getTime()
	createQrCodeImg(encrypt(word));


	  encrypt: function (word) {
		var key = aes.CryptoJS.enc.Utf8.parse("5454395434473454");   //十六位十六进制数作为秘钥
		var iv = aes.CryptoJS.enc.Utf8.parse('6916665466156476');  //十六位十六进制数作为秘钥偏移量
		var srcs = aes.CryptoJS.enc.Base64.parse(word);
		var encrypted = aes.CryptoJS.AES.encrypt(srcs, key, { iv: iv, mode: aes.CryptoJS.mode.CBC, padding: aes.CryptoJS.pad.Pkcs7 });
		var word = encrypted.ciphertext.toString(aes.CryptoJS.enc.Base64);
		console.log(word.length);
		return word;
	  }