import Vue from 'vue'
import App from './App.vue'
import store from './store'
import G6 from "@antv/g6"

require('@/assets/main.scss');
console.log(G6.version)

Vue.config.productionTip = false

new Vue({
  store,
  render: h => h(App)
}).$mount('#app')
