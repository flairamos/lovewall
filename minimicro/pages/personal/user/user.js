const app = getApp();
const http = require("./../../../utils/http.js");

Page({
  data: {
    user: {
      avatar: '',
      name: '',
      background_img: '',  //签名，不要在意细节
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

  onLoad(option) {
    // 页面加载时执行
    // 获取用户信息
    // 获取关注信息
    this.findUser(option.phone)

    this.findFollowStatInfo(option.phone)
    // 判断是否关注
    // 自己不能关注自己
    this.isUserFollow(option.phone)
    // 获取表白
    this.getWallList(option.phone)
   // this.checkFollowStatus();
  },

  checkFollowStatus() {
    // 检查关注状态
    // 这里模拟从服务器获取关注状态
    this.setData({
      isFollowing: false
    });
  },

  // 获取用户信息
  findUser(phone){
    http.get(`/user/form?phone=${phone}`,'',(res)=>{
      if (res.data.code){
        this.setData({
          user: res.data.data
        })
        return
      }else{
        wx.showToast({
          title: res.data.msg,
          icon: 'none'
        })
        return
      }
    })
  },
  
  isUserFollow(collect_phone){
    let user = wx.getStorageSync('user')
    if (user.phone == collect_phone){
      this.setData({
        isFollowing: true
      })
      return
    }
    http.get(`/follow/get?collect_phone=${collect_phone}`,'',(res)=>{
      if (res.data.code){
        this.setData({
          isFollowing: res.data.data
        })
        return
      }else{
        wx.showToast({
          title: res.data.msg,
          icon:'none'
        })
        return
      }
    })
  },


  followUser(collect_phone) {
    // 登录后才可以关注
    let user = wx.getStorageSync('user')
    if (Object.keys(user).length <= 0){
      wx.showToast({
        title: '请先登录',
        icon:'none'
      })
      return
    }
    // 不能对自己账号操作
    if (user.phone == collect_phone){
      wx.showToast({
        title: '包取关不了的',
        icon:'none'
      })
      return
    }
    // 关注/取消关注用户
    const isFollowing = !this.data.isFollowing;
    this.setData({
      isFollowing,
      'stats.fans': isFollowing ? this.data.stats.fans+ 1 : this.data.stats.fans - 1
    });

    if (!isFollowing){
      http.post("/follow/delete",{"collect_phone":this.data.user.phone},(res)=>{
        if (res.data.code){
          wx.showToast({
            title: res.data.msg,
            icon: 'none',
          });
        }else{
          wx.showToast({
            title: res.data.msg,
            icon: 'none',
          });
        }
      })
      return
    }else{
      http.post("/follow/add",{"collect_phone":this.data.user.phone,"collect_id":this.data.user.id},(res)=>{
        if (res.data.code){
          wx.showToast({
            title: res.data.msg,
            icon: 'none',
          });
        }else{
          wx.showToast({
            title: res.data.msg,
            icon: 'none',
          });
        }
      })
    }
    return
  },


    /**
   * 点赞
   */
  // praise: function (event) {
  //   let user = wx.getStorageSync('user')
  //   if (Object.keys(user).length == 0){
  //     wx.showToast({
  //       title: '请先登录',
  //       icon:'none'
  //     })
  //     return
  //   }
  //   let objId = event.target.dataset.obj;
  //   let objType = 1;
  //   this.setData({
  //     show: 0,
  //     hidden: false,
  //     showCommentInput: false
  //   });

  //   http.post(`/posting/praise`, {
  //     id: objId,
  //     obj_type: objType
  //   }, res => {
  //     let postList = this.data.posts;
  //     if (res.data.code) {
  //       let newPostList = postList.map(item => {
  //         if (objId == item.id) {
  //           item.is_like = !item.is_like
  //           if (item.is_like){
  //             item.liker.push({id:user.id, phone:user.phone, name:user.name,avatar:user.avatar });
  //           }else{
  //             let newLiker = item.liker.filter(item => item.id!== user.id);
  //             item.liker = newLiker
  //           }
  //         }
  //         // 返回处理后的元素
  //         return item;
  //       });
  //       //重新赋值，更新数据列表
  //       this.setData({
  //         posts: newPostList
  //       });
  //     }
  //   });
  // },



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

  sendMessage() {
    // 登录后才可以关注
    let user = wx.getStorageSync('user')
    if (Object.keys(user).length <= 0){
      wx.showToast({
        title: '请先登录',
        icon:'none'
      })
      return
    }
    // 不能对自己账号操作
    if (user.phone == collect_phone){
      wx.showToast({
        title: '包私信不了的',
        icon:'none'
      })
      return
    }
    // 发送私信
    wx.navigateTo({
      url: '/pages/chat/chat?userId=123'
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
    // 这里可以实现加载更多内容的逻辑
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
  getWallList: function (phone) {
    http.get(`/posting/get?limit=${this.data.limit}&page=${this.data.page}&phone=${phone}`, {}, res => {
      if (res.data.data.data.length > 0) {
        this.setData({
          posts: res.data.data.data,
          page: this.data.page + 1
        });
      }
    });
  },

  getMoreWallList: function (phone) {
    http.get(`/posting/get?limit=${this.data.limit}&page=${this.data.page}&phone=${phone}`, {}, res => {
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

    /**
   * 删除帖子
   */
  deletePost: function (e) {
    let objId = Number(e.target.id);
    wx.showModal({
      title: '提示',
      content: '确定删除吗?',
      success: res => {
        if (res.confirm) {
          http.post(`/posting/delete`, {id:objId}, res => {
            if (res.data.code) {
              let newPosts = this.data.posts.filter((item, index) => {
                if (item.id != objId) {
                  return item;
                }
              });
              this.setData({
                posts: newPosts
              });
            }
          });
        } else if (res.cancel) {
          console.log('用户点击取消')
        }
      }
    })
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