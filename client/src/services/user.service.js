import AuthHeaderService from "@/services/auth-header";
import {instance} from "@/utils/axios";

class UserService {
    getUser() {
        return instance.get('/user/current', AuthHeaderService.getToken());
    }
}

export default new UserService();