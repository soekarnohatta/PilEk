import axios from "axios";
import NProgress from "nprogress";
import AuthService from "@/services/auth.service";

export const instance = axios.create({
    baseURL: '/api/'
});

instance.interceptors.response.use(
    response => response,
    error => {
        if ((error.response && error.response.data && error.response.data.errors.body) === "invalid or expired jwt") {
            NProgress.done()
            AuthService.logout()
            window.location.href = './login';
        }
    }
)