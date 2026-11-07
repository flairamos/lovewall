const { imgUrl } = require("../../../config.js");
const varable = require("./../../../utils/enum.js");
const http = require("./../../../utils/http.js");
Page({
  data: {
    user :'',
    title: '',
    content: '',
    uploadedImages: [],
    copyData: '',
    selectedAction: 'love', // Tracks which action button is selected
    tmpImgPath : '', // 评论图片的临时地址
    tmpImgType : 'png' // 临时图片类型
  },

  onLoad(){
    let userStorage = wx.getStorageSync('user')
    if (Object.keys(userStorage).length === 0){
      wx.switchTab({
        url: '/pages/personal/login/login',
      }).catch((err)=>{
        // console.log("===",err)
      })
      return
    }
    this.setData({
      user: userStorage
    })
  },

  // Get title input
  getTitle(e) {
    this.setData({
      title: e.detail.value
    });
  },

  // Get content input
  getTextContent(e) {
    this.setData({
      content: e.detail.value
    });
  },

  // Handle action button selection
  selectAction(e) {
    const type = e.currentTarget.dataset.type;
    this.setData({
      selectedAction: this.data.selectedAction === type ? '' : type
    });
  },

  // Choose and upload image
  chooseImage() {
    const that = this;
    if (this.data.uploadedImages.length >= 3) {
      wx.showToast({
        title: '最多只能上传3张图片',
        icon: 'none'
      });
      return;
    }

    wx.chooseMedia({
      count: 3 - that.data.uploadedImages.length,
      mediaType: ['image'],
      sourceType: ['album', 'camera'],
      success(res) {
        const tempFiles = res.tempFiles;
        const newImages = tempFiles.map(file => file.tempFilePath);
        
        that.setData({
          uploadedImages: [...that.data.uploadedImages, ...newImages]
        });
      }
    });
  },

  // Delete image
  deleteImage(e) {
    const index = e.currentTarget.dataset.index;
    const images = this.data.uploadedImages;
    images.splice(index, 1);
    this.setData({
      uploadedImages: images
    });
  },

  // Post content
  post() {
    // if (!this.data.title.trim()) {
    //   wx.showToast({
    //     title: '请输入标题',
    //     icon: 'none'
    //   });
    //   return;
    // }

    if (!this.data.content.trim()) {
      wx.showToast({
        title: '请输入正文',
        icon: 'none'
      });
      return;
    }

    // Here you would typically send the data to your server
    // console.log('Posting:', {
    //   title: this.data.title,
    //   content: this.data.content,
    //   images: this.data.uploadedImages,
    //   selectedAction: this.data.selectedAction
    // });
    let postType = varable.posting.get(this.data.selectedAction)

    // 多文件上传
    if (this.data.uploadedImages.length != 0){
      let result = http.uploadFiles(this.data.uploadedImages)
      result.then((res)=>{
        //console.log('-------',res)
        let images = []
        res.forEach(element => {
          images.push(imgUrl+"/"+element.data.fileName)
        });
  
        let postadd = {
          "phone": this.data.user.phone,
          "user_id": this.data.user.id,
          "type": postType,
          "content": this.data.content,
          "photos": images,
          "ref_comment": 0
        }
        http.post("/posting/add", postadd, res => {
          if (res.data.code){
            wx.showToast({
              title: '评论成功！',
            })
            wx.navigateBack({
              delta: 1, // 返回的页面数，如果 delta 大于现有页面数，则返回到首页
              success: function() {
                  console.log('返回上一页成功');
              },
              fail: function(err) {
                  console.error('返回上一页失败', err);
              },
              complete: function() {
                  console.log('返回上一页操作完成');
              }
          });
          }else{
            wx.showToast({
              title: res.data.msg,
            })
          }
        })
      })
    
      // Show success message
      wx.showToast({
        title: '发布成功',
        icon: 'success'
      });
    }else{
      let postadd = {
        "phone": this.data.user.phone,
        "user_id": this.data.user.id,
        "type": postType,
        "content": this.data.content,
        "photos": [],
        "ref_comment": 0
      }
      http.post("/posting/add", postadd, res => {
        if (res.data.code){
          wx.showToast({
            title: '评论成功！',
          })
          wx.navigateBack({
            delta: 1, // 返回的页面数，如果 delta 大于现有页面数，则返回到首页
            success: function() {
                console.log('返回上一页成功');
            },
            fail: function(err) {
                console.error('返回上一页失败', err);
            },
            complete: function() {
                console.log('返回上一页操作完成');
            }
        });
        }else{
          wx.showToast({
            title: res.data.msg,
          })
        }
      })
      // Show success message
      wx.showToast({
        title: '发布成功',
        icon: 'success'
      });
    }

  }
});