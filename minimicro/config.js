const config={
  dev:{//开发环境
    domain:"http://192.168.1.11:8000/api/wechat",//后台接口地址
    imageUrl:"http://xvwen.cn:8080/api/public/dl/your-share-id",  // 图片前缀
    wss :"ws://192.168.1.11:8000"  // ws地址
  },
  prod:{//生产环境
    domain:"https://xvwen.cn/api/wechat",//后台接口地址
    imageUrl:"http://xvwen.cn:8080/api/public/dl/your-share-id",  // 图片前缀
    wss :"ws://xvwen.cn"  // ws地址
  }
}

const domain = config.prod.domain;
const imgUrl = config.prod.imageUrl
const wss = config.prod.wss
const region = 'SCN';

module.exports = {
  domain, imgUrl,wss,region,
}