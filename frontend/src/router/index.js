import Vue from 'vue';
import VueRouter from 'vue-router';
import ViewProcess from '../views/ViewProcess';
import ViewBash from '../views/ViewBash';
import ViewSystemInfo from '../views/ViewSystemInfo';

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        redirect: {name: 'ViewProcess'}
    },
    {
        path: '/process',
        name: 'ViewProcess',
        component: ViewProcess
    },
    {
        path: '/bash',
        name: 'ViewBash',
        component: ViewBash
    },
    {
        path: '/system-info',
        name: 'ViewSystemInfo',
        component: ViewSystemInfo
    }
];

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
});

export default router;