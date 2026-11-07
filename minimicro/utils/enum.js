const config = require("../config.js");
const app = getApp();


const posting = new Map();
posting.set("love",1)
posting.set("find",2)
posting.set("vote",3)
posting.set("picture",4)



const goods = new Map();
goods.set("教材书籍",'books')
goods.set("数码产品",'digital')
goods.set("服饰装备",'clothes')
goods.set("生活用品",'daily')
goods.set("其他",'others')




module.exports = {posting,goods}