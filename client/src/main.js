// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import Home from '@/components/Home'
// import View from '@/components/View'
import Login from '@/components/Login'
import Signup from '@/components/Signup'
import SecretQuote from '@/components/SecretQuote'
import UserInfo from '@/components/UserInfo'

// custom components
import TreeView from '@/components/TreeView'

import VueResource from 'vue-resource'
Vue.use(VueResource)

import VueRouter from 'vue-router'
Vue.use(VueRouter)

Vue.config.productionTip = true

import auth from './auth'

function requireAuth (to, from, next) {
  if (!auth.isAuthenticated()) {
    this.$router.replace('/login')
  } else {
    next()
  }
}

const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    {
      path: '/',
      component: Home
    },
    {
      path: '/home',
      name: 'home',
      component: Home
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/view',
      name: 'view',
      component: TreeView,
      // beforeEnter: showVue,
      props: true
    },
    // dynamic segments start with a colon
    {
      path: '/view/:id',
      name: 'view.detail',
      components: {
        default: TreeView,
        sidebar: TreeView,
        content: TreeView
      },
      // beforeEnter: showVue,
      props: true
    },
    {
      path: '/signup',
      name: 'signup',
      component: Signup
    },
    {
      path: '/secretquote',
      name: 'secretquote',
      component: SecretQuote,
      beforeEnter: requireAuth
    },
    {
      path: '/userinfo',
      name: 'userinfo',
      component: UserInfo,
      beforeEnter: requireAuth
    }
  ]
})

router.beforeEach((to, from, next) => {
  window.console.log(to)
  next()
})
/* eslint-disable no-new */
const app = new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
}).$mount('#app')

window.addEventListener('popstate', () => {
  app.currentRoute = window.location.pathname
  window.console.log(app.currentRoute)
})
