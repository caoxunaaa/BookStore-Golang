<template>
  <div id="content_overview_paging">
    <div style="width: 60%; display: flex; flex-wrap: wrap; flex-direction: column;">
      <div class="list" :key="key" v-for="(content_display,key) in contents_display" style="width: 14%;margin: 0 3%">
          <el-tag type="success" style="margin: 3px" @click="drawer = true;selectContent(content_display.ChapterContent)">
            第{{ content_display.ChapterNum }}章 {{ content_display.ChapterName }} <i
            class="el-icon-time">{{ content_display.CreateTime }}</i>
          </el-tag>
      </div>
    </div>
    <el-pagination align='center'
                   @current-change="handleCurrentChange"
                   :page-size="pageSize"
                   :current-page="currentPage"
                   layout="total, prev, pager, next, jumper"
                   :total="contents.length">
    </el-pagination>
    <el-drawer
      title=""
      :visible.sync="drawer"
      :with-header="false"
      v-if="drawer" :size="size">
      <div style="white-space: pre-wrap;text-align: left">{{body_content}}</div>
<!--       {{ body_content }}-->
    </el-drawer>
  </div>
</template>

<script>
export default {
  name: 'ContentOverViewPaging',
  props: [
    'book_id'
  ],
  data () {
    return {
      contents: [],
      contents_display: [],
      currentPage: 0,
      pageSize: 20,
      body_content: '',
      drawer: false,
      size: '80%'
    }
  },
  mounted () {
    this.get_content()
  },
  methods: {
    get_content () {
      let that = this
      that.$axios({
        method: 'get',
        url: '/api/book/content/',
        params: {
          bookId: this.book_id
        }
      }).then(function (response) {
        const res = response.data.book_contents_reply
        console.log(res)
        for (let i = 0; i < res.length; i++) {
          that.contents.push({
            Id: res[i].id,
            BookId: res[i].book_id,
            ChapterContent: res[i].chapter_content,
            ChapterName: res[i].chapter_name,
            ChapterNum: res[i].chapter_num,
            CreateTime: res[i].create_time
          })
        }
        if (that.contents.length < 20) {
          that.contents_display = that.contents.slice(0, that.contents.length)
        } else {
          that.contents_display = that.contents.slice(0, 20)
        }
      }).catch(function (error) {
        console.log(error)
        alert('暂无内容章节')
      })
    },
    handleCurrentChange (page) {
      if (this.contents.length < 20 * page) {
        this.contents_display = this.contents.slice(20 * (page - 1), this.contents.length)
      } else {
        this.contents_display = this.contents.slice(20 * (page - 1), 20 * page)
      }
    },
    selectContent (text) {
      this.body_content = text
    }
  }
}
</script>

<style scoped>

</style>
