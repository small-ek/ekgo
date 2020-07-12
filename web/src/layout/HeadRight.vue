<template>
    <div>
        <!--首页-->
        <el-tooltip :content="$t('home')" placement="bottom">
            <span class="header-btn" @click="home">
                <i class="fa fa-home fa-lg"></i>
            </span>
        </el-tooltip>
        <!--主题-->
        <el-tooltip :content="$t('theme')" placement="bottom">
            <span class="header-btn" @click="theme=true">
                 <i class="fa fa-home fa-tachometer"></i>
            </span>
        </el-tooltip>
        <!--搜索-->
        <el-tooltip :content="$t('search_menu')" placement="bottom">
            <span class="header-btn hidden-xs-only" @click="showSearch = !showSearch">
                <i class="fa fa-search fa-lg"></i>
            </span>
        </el-tooltip>
        <span class="header-btn hidden-xs-only" v-show="showSearch">
                <el-select ref="headerSearchSelect" v-model="search" filterable :placeholder="$t('search')" class="header-search-select" @change="change" @blur="showSearch=false">
                    <el-option v-for="(item,index) in menuOptions" :key="index" :value="item.path" :label="item.title.join(' > ')"/>
                </el-select>
        </span>
        <!--刷新-->
        <el-tooltip :content="$t('refresh')" placement="bottom">
                <span class="header-btn" @click="reload">
                    <i class="fa fa-refresh fa-lg"></i>
                </span>
        </el-tooltip>
        <!--全屏-->
        <el-tooltip :content="$t('full_screen')" placement="bottom">
                <span class="header-btn hidden-xs-only" @click="FullScreen">
                    <i class="fa fa-arrows-alt fa-lg"></i>
                </span>
        </el-tooltip>
        <!--布局大小-->
        <el-tooltip :content="$t('layout_size')" placement="bottom" v-show="this.$Config.isLayoutSize">
            <el-dropdown trigger="click" @command="setSize">
                    <span class="header-btn hidden-xs-only el-dropdown-link">
                        <i class="fa fa-text-height fa-lg"></i>
                    </span>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item v-for="item of sizeOptions" :key="item.value" :disabled="size==item.value" :command="item.value">
                        <template v-if="language=='中文'">
                            {{item.cnLabel }}
                        </template>
                        <template v-else>
                            {{item.label }}
                        </template>
                    </el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </el-tooltip>
        <!--语言-->
        <el-dropdown class="header-btn" v-show="this.$Config.isLanguage">
            <span>{{language}}</span>
            <el-dropdown-menu slot="dropdown">
                <el-dropdown-item v-for="item in languageList" :key="item.value" :disabled="language==item.title" @click.native="lang(item.value)">
                    {{item.title}}
                </el-dropdown-item>
            </el-dropdown-menu>
        </el-dropdown>
        <!--用户信息-->
        <el-dropdown class="header-btn">
            <span>{{user.username}}
                <i class="el-icon-arrow-down el-icon--right"></i>
            </span>
            <el-dropdown-menu slot="dropdown">
                <el-dropdown-item @click.native="update_password()">
                    <i class="fa fa-user-circle-o"></i>
                    {{$t('change_password')}}
                </el-dropdown-item>
                <el-dropdown-item @click.native="logout()">
                    <i class="fa fa-key"></i>
                    {{$t('exit_system')}}
                </el-dropdown-item>
            </el-dropdown-menu>
        </el-dropdown>
        <!--修改密码-->
        <template>
            <el-dialog :fullscreen="this.$Func.isFullscreen()" title="修改密码" lock-scroll :visible.sync="dialog" style="text-align: left">
                <el-form :model="form" ref="form">
                    <el-form-item label="原有密码" prop="old_password">
                        <el-input type="password" v-model="form.old_password" autocomplete="off" show-password></el-input>
                    </el-form-item>
                    <el-form-item label="新密码" prop="password">
                        <el-input type="password" v-model="form.password" autocomplete="off" show-password></el-input>
                    </el-form-item>
                </el-form>
                <div slot="footer">
                    <el-button @click="dialog = false">{{$t('cancel')}}</el-button>
                    <el-button type="primary" @click="onSubmit('form')" :loading="loading">{{ loading ? $t('wait') : $t('submit') }}</el-button>
                </div>
            </el-dialog>
        </template>
        <!--主题设置-->
        <el-drawer :title="$t('set_temem')" :visible.sync="theme" :size="this.$Func.searchWidth()" style="text-align: left">
            <div class="search-box">
                <div class="search-head">
                    <el-button @click="reset('form')">{{$t('reset')}}</el-button>
                    <el-button type="primary" @click="submit('config')" :loading="loading">{{ loading ? $t('wait') :
                        $t('submit') }}
                    </el-button>
                </div>
                <div class="search-body">
                    <el-form ref="config" :model="config" label-position="top" label-width="80px">
                        <el-form-item :label="$t('menu_style_setting')">
                            <el-radio v-model="config.menu" label="bright">{{$t('bright_color')}}</el-radio>
                            <el-radio v-model="config.menu" label="dark">{{$t('dark')}}</el-radio>
                        </el-form-item>
                        <el-form-item :label="$t('top_bar_settings')">
                            <el-radio v-model="config.top" label="bright">{{$t('bright_color')}}</el-radio>
                            <el-radio v-model="config.top" label="dark">{{$t('dark')}}</el-radio>
                        </el-form-item>
                        <el-form-item :label="$t('navigation_mode')">
                            <el-radio v-model="config.layout" label="vertical">{{$t('side_menu_layout')}}</el-radio>
                            <el-radio v-model="config.layout" label="horizontal">{{$t('top_menu_layout')}}</el-radio>
                        </el-form-item>
                        <el-form-item :label="$t('multi_label')">
                            <el-switch v-model="config.multi_label"></el-switch>
                        </el-form-item>
                        <el-form-item :label="$t('fixed_head')">
                            <el-switch v-model="config.head"></el-switch>
                        </el-form-item>
                    </el-form>
                </div>
            </div>
        </el-drawer>
    </div>
</template>
<script>
    import {local} from "../common/storage";
    import getMenu from '../common/menu.js';

    export default {
        data() {
            return {
                form: {
                    password: ''
                },
                dialog: false,
                search: '',
                menu: [],
                menuOptions: [],
                language: '',
                showSearch: false,
                languageList: [
                    {
                        'title': '中文',
                        'value': 'zh'
                    }, {
                        'title': 'English',
                        'value': 'en'
                    }
                ],
                size: 'default',
                user: {},
                sizeOptions: [
                    {
                        label: 'Default',
                        cnLabel: '默认',
                        value: 'default'
                    },
                    {
                        label: 'Medium',
                        cnLabel: '中',
                        value: 'medium'
                    },
                    {
                        label: 'Small',
                        cnLabel: '小',
                        value: 'small'
                    },
                    {
                        label: 'Mini',
                        cnLabel: '微小',
                        value: 'mini'
                    }
                ],
                Breadcrumb: [],
                list: [],
                theme: false,
                loading: false,
                config: {
                    menu: 'dark',
                    top: 'bright',
                    layout: 'vertical',
                    multi_label: true,
                    head: false
                }
            };
        },
        created() {
            this.user = local.get("user")
            let self = this;
            if (local.get("theme")) {
                this.config = local.get("theme")
            }
            //获取当前语言
            this.getLanguage();
            //设置下拉框菜单
            this.getMenu().then(function (result) {
                self.menuOptions = self.generateRoutes(result);
            })
            this.set_theme()
        },
        methods: {
            submit() {
                let config = this.config
                //菜单设置
                switch (config.menu) {
                    case "dark":
                        config.menu_background = "#191a23";
                        config.menu_child_background = "#101117";
                        config.menu_text_color = "#C0C0C0";
                        config.menu_active_text_color = "#ffffff";
                        config.menu_active_background = "#5677FC";
                        break
                    case "bright":
                        config.menu_background = "#ffffff";
                        config.menu_child_background = "#ffffff";
                        config.menu_text_color = "#2b2b2b";
                        config.menu_active_text_color = "#5677FC";
                        config.menu_active_background = "#ffffff";
                        break
                }
                //顶栏设置
                switch (config.top) {
                    case "dark":
                        config.top_text_color = "#ffffff";
                        config.top_background = "#191a23";
                        break
                    case "bright":
                        config.top_text_color = "#5a5e66";
                        config.top_background = "#ffffff";
                        break
                }

                local.set("theme", config)
                window.location.reload();
            },
            /**
             * 设置主题
             */
            set_theme() {
                let style = document.createElement("style");
                style.id = "THEME";
                let theme = this.$Config.theme
                if (theme) {
                    style.innerHTML = `
                        .left{
                            background: ${theme.menu_background} !important;
                        }
                        .el-menu{
                            border-right:0px
                        }
                        .el-menu-item:hover{
                            color:${theme.menu_active_text_color} !important;
                        }
                        .logo_title{
                            color:${theme.menu_text_color} !important;
                        }
                        .el-submenu .el-menu-item{
                            height: 45px;
                            line-height: 45px;
                            background: ${theme.menu_child_background} !important;
                        }
                        .el-menu-item.is-active{
                        background:${theme.menu_active_background} !important;
                        }

                        .el-menu-item, .el-submenu__title{
                            height: 50px;
                            line-height: 50px;
                        }
                        .header{
                        background-color:${theme.top_background} !important;
                        }
                        .header-btn{
                        color: ${theme.top_text_color} !important;
                        }
                        .el-breadcrumb__inner{
                        color: ${theme.top_text_color} !important;
                        }
                                       `
                }
                document.getElementsByTagName("head").item(0).appendChild(style);
            },
            /**
             * 重置
             */
            reset() {
                this.config = {
                    menu: 'dark',
                    top: 'bright',
                    layout: 'vertical',
                    multi_label: true,
                    head: false
                }
            },
            /**
             * 路由跳转
             * @param key
             * @param keyPath
             */
            onRoutes(key, keyPath) {
                this.$router.push({path: key})
            },
            //首页
            home() {
                this.$router.push({
                    path: '/',
                })
            },
            /**
             * 获取菜单
             * @returns {Promise<*>}
             */
            async getMenu() {
                return await getMenu.menuList()
            },
            //开关左侧菜单
            switchMenu() {
                this.$bus.$emit('switchMenu', 'head');
            },
            //菜单搜索
            change(val) {
                this.$router.push(val)
                // this.search = '';
                this.showSearch = false
            },
            //过滤菜单合并国际化
            generateRoutes(routes, basePath = '/', prefixTitle = []) {
                let res = []

                for (const router of routes) {
                    const data = {
                        path: router.path,
                        title: [...prefixTitle]
                    }

                    if (router.title) {

                        data.title = [...data.title, router.title]
                        if (router.title && router.path && !router.children) {
                            res.push(data)
                        }

                    }

                    if (router.children) {
                        const tempRoutes = this.generateRoutes(router.children, data.path, data.title)
                        if (tempRoutes.length >= 1) {
                            res = [...res, ...tempRoutes]
                        }
                    }
                }
                return res
            },
            /**
             * 刷新当前页面
             */
            reload() {
                this.$nextTick(() => {
                    this.$router.replace({
                        path: '/redirect' + this.$route.fullPath
                    })
                })
            },
            /**
             * 修改密码
             */
            update_password() {
                this.dialog = true
            },
            onSubmit(formName) {
                const self = this;
                var form = JSON.parse(JSON.stringify(self.form))
                var url = "admin/update_password"

                console.log(form["test"])
                self.$Submit.submitClose(this, formName, url, form, "PUT")
            },
            /**
             * 退出
             */
            logout() {
                localStorage.clear()
                this.$router.push({
                    path: '/login'
                });

            },
            /**
             * 修改全局大小样式
             * @param size
             */
            setSize(size) {
                this.$ELEMENT.size = size
                sessionStorage.setItem("layout_size", size)
                this.size = size
                this.reload()

            },
            /**
             * 改变语言
             * @param language
             */
            lang(language) {
                sessionStorage.setItem("language", language)
                this.$i18n.locale = language
                this.getLanguage()
                this.reload()

            },
            /**
             * 获取语言
             */
            getLanguage() {
                var language = this.$i18n.locale
                switch (language) {
                    case 'zh':
                        this.language = "中文"
                        break;
                    case 'en':
                        this.language = "English"
                        break;
                    default:
                        this.language = "中文"
                        break;
                }
            },
            /**
             * 全屏
             * @constructor
             */
            FullScreen() {
                let element = document.documentElement;
                // 判断是否已经是全屏
                // 如果是全屏，退出
                if (this.fullscreen) {
                    if (document.exitFullscreen) {
                        document.exitFullscreen();
                    } else if (document.webkitCancelFullScreen) {
                        document.webkitCancelFullScreen();
                    } else if (document.mozCancelFullScreen) {
                        document.mozCancelFullScreen();
                    } else if (document.msExitFullscreen) {
                        document.msExitFullscreen();
                    }
                } else { // 否则，进入全屏
                    if (element.requestFullscreen) {
                        element.requestFullscreen();
                    } else if (element.webkitRequestFullScreen) {
                        element.webkitRequestFullScreen();
                    } else if (element.mozRequestFullScreen) {
                        element.mozRequestFullScreen();
                    } else if (element.msRequestFullscreen) {
                        // IE11
                        element.msRequestFullscreen();
                    }
                }
                // 改变当前全屏状态
                this.fullscreen = !this.fullscreen;
            }
        }
    }
</script>
