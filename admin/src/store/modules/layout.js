import config from "../../config/index.config";
//单一状态树
const state = () => ({
    /**
     * 布局方式（整体界面的排版方式）
     * layout-side -- 侧边布局
     * layout-head -- 顶部菜单
     * layout-comp -- 联动布局
     * */
    layout: config.layout == null
            ? "layout-side"
            : localStorage.getItem("layout") == null
            ? config.layout
            : localStorage.getItem("layout"),

    /**
     * 系统主题（整体色调）
     * light -- 白色主题
     * dark -- 暗色主题
     * night -- 夜间主题
     */
    theme: config.theme == null
            ? "theme-dark"
            : localStorage.getItem("theme") == null
            ? config.theme
            : localStorage.getItem("theme"),

    /**
     * 主题颜色(主题颜色)
     */
    color: config.color == null
            ? "theme-green"
            : localStorage.getItem("color") == null
            ? config.color
            : localStorage.getItem("color"),

    /**
     * 侧边状态
     * true  --  隐藏
     * false --  展开
     * */
    collapsed: config.collapsed == null
            ? false
            : localStorage.getItem("collapsed") == null
            ? config.collapsed
            : localStorage.getItem("collapsed"),

    /**
     * 菜单头部
     * true  --  隐藏
     * false --  展开
     * */
    logo: config.logo == null
            ? true
            : localStorage.getItem("logo") == null
            ? config.logo
            : localStorage.getItem("logo"),

    /**
     * 是否开启多标签页
     * true  --  隐藏
     * false --  展开
     * */
    tab: config.tab == null
            ? true
            : localStorage.getItem("tab") == null
            ? config.tab
            : localStorage.getItem("tab"),

    /**
     * 保持状态
     * true -- 是
     * false -- 否
     */
    keepAlive: config.keepAlive == null ? true : localStorage.getItem("keepAlive") == null
        ? config.keepAlive : localStorage.getItem("keepAlive")
    ,
    /**
     * 多标签页样式
     * pear-card-tab
     * pear-dot-tab
     */
    tabType: config.tabType == null
            ? "pear-dot-tab"
            : localStorage.getItem("tabType") == null
            ? config.tabType
            : localStorage.getItem("tabType"),

    /**
     * 侧边菜单栏宽度
     * 单位:px
     * */
    sideWitch: config.sideWidth == null ? 220 : config.sideWidth,
    /**
     * 侧边菜单栏宽度(折叠)
     * 单位:px
     * */
    collapsedWidth: config.collapsedWidth == null ? 60 : config.collapsedWidth,

    /**
     * 固定头部
     * true
     * false
     */
    fixedHeader: config.fixedHeader == null
            ? true
            : localStorage.getItem("fixedHeader") == null
            ? config.fixedHeader
            : localStorage.getItem("fixedHeader"),

    /**
     * 固定侧边
     * true
     * false
     */
    fixedSide: config.fixedSide == null
            ? true
            : localStorage.getItem("fixedSide") == null
            ? config.fixedSide
            : localStorage.getItem("fixedSide"),

    /**
     * 路由动画
     * fadeRight
     * fadeTop
     */
    routerAnimate: config.routerAnimate == null
            ? ""
            : localStorage.getItem("routerAnimate") == null
            ? config.routerAnimate
            : localStorage.getItem("routerAnimate"),

    /**
     * 配色列表
     */
    colorList: config.colorList,

    /**
     * 主题面板状态
     * true
     * false
     */
    setting: {
        opened: false
    },

    // 选项卡内容存储
    panes: sessionStorage.getItem("pear_tabs")
        ? JSON.parse(sessionStorage.getItem("pear_tabs"))
        : [],
    // 当前激活选项卡
    activeKey: "",
    // 当前打开菜单
    openKey: [],
    // 手风琴配置
    muiltOpen: true,
    // 路由刷新辅助变量
    routerActive: true,
    // 浏览器全屏
    fullscreen: false,
    // 路由列表
    routes: [],

    // 是否是移动端
    isMobile: false
})
//通过属性访问用于过滤操作
const getters = {

}
//事件
const mutations = {
    // 修改全屏打开
    updateFullscreen(state) {
        state.fullscreen = !state.fullscreen;
    },
}
//异步操作事件
const actions = {

}
export default { state, getters, mutations, actions }
