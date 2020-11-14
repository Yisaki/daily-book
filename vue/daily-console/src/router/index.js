import Vue from 'vue'
import VueRouter from 'vue-router'

import Daily from '@/views/daily/Daily'

Vue.use(VueRouter)

const routes = [
  {
    path: '/asd',
    name: 'Daily',
    component: Daily
  },

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
