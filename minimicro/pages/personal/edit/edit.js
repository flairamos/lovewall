const http = require("./../../../utils/http.js");
Page({
  data: {
    userInfo: {},
    bioLength: 0,
    isModified: false,
    items: [
      { value: true, name: '男', checked: true },
      { value: false, name: '女', },
      ]
  },

  onLoad(options) {
    // 获取用户信息
    this.getUserInfo();
  },

  getUserInfo() {
    // 这里应该从服务器获取用户信息
    // 这里使用模拟数据
    let user = wx.getStorageSync('user')
    this.setData({
      userInfo : user,
      bioLength: user.background_img.length
    });
    if (user.sex){
      this.setData({
        items: [
          { value: true, name: '男', checked: true },
          { value: false, name: '女', },
          ]
      });
    }else{
      this.setData({
        items: [
          { value: true, name: '男' },
          { value: false, name: '女',checked: true },
          ]
      });
    }
  },

  uploadAvatar(filePath) {
    wx.showLoading({
      title: '上传中...',
    });
    // 这里应该调用上传API
    setTimeout(() => {
      wx.hideLoading();
      wx.showToast({
        title: '上传成功',
        icon: 'success'
      });
    }, 1500);
  },

  chooseAvatar() {
    let that = this;
    wx.chooseMedia({
      count: 1,
      mediaType: ['image'],
      sourceType: ['album', 'camera'],
      success(res) {
        const tempFiles = res.tempFiles[0].tempFilePath;
        that.setData({
          'userInfo.avatar': tempFiles,
          isModified: true
        });
        // 这里可以立即上传头像
        that.uploadAvatar(tempFiles);
      }
    });
  },

  // 性别选项
  radioChange: function(e) {
    this.setData({
    selectedValue: e.detail.value
    });
  },

  onNicknameChange(e) {
    this.setData({
      'userInfo.name': e.detail.value,
      isModified: true
    });
  },

  onBioChange(e) {
    this.setData({
      'userInfo.background_img': e.detail.value,
      bioLength: e.detail.value.length,
      isModified: true
    });
  },

  navigateToAboutUs() {
    wx.navigateTo({
      url: '/pages/about/about'
    });
  },

  navigateToContact() {
    wx.navigateTo({
      url: '/pages/contact/contact'
    });
  },

  saveProfile() {
    if (!this.data.isModified) {
      wx.navigateBack();
      return;
    }

    wx.showLoading({
      title: '保存中...'
    });

    // 这里应该调用保存API
    setTimeout(() => {
      wx.hideLoading();
      wx.showToast({
        title: '保存成功',
        icon: 'success',
        duration: 1500,
        success: () => {
          setTimeout(() => {
            wx.navigateBack();
          }, 1500);
        }
      });
    }, 1500);
  },

  cancelEdit() {
    if (this.data.isModified) {
      wx.showModal({
        title: '提示',
        content: '确定要放弃修改吗？',
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

  onUnload() {
    // 页面卸载时可以做一些清理工作
  }
}); 