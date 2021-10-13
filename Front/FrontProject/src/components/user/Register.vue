<template>
  <div id="register" style="">
    <el-row :gutter="20">
      <el-col :span="8" :offset="8">
        <div class="grid-content bg-purple">
          <el-form ref="form" :model="form" label-width="140px">
            <el-form-item label="用户名">
              <el-input le v-model="form.username"></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="form.password"></el-input>
            </el-form-item>
            <el-form-item label="确认密码">
              <el-input v-model="form.pwAgain"></el-input>
            </el-form-item>
            <el-form-item label="邮箱">
              <el-input v-model="form.email"></el-input>
            </el-form-item>
            <el-form-item label="电话">
              <el-input v-model="form.phone"></el-input>
            </el-form-item>
            <el-form-item label="昵称">
              <el-input v-model="form.nickname"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onSubmit('form')">注册</el-button>
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
  name: 'Register',
  data () {
    return {
      form: {
        username: '',
        password: '',
        pwAgain: '',
        email: '',
        phone: '',
        nickname: ''
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
          if (that.form.password !== that.form.pwAgain) {
            alert('两次密码输入不一致')
            return false
          }
          let formData = new FormData()
          formData.append('username', that.form['username'])
          formData.append('email', that.form['email'])
          formData.append('phone', that.form['phone'])
          formData.append('password', that.form['password'])
          formData.append('pwAgain', that.form['pwAgain'])
          formData.append('nickname', that.form['nickname'])

          that.$axios({
            method: 'post',
            url: '/api/user/register',
            data: formData
          }).then(function (response) {
            const res = response.data
            this.$message('注册成功')
            console.log(res)
            that.$router.push({path: '/user/login'})
            window.location.reload()
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
