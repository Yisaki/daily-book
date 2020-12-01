import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    //0:开始 1:结束
    periodTypeMap:{
      0:'开始',
      1:'结束'
    }
  },
  mutations: {
  },
  actions: {
  },
  modules: {
  }
})
