const http = require("./../../../utils/http.js");
const app = getApp()
Page({
  data: {
    starredFriends: [
    //   {
    //     id: 1,
    //     name: '李文豪',
    //     avatar: '/images/avatar1.png'
    //   },
    //   {
    //     id: 2,
    //     name: '刘文静',
    //     avatar: '/images/avatar2.png'
    //   },
    //   {
    //     id: 3,
    //     name: '李文',
    //     avatar: '/images/avatar3.png'
    //   },
    //   {
    //     id: 4,
    //     name: '潘小明',
    //     avatar: '/images/avatar4.png'
    //   },
    //   {
    //     id: 5,
    //     name: '陈冰冰',
    //     avatar: '/images/avatar5.png'
    //   },
    //   {
    //     id: 6,
    //     name: '陈冰',
    //     avatar: '/images/avatar6.png'
    //   }
    ],
    messages: [
      // {
      //   id: 1,             // 该id是用户id用于进入私聊界面
      //   name: '陈冰冰',
      //   avatar: '/images/avatar5.png',
      //   lastMessage: '你好啊，一起去吃饭吗？',
      //   time: '18:54',
      //   unreadCount: 1
      // },
      // {
      //   id: 2,             // 该id是用户id用于进入私聊界面
      //   name: '李宇',
      //   avatar: '/images/avatar2.png',
      //   lastMessage: '好的，那就这么说定了！',
      //   time: '11:37',
      //   unreadCount: 0
      // },
      // {
      //   id: 3,
      //   name: '李文云',
      //   avatar: '/images/avatar3.png',
      //   lastMessage: '这个项目什么时候开始？',
      //   time: '昨天',
      //   unreadCount: 2
      // },
      // {
      //   id: 4,
      //   name: '王明',
      //   avatar: '/images/avatar4.png',
      //   lastMessage: '周末有空吗？我们讨论一下方案',
      //   time: '昨天',
      //   unreadCount: 0
      // },
      // {
      //   id: 5,
      //   name: '何怡',
      //   avatar: '/images/avatar1.png',
      //   lastMessage: '收到了，谢谢！',
      //   time: '2023/10/21',
      //   unreadCount: 0
      // }
    ],
    page: 1,
    isLoading: false,
    hasMore: true,
    wss : app.globalData.wss,
    active: app.globalData.active,
    timer: null
  },

  onLoad: function() {
    // 初始消息已经在 data 中定义，不需要重新加载
    let user = wx.getStorageSync('user')
    if (Object.keys(user).length == 0){
      wx.navigateTo({
        url: '/pages/personal/login/login',
      })
      return
    }
    this.getFollowList()
    // this.connectSocketServer()
  },

  onShow: function (params) {
      // 开启定时器，每 3 秒请求一次服务器
      this.data.timer = setInterval(() => {
        this.getNewMessge()
    }, 3000);
  },
  onHide() {
    // 页面隐藏时清除定时器，避免不必要的请求
    if (this.data.active){
      return
    }
    clearInterval(this.data.timer);
},


  onUnload: function(params) {
    wx.onSocketClose(function (res) {
      console.log('WebSocket连接已关闭:', res)
    })
  },

  // 获取关注列表
  getFollowList: function (params) {
    http.get('/follow/follow_list','',(res)=>{
      // console.log(res.data)
      if (res.data.code){
        this.setData({
          starredFriends: res.data.data
        })
      }else{
        wx.showToast({
          title: res.data.msg,
          icon: 'none'
        })
      }
    })
  },

  /*
  connectSocketServer(){
    let that = this
    let lastTime = wx.getStorageSync('last_time') 
    wx.connectSocket({
      url: this.data.wss+"/push"+`?last_time=${lastTime}`,
      header:{
        'Authorization': 'Bearer '+ wx.getStorageSync('token')
      },
      success: function (res) {
        console.log('WebSocket连接成功',res)
      },
      fail: function (res) {
        console.log('WebSocket连接失败:', res)
      }
    })
    wx.onSocketOpen(function (res) {
      console.log('WebSocket连接已打开',res)
    })
    wx.onSocketError(function (err) {
      console.log('WebSocket连接打开失败:', err)
    })
    // wx.sendSocketMessage({
    //   data: 'Hello, WebSocket!'
    // })
    wx.onSocketMessage(function (res) {
      // console.log('接收到服务器发送的数据:', res.data)
      // console.log('接收到服务器发送的数据:', typeof(res.data))
      let data = JSON.parse(res.data)
      that.setData({
        messages: data
      })
    })
    // wx.onSocketClose(function (res) {
    //   console.log('WebSocket连接已关闭:', res)
    // })
  },
  */

  getNewMessge: function() {
    let user = wx.getStorageSync('user')
    if (Object.keys(user).length == 0){
      return
    }
    http.get(`/message/get_message`,'',(res)=>{
      if (res.data.code){
        if (res.data.data == null){
          return
        }
        this.setData({
          messages: res.data.data
        })
      }else{
      }
    })
  },


  getMessages: function(isLoadMore = false) {
    if (!this.data.hasMore) return
    this.setData({
      isLoading: true
    })

    // 模拟API请求
    setTimeout(() => {
      const newMessages = [
        // {
        //   id: this.data.messages.length + 1,
        //   name: '新消息' + (this.data.messages.length + 1),
        //   avatar: '/images/avatar' + ((this.data.messages.length % 5) + 1) + '.png',
        //   lastMessage: '这是新加载的消息内容',
        //   time: '刚刚',
        //   unreadCount: Math.floor(Math.random() * 3)
        // },
        // {
        //   id: this.data.messages.length + 2,
        //   name: '新消息' + (this.data.messages.length + 2),
        //   avatar: '/images/avatar' + (((this.data.messages.length + 1) % 5) + 1) + '.png',
        //   lastMessage: '又一条新消息',
        //   time: '刚刚',
        //   unreadCount: Math.floor(Math.random() * 3)
        // }
      ]

      this.setData({
        messages: isLoadMore ? [...this.data.messages, ...newMessages] : newMessages,
        page: this.data.page + 1,
        isLoading: false,
        hasMore: this.data.page < 5 // 模拟只有5页数据
      })
    }, 1000)
  },

  onLoadMore: function() {
    if (!this.data.isLoading && this.data.hasMore) {
      this.getMessages(true)
    }
  },

  onMessageTap: function(e) {
    const messageId = e.currentTarget.dataset.id
    // 跳转到聊天详情页
    wx.navigateTo({
      url: `/pages/message/chat/chat?id=${messageId}`
    })
  },

  onSearch: function(e) {
    const keyword = e.detail.value
    // 实现搜索功能
    console.log('搜索关键词：', keyword)
  },

  onPullDownRefresh: function() {
    // 下拉刷新，重置页码并加载新数据
    this.setData({
      page: 1,
      hasMore: true
    }, () => {
      this.getMessages(false)
    })
    wx.stopPullDownRefresh()
  }
}) 