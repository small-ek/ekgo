<template>
    <el-container>
        <!-- 左侧导航 -->
        <Left></Left>
        <el-container direction="vertical" style="height: 100vh;">
            <!--是否固定头部-->
            <tempalte v-if="this.$Config.theme.head==false">
                <!-- 头部 -->
                <Head></Head>
                <!-- 导航栏 -->
                <!--<NavBar class="hidden-xs-only"></NavBar>-->
                <NavBar></NavBar>
            </tempalte>

            <!-- 主体 -->
            <el-main id="main" v-bind:class="{mains:isHorizontal}">
                <!--是否固定头部-->
                <tempalte v-if="this.$Config.theme.head==true">
                    <!-- 头部 -->
                    <Head></Head>
                    <!-- 导航栏 -->
                    <NavBar class="hidden-xs-only"></NavBar>
                </tempalte>

                <template>
                    <!--顶部菜单-->
                    <el-row :gutter="0" v-if="$Config.theme.layout=='horizontal'" class="header">
                        <el-col :xs="9" :sm="7" :md="8" :lg="12" :xl="15" style="padding-left: 10px">
                            <HeadLeft></HeadLeft>
                        </el-col>
                        <el-col :xs="13" :sm="17" :md="15" :lg="12" :xl="9" style="text-align: right;;padding-right: 10px">
                            <HeadRight></HeadRight>
                        </el-col>
                    </el-row>
                    <transition name="slide-fade">
                        <keep-alive :include="keep_alive">
                            <router-view></router-view>
                        </keep-alive>
                    </transition>
                </template>
            </el-main>
        </el-container>
        <el-backtop target="#main"></el-backtop>
    </el-container>
</template>

<script>
    import Left from './Left.vue'
    import Head from './Head.vue'
    import NavBar from './NavBar.vue'
    import HeadLeft from "./HeadLeft";
    import HeadRight from "./HeadRight";

    export default {
        components: {
            Left,
            Head,
            NavBar,
            HeadLeft,
            HeadRight
        },
        data() {
            return {
                isHorizontal: false,
                keep_alive: [],
            }
        },
        created() {
            this.keep_alive = this.$Config.keep_alive
            this.$Config.theme.layout == "horizontal" ? this.isHorizontal = true : ''
        },
    };
</script>
<style scoped="scoped">
    .el-main {
        padding: 0%;
        background: #f0f2f5;

    }

    .main {
        padding: 1%;
    }

    .slide-fade {
        position: absolute;
        left: 0;
        right: 0;
    }

    .slide-fade-enter-active {
        transition: all 1.2s ease;
    }

    .slide-fade-leave-active {

        transition: all .1s cubic-bezier(2.0, 0.5, 0.8, 1.0);
    }

    .slide-fade-enter,
    .slide-fade-leave-to {
        left: 0;
        right: 0;
        transform: translateX(50px);
        opacity: 0;
    }
</style>
