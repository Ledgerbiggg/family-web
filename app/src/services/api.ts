import axios, {AxiosResponse, AxiosError, InternalAxiosRequestConfig, AxiosHeaders} from 'axios';
import {message} from "ant-design-vue";
import router from "../router";

// 创建 Axios 实例
const api = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL,
    timeout: 10000,  // 设置请求超时
});

// 请求拦截器
api.interceptors.request.use(
    (config: InternalAxiosRequestConfig) => {
        // 如果 headers 不存在，则初始化为 AxiosHeaders 对象
        if (!config.headers) {
            config.headers = new AxiosHeaders(); // 使用 AxiosHeaders 构造函数初始化
        }

        // 可以在这里进行请求的配置，添加认证 token 等
        const token = localStorage.getItem('token');
        if (token) {
            // 将 token 添加到请求头
            config.headers.set('Authorization', `Bearer ${token}`);
        }

        config.withCredentials = true;

        return config; // 必须返回 config，否则请求无法继续
    },
    (error: AxiosError) => {
        // 请求出错时的处理
        console.error('Request error:', error);
        return Promise.reject(error);
    }
);


// 响应拦截器
api.interceptors.response.use(
    (response: AxiosResponse) => {
        // 判断响应类型是否为图片（Blob 类型）
        const isImage = response.headers['content-type'] && response.headers['content-type'].includes('image');

        if (isImage) {
            // 如果是图片（Blob 类型），直接返回原始响应数据
            return response.data;
        }
        if (response.status === 200 && response.data) {
            // 判断是否返回了业务错误（code != '10000'）
            if (response.data.code !== '10000') {
                const isBlob = response.data instanceof Blob;
                if (isBlob) {
                    return;
                }
                // 如果是错误信息，显示警告消息
                message.warn(response.data.message);

                // 登录过期需要重新登录
                if (response.data.code === '20008') {
                    router.push({name: 'Login'});
                }
                return Promise.reject(new Error(response.data.message));
            }

            // 如果有 token 返回，存储到本地
            if (response.headers.token) {
                localStorage.setItem('token', response.headers.token);
            }

            // 返回正常的业务数据
            return response.data.data || 'success';
        }

        return Promise.reject(new Error('Response status is not success'));
    },
    (error: AxiosError) => {
        var showError: number | string = ''
        // 响应出错时的处理
        if (error.response) {
            // 服务器响应错误
            console.error('Response error status:', error.response.status);
            showError = error.response.status
        } else {
            // 请求没有发出去，或者没有响应
            console.error('Response error:', error.message);
            showError = error.message
        }
        message.error('请求失败 状态码: ' + showError);
        return Promise.reject(error); // 可以根据需要修改，返回不同的错误信息
    }
);


// 定义请求参数的类型
interface Params {
    [key: string]: any;
}

// 定义请求配置的类型
interface RequestConfig {
    method: 'get' | 'post' | 'put' | 'delete';
    url: string;
    responseType?: 'json' | 'blob' | 'text';
    params?: Params;
    data?: Params;
}

// 定义一个返回的请求类型
interface RequestResponse<T = any> {
    data: T;
    status: number;
    statusText: string;
    headers: Record<string, string>;
    config: RequestConfig;
}

const http = {
    /**
     * GET 请求
     * @param url 请求地址
     * @param params 请求参数
     * @param responseType 响应数据类型，默认是 'json'
     * @returns 返回一个 Promise 对象，结果是请求的响应数据
     */
    get<T>(url: string, params?: Params, responseType: 'json' | 'blob' | 'text' = 'json'): Promise<RequestResponse<T>> {
        const config: RequestConfig = {
            method: 'get',
            url: url,
            responseType: responseType, // 设置响应类型
        };
        if (params) config.params = params;
        return api(config);
    },

    /**
     * POST 请求
     * @param url 请求地址
     * @param params 请求参数
     * @param responseType 响应数据类型，默认是 'json'
     * @returns 返回一个 Promise 对象，结果是请求的响应数据
     */
    post<T>(url: string, params?: Params, responseType: 'json' | 'blob' | 'text' = 'json'): Promise<RequestResponse<T>> {
        const config: RequestConfig = {
            method: 'post',
            url: url,
            responseType: responseType, // 设置响应类型
        };
        if (params) config.data = params;
        return api(config);
    },

    /**
     * PUT 请求
     * @param url 请求地址
     * @param params 请求参数
     * @param responseType 响应数据类型，默认是 'json'
     * @returns 返回一个 Promise 对象，结果是请求的响应数据
     */
    put<T>(url: string, params?: Params, responseType: 'json' | 'blob' | 'text' = 'json'): Promise<RequestResponse<T>> {
        const config: RequestConfig = {
            method: 'put',
            url: url,
            responseType: responseType, // 设置响应类型
        };
        if (params) config.params = params;
        return api(config);
    },

    /**
     * DELETE 请求
     * @param url 请求地址
     * @param params 请求参数
     * @param responseType 响应数据类型，默认是 'json'
     * @returns 返回一个 Promise 对象，结果是请求的响应数据
     */
    delete<T>(url: string, params?: Params, responseType: 'json' | 'blob' | 'text' = 'json'): Promise<RequestResponse<T>> {
        const config: RequestConfig = {
            method: 'delete',
            url: url,
            responseType: responseType, // 设置响应类型
        };
        if (params) config.params = params;
        return api(config);
    },
};

// 导出 http 实例
export default http;

