const util = require("./../../../utils/util.js");
const http = require("./../../../utils/http.js");
const app = getApp()

Page({
  data: {
    userInfo: {},
    posts: [],
    postType: 0, // 帖子类型
    baseImageUrl: app.globalData.apiUrl,
    show: 0,
    hidden: false,
    commentObjId: '',  // 评论人
    commentType: '',   // 品论类型
    refcommentId: '',  // 评论对象
    filter: '',
    pageSize: 2,
    pageNumber: 1,
    currentTime: '',
    select: 0,
    topic: '',
    showTopic: false,
    isLike: false ,  // 是否点赞
    showComments: false, // 评论框
    commentId: 0, // 评论组件属性- 评论ID
    commentName: '', // 
    commentType: 0 // // 评论组件属性- 评论类型
  },

  onLoad: function (e) {
    wx.showLoading({
      title: '加载中',
    });
    this.topic();
    this.getPost();
  },

  onShow: function (option) {
  },
  // 关闭评论区
  closeComment:function(val) {
    this.setData({
      showComments: false
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


  /**
   * 获取具体类型的贴子
   */
  selected(e) {
    let objType = e.target.dataset.type;
    // this. this.topic();
    // 标签切换和帖子类型切换
    // 页数置为0
    // 初始帖子数组置为空
    this.setData({
      select : objType,
      postType: objType,
      pageNumber: 1,
      posts: []
    })
    let thisTopic = this.data.topic;
    if (thisTopic != null){
      this.setData({
        showTopic: true,
      });
    }else{
      this.setData({
        showTopic: false,
      });
    }
    this.getNewPost(this.data.postType,'');
  },

  /**
   * 搜索
   */
  search: function (e) {
    wx.showLoading({
      title: '搜索中...',
    });
    // console.log("------",this.data.filter)
    this.getNewPost(this.data.postType,this.data.filter)
    wx.hideLoading()
  },

    /**
   * 获取搜索框的内容
   */
  getFilter: function (event) {
    let content = event.detail.value;
    this.setData({
      filter: content
    })
  },

  /**
   * 下拉刷新，获取最新的贴子
   */
  onPullDownRefresh: function () {
    let limit = this.data.pageNumber
    limit = limit + 1
    this.setData({
      pageNumber: limit,
    })
    this.getNewPost(this.data.postType);
  },

  /**
   * 上拉加载更多
   */
  onReachBottom: function () {
    let limit = this.data.pageNumber
    limit = limit + 1
    this.setData({
      pageNumber: limit,
    })
    // 该方法中没有新数据时需要将页数减小1
    this.getNewPost(this.data.postType);
  },

  /**
   * 获取用户信息
   */
  getUserInfo: function (e) {
    app.globalData.userInfo = e.detail.userInfo
    this.setData({
      userInfo: e.detail.userInfo,
    })
  },

  /** 
   * 进入发表页面
   */
  postView: function () {
    const user = wx.getStorageSync('user')
    if (Object.keys(user).length === 0){
      wx.showToast({
        title: '请先登录',
        icon: 'none'
      })
      return
    }
    wx.navigateTo({
      url: '/pages/home/post/post'
    })
  },


  /**
   * 获取贴子
   */
  // postingType 为数字 
  // search 为字符串
  getPost: function (postingType, search) {
    if (postingType === null || postingType === undefined) {
      postingType = this.data.postType
    }
    http.get(`/posting/get?limit=${this.data.pageSize}&page=${this.data.pageNumber}&type=${postingType}&name=${this.data.filter}`, {},
      res => {
        setTimeout(t => {
          wx.stopPullDownRefresh();
        }, 700)
        wx.hideLoading();
        if (res.data.data.data.length == 0){
          wx.showToast({
            title: '没有更多内容了',
            icon:"none"
          })
          return 
        }
        this.setData({
          posts: res.data.data.data
        })
      });
  },

  getNewPost: function (postingType, search) {
    if (postingType === null || postingType === undefined) {
      postingType = this.data.postType
    }
    http.get(`/posting/get?limit=${this.data.pageSize}&page=${this.data.pageNumber}&type=${postingType}&name=${this.data.filter}`, {},
      res => {
        setTimeout(t => {
          wx.stopPullDownRefresh();
        }, 700)
        wx.hideLoading();
        if (res.data.data.data.length == 0){
          wx.showToast({
            title: '没有更多内容了',
            icon:"none"
          })
          this.setData({
            pageNumber: this.data.pageNumber - 1
          })
          return
        }
        this.setData({
          posts: this.data.posts.concat(res.data.data.data)
        })
      });
  },

  /**
   * 预览图片
   */
  previewImage: function (event) {
    let url = event.target.id;
    wx.previewImage({
      current: '',
      urls: [url]
    })
  },

  /**
   * 预览图片
   */
  previewMoreImage: function (event) {
    let images = event.currentTarget.dataset.obj.map(item => {
      return this.data.baseImageUrl + item;
    });
    let url = event.target.id;
    wx.previewImage({
      current: url,
      urls: images
    })
  },

  /**
   * 显示评论控制面板
   */
  showComment: function (event) {
    this.setData({
      show: 0
    });
    let id = event.target.id;
    let hidden = event.target.dataset.show;
    if (!hidden) {
      this.setData({
        show: id,
        showCommentInput: true
      });
    } else {
      this.setData({
        show: 0,
        showCommentInput: false
      });
    }
  },

  /**
   * 触摸屏幕后移动触发一些隐藏操作
   */
  hiddenComment: function () {
    this.setData({
      show: 0,
      hidden: false,
      showCommentInput: false
    });
  },

  /**
   * 点赞
   */
  praise: function (event) {
    let user = wx.getStorageSync('user')
    if (Object.keys(user).length == 0){
      wx.showToast({
        title: '请先登录',
        icon:'none'
      })
      return
    }
    let objId = event.target.dataset.obj;
    let objType = 1;
    this.setData({
      show: 0,
      hidden: false,
      showCommentInput: false
    });

    http.post(`/posting/praise`, {
      id: objId,
      obj_type: objType
    }, res => {
      let postList = this.data.posts;
      if (res.data.code) {
        let newPostList = postList.map(item => {
          if (objId == item.id) {
            item.is_like = !item.is_like
            if (item.is_like){
              item.liker.push({id:user.id, phone:user.phone, name:user.name,avatar:user.avatar });
            }else{
              let newLiker = item.liker.filter(item => item.id!== user.id);
              item.liker = newLiker
            }
          }
          // 返回处理后的元素
          return item;
        });
        //重新赋值，更新数据列表
        this.setData({
          posts: newPostList
        });
      }
    });
  },

  /**
   * 点赞话题
   */
  // 取消点赞今日话题
  cancelLike: function (params) {
    let user = wx.getStorageSync('user')
    if (Object.keys(user).length == 0){
      wx.showToast({
        title: '请先登录',
        icon: 'none'
      })
      return
    }
    let id = params.currentTarget.dataset.obj
    let newTopic = this.data.topic
    newTopic.is_like = !this.data.topic.is_like
    newTopic.like_number = this.data.topic.like_number - 1
    http.get(`/topic/praise?id=${id}`,{},(res)=>{
      if (res.data.code){
        this.setData({
          topic: newTopic
        })
      }
    })
  },

  // 点赞今日话题
  addLike: function (params) {
    let user = wx.getStorageSync('user')
    if (Object.keys(user).length == 0){
      wx.showToast({
        title: '请先登录',
        icon: 'none'
      })
      return
    }
    let id = params.currentTarget.dataset.obj
    let newTopic = this.data.topic
    newTopic.is_like = !this.data.topic.is_like
    newTopic.like_number = this.data.topic.like_number + 1
    http.get(`/topic/praise?id=${id}`,{},(res)=>{
      if (res.data.code){
        this.setData({
          topic: newTopic
        })
      }
    })
  },

  /**
   * 回复别人
   */
  commentOtherComment: function (e) {
    let objId = e.currentTarget.dataset.objid;
    let type = e.currentTarget.dataset.objtype;
    let refcommentId = e.currentTarget.dataset.refid;
    this.setData({
      commentObjId: objId,
      commentType: type,
      show: 0,
      hidden: false,
      showCommentInput: true,
      refcommentId: refcommentId
    });
  },

  /**
   * 删除评论
   */
  deleteComment: function (e) {
    let objId = e.currentTarget.dataset.objid;
    let commentId = e.currentTarget.dataset.refid;
    wx.showModal({
      title: '提示',
      content: '确认删除该评论?',
      success: res => {
        if (res.confirm) {
          http.post(`/posting/delete`, {id:objId}, res => {
            if (res.data.data == 1) {
              let newPostList = this.data.posts.map(item => {
                if (objId == item.id) {
                  let newComment = item.comments.filter((item, index) => {
                    if (item.id != commentId) {
                      return item;
                    }
                  });
                  item.comments = newComment;
                }
                return item;
              });
              this.setData({
                posts: newPostList
              });
            }
          });

        } else if (res.cancel) {
          console.log('用户点击取消')
        }
      }
    })
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

  /**
   * 关注
   */
  follow: function (e) {
    let objId = e.target.dataset.obj;
    http.post('/follow', {
      obj_id: objId,
      obj_type: 1
    }, res => {
      let follow = res.data.data;
      let post = this.data.posts;
      let newPost = post.map(item => {
        if (item.id == follow.obj_id) {
          item.follow = true;
        }
        return item;
      });

      this.setData({
        posts: newPost
      });
    });
  },

  /**
   * 取消关注
   */
  cancelFolllow: function (e) {
    let objId = e.target.dataset.obj;
    http.put(`/cancel/${objId}/follow/1`, {}, res => {
      let follow = res.data.data;
      let post = this.data.posts;
      let newPost = post.map(item => {
        if (item.id == objId) {
          item.follow = false;
        }
        return item;
      });
      this.setData({
        posts: newPost
      });
    });

  },

  topic: function () {
    http.get(`/topic/get`, {}, res => {
      let topicShow = res.data != null ? true : false;
      this.setData({
        topic: res.data.data,
        showTopic: topicShow
      });
    });
  },


  openUserInfo: function (e) {
    let phone = e.currentTarget.dataset.phone;
    wx.navigateTo({
      url: '/pages/personal/user/user?phone=' + phone
    }).catch((err)=>{
      console.log(err)
    })
  }

})