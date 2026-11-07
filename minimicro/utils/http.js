const config = require("./../config.js");
const app = getApp();

/**
 * get
 */
const get=function ( _url, _data, callback) {
  httpRequest("GET",_url,_data,callback);
}

/**
 * post
 */
const post = function (_url, _data, callback) {
  httpRequest("POST", _url, _data, callback);
}

/**
 * put
 */
const put = function (_url, _data, callback) {
  httpRequest("PUT", _url, _data, callback);
}

/**
 * delete
 */
const httpDelete = function (_url, _data, callback) {
  httpRequest("DELETE", _url, _data, callback);
}

/**
 * patch
 */
const patch = function (_url, _data, callback) {
  httpRequest("PATCH", _url, _data, callback);
}

/** 
* 封装微信http请求
*/
const httpRequest=function (_method, _url, _data, callback) {
  let token = wx.getStorageSync('token');
  let _this = this;
  wx.request({
    url: config.domain + _url,
    header: {
      'content-type': 'application/json',
      'Authorization': 'Bearer ' + token
    },
    method: _method,
    data: _data,
    success: function (res) {
      if (res.data.error_code == '5000') {
        app.globalData.authStatus = true;
        callback(res);
        wx.showToast({
          title: res.data.error_message,
          icon:"none"
        })
        setTimeout(res=>{
          wx.switchTab({
            url: '/pages/personal/login/login'
          })
        },1500)
    
      } else {
        if (res.data.error_code != 0) {
          wx.showToast({
            title: res.data.error_message,
            icon:"none"
          })
        }
        callback(res);
      }
    },
    fail: function (res) {
      console.log(res);
    }
  })
}




const uploadFile = function(_tmpImgPath,_tmpImgType) {
    let token = wx.getStorageSync('token');
    return new Promise((resolve,reject)=>{
      wx.uploadFile({
        filePath: _tmpImgPath,
        name: 'file',
        url: config.domain+'/images/upload',
        header:{
          'file_type':  _tmpImgType,
          'Authorization': 'Bearer ' + token
        },
        fail(err){
          reject(err)
        },
        success(res){
          let resp =  JSON.parse(res.data)
          resolve(resp)
        }
      })
    })
} 


  // 上传多个文件
const  uploadFiles =  async function(_tmpImgPath) {
  try {
      const uploadPromises = _tmpImgPath.map(filePath => uploadFile(filePath,'png'));
      const results = await Promise.all(uploadPromises);
      // console.log('所有文件上传成功', results);
      return results
  } catch (error) {
      console.error('文件上传失败', error);
      return err
  }
}



module.exports = { get, post, patch, put, httpDelete, httpRequest,uploadFile,uploadFiles}

