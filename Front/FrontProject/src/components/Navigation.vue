<template>
  <div id="navigation">
    <el-menu
      :default-active="activeIndex"
      :router="true"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b">
      <el-menu-item>
        <router-link tag="li" :to="{path:'/book/overview'}">首页</router-link>
      </el-menu-item>
      <el-submenu index="2">
        <template slot="title">我的工作台</template>
        <el-menu-item index="2-1">选项1</el-menu-item>
        <el-menu-item index="2-2">选项2</el-menu-item>
        <el-menu-item index="2-3">选项3</el-menu-item>
        <el-submenu index="2-4">
          <template slot="title">选项4</template>
          <el-menu-item index="2-4-1">选项1</el-menu-item>
          <el-menu-item index="2-4-2">选项2</el-menu-item>
          <el-menu-item index="2-4-3">选项3</el-menu-item>
        </el-submenu>
      </el-submenu>

      <el-menu-item index="3" disabled>消息中心</el-menu-item>

      <div style="float: right" v-if="logged">
        <el-submenu index="5" style="float: right">
          <template slot="title">{{ nickname }}</template>
          <el-menu-item v-if="nickname==='书店老板'">
            <el-button size="small" type="danger" @click="start_hot_sale">开始热卖</el-button>
          </el-menu-item>
          <el-menu-item>
            <router-link tag="li" :to="{path:'/book/myself'}">我的书籍</router-link>
          </el-menu-item>
          <el-menu-item index="4-2" @click="logout">退出登录</el-menu-item>
        </el-submenu>
<!--        <el-menu-item index="4" style="float: right">-->
<!--          <router-link tag="li" :to="{path:'/book/upload'}">上传书籍</router-link>-->
<!--        </el-menu-item>-->
      </div>
      <div v-else>
        <el-menu-item style="float: right">
          <router-link tag="li" :to="{path:'/user/login'}">登录</router-link>
        </el-menu-item>
        <el-menu-item style="float: right">
          <router-link tag="li" :to="{path:'/user/register'}">注册</router-link>
        </el-menu-item>
      </div>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: 'Navigation',
  data () {
    return {
      activeIndex: '1',
      logged: false,
      nickname: '',
      timer: ''
    }
  },
  mounted () {
    this.getUser()
  },
  methods: {
    start_hot_sale () {
      let that = this
      that.$axios({
        url: '/api/order/start-order-handle',
        method: 'post'
      }).then(function (response) {
        const res = response.data
        console.log(res)
        alert('开始热卖')
      }).catch(function (err) {
        console.log(err.response)
      })
    },
    handleSelect (key, keyPath) {
      this.activeIndex = key
    },
    getUser () {
      let that = this
      that.logged = localStorage.getItem('Username') !== null
      if (that.logged === true) {
        that.nickname = localStorage.getItem('Nickname')
      }
    },
    logout () {
      localStorage.clear()
      this.logged = false
    }
  }
}
</script>

<style>
#navigation {
  margin-bottom: 20px;
}
</style>
