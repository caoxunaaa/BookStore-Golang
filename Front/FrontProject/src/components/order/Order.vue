<template>
  <div id="order">
    <div v-if="need_line_up">
      <h3>购买人数太多,排队中...</h3>
    </div>
    <div v-else>
      <el-form ref="form" :model="order_info" label-width="80px">
        <el-form-item label="书籍名称">
          <el-input v-model="order_info.bookName" disabled></el-input>
        </el-form-item>
        <el-form-item label="购买者">
          <el-input v-model="order_info.buyerName" disabled></el-input>
        </el-form-item>
        <el-form-item label="费用">
          <el-input v-model="order_info.cost" disabled></el-input>
        </el-form-item>
        <el-form-item label="订单号">
          <el-input v-model="order_info.orderNum" disabled></el-input>
        </el-form-item>
        <el-form-item label="订单时间">
          <el-input v-model="order_info.orderTime" disabled></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="confirmOrderAndPay(order_info.orderNum)">确认订单并支付</el-button>
          <el-button type="warning" @click="deleteOrder(order_info.orderNum)">取消订单</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  components: {
  },
  name: 'Order',
  data () {
    return {
      need_line_up: true,
      timer_to_get_order_info: null,
      order_info: {bookId: 0, bookName: this.bookName, buyerId: 0, buyerName: '', cost: 0, id: 0, orderNum: '', orderStatus: '', orderTime: ''}
    }
  },
  props: [
    'bookId'
  ],
  mounted () {
    this.setTimerToGetOrderInfo()
  },
  beforeDestroy () {
    this.stopTimerToGetOrderInfo()
  },
  methods: {
    // 定时查询订单情况
    setTimerToGetOrderInfo () {
      let that = this
      let bookId = this.bookId
      that.timer_to_get_order_info = setInterval(func, 5000)
      func()
      function func () {
        that.$axios({
          url: '/api/order/not-paid-order-info',
          method: 'get',
          params: {
            buyerId: localStorage.getItem('UserId'),
            bookId: bookId
          }
        }).then(function (response) {
          const res = response.data
          console.log('not-paid-order-info', res)
          if (res.code === 2000) {
            that.need_line_up = false
            // that.stopTimerToGetOrderInfo()
            that.order_info = {
              bookId: res.message.bookId,
              bookName: res.bookName,
              buyerId: res.message.buyerId,
              buyerName: localStorage.getItem('Nickname'),
              cost: res.message.cost,
              id: res.message.id,
              orderNum: res.message.orderNum,
              orderStatus: res.message.orderStatus,
              orderTime: res.message.orderTime
            }
          } else if (res.code === 2002) {
            that.need_line_up = true
          } else if (res.code === 2003) {
            that.need_line_up = false
            alert(res.message)
            that.$router.go(0)
          } else if (res.code === 2001) {
            that.need_line_up = false
            that.stopTimerToGetOrderInfo()
            alert(res.message)
            that.$router.go(0)
          } else {
            alert(res.message)
            that.$router.go(0)
          }
        }).catch(function (error) {
          console.log(error)
          that.need_line_up = false
          that.$router.go(0)
        })
      }
    },
    stopTimerToGetOrderInfo () {
      clearInterval(this.timer_to_get_order_info)
    },
    confirmOrderAndPay (orderNum) {
      let that = this
      let formData = new FormData()
      formData.append('orderNum', orderNum)
      that.$axios({
        method: 'post',
        url: '/api/order/pay-for-order',
        data: formData
      }).then(function (response) {
        const res = response.data
        console.log(res)
        alert(res.message)
        if (res.code === 2000) {
          that.$router.go(0)
        }
      }).catch(function (error) {
        console.log(error)
        alert(error.response.data.message)
      })
    },
    deleteOrder (orderNum) {
      let that = this
      let formData = new FormData()
      formData.append('orderNum', orderNum)
      that.$axios({
        url: '/api/order/cancel-order',
        method: 'delete',
        data: formData
      }).then(function (response) {
        const res = response.data
        alert('已经取消订单')
        that.$router.go(0)
        console.log(res)
      }).catch(function (err) {
        console.log(err.response)
        alert('取消订单失败')
      })
    }
  }
}
</script>

<style scoped>

</style>
