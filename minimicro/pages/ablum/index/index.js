const util = require("../../../utils/util.js");
const http = require("../../../utils/http.js");

const app = getApp();
Page({
  data: {
    baseImageUrl: app.globalData.imageUrl,
    pageSize: 4,
    pageNumber: 1,
    select: 1,
    leftList: [],
    rightList: [],
    balloons: [],   // 点击特效
    scrollTop: 0, // 存储页面滚动距离
  },
  onLoad: function () {
    wx.showLoading({
      title: '加载中',
    });
    this.getList();
  },

  onShow:function(){
      this.setData({
        leftList: [],
        rightList: [],
        leftHeight: 0,
        rightHeigt: 1,
        pageNumber:1
      })
      this.getList();
  },

  /**
   * 获取具体类型的贴子
   */
  selected:function(e) {
    let objType = e.target.dataset.type;
    this.setData({
      select: objType,
      pageNumber: 1
    })
    this.setData({
      pageNumber: this.data.pageNumber,
      leftList: [],
      rightList: [],
      leftHeight: 0,
      rightHeigt: 1
    });
    this.getList();
  },

  /**
   * 进入发表页面
   */
  post: function () {
    // 判断是否登录
    let user = wx.getStorageSync('user')
    if (Object.keys(user).length === 0){
      wx.showToast({
        title: '请先登录',
        icon: 'none'
      })
      return
    }
    wx.navigateTo({
      url: '/pages/ablum/post/post'
    }).catch((err)=>{
      console.log(err)
    })
  },

  /**
   * 获取贴子列表
   */
  getList:function(){
    let objType = this.data.select;
    http.get(`/ablum/get?limit=${this.data.pageSize}&page=${this.data.pageNumber}&type=${objType}`,{},res => {
      wx.hideLoading();
      if (res.data.code){
        this.setData({
          leftList: res.data.data.leftList,
          rightList:res.data.data.rightList
        })
      }
    });
  },

/*
    追加的方式获取帖子
 */
  getMoreList:function(){
    let objType = this.data.select;
    http.get(`/ablum/get?limit=${this.data.pageSize}&page=${this.data.pageNumber}&type=${objType}`,{},res => {
      wx.hideLoading();
      if (res.data.code){
        if (res.data.data.leftList.length === 0 && res.data.data.rightList.length === 0){
          wx.showToast({
            title: '没有更多内容了',
            icon: 'none'
          })
          let page =  this.data.pageNumber
          page = page - 1
          this.setData({
            pageNumber: page
          })
          return
        }else{
          this.setData({
            leftList: this.data.leftList.concat(res.data.data.leftList),
            rightList: this.data.rightList.concat(res.data.data.rightList)
          })
        }
      }
    });
  },

  /**
  * 下拉刷新，获取最新的贴子
  */
  onPullDownRefresh: function () {
    this.setData({
      pageNumber: 1
    })
    this.getList();
  },

  /**
   * 上拉加载跟多
  */
  onReachBottom: function () {
    let page = this.data.pageNumber
    page = page + 1
    this.setData({
      pageNumber: page
    })
    wx.showLoading({
      title:'加载中...'
    })
    this.getMoreList();
  },

  /**
   * 点赞
   */
  praise:function(e){
    // 点赞特效
    this.handleTap(e);

    let objId = e.currentTarget.dataset.objid;
    let leftList = this.data.leftList;
    let rightList = this.data.rightList;

    const leftListIndex = leftList.findIndex(item => item.id === objId);
    const rightListIndex = rightList.findIndex(item => item.id === objId);
    if (leftListIndex !== -1){
      leftList[leftListIndex].like_num ++ ;
      this.setData({
        leftList: leftList,
      })
    }
    if (rightListIndex !== -1){
      rightList[rightListIndex].like_num ++ ;
      this.setData({
        rightList: rightList,
      })
    }
    http.get(`/ablum/praise`, {'id': objId }, res => {
      if (res.data.code){
        // wx.showToast({
        //   title: '点赞成功',
        //   icon: 'none'
        // })
      }else{
        wx.showToast({
          title: '点赞失败',
          icon: 'none'
        })
      }
    });
  },

  // 点赞特效
  handleTap: function(e) {
    const that = this;
    const touch = e.touches[0];
    
    // 获取当前滚动距离
    const scrollTop = this.data.scrollTop || 0;
    
    // 计算正确的 Y 坐标（pageY - scrollTop）
    const correctedY = touch.pageY - scrollTop;

    // 随机颜色
    const colors = ['#FF5252', '#FF4081', '#E040FB', '#7C4DFF', '#448AFF'];
    const randomColor = colors[Math.floor(Math.random() * colors.length)];
    
    // 新气球
    const newBalloon = {
      id: Date.now(),
      x: touch.pageX,
      y: correctedY, // 使用修正后的 Y 坐标
      color: randomColor
    };
    
    // 添加气球
    this.setData({
      balloons: this.data.balloons.concat([newBalloon])
    });
    
    // 1秒后移除气球
    setTimeout(() => {
      that.setData({
        balloons: that.data.balloons.filter(b => b.id !== newBalloon.id)
      });
    }, 1000);
  },
  
  // 监听页面滚动
  onPageScroll(e) {
    this.setData({ scrollTop: e.scrollTop }); 
  },


  openUserInfo:function(e){
    let id = e.currentTarget.dataset.id;
    wx.navigateTo({
      // url: '?id=' + id
    })
  }
})