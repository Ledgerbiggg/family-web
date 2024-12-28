// 获取相册照片数据
import api from "@/services/api.ts";

export const albumPhotoService = (): any => {
    return api.get("/album/category-list");
}