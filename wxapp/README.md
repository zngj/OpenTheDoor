#weapp

后台返回的凭证为 rsa 加密过的 bytes


混淆算法：

混淆值=
	 2 bytes (随机数)
	 + n bytes (rsa)
	 + 1 byte (1 -入 2-出 0-不知道)
	 + 4 bytes (时间戳,低位在前，高位在后)


二维码加密算法
	aes( 混淆值 , "5454395434473454" ,"6916665466156476");

	

新算法：
	
var evidence =  evidenceFromServer;  //后台返回的凭证
var bytes = base64ToBytes(evidence).concat(byte(入站 ? 1 : 2)); //后台凭证转为byte后，再添加一个进出站byte
var encryptBytes = aes(bytes, key=5454395434473454 , iv = 6916665466156476) ; //aes加密
var []splitBytes = split(encryptBytes,2);//将加密完的bytes分成2段，每个段前面增加一个byte用于标识段数（左移4位将段数置于高位）
var []qr =  createQrCode(splitBytes); //将每一段的bytes分别toBase64,然后加上^前缀及$后缀后生成二维码



二维码逻辑:
	createQrCode: function(splitBytes) {
		for (var i = 0; i < splitBytes.length; i++) {
		  qrImgs[i] = wxqrcode.createQrCodeImg("^" + bytesToBase64(splitBytes[i]) + "$", { 'size': 200 });
		}
	}
分段逻辑
  split: function (wholeBytes, segCount) {
    var byteArr = [];
    var newLen = wholeBytes.length / segCount;
    for (var i = 0; i < segCount; i++) {
      byteArr[i] = [(segCount << 4) + i].concat(wholeBytes.slice(i * newLen, (i == segCount - 1) ? wholeBytes.length : newLen * (i + 1)));
    }
    return byteArr;
  },