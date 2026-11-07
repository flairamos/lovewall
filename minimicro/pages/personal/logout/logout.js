const http = require("./../../../utils/http.js");
Page({

  /**
   * 页面的初始数据
   */
  data: {
  phone:'',
  code:'',
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {

  },
  phone(e){
    this.setData({
      phone:e.detail.value
    })
  },
  code(e){
    this.setData({
      code:e.detail.value
    })
  },

  sendCode() {
    if (this.data.countdown > 0) return;
    
    // 这里添加发送验证码的API调用代码
    if(this.data.phone == ''){
      wx.showToast({
        title: '手机号不能为空',
        icon:'none'
      })
    }

    let phone={
      name :this.data.phone
    }
    http.post('/public/send_code',phone, res=>{
      if (!res.code){
        wx.showToast({
          title: res.msg,
          icon: 'none'
        })
      }else{
        wx.showToast({
          title: '发送成功',
          icon: 'none'
        })
      }
    })

    // 开始倒计时
    this.setData({
      countdown: 60
    });
    
    this.data.timer = setInterval(() => {
      let current = this.data.countdown - 1;
      this.setData({
        countdown: current
      });
      
      if (current <= 0) {
        clearInterval(this.data.timer);
      }
    }, 1000);
  },

  // 注册
  register(){
    if(this.data.phone==''){
      wx.showToast({
        title: '邮箱不能为空呀',
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
  
    if(this.data.code == ''){
      wx.showToast({
        title: '验证码不能为空',
        icon:'none'
      })
      return
    }
    // 验证
    wx.login({
      success: (res) => {
        if (res.code){
          let param = {
            name: this.data.phone,
            password: this.strToHex('123456'),
            code: this.data.code,
            js_code: res.code
          }
          http.post('/public/register',param, resp=>{
            if (resp.data.code){
              wx.showToast({
                title: '注册成功',
                icon: 'none'
              })
              this.openLogin()
              return
            }else{
               wx.showToast({
                title: resp.data.msg,
                icon: 'none'
              })
            }
          })
        }else{
          console.log('登录失败！' + res.errMsg)
        }
      },
    })
  },

  // 加密
  strToHex(str){
      let result = ""
      for(let i = 0;i<str.length;i++){
        result += str.charCodeAt(i).toString(16)
      }
      return result
  },
  openLogin: function () {
    wx.navigateTo({
      url: '/pages/personal/login/login'
    })
  },
})
