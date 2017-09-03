// find.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    "longitude": 112.984631,
    "latitude": 28.194967,
    scale: 11,
    polyline: [{
      points: [
        {
          latitude: 28.148049,
          longitude: 113.085469
        }
        , {

          latitude: 28.14639,
          longitude: 113.05686
        }
      ],

      color: "#FF0000DD",
      width: 2,
      dottedLine: false
    }],
    includePoints: [],
    markers: [{
      iconPath: "/images/my.png",
      id: 0,
      latitude: 28.148049,
      longitude: 113.085469,
      width: 20,
      height: 20
    }],
    controls: [{
      id: 1,
      iconPath: '/images/recharge.png',
      position: {
        left: 0,
        top: 300 - 50,
        width: 20,
        height: 20
      },
      clickable: true
    }],
    subway: [
      {
        name:'地铁1号线',
        stations: [
          { "name": "开福区政府", "longitude": "112.985325", "latitude": "28.261984", "uid": "11247311063686400543" },
          { "name": "马厂", "longitude": "112.988894", "latitude": "28.251105", "uid": "1475835411160154032" },
          { "name": "北辰三角洲", "longitude": "112.983229", "latitude": "28.235642", "uid": "15699920661210282382" },
          { "name": "开福寺", "longitude": "112.979697", "latitude": "28.222389", "uid": "11614230111235277703" },
          { "name": "文昌阁", "longitude": "112.978174", "latitude": "28.212604", "uid": "1722775984414661037" },
          { "name": "培元桥", "longitude": "112.976455", "latitude": "28.206193", "uid": "1485693782423632739" },
          { "name": "五一广场", "longitude": "112.976474", "latitude": "28.194772", "trans_lines": "地铁2号线;", "uid": "3085815401333131216" },
          { "name": "黄兴广场", "longitude": "112.976371", "latitude": "28.189436", "uid": "13015907022395028533" },
          { "name": "南门口", "longitude": "112.976178", "latitude": "28.183623", "uid": "11491402381527529330" },
          { "name": "侯家塘", "longitude": "112.982415", "latitude": "28.173705", "uid": "16675848450185574099" },
          { "name": "南湖路", "longitude": "112.985542", "latitude": "28.165997", "uid": "14083744060577101567" },
          { "name": "黄土岭", "longitude": "112.985546", "latitude": "28.160224", "uid": "11779576342339159644" },
          { "name": "涂家冲", "longitude": "112.984661", "latitude": "28.153198", "uid": "4598144884783373670" },
          { "name": "铁道学院", "longitude": "112.987204", "latitude": "28.136235", "uid": "17785856099803378425" },
          { "name": "友谊路", "longitude": "112.986816", "latitude": "28.123802", "uid": "9190923104948045027" },
          { "name": "省政府·清风", "longitude": "112.987500", "latitude": "28.111274", "uid": "6898474052995847997" },
          { "name": "桂花坪", "longitude": "112.988329", "latitude": "28.099194", "uid": "17087974960429038598" },
          { "name": "大托", "longitude": "112.989222", "latitude": "28.082624", "uid": "139518813390161443" },
          { "name": "中信广场", "longitude": "112.989729", "latitude": "28.073688", "uid": "13239681101490185556" },
          { "name": "尚双塘", "longitude": "112.992345", "latitude": "28.062288", "uid": "17135243337517493928" }
        ]
      },
      {
        'name': '地铁2号线',
        'stations': [
          { "name": "梅溪湖西", "longitude": 112.881827, "latitude": 28.185774, "uid": "11659242893870332999" }
          , { "name": "麓云路", "longitude": 112.892211, "latitude": 28.190743, "uid": "7721243764382822163" }
          , { "name": "文化艺术中心", "longitude": 112.900330, "latitude": 28.196644, "uid": "9239846912639375797" }
          , { "name": "梅溪湖东", "longitude": 112.907976, "latitude": 28.201830, "uid": "4997840561799352191" }
          , { "name": "望城坡", "longitude": 112.914823, "latitude": 28.207478, "uid": "10266310351938291968" }
          , { "name": "金星路", "longitude": 112.928475, "latitude": 28.206031, "uid": "12615443586153310111" }
          , { "name": "西湖公园", "longitude": 112.940163, "latitude": 28.202551, "uid": "2166578440427736594" }
          , { "name": "溁湾镇", "longitude": 112.951420, "latitude": 28.197627, "uid": "12223607536123808274" }
          , { "name": "橘子洲", "longitude": 112.962543, "latitude": 28.195696, "uid": "5290261544507460275" }
          , { "name": "湘江中路", "longitude": 112.969346, "latitude": 28.195251, "uid": "13911801601409075896" }
          , { "name": "五一广场", "longitude": 112.976456, "latitude": 28.195298, "trans_lines": "地铁1号线;", "uid": "2399355950936030474" }
          , { "name": "芙蓉广场", "longitude": 112.984631, "latitude": 28.194967, "uid": "4070001542621318027" }
          , { "name": "迎宾路口", "longitude": 112.992091, "latitude": 28.194825, "uid": "10210138934798551043" }
          , { "name": "袁家岭", "longitude": 113.000682, "latitude": 28.194488, "uid": "7453847862944904106" }
          , { "name": "长沙火车站", "longitude": 113.010917, "latitude": 28.193723, "uid": "1332072884274795442" }
          , { "name": "锦泰广场", "longitude": 113.017599, "latitude": 28.191957, "uid": "7046742010770352847" }
          , { "name": "万家丽广场", "longitude": 113.030456, "latitude": 28.191522, "uid": "2281018118243600345" }
          , { "name": "人民东路", "longitude": 113.038747, "latitude": 28.184234, "uid": "3399889797179418479" }
          , { "name": "长沙大道", "longitude": 113.044422, "latitude": 28.168146, "uid": "6058723314550796468" }
          , { "name": "沙湾公园", "longitude": 113.044847, "latitude": 28.157682, "uid": "5569163928479260326" }
          , { "name": "杜花路", "longitude": 113.056862, "latitude": 28.146388, "uid": "16929786566968102070" }
          , { "name": "长沙火车南站", "longitude": 113.065220, "latitude": 28.147177, "uid": "5166232884659095603" }
          , { "name": "光达", "longitude": 113.085469, "latitude": 28.148049, "uid": "14727431796250397646" }
        ]
      }
    ]
  },
  controltap(e) {
    console.log(e.controlId)
  },
  markertap(e) {
    console.log(e.markerId)
  },
  locate: function () {
    wx.chooseLocation({
      success: function (s) {
        console.log(s);
      }
    });
    // wx.openLocation({
    //   latitude: 28.14639,
    //   longitude: 113.05686
    // })
  },
  scaleUp: function () {
    if (this.data.scale < 18) {
      this.setData({ scale: this.data.scale + 1 });
    }
    // 5~18
  },
  scaleDown: function () {
    if (this.data.scale > 5) {
      this.setData({ scale: this.data.scale - 1 });
    }
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this.setData({
      polyline: [{
        points: this.data.subway[1].stations,
        color: "#FF0000DD",
        width: 2,
        dottedLine: false
      }, {
        points: this.data.subway[0].stations,
        color: "#0000FFDD",
        width: 2,
        dottedLine: false
      }]
      , includePoints: this.data.subway[0].stations
    });

  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})