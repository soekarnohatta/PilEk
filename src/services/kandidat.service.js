import axios from 'axios';

const API_URL = '/api/kandidat/'

class KandidatService {
    getKandidat() {
        return axios.get(API_URL + 'get')
    }

    getAll() {
        return axios.get(API_URL + 'getall')
    }
}

export default new KandidatService();