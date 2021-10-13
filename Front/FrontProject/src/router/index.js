import Vue from 'vue'
import Router from 'vue-router'
import User from '@/components/user/User'
import Login from '@/components/user/Login'
import Register from '@/components/user/Register'
import HelloWorld from '@/components/HelloWorld'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: HelloWorld
    },
    {
      path: '/user',
      component: User,
      children: [
        {path: 'login', component: Login},
        {path: 'register', component: Register}
      ]
    }
  ]
})
