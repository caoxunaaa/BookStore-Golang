<template>
  <div id="books_upload">
    <el-row :gutter="20">
      <el-col :span="8" :offset="8">
        <div class="grid-content bg-purple">
          <el-form ref="form" :model="form">
            <el-form-item label="书籍名称">
              <el-input v-model="form.name"></el-input>
            </el-form-item>
            <el-form-item label="书籍作者">
              <el-input v-model="form.author"></el-input>
            </el-form-item>
            <el-form-item label="上传图片"><input type="file" id="book_image" accept=".jpg,.png,.jpeg"></el-form-item>
            <el-form-item>
              <el-button @click="create_book('form')"
                         style="background-color: #409eff; color: #fff; height: 40px; width:120px">新增
              </el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'BooksUpload',
  data () {
    return {
      form: {
        name: '',
        author: '',
        storageTime: '',
        storageUserId: 0
      }
    }
  },
  methods: {
    create_book (form) {
      this.$refs[form].validate((valid) => {
        if (valid) {
          let that = this
          let bookImage = document.getElementById('book_image').files[0]
          let formData = new FormData()
          formData.append('name', this.form['name'])
          formData.append('author', this.form['author'])
          formData.append('storageTime', this.getNowFormatDate())
          formData.append('storageUserId', localStorage.getItem('UserId'))
          if (bookImage !== undefined) {
            formData.append('image', bookImage, bookImage.name)
          } else {
            formData.append('image', '')
          }

          that.$axios({
            method: 'post',
            url: '/api/book/',
            data: formData
          }).then(function (response) {
            const res = response.data
            console.log(res)
            that.$message({message: '上传成功', duration: 1000})
            setTimeout(function () {
              that.$router.push({path: '/book/overview'})
              window.location.reload()
            }, 1000)
          }).catch(function (error) {
            console.log(error)
            alert('上传书籍失败')
          })
        } else {
          console.log('error create book!!')
          return false
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
