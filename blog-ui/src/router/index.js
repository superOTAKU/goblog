import {createRouter, createWebHistory} from 'vue-router'

import Home from '../components/Home'

//单页应用的路由管理
const routes = [
    {
        path: '/',
        component: Home
    }
]
let router = createRouter({
    history: createWebHistory(),
    routes
})

export default router