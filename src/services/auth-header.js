class AuthHeaderService {
    getToken() {
        let user = JSON.parse(localStorage.getItem('user'));
        if (user) {
            return {
                headers: {'Authorization': 'Bearer ' + user.accessToken},
            };
        }
    }
}

export default new AuthHeaderService();
