<template>
    <div class="left" v-if="$Config.theme.layout=='vertical'" style="height: 100vh">
        <div class="logo-container-vertical hidden-sm-and-down" v-if="$Config.logo" v-show="isCollapse==false">
            <img src="../assets/images/ek.jpg" class="logo_img" alt="">
            <h2 class="logo_title">
                {{$Config.logoTitle}}
            </h2>
        </div>
        <div class="el-loading-mask" style="background-color: rgba(0, 0, 0, 0.8);z-index: 5" v-show="isBackground" @click="hideMenu">
        </div>
        <el-menu v-bind:class="{smallmenu:isMobile,smallHide:isHidden}" unique-opened :background-color="menu_background" :text-color="menu_text_color" :active-text-color="menu_active_text_color" class="el-menu-vertical" :default-active="onRoutes" router :collapse="isCollapse" collapse-transition="true">
            <!--菜单内容-->
            <Menu :list="list"></Menu>
        </el-menu>

    </div>
</template>

<script>
    import Menu from './Menu';
    import getMenu from '../common/menu.js';

    export default {
        data() {
            return {
                screenWidth: document.documentElement.clientWidth,
                isMobile: false,//手机
                isBackground: false,//背景色
                isHidden: false,//是否隐藏手机菜单
                list: [],//数据
                isCollapse: false,//是否缩放
                menu_background: '',
                menu_text_color: '',
                menu_active_text_color: '',
                menu_active_background: '',
            };
        },
        components: {Menu},
        created() {
            let self = this;

            getMenu.setMenuTheme(self)

            this.getMenu().then(function (result) {
                self.list = result
            })

            this.loadMenuStyle(this.screenWidth)

            this.$bus.$on('switchMenu', data => {
                //如果是手机
                if (this.isMobile && data == "head") {
                    this.isBackground = this.isBackground == true ? false : true
                    this.isHidden = this.isHidden == true ? false : true
                    this.isCollapse = false
                } else if (data == "head") {
                    this.isBackground = false
                    this.isHidden = false
                    this.isCollapse = this.isCollapse == true ? false : true;
                }
                if (this.isMobile && data == "left") {
                    this.isBackground = false
                    this.isHidden = true
                }
            });

        },
        computed: {
            onRoutes() {
                return this.$route.path;
            },
        },
        methods: {
            /**
             *获取菜单
             */
            async getMenu() {
                return await getMenu.menuList()
            },
            /**
             * 隐藏菜单
             */
            hideMenu() {
                this.isBackground = false
                this.isHidden = true
            },
            /**
             * 加载手机样式
             * @param width
             */
            loadMenuStyle(width) {
                //手机屏幕
                if (width < 500) {
                    this.isMobile = true
                    this.isBackground = true
                    this.isHidden = false
                    this.isCollapse = false
                } else if (width > 500 && width < 992) {//小尺寸
                    this.isMobile = false
                    this.isBackground = false
                    this.isCollapse = true
                    this.isHidden = false
                } else {
                    this.isMobile = false
                    this.isBackground = false
                    this.isCollapse = false
                    this.isHidden = false
                }
            }
        },
        mounted() {
            const self = this
            window.onresize = () => {
                return (() => {
                    self.screenWidth = document.documentElement.clientWidth
                })()
            }
        },
        watch: {
            screenWidth(val) {
                this.loadMenuStyle(val)
            }
        }
    }
</script>

<style>
    .logo_title {
        color: #fff !important;
        font-weight: 600;
        font-size: 24px;
        vertical-align: middle;
        display: inline-block;
        margin-left: 25px;
    }

    .logo_img {
        width: 35px;
        height: 35px;
        vertical-align: middle;
        background: #fff;
        border-radius: 50%;
        padding: 3px;
        margin-left: 20px;
    }

    .logo-container-vertical {
        position: relative;
        height: 55px;
        overflow: hidden;
        line-height: 55px;
        padding: 0px 2%;
    }

    .logo-container-vertical a {
        color: white;
    }

    .el-menu .fa {
        vertical-align: middle;
        margin-right: 5px;
        width: 20px;
        text-align: center;
    }

    .el-menu-vertical:not(.el-menu--collapse) {
        width: 240px;
        min-height: 400px;
    }

    .el-menu--collapse span,
    .el-menu--collapse i.el-submenu__icon-arrow {
        height: 0;
        width: 0;
        overflow: hidden;
        visibility: hidden;
        display: inline-block;
    }

    .smallmenu {
        background: #1F262D;
        position: absolute;
        z-index: 100;
        height: 100vh;
    }

    .smallHide {
        display: none;
    }
</style>
