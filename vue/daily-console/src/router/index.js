import Vue from 'vue'
import VueRouter from 'vue-router'

import Period from '@/views/period/Period'
import PeriodList from '@/views/period/PeriodList'

Vue.use(VueRouter)

const routes = [
  {
    path: '/period',
    name: 'Period',
    component: Period
  },
  {
    path: '/periodList',
    name: 'PeriodList',
    component: PeriodList
  },

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
