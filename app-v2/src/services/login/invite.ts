import api from "@/services/api.ts";

export interface InviteInfo {
    inviterPhone: string
    inviterRealName: string
    invitedAdmin: boolean
    description: string
}


export interface InviteUser {
    inviteUid: string
    realName: string
    username: string
    captcha: string
}
/**
 * 邀请注册服务
 * @param inviteUser 一个 InviteUser 对象，包含邀请人id、真实姓名、用户名、验证码信息
 * @returns 一个 Promise 对象，结果是请求的响应数据
 */
export const inviteRegisterService = (inviteUser: InviteUser) => {
    return api.post("/invite/register", {...inviteUser})
}

/**
 * 获取邀请信息服务
 * @param uid 邀请link的id
 * @returns 一个 Promise 对象，结果是请求的响应数据，包含邀请人的真实姓名、电话号码、是否是管理员
 */
export const inviteInfoService = (uid: string) => {
    return api.get("/invite/info", {uid})
}
