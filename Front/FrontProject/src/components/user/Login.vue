<template>
  <div id="login" style="">
    <el-row :gutter="20">
      <el-col :span="8" :offset="8">
        <div class="grid-content bg-purple">
          <el-form ref="form" :model="form" label-width="140px">
            <el-form-item label="用户名或邮箱或电话">
              <el-input le v-model="form.username"></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="form.password" show-password></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onSubmit('form')">登录</el-button>
              <el-button>取消</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data () {
    return {
      form: {
        username: '',
        password: ''
      }
    }
  },
  mounted () {
  },
  methods: {
    onSubmit (form) {
      this.$refs[form].validate((valid) => {
        if (valid) {
          let that = this
          let username = ''
          let email = ''
          let phone = ''
          if (/^1\d{10}$/.test(that.form['username'])) {
            phone = that.form['username']
          } else if (/^(\w-*\.*)+@(\w-?)+(\.\w{2,})+$/.test(that.form['username'])) {
            email = that.form['username']
          } else {
            username = that.form['username']
          }

          let formData = new FormData()
          formData.append('username', username)
          formData.append('email', email)
          formData.append('phone', phone)
          formData.append('password', that.form['password'])

          that.$axios({
            method: 'post',
            url: '/api/user/login',
            data: formData
          }).then(function (response) {
            const res = response.data
            console.log(res)
            localStorage.setItem('Token', res['AccessToken'])
            localStorage.setItem('Username', res['Name'])
            localStorage.setItem('Nickname', res['NickName'])
            that.$message({message: '登录成功', duration: 1000})
            setTimeout(function () {
              that.$router.push({path: '/'})
              window.location.reload()
            }, 1000)
          }).catch(function (error) {
            console.log(error)
            alert('用户不存在')
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
