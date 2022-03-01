import Vuex from 'vuex';
import Vue from "vue";
import {auth} from './auth.module';
import {vote} from './vote.module';

Vue.use(Vuex)
export default new Vuex.Store({
    state: {
        version: '',
        count: 1
    },
    getters: {},
    modules: {
        auth,
        vote,
    }
});

export function currentUser() {
    let user = JSON.parse(localStorage.getItem('user'))
    if (user) {
        return user
    }
   return null
}
