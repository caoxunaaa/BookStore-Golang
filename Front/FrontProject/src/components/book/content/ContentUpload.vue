<template>
  <div id="content_upload">
    <el-row :gutter="20">
      <el-col :span="16" :offset="4">
        <div class="grid-content bg-purple">
          <el-form ref="form" :model="form">
            <el-form-item label="书名">
              <el-input disabled v-model="form.BookName"></el-input>
            </el-form-item>
            <el-form-item label="章节数">
              <el-input v-model="form.ChapterNum"></el-input>
            </el-form-item>
            <el-form-item label="章节名">
              <el-input v-model="form.ChapterName"></el-input>
            </el-form-item>
            <el-form-item label="上传txt文件"><input type="file" id="chapter_content" accept=".txt" @change="read_file">
            </el-form-item>
            <el-form-item label="章节名">
              <el-input type="textarea" :rows="2" v-model="form.ChapterContent"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button @click="upload_content('form')"
                         style="background-color: #409eff; color: #fff; height: 40px; width:120px">上传
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
  name: 'ContentUpload',
  data () {
    return {
      form: {
        BookName: '',
        ChapterNum: 1,
        ChapterName: '',
        ChapterContent: '',
        CreateTime: ''
      }
    }
  },
  props: [
    'book_id',
    'book_name',
    'chapter_count'
  ],
  mounted () {
    this.get_book_name_and_chapter_num()
  },
  methods: {
    get_book_name_and_chapter_num () {
      this.form.BookName = this.book_name
      this.form.ChapterNum = this.chapter_count + 1
    },
    upload_content (form) {
      this.$refs[form].validate((valid) => {
        if (valid) {
          let that = this
          let formData = new FormData()
          formData.append('bookId', this.book_id)
          formData.append('chapterNum', that.form.ChapterNum)
          formData.append('chapterName', that.form.ChapterName)
          formData.append('chapterContent', that.form.ChapterContent)
          formData.append('createTime', that.getNowFormatDate())
          that.$axios({
            method: 'post',
            url: '/api/book/content/',
            data: formData
          }).then(function (response) {
            const res = response.data
            console.log(res)
            that.$message({message: '上传章节成功', duration: 1000})
            setTimeout(function () {
              window.location.reload()
            }, 1000)
          }).catch(function (error) {
            console.log(error)
            alert('上传书籍失败或者章节数重复')
          })
        } else {
          console.log('error create book!!')
          return false
        }
      })
    },
    read_file () {
      let that = this
      let content = document.getElementById('chapter_content').files[0]
      const reader = new FileReader()
      reader.readAsText(content, 'GB2312')
      setTimeout(function () {
        that.form.ChapterContent = reader.result
      }, 10)
    }
  }
}
</script>

<style scoped>

</style>
