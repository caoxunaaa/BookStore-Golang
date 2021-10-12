import Vue from 'vue'
import Router from 'vue-router'
import User from '@/components/user/User'
import Login from '@/components/user/Login'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'book'
    },
    {
      path: '/user',
      component: User,
      children: [
        {path: 'login', component: Login}
      ]
    }
  ]
})
