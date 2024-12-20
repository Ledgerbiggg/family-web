// 封装验证重置密码的函数
import api from "../api.ts";

export interface User {
    username: string;
    password: string;
    captcha: string;
}

export const loginService =  (user: User) => {
    return  api.post('/login', {...user});
}

export const resetPasswordService =  (user: User) => {
    return api.post('/verify', {...user});
};














