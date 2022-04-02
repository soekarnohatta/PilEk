import {instance} from "@/utils/axios";

class AuthService {
    login(user) {
        return instance
            .post('login', {
                user: {
                    username: user.username,
                    password: user.password
                }
            })
            .then(response => {
                if (response.data.userlogin && response.data.userlogin.aktif && response.data.userlogin.accessToken) {
                    localStorage.setItem('user', JSON.stringify(response.data.userlogin));
                }

                return response.data.userlogin;
            });
    }

    logout() {
        localStorage.removeItem('user');
    }
}

export default new AuthService();
