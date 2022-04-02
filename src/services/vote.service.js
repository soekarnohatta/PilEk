import AuthHeaderService from "./auth-header";
import {instance} from "@/utils/axios";

class VoteService {
    voteKandidat(id) {
        return instance
            .put('/user/vote?idkandidat=' + id, {}, AuthHeaderService.getToken())
            .then(response => {return response.data})
    }
}

export default new VoteService();