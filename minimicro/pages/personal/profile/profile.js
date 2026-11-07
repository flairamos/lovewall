const app = getApp();
const http = require("./../../../utils/http.js");

Page({
  data: {
    user: {
      avatar: '',
      name: '',
      background_img: '',
      sex: false,
    },
    stats: {
      follows: 0,
      fans: 0,
      likes: 0
    },
    currentTab: 'love',
    isFollowing: false,
    posts:[],
    page: 1,
    limit: 2,
    showComments: false, // 评论框
    commentId: 0, // 评论组件属性- 评论ID
    commentName: '', // 
    commentType: 0 // // 评论组件属性- 评论类型
  },

  onLoad() {
    // 页面加载时执行
    // 获取用户信息
    let userStorage = wx.getStorageSync('user');
    if (Object.keys(userStorage).length !== 0){
      this.setData({
        user: userStorage
      })
    }else{
      // wx.navigateTo({
      //   url: '/pages/personal/login/login',
      // })
      return
    }
    // 获取关注信息
    this.findFollowStatInfo(userStorage.phone)
    this.getWallList()
    // this.checkFollowStatus();
  },

  onShow(){
    let user = wx.getStorageSync('user')
    if (Object.keys(user).length == 0){
      wx.navigateTo({
        url: '/pages/personal/login/login',
      })
    }
  },

  checkFollowStatus() {
    // 检查关注状态
    // 这里模拟从服务器获取关注状态
    this.setData({
      isFollowing: true
    });
  },

  // 获取关注列表
  findFollowStatInfo(phone){
    http.get(`/follow/info?&phone=${phone}`,'',(res)=>{
      this.setData({
        stats:{
          follows: res.data.data.follows,
          fans: res.data.data.fans,
          likes: res.data.data.likes
        }
      })
    })
  },


  // followUser() {
  //   // 关注/取消关注用户
  //   const isFollowing = !this.data.isFollowing;
  //   this.setData({
  //     isFollowing,
  //     'stats.fans': isFollowing ? '1.2w' : '1.1w'
  //   });
    
  //   wx.showToast({
  //     title: isFollowing ? '关注成功' : '已取消关注',
  //     icon: 'success',
  //     duration: 1500
  //   });
  // },

  sendMessage() {
    // 发送私信
    wx.navigateTo({
      url: '/pages/chat/chat?userId=123'
    });
  },

  editProfile() {
    wx.navigateTo({
      url: '/pages/personal/edit/edit'
    });
  },

  switchTab(e) {
    const tab = e.currentTarget.dataset.tab;
    this.setData({
      currentTab: tab
    });
    this.filterContent(tab);
  },

  filterContent(tab) {
    // 根据标签筛选内容
    let filteredContents = this.data.contents;
    if (tab === 'images') {
      filteredContents = this.data.contents.filter(item => item.type === 'image');
    } else if (tab === 'videos') {
      filteredContents = this.data.contents.filter(item => item.type === 'video');
    }
    
    this.setData({
      filteredContents
    });
  },

  onPullDownRefresh() {
    // 下拉刷新
    setTimeout(() => {
      wx.stopPullDownRefresh();
    }, 1500);
  },

  onReachBottom() {
    // 上拉加载更多
    // 记录旧数据长度
    let oldLen = this.data.posts.length
    this.setData({
      page: this.data.page + 1
    })
    this.getMoreWallList()
    if (this.data.posts.length == oldLen){
      this.setData({
        page: this.data.page - 1
      })
    }
  },

      /**
   * 获取贴子
   */
  getWallList: function () {
    let token = wx.getStorageSync('token')
    if (token != ""){
      http.get(`/posting/get?limit=${this.data.limit}&page=${this.data.page}&phone=${this.data.user.phone}`, {}, res => {
        if (res.data.data.data.length > 0) {
          this.setData({
            posts: res.data.data.data,
            page: this.data.page + 1
          });
        }
      });
    }
  },

  getMoreWallList: function () {
    http.get(`/posting/get?limit=${this.data.limit}&page=${this.data.page}&phone=${this.data.user.phone}`, {}, res => {
      if (res.data.data.data.length > 0) {
        this.setData({
          posts: this.data.posts.concat(res.data.data.data),
        });
      }else{
        wx.showToast({
          title: '没有更多内容了',
          icon:'none'
        })
      }
    });
  },

  // 退出登录
  logout(){
    try {
      wx.switchTab({
        url: '/pages/home/index/index',
      })
      wx.clearStorageSync()
      // this.onLoad()
      wx.reLaunch({
        url:"/pages/home/index/index"
      })
    } catch(err){
      console.error('所有存储数据清除失败', err);
    }
  },


  // 显示评论框
  showCommentSection(e) {
    // 获取点击评论图标的编号
    let param = e.currentTarget.dataset
    this.setData({
      showComments: true,
      commentId: param.objid,
      commentName: param.objname,
      commentType: param.objtype
    });
  },

  // 关闭评论区
  closeComment:function(val) {
    this.setData({
      showComments: false
    })
  },
}); 