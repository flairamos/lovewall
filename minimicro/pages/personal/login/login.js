const app = getApp();
const http = require("./../../../utils/http.js");
Page({

  /**
   * 页面的初始数据
   */
  data: {
    phone:'',
    password:'',
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {

  },
  // onHide(){
  //   let user = wx.getStorageSync('user')
  //   if (Object.keys(user).length == 0){
  //     wx.switchTab({
  //       url: '/pages/home/index/index',
  //       success: (res)=>{
  //         console.log(res)
  //       },
  //       fail:(res)=>{
  //         console.log("err:",res)
  //       }
  //     })
  //   }
  // },
  onUnload(){
    wx.switchTab({
      url: '/pages/home/index/index',
      success: (res)=>{
        console.log(res)
      },
      fail:(res)=>{
        console.log("err:",res)
      }
    })
  },
  phone(e){
    this.setData({
      phone:e.detail.value
    })
  },
  password(e){
    this.setData({
      password:e.detail.value
    })
  },

  sendCode(){
    if(this.data.phone!=''){
      //发送验证码
      let phone={
        phone:this.data.phone
      }
     code(phone).then(res=>{
      this.setData({
        code:res.data
      })
     })
    }else{
      wx.showToast({
        title: '手机号不能为空',
        icon:'none'
      })
    }
  },

     // 加密
  strToHex(str){
      let result = ""
      for(let i = 0;i<str.length;i++){
        result += str.charCodeAt(i).toString(16)
      }
      return result
  },

  login(){
    if(this.data.phone == '' || this.data.password == ''){
      wx.showToast({
        title: '邮箱账号或密码不能为空呀',
        icon:'none'
      })
      return
    }
    var pattern = /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
    let flag = pattern.test(this.data.phone)
    if (flag === false) {
      wx.showToast({
        title: '邮箱账号格式不对呢',
        icon:'none'
      })
      return
    }
    // 密码加密转16进制
    let pass = this.strToHex(this.data.password)
    http.post("/public/login",{name:this.data.phone, password:pass}, res=>{
      if (res.data.code){
        wx.setStorageSync('token', res.data.data.token);
        let user = {
          id: res.data.data.user_info.id,
          phone: res.data.data.user_info.phone,
          name: res.data.data.user_info.name,
          sex:res.data.data.user_info.sex,
          auth:res.data.data.user_info.authentication,
          auth_name:res.data.data.user_info.authentication_name,
          background_img:res.data.data.user_info.background_img,
          avatar:res.data.data.user_info.avatar,
        }
        wx.setStorageSync('user', user);
        // 设置全局变量
        app.globalData.userInfo = user
        app.globalData.authStatus = true
        // 跳转到我的页面
        wx.switchTab({
          url: '/pages/personal/profile/profile',
          fail: function (err) {
            console.error('页面跳转失败:', err);
          },
          success: function (e) {
            // 页面自动刷新
            var page = getCurrentPages().pop();
            if (page == undefined || page == null) return;
            page.onLoad();
          }
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

  /**
   * 跳转到注册页面
   */
  openRegisterView:function(params) {
    wx.navigateTo({
      url: '/pages/personal/logout/logout',
    })
  }
})