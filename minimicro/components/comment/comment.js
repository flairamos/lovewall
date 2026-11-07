const http = require("./../../utils/http.js");
const app = getApp()
Component({
  properties: {
    commentId: {
      type: Number,
      value: 0
    },
    commentName: {
      type: String,
      value: "",
    },
    commentType: {
      type: Number,
      value: 0
    },
  },

  data: {
    comments: [],
    inputValue: '',
    flatComments: [], // 添加平铺后的评论数组
    isReply: false, // 是否是回复
    replyObj: null, // 回复的评论obj
    hasImage: false, // 评论框是否有图片
    tmpImgPath : '', // 评论图片的临时地址
    tmpImgType : '', // 临时图片类型
    inputHeight: 72, // 初始高度，输入框自适应变化
    isRefreshing: false,
    emojiList: ['😊', '😂', '🤣', '😍', '😒', '😘', '🙄', '😭', '😡', '🤔', '👍', '❤️', '💔', '🌹', '✨'],
    page: 1,
    limit: 20,
    total: 0,
    baseUrl: app.globalData.apiUrl,
  },

  methods: {
    hideCommentSection() {
      this.triggerEvent("closeComment", false)
    },
    preventBubble() {
      // 阻止冒泡
    },

    onLoadMore() {
      if (this.data.comments.length >= this.data.total) {
        return
      }
      this.setData({
        page: this.data.page + 1
      });
      this.loadComments(this.data.page);
    },

    async onRefresh() {
      this.setData({
        isRefreshing: true
      });
      const newComments = await this.loadComments(1);
      this.setData({
        comments: newComments,
        page: 1,
        isRefreshing: false,
        replyTo: null,
        showReplyInput: false,
        selectedImages: []
      });
    },

    loadComments(page) {
      http.get(`/posting/get_comments?limit=${this.data.limit}&page=${ page }&type=${+this.data.commentType}&id=${this.data.commentId}`, {},
        res => {
          this.setData({
            total: res.data.data.total
          })
          const commentsWithParentId = this.addParentId(res.data.data.data).map(it => {
            it.flatComments = this.flattenComments(it?.children || [])
            return it
          })
          let nowComments = this.data.comments || []
          this.setData({
            comments: [...nowComments, ...commentsWithParentId]
          });
        });
    },

    // 选择图片
    async chooseImage() {
      // 判断是否登录
      let user = wx.getStorageSync('user')
      if (Object.keys(user).length === 0){
        wx.showToast({
          title: '请先登录',
          icon: 'none'
        })
        return
      }
      let that = this
      try {
        const res = await wx.chooseMedia({
          count: 1,
          mediaType: ['image'],
          sizeType: ['compressed'],
          sourceType: ['album', 'camera'],
          success (res) {
            // tempFilePath可以作为img标签的src属性显示图片
            const tempFilePaths = res.tempFiles[0].tempFilePath;
            // console.log("-----",tempFilePaths)
            // 获取文件类型
            const match = tempFilePaths.match(/\.[^.]*$/);
            let extension = ''
            if (match) {
                extension = match[0].substring(1);
            }
            that.setData({
              hasImage: true,
              tmpImgPath: tempFilePaths,
              tmpImgType: extension
            });
          }
        });
      } catch (error) {
        wx.showToast({
          title: '选择图片失败'+ error,
        })
      }
    },

        // 预览图片
    previewImage(e) {
      const {
        current,
        urls
      } = e.currentTarget.dataset;

      wx.previewImage({
        current,
        urls
      });
    },

    // 删除缓存图片
    deleteImage() {
      this.setData({
        hasImage: false,
        imageUrl: ''
      });
    },

    bindKeyInput(e) {
      this.setData({
        inputValue: e.detail.value,
      })
    },

    // 发送评论 - 按钮
    submitComment() {
      // 判断是否登录
      let user = wx.getStorageSync('user')
      if (Object.keys(user).length === 0){
        wx.showToast({
          title: '请先登录',
          icon:"none"
        })
        return
      }
      if (this.data.inputValue.length === 0){
        wx.showToast({
          title: '不能提交空信息',
        })
        return
      }
      // console.log(this.data.inputValue, 'inputValue', 'type',this.data.commentId, this.data.commentType)
      // console.log(this.data.commentName, 'commentObj')
      if (!this.data.isReply) {
        // console.log('给帖子的评论')
              // 上传图片
      http.uploadFile(this.data.tmpImgPath,this.data.tmpImgType)
      .then((res)=>{
        if (res.code){
          let photos = []
          photos.push(app.globalData.imgUrl + '/'+res.data.fileName)
          let postingadd = {
            "phone": user.phone,
            "user_id": user.id,
            "type": this.data.commentType,
            "content": this.data.inputValue,
            "photos": photos,
            "ref_comment": this.data.commentId
          }
          http.post("/posting/add", postingadd, res => {
            if (res.data.code){
              wx.showToast({
                title: '评论成功！',
              })
              this.onLoadMore()
            }else{
              wx.showToast({
                title: res.data.msg,
              })
            }
          })
        }
        
      }).catch((err)=>{
        wx.showToast({
          title: "服务器繁忙请稍后再试！",
        })
      })
      } else {
        // 判断评论类型 (0 , 1, 2)
        // console.log('回复给别人的评论')
        http.uploadFile(this.data.tmpImgPath,this.data.tmpImgType).then((res)=>{
          let photos = []
          photos.push(app.globalData.imgUrl + '/'+res.data.fileName)
                  // 构造评论参数
          let postingAdd = {
            "phone": user.phone,
            "user_id": user.id,
            "type": this.data.commentType,  // 评论帖子类型 （0，1，2）
            "content":this.data.inputValue,
            "photos": photos,
            "ref_comment": this.data.replyObj.id
          }
          http.post("/posting/add", postingAdd, res => {
            if (res.data.code){
              wx.showToast({
                title: '评论成功！',
              })
              this.onLoadMore()
            }else{
              wx.showToast({
                title: res.data.msg,
              })
            }
          })
        }).catch((err)=>{
          wx.showToast({
            title: "服务器繁忙请稍后再试！",
          })
        })
      }
      // 页面刷新
      var page = getCurrentPages().pop();
      if (page == undefined || page == null) return;
      page.onLoad();
    },

    // 发送评论 - 键盘
    sendReply(e) {
      //console.log('---',e)
    },

    // 打开回复弹窗
    openReply(e) {
      this.setData({
        isReply: true,
        replyObj: e.currentTarget.dataset.item
      })
      // console.log(e.currentTarget.dataset.item)
    },

    // 添加处理parentId的辅助方法
    addParentId(comments, parentItem = null) {
      return comments.map(comment => {
        const newComment = {
          ...comment
        };
        if (parentItem !== null) {
          newComment.parentId = parentItem.id;
          newComment.recoverName = parentItem.name
        }
        if (newComment.children && newComment.children.length > 0) {
          newComment.children = this.addParentId(newComment.children, newComment);
        }
        return newComment;
      });
    },

    // 将树形评论数据平铺成一维数组
    flattenComments(comments, level = 0) {
      let flattened = [];

      comments.forEach(comment => {
        // 创建新对象，避免修改原始数据
        const flatComment = {
          ...comment
        };
        // 添加层级信息
        flatComment.level = level;
        // 暂存子评论并删除，避免重复
        const children = flatComment.children;
        delete flatComment.children;

        // 将当前评论添加到结果数组
        flattened.push(flatComment);

        // 递归处理子评论
        if (children && children.length > 0) {
          flattened = flattened.concat(this.flattenComments(children, level + 1));
        }
      });

      return flattened;
    },
  },

  lifetimes: {
    attached: function () {
      // 在组件实例进入页面节点树时执行
      this.loadComments(1)
    },
    detached: function () {
      // 在组件实例被从页面节点树移除时执行
    },
  },
});