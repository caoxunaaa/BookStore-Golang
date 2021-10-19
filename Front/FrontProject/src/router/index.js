import Vue from 'vue'
import Router from 'vue-router'
import User from '@/components/user/User'
import Login from '@/components/user/Login'
import Register from '@/components/user/Register'
import HelloWorld from '@/components/HelloWorld'
import Book from '@/components/book/Book'
import BooksOverView from '@/components/book/BooksOverViewPaging'
import BooksUpload from '@/components/book/BooksUpload'
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
    },
    {
      path: '/book',
      component: Book,
      children: [
        {path: 'overview', component: BooksOverView},
        {path: 'upload', component: BooksUpload}
      ]
    }
  ]
})
