import {session,local} from "../../utils/storage";
import config from "../../config/index.config";
//单一状态树
const state = () => ({
    //项目名称
    title: config.title,
    /**
     * 布局方式（整体界面的排版方式）
     * layout-side -- 侧边布局
     * layout-head -- 顶部菜单
     * layout-comp -- 联动布局
     * */
    layout: config.layout == null
        ? "layout-side"
        : local.get("layout") == null
            ? config.layout
            : local.get("layout"),

    /**
     * 系统主题（整体色调）
     * light -- 白色主题
     * dark -- 暗色主题
     * night -- 夜间主题
     */
    theme: config.theme == null
        ? "theme-dark"
        : local.get("theme") == null
            ? config.theme
            : local.get("theme"),

    /**
     * 主题颜色(主题颜色)
     */
    color: config.color == null
        ? "theme-green"
        : local.get("color") == null
            ? config.color
            : local.get("color"),

    /**
     * 侧边状态
     * true  --  隐藏
     * false --  展开
     * */
    collapsed: config.collapsed == null
        ? false
        : local.get("collapsed") == null
            ? config.collapsed
            : local.get("collapsed"),

    /**
     * 菜单头部
     * true  --  隐藏
     * false --  展开
     * */
    logo: config.logo == null
        ? true
        : local.get("logo") == null
            ? config.logo
            : local.get("logo"),

    /**
     * 是否开启多标签页
     * true  --  隐藏
     * false --  展开
     * */
    tab: config.tab == null
        ? true
        : local.get("tab") == null
            ? config.tab
            : local.get("tab"),

    /**
     * 保持状态
     * true -- 是
     * false -- 否
     */
    keepAlive: config.keepAlive == null ? true : local.get("keepAlive") == null
        ? config.keepAlive : local.get("keepAlive")
    ,
    /**
     * 多标签页样式
     * pear-card-tab
     * pear-dot-tab
     */
    tabType: config.tabType == null
        ? "pear-dot-tab"
        : local.get("tabType") == null
            ? config.tabType
            : local.get("tabType"),

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
        : local.get("fixedHeader") == null
            ? config.fixedHeader
            : local.get("fixedHeader"),

    /**
     * 固定侧边
     * true
     * false
     */
    fixedSide: config.fixedSide == null
        ? true
        : local.get("fixedSide") == null
            ? config.fixedSide
            : local.get("fixedSide"),

    /**
     * 路由动画
     * fadeRight
     * fadeTop
     */
    routerAnimate: config.routerAnimate == null
        ? ""
        : local.get("routerAnimate") == null
            ? config.routerAnimate
            : local.get("routerAnimate"),

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
    panes: session.get("tabs")
        ? JSON.parse(session.get("tabs"))
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

    // 是否是移动端
    isMobile: false
})
//通过属性访问用于过滤操作
const getters = {}
//事件
const mutations = {
    // 修改全屏打开
    updateFullscreen(state) {
        state.fullscreen = !state.fullscreen;
    },
    //修改菜单
    updateCollapsed(state) {
        state.collapsed = !state.collapsed;
    },
    //修改路由刷新
    updateRouterActive(state) {
        state.routerActive = !state.routerActive;
    },
    // 是否开启选项卡模式
    updateTab(state) {
        state.tab = !state.tab;
    },
    //是否固定侧边
    updateFixedSide(state) {
        state.fixedSide = !state.fixedSide;
    },
    //是否固定头部
    updateFixedHeader(state) {
        state.fixedHeader = !state.fixedHeader;
    },
    // 是否显示LOGO
    updateLogo(state) {
        state.logo = !state.logo;
    },
    // 设置面板是否打开
    updateSettingOpen(state) {
        state.setting.opened = !state.setting.opened;
    },
    //修改菜单打开项
    updateOpenKey(state, { openKeys }) {
        //手风琴模式, 只保留当前打开的节点
        let length = openKeys.length;
        if (length > 0 && state.muiltOpen) {
            //除了最后打开的节点, 其他的节点可能包含父节点
            let otherKeys = openKeys.slice(0, length - 1);
            //最后一次打开的节点
            let lastKey = openKeys[length - 1];
            state.openKey = [
                ...otherKeys.filter(ok => lastKey.startsWith(ok)),
                lastKey
            ];
        } else {
            state.openKey = openKeys;
        }
        console.log(state.openKey)
    },
    updateLanguage(state, payload) {
        state.language = payload
        local.set('pear_lang', payload)
    },
    updateLayout(state, layout) {
        state.layout = layout;
    },
    updateTheme(state, theme) {
        state.theme = theme;
    },
    updateRouterAnimate(state, routerAnimate) {
        state.routerAnimate = routerAnimate;
    },
    updateKeepAlive(state) {
        state.keepAlive = !state.keepAlive;
    },
    updateColor(state, color) {
        state.color = color;
    },
    updateTabType(state, tabType) {
        state.tabType = tabType;
    },
    updateIsMobile(state, isMobile) {
        state.isMobile = isMobile;
    },
    updatePanes(state, panes) {
        state.panes = panes;
    },
}
//异步操作事件
const actions = {}
export default {state, getters, mutations, actions}
