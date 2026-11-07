const config = require("./config.js");

let SCREEN_WIDTH = 750
let RATE = wx.getSystemSetting().windowHeight / wx.getSystemSetting().windowWidth

App({
  onLaunch: function () {
    this.globalData.apiUrl = config.domain;
    this.globalData.imgUrl = config.imgUrl;
    this.globalData.wss = config.wss;
    this.globalData.active = true;
  },


  onShow: function (params) {
    this.globalData.active = true
  },

  onHide: function (params) {
    this.globalData.active = false
  },


  onUnload: function (params) {
    this.globalData.active = false
  },

  globalData: {
    appId:null,
    userInfo: null,
    apiUrl: null,
    wss: null,
    imgUrl: null,
    color: '0aecc3',
    active : false,
    windowHeight:SCREEN_WIDTH * RATE
  }
})