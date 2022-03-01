import {currentUser} from "@/store";

class AuthHeaderService {
    getToken() {
       // if (user) {
        // this.user = null;
            return {
                headers: {'Authorization': 'Bearer ' + currentUser().accessToken},
            };
       // }
    }
}

export default new AuthHeaderService();
