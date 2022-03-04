import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import NProgress from 'nprogress/nprogress'
import 'nprogress/nprogress.css'
import axios from "axios";

Vue.use(Router);
export const router = new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home
        },
        {
            path: '/home',
            component: Home
        },
        {
            path: '/login',
            component: () => import('./views/Login.vue')
        },
        {
            path: '/profile',
            name: 'profile',
            component: () => import('./views/Profile.vue')
        },
        {
            path: '/vote',
            name: 'voting',
            component: () => import('./views/Vote')
        },
        {
            path: '/sudah',
            name: 'Thanks',
            component: () => import('./views/Thanks')
        },
        {
            path: "/:catchAll(.*)",
            name: 'error',
            component: () => import('./views/Error.vue')
        },
    ]
});

router.beforeEach(async (to, from, next) => {
    NProgress.start()

    const publicPages = ['/login', '/home'];
    const authRequired = !publicPages.includes(to.path);
    const loggedIn = localStorage.getItem('user');

    if (authRequired && !loggedIn) {
        next('/login');
    } else {
        next();
    }
});

axios.interceptors.response.use(
    response => response,
    error => {
        if ((error.response && error.response.data && error.response.data.errors.body) === "invalid or expired jwt") {
            NProgress.done()
            window.location.href = './login';
        }

    })