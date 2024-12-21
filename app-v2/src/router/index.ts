// src/router/index.ts
import {
    createRouter,
    createWebHistory,
    NavigationGuardNext,
    RouteLocationNormalized,
    RouteLocationNormalizedLoaded
} from 'vue-router'

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
    // 相册
    {
        path: '/album',
        name: 'Album',
        component: () => import('../views/home/album/CategoryView.vue'),
        // beforeEnter: (_: RouteLocationNormalized, from: RouteLocationNormalizedLoaded, next: NavigationGuardNext) => {
        //     checkFromHome(from, next); // 调用抽离的函数
        // },
    },
    // 照片
    {
        path: '/photo/:category',
        name: 'Photo',
        component: () => import('../views/home/album/PhotoView.vue')
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


// src/utils/routerGuard.ts
export const checkFromHome = (from: any, next: any): void => {
    // 判断来源路径是否是 '/home'
    if (from.path === '/home') {
        next(); // 允许访问
    } else {
        next({name: 'Home'}); // 如果不是从 /home 进入，则重定向到首页
    }
};

// 创建路由实例
const router = createRouter({
    history: createWebHistory(), // 使用 HTML5 历史模式
    routes
});


// 路由守卫(放行登录+注册+忘记密码)
router.beforeEach((_: RouteLocationNormalized, __: RouteLocationNormalizedLoaded, next: NavigationGuardNext) => {
    return next()
})


// 导出路由
export default router
