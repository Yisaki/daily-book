import Vue from 'vue'
import VueRouter from 'vue-router'

import Period from '@/views/period/Period'
import PeriodList from '@/views/period/PeriodList'

import Pet from '@/views/pet/Pet'
import PetList from '@/views/pet/PetList'

import Index from '@/views/Index'



Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Index',
    component: Index
  },  
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
  {
    path: '/pet',
    name: 'Pet',
    component: Pet
  },
  {
    path: '/petList',
    name: 'PetList',
    component: PetList
  }, 


]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
