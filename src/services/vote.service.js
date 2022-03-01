import axios from 'axios';
import AuthHeaderService from "./auth-header";

const API_URL = '/api/user/vote'

class VoteService {
    voteKandidat(id) {
        return axios.put(API_URL + "?idkandidat=" + id, {}, AuthHeaderService.getToken())
            .then(response => {
                return response.data
            })
    }
}

export default new VoteService();