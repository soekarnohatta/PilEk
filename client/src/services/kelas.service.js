import axios from 'axios';

const API_URL = '/api/kelas/'

class KelasService {
    getKelas() {
        return axios.get(API_URL + 'get');
    }

    getAll() {
        return axios.get(API_URL + 'getall',);
    }
}

export default new KelasService();