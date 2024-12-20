// 定义 InviteInfo 对象的类型
import api from "../api.ts";

export interface InviteInfo {
    inviterPhone: string
    inviterRealName: string
    invitedAdmin: boolean
    description: string
}

// 定义 InviteUser 对象的类型
export interface InviteUser {
    inviteUid: string
    realName: string
    username: string
    captcha: string
}

export const inviteRegisterService = (inviteUser: InviteUser) => {
    return api.post("/invite/register", {...inviteUser})
}

export const inviteInfoService = (uid: string) => {
    return api.get("/invite/info", {uid})
}
