import Vue from 'vue'
import axios from './plugins/axios'

Vue.config.productionTip = false

new Vue({
  el: '#app',
  data () {
    return {
      info: null
    }
  },
  mounted () {
    axios
      .get('https://localhost:1323/standard_codes')
      .then(response => (this.info = response))
  }
})
