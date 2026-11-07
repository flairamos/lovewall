const http = require("./../../../utils/http.js");
const enums = require("./../../../utils/enum.js");
const app = getApp()
Page({
  data: {
    images: '', //临时图片地址
    showOriginalPrice: false,
    categoryIndex: -1,
    categories: ['教材书籍', '数码产品', '服饰装备', '生活用品', '其他'],
    formData: {
      title: '',
      description: '',
      price: '',
      original_price: '',
      category: '',
      images:''
    },
    canSubmit: false,
  },

  onLoad() {
    // 页面加载时执行
  },

  // 选择图片
  chooseImage() {
    let that = this;
    wx.chooseMedia({
      count: 1,
      mediaType: ['image'],
      sourceType: ['album', 'camera'],
      success(res) {
        const tempFiles = res.tempFiles[0].tempFilePath;
        that.setData({
          images: tempFiles
        });
        that.checkCanSubmit();
      }
    });
  },

  // 删除图片
  deleteImage(e) {
    // const index = e.currentTarget.dataset.index;
    // const images = this.data.images;
    this.setData({
      images: ''
    });
    this.checkCanSubmit();
  },

  // 标题输入
  onTitleInput(e) {
    this.setData({
      'formData.title': e.detail.value
    });
    this.checkCanSubmit();
  },

  // 描述输入
  onDescInput(e) {
    this.setData({
      'formData.description': e.detail.value
    });
    this.checkCanSubmit();
  },

  // 价格输入
  onPriceInput(e) {
    this.setData({
      'formData.price': e.detail.value
    });
    this.checkCanSubmit();
  },

  // 原价输入
  onOriginalPriceInput(e) {
    this.setData({
      formData: {
        ...this.data.formData,
        original_price: e.detail.value
      }
    });
  },

  // 切换显示原价输入框
  toggleOriginalPrice() {
    this.setData({
      showOriginalPrice: !this.data.showOriginalPrice
    });
    if (!this.data.showOriginalPrice) {
      this.setData({
        formData: {
          original_price: '',
        }
      });
    }
  },

  // 选择分类
  onCategoryChange(e) {
    const index = e.detail.value;
    if (index == -1  && this.data.categories[index] == undefined ||this.data.categories[index] == null ) {
      wx.showToast({
        title: '请选择分类',
        icon:'none'
      })
      return
    }
    this.setData({
      categoryIndex: index,
      formData:{
        ...this.data.formData,
        category: this.data.categories[index]
      } 
    });
    this.checkCanSubmit();
  },

  // 检查是否可以提交
  checkCanSubmit() {
    const { title, price, category } = this.data.formData;
    const user = wx.getStorageSync('user')
    if (Object.keys(user).length === 0){
      wx.showToast({
        title: '请先登录',
        icon: 'none'
      })
      return
    }

    const canSubmit = 
      this.data.images.length > 0 && 
      title.length > 0 && 
      price.length > 0 && 
      category.length > 0 ;
    this.setData({
      canSubmit
    });
  },

  // 取消发布
  cancelPublish() {
    if (this.data.images.length > 0 || this.data.formData.title || this.data.formData.description) {
      wx.showModal({
        title: '提示',
        content: '确定要放弃发布吗？',
        success: (res) => {
          if (res.confirm) {
            wx.navigateBack();
          }
        }
      });
    } else {
      wx.navigateBack();
    }
  },

  // 提交发布
  submitPublish() {
    if (!this.data.canSubmit) return;

    wx.showLoading({
      title: '发布中...'
    });

    // 这里应该调用实际的发布API
    this.data.formData.category = enums.goods.get(this.data.formData.category)
    http.uploadFile(this.data.images,'png').
    then((res)=>{
      if (res.code){
        let imageurl = app.globalData.imgUrl + '/'+res.data.fileName
        let param = this.data.formData
        param.images = imageurl
        http.post('/goods/add',param,(res)=>{
          if (res.data.code){
            setTimeout(() => {
              wx.hideLoading();
              wx.showToast({
                title: '发布成功',
                icon: 'success',
                duration: 2000,
                success: () => {
                  setTimeout(() => {
                    wx.navigateBack();
                  }, 2000);
                },
                fail: (err)=>{
                  console.log(err)
                  wx.navigateBack();
                }
              });
            }, 1500);
          }
        })
      }else{
        wx.showToast({
          title: res.msg,
        })
        return
      }
    }).catch((err)=>{
      console.log(err)
    })
  }
}); 