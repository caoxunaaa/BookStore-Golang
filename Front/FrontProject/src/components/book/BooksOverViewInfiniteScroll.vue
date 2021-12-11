<template>
  <div id="books_overview" style="width: 60%; display: flex;flex-wrap: wrap;">
    <div class="list" :key="key" v-for="(book_display,key) in books_display" style="width: 14%;margin: 0 3%">
      <el-image style="width: 80%;" :src="book_display.Image" lazy></el-image>
      <el-tag type="success" style="margin: 3px">{{ book_display.Name }}</el-tag>
      <el-tag type="info" style="margin: 3px"><i class="el-icon-user">{{ book_display.Author }}</i></el-tag>
      <el-tag type="warning" style="margin: 3px"><i class="el-icon-time">{{ book_display.Time }}</i></el-tag>
    </div>
    <infinite-loading @infinite="onInfinite" ref="infiniteLoading"><span slot="no-more"></span></infinite-loading>
  </div>
</template>

<script>
import InfiniteLoading from 'vue-infinite-loading'

export default {
  name: 'BooksOverView',
  components: {InfiniteLoading},
  data () {
    return {
      books: [],
      books_display: [],
      book_current_display_count: 0
    }
  },
  mounted () {
    this.get_all_books()
  },
  updated () {
    this.displayBeforeShowScrollbar()
  },
  methods: {
    onInfinite ($state) {
      this.book_current_display_count++
      if (this.books.length >= this.book_current_display_count) {
        setTimeout(() => {
          this.books_display.push({
            Name: this.books[this.book_current_display_count - 1].Name,
            Image: this.books[this.book_current_display_count - 1].Image,
            Author: this.books[this.book_current_display_count - 1].Author,
            Time: this.books[this.book_current_display_count - 1].Time
          })
          $state.loaded()
        }, 500)
      } else {
        $state.complete()
      }
      console.log(this.book_current_display_count)
    },
    displayBeforeShowScrollbar () {
      // 不断展示书籍直到出现滚动条
      if (document.body.scrollHeight < document.documentElement.clientHeight) {
        setTimeout(() => {
          console.log(this.books_display)
          if (this.books.length === this.book_current_display_count) {
            return
          }
          this.books_display.push(this.books[this.book_current_display_count])
          this.book_current_display_count++
        }, 100)
      }
      // console.log(document.body.clientHeight, document.body.scrollHeight, document.body.offsetHeight, document.documentElement.clientHeight, window.screen.height, window.screen.availHeight)
      // console.log(this.books.length, this.book_current_display_count)
    },
    get_all_books () {
      let that = this
      that.$axios({
        method: 'get',
        url: '/api/book/'
      }).then(function (response) {
        const res = response.data.booksBasicInfo
        console.log(res)
        for (let i = 0; i < res.length; i++) {
          that.books.push({
            Name: res[i].name,
            Image: 'http://114.115.169.233:8002/' + res[i].image,
            Author: res[i].author,
            Time: res[i].storage_time
          })
        }
        if (that.books.length > 0) {
          that.books_display.push(that.books[0])
          that.book_current_display_count++
        }
      }).catch(function (error) {
        console.log(error)
        alert('获取所有书籍失败')
      })
    }
  }
}
</script>

<style scoped>

</style>
