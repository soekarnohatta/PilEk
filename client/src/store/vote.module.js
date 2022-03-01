import VoteService from '@/services/vote.service';
import {initialState} from "@/store/auth.module";

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
