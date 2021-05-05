/**
 * Project default config.
 *
 * 项目默认配置信息
 *
 * store 初始化时决定使用 localstore 本地 / index.config.js 默认配置
 */
export default {
    //请求地址
    url: "http://127.0.0.1:88/",
    //请求超时
    timeout: 60000,
    //请求参数key
    tokenName: "authorization",
    //登录超时时间
    loginTimeout: 86400,
    //登录失效code
    invalidCode: 402,
    //网站图标
    image: "https://gitee.com/pear-admin/Pear-Admin-Layui/raw/master/admin/images/logo.png",
    //网站名称
    title: "Ekgo Admin",
    //默认使用的布局
    layout: "layout-side",
    //默认的主题
    theme: "theme-dark",
    //当前菜单的状态
    menuCollapsed: false,
    //是否本地菜单
    isLocalMenu: false,
    //是否显示菜单头
    logo: true,
    //是否开启多标签页
    tab: true,
    /**
     * 参数: 侧边菜单栏宽度
     * 单位: px
     * */
    sideWidth: 240,
    /**
     * 参数: 侧边菜单栏宽度(折叠)
     * 单位: px
     * */
    collapsedWidth: 70,
    /**
     * 参数: 是否固定侧边
     * true
     * false
     * */
    fixedSide: true,
    /**
     * 参数：是否开启全局 keep - Alive 状态缓存
     * true
     * false
     * */
    keepAlive: true,
    /**
     * 参数: 是否固定顶部
     * true
     * false
     */
    fixedHeader: true,
    /**
     * 参数: 选项卡样式
     * pear-card-tab
     * pear-dot-tab
     */
    tabType: "pear-dot-tab",
    /**
     * 参数: 主题颜色
     * color - list
     */
    color: "#36b368",
    /**
     * 参数: 可选的主题颜色列表
     * color - key
     */
    colorList: [
        "#2d8cf0",
        "rgb(255, 184, 0)",
        "rgb(255, 87, 34)",
        "#36b368",
        "black",
        "gray"
    ],
    /**
     * 参数: 路由动画
     * fade-right
     * fade-top
     */
    routerAnimate: "fade-top",
    /**
     * 参数: 通栏
     * true
     * false
     */
    passbar: true,
    /**
     * 参数：默认的tab栏
     */
    defaultTab: "dashboard-console",
    /**
     * 参数：国际化默认语言
     *
     * zh-CN 中文
     * en-US 英文
     */
    defaultLanguage: 'zh-CN'
};
