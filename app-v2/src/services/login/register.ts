import api from "../api.ts";

export interface RegisterUser {
    username: string;
    password: string;
    confirmPassword: string;
    captcha: string;
}

export const registerService =  (user: RegisterUser) => {
    return api.post("/register", { ...user });  // 返回实际的数据
}
