import Vue from 'vue'

import Success from '../components/message/Success'
import Warning from '../components/message/Warning'
import Info from '../components/message/Info'
import Error from '../components/message/Error'

Vue.component('message-success', Success)
Vue.component('message-warning', Warning)
Vue.component('message-info', Info)
Vue.component('message-error', Error)
