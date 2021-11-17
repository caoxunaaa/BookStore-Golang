<template>
  <div class="comment">
    <div class="showClass">
      <el-button type="warning" @click="get_book_raply">评论文章</el-button>
      <div v-for="(item, index) in commentData['CommentsTree']" :key="index" class="fatherClass" style="width: 95%;border:1px solid green;margin: 10px">
        <div style="width: 100%;height:100%;display: -webkit-flex;">
          <div style="width:20%;">
            <div style="background: yellow;display:block">
              <span style="font-size:15px;font-weight:bold;">{{item.Comments.CommentByNickname}}</span><br>
            </div>
          </div>
          <div style="width:80%;">
            <div style="display: -webkit-flex; margin-bottom: 10px">
              <div style="width:80%;font-size:20px;font-weight:bold;text-align:left">
                {{item.Comments.Comment}}
              </div>
              <div style="width:20%">
                <el-button style="border: 0; background-color: transparent;padding:6px" @click="get_reply(item.Comments)"><i class="iconfont el-icon-s-comment"></i><span>回复</span></el-button>
              </div>
            </div>

            <div style="width:100%;">
              <div v-for="(items, indexs) in item.CommentsNode" :key="indexs" style="display: -webkit-flex;">
                <div style="width:80%;text-align:left">
                  <span><b>{{items.Comments.CommentByNickname}}</b>回复<b>{{items.Comments.CommentToNickname}}:</b></span>
                  <span>{{items.Comments.Comment}}</span>
                </div>
                <div style="width:20%;float:right;">
                  <el-button style="border: 0; background-color: transparent;" @click="get_reply(items.Comments)"><i class="iconfont el-icon-s-comment"></i><span>回复</span></el-button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div v-if="replyColseOpen" style="width: 100%; height: 100px; border:1px solid red">
         <el-input v-model="replyInputData" :placeholder="reply"></el-input><el-button @click="get_comment_button">评论</el-button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  components: {
  },
  name: 'comment',
  data () {
    return {
      user: {
        CommentByUserId: localStorage.getItem('UserId'),
        CommentByNickname: localStorage.getItem('Nickname')
      },
      commentData: [],
      replyData: [],
      replyMessage: {},
      replyInputData: '',
      replyColseOpen: false,
      reply: ''
    }
  },
  props: [
    'contentId'
  ],
  methods: {
    // 回复评论按钮
    get_reply (data) {
      console.log(data)
      this.replyMessage = data
      this.replyInputData = ''
      this.replyColseOpen = true
      this.reply = ':回复' + this.replyMessage.CommentByNickname
    },
    // 回复文章
    get_book_raply () {
      this.replyMessage = {
        Id: 0,
        BookContentId: this.contentId,
        CommentByUserId: this.user.CommentByUserId,
        CommentByNickname: this.user.CommentByNickname
      }
      this.replyInputData = ''
      this.replyColseOpen = true
      this.reply = ':回复文章'
    },
    // 提交按钮
    get_comment_button () {
      console.log('提交')
      let that = this
      let formData = new FormData()
      formData.append('parentId', that.replyMessage['Id'])
      formData.append('bookContentId', this.contentId)
      formData.append('comment', that.replyInputData)
      formData.append('commentToUserId', that.replyMessage['CommentByUserId'])
      formData.append('commentToNickname', that.replyMessage['CommentByNickname'])
      formData.append('commentByUserId', that.user.CommentByUserId)
      formData.append('commentByNickname', that.user.CommentByNickname)
      that.$axios({
        url: '/api/action/comment/',
        method: 'POST',
        data: formData
      }).then(function (response) {
        console.log(response)
        if (response.status === 200) {
          that.$message({
            type: 'success',
            message: '评论成功！'
          })
          that.get_data()
          that.replyInputData = ''
        }
      })
    },
    // 获取当前章节所有的评论
    get_data () {
      console.log('成功this.contentId', this.contentId)
      let that = this
      that.$axios({
        url: '/api/action/comment/by-book-content-id',
        method: 'get',
        params: {
          bookContentId: this.contentId
        }
      }).then(function (response) {
        that.commentData = response.data
        console.log('成功that.commentData', that.commentData)
      })
    }
  },
  mounted () {
    this.get_data()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.fatherClass{
  display: -webkit-flex;
  display: flex;
  -webkit-flex-direction: column;
  flex-direction: column;
  width: 400px;
  background-color: lightgrey;
}
</style>
