import Vue from 'vue';
import App from './App.vue';
import {router} from './utils/router';
import store from './store';
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.config.productionTip = false;

new Vue({
    store,
    router,
    render: h => h(App)
}).$mount('#app');
