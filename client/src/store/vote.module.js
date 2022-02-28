import VoteService from '../services/vote.service';

const user = JSON.parse(localStorage.getItem('user'));
const initialState = user
    ? {statusUser: {vote: user.status}, user}
    : {statusUser: {vote: "Belum Memilih"}, user: null};

export const vote = {
    namespaced: true,
    state: initialState,
    actions: {
        vote({commit}, id) {
            return VoteService.voteKandidat(id).then(resp => {
                commit('successVote')
                return Promise.resolve(resp);
            }, error => {
                commit('unsuccessVote')
                return Promise.reject(error);
            })
        },
    },
    mutations: {
        successVote(state) {
            state.statusUser.vote = "Sudah Memilih";
            state.user.status = "Sudah Memilih";
        },
        unsuccessVote(state) {
            state.statusUser.vote = "Belum Memilih";
        },
    }
};
