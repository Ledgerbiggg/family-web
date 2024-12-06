// src/router/index.ts
import {createRouter, createWebHistory} from 'vue-router'
import {message} from "ant-design-vue";

// 定义路由
const routes = [
    // 登录界面
    {
        path: '/login',
        name: 'Login',
        component: () => import('../views/login/LoginView.vue')
    },
    // 注册界面
    {
        path: '/register',
        name: 'Register',
        component: () => import('../views/login/RegisterView.vue')
    },
    // 忘记密码
    {
        path: '/forgot-password',
        name: 'ForgotPassword',
        component: () => import('../views/login/ForgotPasswordView.vue')
    },
    // 邀请注册
    {
        path: '/invite-register',
        name: 'InviteRegister',
        component: () => import('../views/login/InviteRegister.vue')
    },
    // Home界面
    {
        path: '/',
        redirect: '/home'
    },
    {
        path: '/home',
        name: 'Home',
        component: () => import('../views/HomeView.vue')
    },
    {
        path: '/album',
        name: 'Album',
        component: () => import('../views/home/AlbumView.vue')
    },
    // 404
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: () => import('../views/other/NotFoundView.vue')
    },
    // test
    {
        path: '/test',
        name: 'Test',
        component: () => import('../views/other/TestView.vue')
    }
]


// 创建路由实例
const router = createRouter({
    history: createWebHistory(), // 使用 HTML5 历史模式
    routes
});


// 路由守卫(放行登录+注册+忘记密码)
router.beforeEach((to, from, next) => {
    return  next()
})


// 导出路由
export default router
