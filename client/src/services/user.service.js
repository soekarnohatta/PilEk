import axios from 'axios';
import AuthHeaderService from "@/services/auth-header";

const API_URL = '/api/user/'

class UserService {
    getUser() {
        return axios.get(API_URL + 'current', AuthHeaderService.getToken());
    }
}

export default new UserService();