// src/router/index.ts
import { createRouter, createWebHistory } from 'vue-router'

// 定义路由
const routes = [
    // 登录界面
    {
        path: '/login',
        name: 'Login',
        component: () => import('../views/Login/LoginView.vue')
    },
    // 注册界面
    {
        path: '/register',
        name: 'Register',
        component: () => import('../views/Login/RegisterView.vue')
    },
    // 忘记密码
    {
        path: '/forgot-password',
        name: 'ForgotPassword',
        component: () => import('../views/Login/ForgotPasswordView.vue')
    }
]


// 创建路由实例
const router = createRouter({
    history: createWebHistory(), // 使用 HTML5 历史模式
    routes
});

// 导出路由
export default router
