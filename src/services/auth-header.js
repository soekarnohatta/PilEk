import {currentUser} from "@/store";

class AuthHeaderService {
    getToken() {
        return {
            headers: {'Authorization': 'Bearer ' + currentUser().accessToken},
        };
    }
}

export default new AuthHeaderService();
