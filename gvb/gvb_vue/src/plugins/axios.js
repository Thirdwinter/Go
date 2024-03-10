import axios from 'axios';
import { router } from '../router/index';

const as = axios.create({
    baseURL: '/Api',
    timeout: 5000,
    headers: {
        'Content-Type': 'application/json;charset=utf-8',
        //'Authorization':'',
        //'RefreshAuthorization':'',
    },
});
let atoken = localStorage.getItem('atoken');
let rtoken = localStorage.getItem('rtoken');

// 请求拦截
as.interceptors.request.use(
    // config => {
    //     let atoken = localStorage.getItem('atoken');
    //     let rtoken = localStorage.getItem('rtoken');
    //         config.headers.Authorization = localStorage.getItem('atoken'); // 添加 Authorization 头部
    //         //config.headers.RefreshAuthorization = localStorage.getItem('rtoken'); // 添加 Authorization 头部
    //         console.log("已添加 Authorization 头部:", config.headers.Authorization);
    //         console.log("已添加 RefreshAuthorization 头部:", config.headers.RefreshAuthorization);
    //         return config;

    //     // return config;
    // },
    // error => {
    //     return Promise.reject(error);
    // }
);


// 响应拦截
as.interceptors.response.use(
    response => {
        let atoken = localStorage.getItem('atoken');
        let rtoken = localStorage.getItem('rtoken');
        const newConfig = { ...response.config };

        const responseData = response.data;
        if (responseData && responseData.code === 200) {
            // 如果响应中的 code 为 200，则直接返回原始响应
            return response;
        } else {
            switch (responseData.code) {
                case 4001:
                    localStorage.removeItem("atoken")
                    localStorage.removeItem("rtoken")
                    console.log("4001:",responseData.code)
                    // 处理 code 为 4001 的情况
                    router.replace({
                        path: '/login',
                        query: {
                            redirect: router.currentRoute.value.fullPath
                        }
                    });
                    break;
                case 4003:
                    console.log("4003:",responseData.code)
                    localStorage.removeItem("atoken")
                    localStorage.removeItem("rtoken")
                    router.replace({
                        path: '/login',
                        query: {
                            redirect: router.currentRoute.value.fullPath
                        }
                    });
                    break;
                case 4004:
                    console.log("4004:",responseData.code)
                    localStorage.removeItem("atoken")
                    localStorage.removeItem("rtoken")
                    router.replace({
                        path: '/login',
                        query: {
                            redirect: router.currentRoute.value.fullPath
                        }
                    });
                    break;
                case 4007:
                    console.log("4007:",responseData.code)
                    localStorage.removeItem("atoken")
                    localStorage.removeItem("rtoken")
                    router.replace({
                        path: '/login',
                        query: {
                            redirect: router.currentRoute.value.fullPath
                        }
                    });
                    break;
                case 4002:
                    console.log("4002:", responseData.code);
                    // // 在请求头中加入 RefreshAuthorization 字段
                    // newConfig.headers.RefreshAuthorization = localStorage.getItem('rtoken');
                    // // 重新发送请求
                    // // 注意：返回新的配置，避免再次进入响应拦截器
                // return as(newConfig);
                    localStorage.removeItem("atoken")
                    localStorage.removeItem("rtoken")
                    router.replace({
                        path: '/login',
                        query: {
                            redirect: router.currentRoute.value.fullPath
                        }
                    });
                    break;
                case 4005:
                    console.log("4005:",responseData.code)
                    // // localStorage.removeItem("atoken");
                    // // localStorage.removeItem("rtoken");
                    // localStorage.setItem("atoken",responseData.atoken);
                    // localStorage.setItem("rtoken",responseData.rtoken);
                    // return as(response.config);
                    localStorage.removeItem("atoken")
                    localStorage.removeItem("rtoken")
                    router.replace({
                        path: '/login',
                        query: {
                            redirect: router.currentRoute.value.fullPath
                        }
                    });
                    break;
                default:
                    console.log("default:")
                    // 处理其他情况
                    break;
            }
        }
        // 如果需要对响应进行修改，请在这里返回修改后的响应
        return response;
    },
    error => {
        // 处理请求错误
        return Promise.reject(error);
    }
);
export default as;
