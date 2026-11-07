const http = require("./../../../utils/http.js");
Page({
  data: {
    currentCategory: 'all',  // 二手市场货物类型
    isLoading: false,
    goodsList: [],
    title:'', // 二手市场名称
    page: 1,
    limit:3
  },

  onLoad() {
    // 页面加载时执行
    this.loadGoods();
  },

  // 切换分类
  switchCategory(e) {
    const category = e.currentTarget.dataset.category;
    this.setData({
      currentCategory: category,
      goodsList: [] // 清空列表，重新加载
    });
    this.loadGoods(category);
  },

  // 搜索商品
  onSearch(e) {
  
    const keyword = e.detail.value;
    // 这里实现搜索逻辑
    this.searchGoods(keyword);
  },

  // 加载商品列表
  loadGoods() {
    this.setData({ isLoading: true });
    // 模拟API请求
    setTimeout(() => {
      // 这里应该调用实际的API
      http.get(`/goods/get?title=${this.data.title}&category=${this.data.currentCategory}&page=${this.data.page}&limit=${this.data.limit}`,{},(res)=>{
        if (res.data.code){
          this.setData({
            goodsList: res.data.data.data
          })
        }else{
          wx.showToast({
            title: res.data.msg,
            icon: 'none'
          })
        }
      })
      this.setData({ isLoading: false });
    }, 1000);
  },

  // 搜索商品
  searchGoods(keyword) {
    if (!keyword) {
      this.loadGoods(this.data.currentCategory);
      return;
    }

    // 这里实现搜索逻辑
    const filteredGoods = this.data.goodsList.filter(item => 
      item.title.includes(keyword) || 
      item.description.includes(keyword)
    );

    this.setData({
      goodsList: filteredGoods
    });
  },

  // 加载更多
  loadMore() {
    if (this.data.isLoading) return;
    this.loadGoods(this.data.currentCategory);
  },

  // 查看商品详情
  viewGoodsDetail(e) {
    const id = e.currentTarget.dataset.id;
    wx.navigateTo({
      url: `/pages/goods-detail/goods-detail?id=${id}`
    });
  },

  // 发布商品
  navigateToPublish() {
    const user = wx.getStorageSync('user')
    if (Object.keys(user).length === 0){
      wx.showToast({
        title: '请先登录',
        icon: 'none'
      })
      return
    }
    wx.navigateTo({
      url: '/pages/market/post/post'
    });
  },

  onPullDownRefresh() {
    // 下拉刷新
    this.setData({
      goodsList: []
    });
    this.loadGoods();
    wx.stopPullDownRefresh();
  },

  onReachBottom() {
    // 上拉加载更多
    this.loadMore();
  }
});