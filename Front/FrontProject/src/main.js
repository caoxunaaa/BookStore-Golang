// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

Vue.config.productionTip = false
Vue.prototype.$axios = axios
Vue.use(ElementUI)

Vue.prototype.getNowFormatDate = function () {
  const date = new Date()
  let year = date.getFullYear()
  let month = date.getMonth() + 1
  let day = date.getDate()
  let hour = date.getHours()
  let min = date.getMinutes()
  let sec = date.getSeconds()
  if (month >= 1 && month <= 9) {
    month = '0' + month
  }
  if (day >= 0 && day <= 9) {
    day = '0' + day
  }
  return year + '-' + month + '-' + day + ' ' + hour + ':' + min + ':' + sec
}

axios.interceptors.request.use((config) => {
  config.headers.authorization = localStorage.getItem('Token')
  return config
}, error => {
  return Promise.reject(error)
})

/* eslint-disable no-new */
const Bus = new Vue()
new Vue({
  el: '#app',
  router,
  components: {App},
  data () {
    return {
      Bus
    }
  },
  template: '<App/>'
})
