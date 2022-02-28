import AuthService from '../services/auth.service';

const user = JSON.parse(localStorage.getItem('user'));
const initialState = user
    ? {statusUser: {loggedIn: true, vote: user.status}, user}
    : {statusUser: {loggedIn: false, vote: "Belum Memilih"}, user: null};

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
