function formatTime(date) {
  var year = date.getFullYear()
  var month = date.getMonth() + 1
  var day = date.getDate()

  var hour = date.getHours()
  var minute = date.getMinutes()
  var second = date.getSeconds()


  return [year, month, day].map(formatNumber).join('-') + ' ' + [hour, minute, second].map(formatNumber).join(':')
}

function formatNumber(n) {
  n = n.toString()
  return n[1] ? n : '0' + n
}

function trimSure(str) {
  return (str == null || str == undefined) ? '' : str.replace(/(^\s*)|(\s*$)/g, '');
}

function redirectTo(url) {
  wx.redirectTo({
    url: url
  })
}

function formatMoney(money) {
  if (money == null || money == undefined) {
    return '0.00';
  }
  var floatMoney = parseFloat(money);
  return floatMoney.toFixed(2);
}

function isMobile(mobile) {
  if (!mobile) {
    return false;
  }
  var reg = /^(((13[0-9]{1})|(14[0-9]{1})|(15[0-9]{1})|(17[0-9]{1})|(18[0-9]{1}))+\d{8})$/;
  if (!reg.test(mobile)) {
    return false;
  }
  return true;
}

function isValidCardNO(id) {
  if (id == null) {
    return false;
  }
  id = trimSure(id).toUpperCase();
  if (id.length != 18) {
    return false;
  }
  var xnum = 0;
  // 1.将身份证号码前面的17位数分别乘以不同的系数。
  // 从第一位到第十七位的系数分别为：7 9 10 5 8 4 2 1 6 3 7 9 10 5 8 4 2
  var factor = [7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2 ];
  var sum = 0;
  for (var i = 0; i < factor.length; i++) {
    var ch = id.charAt(i);
    if ((ch < '0' || ch > '9') && ch != 'X') {
      return false;
    }
    if (ch == 'X') {
      return false;
    }
    // 2.将这17位数字和系数相乘的结果相加。
    sum += parseInt(ch) * factor[i];
  } // end for (int i = 0; i < factor.length; i++)
  if (id.charAt(id.length - 1) == 'X') {
    xnum++;
  }
  if (xnum > 1) {
    // 超过两个x
    return false;
  }
  // 3.用加出来和除以11，看余数是多少？
  var mod = sum % 11;
  // 4.余数只可能有0 1 2 3 4 5 6 7 8 9 10这11个数字。
  // 余数分别对应的最后一位身份证的号码为1 0 X 9 8 7 6 5 4 3 2。
  var suffix = [ '1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2' ];
  return suffix[mod] == id.charAt(id.length - 1);
}

module.exports = {
  formatTime: formatTime,
  trimSure: trimSure,
  redirectTo: redirectTo,
  formatMoney: formatMoney,
  isMobile: isMobile,
  isValidCardNO: isValidCardNO
}

