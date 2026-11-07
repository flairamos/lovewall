const http = require("./../../../utils/http.js");
const util = require("./../../../utils/util.js");
const app = getApp()
Page({
  data: {
    // chatInfo: null,
    chatInfo: { // 对方头像信息
      // id: 'he',
      // name: '他',
      // avatar: ''
    },
    userInfo: { // 我方头像信息
      // id: 'self',
      // name: '我',
      // avatar: ''
    },
    messages: [],
    inputMessage: '',
    scrollToMessage: '',
    page: 1,
    limit: 10,
    timer: null
  },

  onLoad: function (options) {
    // 获取聊天对象信息
    this.getChatInfo(options.id)
    // 获取自己的头像信息
    this.getItSelfInfo(options.id)
    // 获取历史消息
    this.getHistoryMessages(options.id)
  },


  onShow: function (params) {
    // this.data.timer = setInterval(() => {
      this.getMoreHistoryMessages(this.data.chatInfo.id)
      // this.clearUnread(this.data.chatInfo.id)
    // }, 3000);
  },

  onHide: function (params) {
    clearInterval(this.data.timer);
  },

  getItSelfInfo() {
    let user = wx.getStorageSync('user')
    let userInfo = {
      id: user.id,
      name: user.name,
      avatar: user.avatar
    }
    this.setData({
      userInfo: userInfo
    })
  },

  getChatInfo: function (chatId) {
    // 模拟获取聊天对象信息
    http.get(`/user/getbyid?id=${chatId}`, '', (res) => {
      if (res.data.code) {
        this.setData({
          chatInfo: res.data.data
        })
      } else {
        wx.showToast({
          title: '获取用户信息失败',
          icon: 'none'
        })
      }
    })
    // const chatInfo = {
    //   id: chatId,
    //   name: '张三',
    //   avatar: '/images/avatar1.png'
    // }
    // this.setData({ chatInfo })
  },

  getHistoryMessages: function (user_id) {
    let lastTime = this.getDateZero();
    let that = this
    http.get(`/message/unread_message?opponentId=${user_id}&lastMessageTime=${lastTime}`, '', (res) => {
      if (res.data.code) {
        that.setData({
          messages: res.data.data
        })
      }
    })
  },

  getDateZero() {
    let now = util.formatTime(new Date())
    let str = now.split(" ")[0]
    return str + " " + '00:00:00'
  },


  clearUnread() {
    let arrids = []
    this.data.messages.forEach(item => {
      arrids.push(item.id)
    })
    http.post(`/message/clear_unread_message?user_id=${this.data.chatInfo.id}`, {
      ids: arrids
    }, (res) => {
      if (res.data.code) {} else {
        wx.showToast({
          title: '消除未读消息失败',
          icon: 'none'
        })
      }
    })
  },


  getMoreHistoryMessages: function (user_id) {
    if(!this.data.messages.length){ return }
    let lastArr = this.data.messages.at(-1)
    let that = this
    http.get(`/message/unread_message?opponentId=${user_id}&lastMessageTime=${lastArr.time}`, '', (res) => {
      if (res.data.code) {
        if (res.data.data != null) {
          let arr = []
          res.data.data.forEach(element => {
            if (element.id > lastArr.id) {
              arr.push(element)
            }
          });
          let allData = that.data.messages
          allData.push(...arr)
          this.setData({
            messages: allData
          })
        }
      }
    })
  },

  onInputChange: function (e) {
    this.setData({
      inputMessage: e.detail.value
    })
  },


  // 发送消息
  // sendMessageSocketServer(){
  //   let that = this
  //   wx.connectSocket({
  //     url: this.data.wss+"/ws",
  //     header:{
  //       'Authorization': 'Bearer '+ wx.getStorageSync('token')
  //     },
  //     success: function (res) {
  //       console.log('WebSocket连接成功',res)
  //     },
  //     fail: function (res) {
  //       console.log('WebSocket连接失败:', res)
  //     }
  //   })
  //   wx.onSocketOpen(function (res) {
  //     console.log('WebSocket连接已打开',res)
  //   })
  //   wx.onSocketError(function (err) {
  //     console.log('WebSocket连接打开失败:', err)
  //   })
  //   // wx.sendSocketMessage({
  //   //   data: 'Hello, WebSocket!'
  //   // })
  //   wx.onSocketMessage(function (res) {
  //     // console.log('接收到服务器发送的数据:', res.data)
  //     // console.log('接收到服务器发送的数据:', typeof(res.data))
  //     let data = JSON.parse(res.data)
  //     that.setData({
  //       messages: data
  //     })
  //   })
  //   // wx.onSocketClose(function (res) {
  //   //   console.log('WebSocket连接已关闭:', res)
  //   // })
  // },

  onSendMessage: function () {
    if (!this.data.inputMessage.trim()) return
    const newMessage = {
      id: this.data.messages.length + 1,
      content: this.data.inputMessage,
      time: util.formatTime(new Date()),
      isSelf: true,
      showTime: true
    }
    let save = {
      sender_id: this.data.userInfo.id,
      receiver_id: this.data.chatInfo.id,
      group_id: 0,
      type: 'private',
      message_text: this.data.inputMessage
    }

    http.post('/message/send_message', save, (res) => {
      if (res.data.code) {} else {
        this.setData({
          messages: this.data.messages.pop()
        })
      }
    })

    this.setData({
      messages: [...this.data.messages, newMessage],
      inputMessage: '',
      scrollToMessage: `msg-${newMessage.id}`
    })
  },

  onLoadMore: function () {
    this.getHistoryMessages(this.data.chatInfo.id)
  },

  onBack: function () {
    wx.navigateBack()
  },

  // 预览自己头像
  onAvatarTap: function (e) {
    const userId = e.currentTarget.dataset.userId
    let chatId = this.data.userInfo.id
    const avatar = userId === chatId ? this.data.userInfo.avatar : this.data.chatInfo.avatar
    // 预览头像大图
    wx.previewImage({
      urls: [avatar],
      current: avatar
    })
  }
})