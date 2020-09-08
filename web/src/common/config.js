import {local} from "../common/storage";

var theme = local.get('theme') || {};
export const config = {
    //项目名称,用于代码生成包导入
    projectName: 'ekgo',
    //版本
    version: 1.0,
    //token名称
    tokenName: "Authorization",
    //浏览器标题
    title: "后台管理",
    //是否显示logo
    logo: true,
    //菜单名字
    logoTitle: '后台管理',
    //是否加载本地菜单 本地菜单:src\menu\index.json,服务端接口:src\common\menu.js
    isLocalMenu: false,
    //是否显示布局大小
    isLayoutSize: true,
    //默认布局大小
    layoutSize: 'medium',
    //不经过token校验的路由
    routesWhiteList: ["/login", "/404", "/401"],
    //组件缓存
    keep_alive: [],
    //是否加载多语言
    isLanguage: true,
    //默认语言 zh和en
    language: 'zh',
    //最长请求时间
    requestTimeout: 10000,
    //操作正常code
    successCode: 200,
    //登录失效code
    invalidCode: 402,
    //无权限code
    noPermissionCode: 401,
    //图片预览地址
    pictureUrl: 'https://ek-shop.oss-cn-shanghai.aliyuncs.com/',
    //本地环境请求地址 npm run serve
    localUrl: 'http://127.0.0.1:95/',
    //测试环境请求地址 npm run test
    testUrl: 'http://127.0.0.1:95/',
    //正式环境正式地址 npm run build
    url: 'https://stage-admin.shoplitlive.com/',
    //主题
    theme: {
        //横纵布局 horizontal vertical
        layout: theme["layout"] || 'vertical',
        //固定头部
        head: theme["head"] || false,
        //多标签
        multi_label: theme["multi_label"] == false ? false : true,
        //菜单背景色
        menu_background: theme["menu_background"] || '#191a23',
        //菜单子集打开背景色
        menu_child_background: theme["menu_child_background"] || '#101117',
        //菜单字体颜色
        menu_text_color: theme["menu_text_color"] || '#b5b5b5',
        //菜单字体选中颜色
        menu_active_text_color: theme["menu_active_text_color"] || '#ffffff',
        //菜单选中颜色
        menu_active_background: theme["menu_active_background"] || '#5677FC',
        //顶栏文字
        top_text_color: theme["top_text_color"] || '#5a5e66',
        //顶栏背景
        top_background: theme["top_background"] || '#ffffff',
    }
}