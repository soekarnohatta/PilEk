import axios from 'axios';

const API_URL = '/api/';

class AuthService {
    login(user) {
        return axios
            .post(API_URL + 'login', {
                user: {
                    username: user.username,
                    password: user.password
                }
            })
            .then(response => {
                if (response.data.userlogin.accessToken) {
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
