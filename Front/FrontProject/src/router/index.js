import Vue from 'vue'
import Router from 'vue-router'
import User from '@/components/user/User'
import Login from '@/components/user/Login'
import Register from '@/components/user/Register'
import HelloWorld from '@/components/HelloWorld'
import Book from '@/components/book/Book'
import BooksOverView from '@/components/book/BooksOverViewPaging'
import BooksMyself from '@/components/book/BooksMyself'
import BooksUpload from '@/components/book/BooksUpload'
import Content from '@/components/book/content/Content'
import ContentOverView from '@/components/book/content/ContentOverViewPaging'
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
        {name: 'login', path: 'login', component: Login},
        {path: 'register', component: Register}
      ]
    },
    {
      path: '/book',
      component: Book,
      children: [
        {path: 'overview', component: BooksOverView},
        {path: 'myself', component: BooksMyself},
        {path: 'upload', component: BooksUpload},
        {
          path: ':book_id/content',
          component: Content,
          props: true,
          children: [
            {path: 'overview', component: ContentOverView}
          ]
        }

      ]
    }
  ]
})
