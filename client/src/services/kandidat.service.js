import {instance} from "@/utils/axios";

class KandidatService {
    getKandidat() {
        return instance.get( '/kandidat/get')
    }

    getAll() {
        return instance.get('/kandidat/getall')
    }
}

export default new KandidatService();