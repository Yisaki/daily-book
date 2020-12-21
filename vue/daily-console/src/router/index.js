import Vue from 'vue'
import VueRouter from 'vue-router'

import Period from '@/views/period/Period'
import PeriodList from '@/views/period/PeriodList'

import Pet from '@/views/pet/Pet'


Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Period',
    component: Period
  },
  {
    path: '/periodList',
    name: 'PeriodList',
    component: PeriodList
  },
  {
    path: '/pet',
    name: 'Pet',
    component: Pet
  },


]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
