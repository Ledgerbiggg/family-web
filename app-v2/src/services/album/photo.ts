import api from "@/services/api.ts";


/**
 * 获取指定相册下的所有照片
 * @param category 相册的id
 * @returns 一个 Promise 对象，结果是请求的响应数据
 */
export const albumPhotoListService = (category: number) => {
    return api.get("/album/" + category + "/photos");
}

/**
 * 获取相册照片
 * @param params 请求参数，包含相册和照片的相关信息
 * @returns 一个 Promise 对象，结果是请求的响应数据，响应类型为 'blob'
 */
export const albumPhotoService = (params: any) => {
    return api.get("/album/photo", params, 'blob');
}

