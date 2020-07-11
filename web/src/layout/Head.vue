<template>
    <el-header directio="horizontal" class="header">
        <el-row :gutter="0">
            <tempalte v-if="$Config.theme.layout=='vertical'">
                <el-col :xs="8" :sm="7" :md="8" :lg="12" :xl="15" style="padding-left: 10px">
                    <el-col :xs="9" :sm="4" :md="3" :lg="2" :xl="1">
                    <span class="header-btn" @click="switchMenu()">
                        <i class="fa fa-navicon fa-lg"></i>
                    </span>
                    </el-col>
                    <el-col :xs="20" :sm="20" :md="21" :lg="21" :xl="22" class="hidden-xs-only">
                        <HeadLeft></HeadLeft>
                    </el-col>
                </el-col>
                <el-col :xs="16" :sm="17" :md="16" :lg="12" :xl="9" style="text-align: right;;padding-right: 10px">
                    <HeadRight></HeadRight>
                </el-col>
            </tempalte>
            <tempalte v-else>
                <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
                    <el-menu class="el-menu-demo" :default-active="onRoute" :background-color="menu_background" :text-color="menu_text_color" :active-text-color="menu_active_text_color" mode="horizontal" @select="onRoutes">
                        <Menu :list="list"></Menu>
                    </el-menu>
                </el-col>
            </tempalte>

        </el-row>
    </el-header>
</template>

<script>
    import Menu from './Menu';
    import getMenu from '../common/menu.js';
    import HeadRight from "./HeadRight";
    import HeadLeft from "./HeadLeft";

    export default {
        data() {
            return {
                menuOptions: [],
                list: [],
                menu_background: '',
                menu_text_color: '',
                menu_active_text_color: '',
                menu_active_background: '',
            };
        },
        components: {Menu, HeadLeft, HeadRight},
        created() {
            let self = this;
            getMenu.setMenuTheme(self)
            //设置下拉框菜单
            if (this.$Config.theme.layout == "horizontal") {
                this.getMenu().then(function (result) {
                    self.list = result
                })
            }
        },
        computed: {
            onRoute() {
                return this.$route.path;
            },
        },
        methods: {
            /**
             * 菜单路由跳转
             * @param key
             * @param keyPath
             */
            onRoutes(key, keyPath) {
                this.$router.push({path: key})
            },
            async getMenu() {
                return await getMenu.menuList()
            },
            //开关左侧菜单
            switchMenu() {
                this.$bus.$emit('switchMenu', 'head');
            },
        }
    }
</script>

<style lang="scss">
    .el-header {
        height: auto !important;
        padding: 0px;
        background: white;
    }

    .header-btn {
        width: auto;
        padding: 0 7px;
        overflow: hidden;
        height: 51px;
        display: inline-block;
        text-align: center;
        line-height: 51px;
        cursor: pointer;
        color: #5a5e66;
        font-size: 13px;
        transition: all .2s ease-in-out;
    }

    .header-btn:hover {
        opacity: 0.8;
        filter: alpha(opacity=80);
    }

    .header-search-select {
        font-size: 18px;
        transition: width 0.2s;
        overflow: hidden;
        background: transparent;
        border-radius: 0;
        display: inline-block;
        vertical-align: middle;
        text-align: center;
        line-height: 45px;

        /deep/ .el-input__inner {
            border-radius: 0;
            border: 0;
            padding-left: 0;
            padding-right: 0;
            box-shadow: none !important;
            border-bottom: 1px solid #d9d9d9;
            vertical-align: middle;
        }
    }
</style>
