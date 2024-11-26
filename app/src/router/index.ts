// src/router/index.ts
import {createRouter, createWebHistory} from 'vue-router'

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
    // Home界面
    {
        path: '/home',
        name: 'Home',
        component: () => import('../views/home/HomeView.vue')
    }
]


// 创建路由实例
const router = createRouter({
    history: createWebHistory(), // 使用 HTML5 历史模式
    routes
});

// 路由守卫(放行登录+注册+忘记密码)
router.beforeEach((to, from, next) => {
    //如果用户访问的登录页，直接放行
    if (to.path === '/login' || to.path === '/register' || to.path === '/forgot-password') return next()
    // 如果token值存在,直接放行
    const token = localStorage.getItem('token')
    if (token) return next()
    // 如果token值不存在，跳转到登录页
    next('/login')
})


// 导出路由
export default router
