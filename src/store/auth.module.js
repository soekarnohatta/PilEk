import AuthService from '@/services/auth.service';

let user = JSON.parse(localStorage.getItem('user'));
export let initialState = user
    ? {statusUser: {loggedIn: true, vote: user.status, isAdmin: user.isAdmin}, user}
    : {statusUser: {loggedIn: false, vote: "Belum Memilih", isAdmin: false}, user: null};

export const auth = {
    namespaced: true,
    state: initialState,
    actions: {
        login({commit}, user) {
            return AuthService.login(user).then(
                user => {
                    commit('loginSuccess', user);
                    return Promise.resolve(user);
                },
                error => {
                    commit('loginFailure');
                    return Promise.reject(error);
                }
            );
        },
        logout({commit}) {
            AuthService.logout();
            commit('logout');
        },
    },
    mutations: {
        loginSuccess(state, user) {
            state.statusUser.loggedIn = true;
            state.statusUser.isAdmin = user.isAdmin
            state.user = user;
        },
        loginFailure(state) {
            state.statusUser.loggedIn = false;
            state.user = null;
        },
        logout(state) {
            state.statusUser.loggedIn = false;
            state.user = null;
        },
    }
};
