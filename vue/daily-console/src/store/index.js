import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    //0:开始 1:结束
    periodTypeMap:{
      0:'开始',
      1:'结束'
    },
    petTypeMap:{
      0:'拉屎',
      1:'上药',
      2:'洗澡',
      3:'吃营养膏'
    }
  },
  mutations: {
  },
  actions: {
  },
  modules: {
  }
})
