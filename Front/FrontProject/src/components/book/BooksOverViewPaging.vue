<template>
  <div id="books_overview">
    <div style="width: 100%; display: flex;flex-wrap: wrap;">
      <div class="list" :key="key" v-for="(book_display,key) in books_display" style="width: 10%;margin: 0 3%">
        <router-link tag="span" :to="{path:'/book/'+book_display.Id + '/content/overview'}"><el-image style="width: 80%; height: 50%" :src="book_display.Image" lazy></el-image></router-link>
        <el-tag type="success" style="margin: 3px">{{ book_display.Name }}</el-tag>
        <el-tag type="success" style="margin: 3px">阅读量: {{ book_display.TrafficStatistic ? book_display.TrafficStatistic : 0}}</el-tag>
        <el-tag type="info" style="margin: 3px"><i class="el-icon-user">{{ book_display.Author }}</i></el-tag>
        <el-tag type="warning" style="margin: 3px"><i class="el-icon-time">{{ book_display.Time }}</i></el-tag>
      </div>
    </div>
    <el-pagination align='center'
                   @current-change="handleCurrentChange"
                   :page-size="pageSize"
                   :current-page="currentPage"
                   layout="total, prev, pager, next, jumper"
                   :total="books.length">
    </el-pagination>
  </div>
</template>

<script>
export default {
  name: 'BooksOverViewPaging',
  data () {
    return {
      books: [],
      books_display: [],
      currentPage: 0,
      pageSize: 20,
      traffic_statistic: {}
    }
  },
  mounted () {
    this.get_all_traffic_statistic_sorted_by_book_id()
    this.get_all_books()
  },
  methods: {
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
            Id: res[i].id,
            TrafficStatistic: that.traffic_statistic[res[i].id],
            Name: res[i].name,
            Image: 'http://172.20.3.111:8002/' + res[i].image,
            Author: res[i].author,
            Time: res[i].storage_time
          })
        }
        if (that.books.length < 20) {
          that.books_display = that.books.slice(0, that.books.length)
        } else {
          that.books_display = that.books.slice(0, 20)
        }
      }).catch(function (error) {
        console.log(error)
        alert('获取所有书籍失败')
      })
    },
    handleCurrentChange (page) {
      if (this.books.length < 20 * page) {
        this.books_display = this.books.slice(20 * (page - 1), this.books.length)
      } else {
        this.books_display = this.books.slice(20 * (page - 1), 20 * page)
      }
    },
    get_all_traffic_statistic_sorted_by_book_id () {
      // 统计所有书籍的访问量
      let that = this
      that.$axios({
        method: 'get',
        url: '/api/action/traffic-statistic/'
      }).then(function (response) {
        const res = response.data
        Object.keys(res).forEach(function (key) {
          let bookId = key.split(':')[0]
          if (that.traffic_statistic[bookId] === undefined) {
            that.traffic_statistic[bookId] = 0
          }
          that.traffic_statistic[bookId] += parseInt(res[key])
        })
        console.log(that.traffic_statistic)
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
