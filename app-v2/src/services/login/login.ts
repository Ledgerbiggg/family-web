// 封装验证重置密码的函数
import api from "@/services/api.ts";

export interface LoginUser {
    username: string;
    password: string;
    captcha: string;
}
/**
 * 登录服务
 * @param user 一个 LoginUser 对象
 * @returns 一个 Promise 对象，结果是请求的响应数据
 */
export const loginService =  (user: LoginUser) => {
    return  api.post('/login', {...user});
}

/**
 * 验证重置密码的服务
 * @param user 一个 LoginUser 对象
 * @returns 一个 Promise 对象，结果是请求的响应数据
 */
export const resetPasswordService =  (user: LoginUser) => {
    return api.post('/verify', {...user});
};














