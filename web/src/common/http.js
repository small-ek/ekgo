import axios from 'axios';
import {config} from './config' //配置
//配置
import {local} from './storage.js'
import {Notification} from "element-ui";
//导入axios组件
var base_url = ''

switch (process.env.NODE_ENV) {
    case 'development'://本地地址
        base_url = config.localUrl //这里是本地的请求url
        break
    case 'alpha': //测试地址
        base_url = config.testUrl
        break
    case 'production'://生产地址
        base_url = config.url //生产环境url
        break
}
var user = local.get("user")
var Authorization = ''

if (user) {
    Authorization = user['Authorization']
}
var headers = {}
headers[config.tokenName] = Authorization
var http = axios.create({
    baseURL: base_url,
    timeout: config.requestTimeout,
    headers: headers
});

// 添加请求拦截器
http.interceptors.request.use(function (config) {
    return config;
}, function (error) {
    // 对请求错误做些什么
    console.log(error);
    return Promise.reject(error);

});

// 添加响应拦截器
http.interceptors.response.use(function (response) {
    // 登录权限不足,重写登录
    if (response.status && response.data.code === config.invalidCode) {
        //跳转
        Notification.error({
            title: "错误提示",
            message: "当前接口权限不足"
        })
        /*window.location.href = '/login'*/
    }
    return response;
}, function (error) {
    if (error.request.status == 0 || error.request.status == 500) {
        Notification.error({
            title: "数据请求失败",
            message: "请联系系统管理员"
        })
    } else {
        Notification.error({
            title: "当前数据请求权限不足",
            message: "参数错误，请联系系统管理员"
        })
    }
    // 对响应错误做点什么
    return Promise.reject(error);
});
export default http;
