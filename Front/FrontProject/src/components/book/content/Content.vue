<template>
  <div id="content">
    <el-button type="primary" @click="drawer_upload_content = true" style="display: flex;flex-wrap: wrap; margin: 20px 0">上传章节</el-button>
    <router-view :book_id="book_id"></router-view>
    <el-drawer
      title= "上传章节"
      :direction="direction"
      :visible.sync="drawer_upload_content">
      <ContentUpload v-if="drawer_upload_content" :book_id="book_id" :book_name="book_name" :chapter_count="chapter_count"></ContentUpload>
    </el-drawer>
  </div>
</template>

<script>
import ContentUpload from '@/components/book/content/ContentUpload'
export default {
  name: 'Content',
  components: {ContentUpload},
  props: [
    'book_id'
  ],
  data () {
    return {
      drawer_upload_content: false,
      direction: 'ltr',
      book_name: '',
      chapter_count: ''
    }
  },
  mounted () {
    this.get_book_name()
    this.get_book_content_count()
  },
  methods: {
    get_book_content_count () {
      this.$root.Bus.$on('book_chapter_count', e => {
        console.log('接收到', e)
        this.chapter_count = e
      })
    },
    get_book_name () {
      let that = this
      that.$axios({
        method: 'get',
        url: '/api/book/id/' + this.book_id
      }).then(function (response) {
        const res = response.data
        console.log('书籍信息', res)
        that.book_name = res.name
      }).catch(function (error) {
        console.log(error)
        alert('获取当前内容书籍失败')
      })
    }
  }
}
</script>

<style scoped>

</style>
