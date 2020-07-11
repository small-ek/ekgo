<template>
    <div class="sudoku-main">
        <el-row>
            <el-col :span="24">
                <div class="welcome">{{$t('welcome')}}</div>
            </el-col>
        </el-row>
        <div class="sudoku">

            <el-row :gutter="20">
                <el-col class="sudoku-item" :xs="24" :sm="12" :md="8" :lg="6" :xl="4" :key="index" v-for="(item,index) in menu">
                    <div v-if="item.status=='true'" :style="'background: '+background[index]" @click="jump(item.path)">
                        {{item.title}}
                    </div>
                </el-col>
            </el-row>
        </div>
    </div>
</template>
<script>
    import menu from '../../common/menu.js';
    import {checkLogin} from "../../common/api";

    export default {
        data() {
            return {
                menu: [],
                background: ['#00adb5', '#b83b5e', '#1b6ca8', '#2d4059', '#7e6bc4', '#071a52', '#f47c7c', '#70a1d7', '#6a65d8',
                    '#00adb5', '#b83b5e', '#1b6ca8', '#2d4059', '#7e6bc4', '#071a52', '#f47c7c', '#70a1d7', '#6a65d8',
                    '#00adb5', '#b83b5e', '#1b6ca8', '#2d4059', '#7e6bc4', '#071a52', '#f47c7c', '#70a1d7', '#6a65d8',
                    '#339966', '#666699', '#3399CC', '#33CCCC', '#996699', '#C24747', '#e2674a', '#FFCC66', '#99CC99', '#669999',
                    '#CC6699',
                    '#339966', '#666699', '#3399CC', '#33CCCC', '#996699', '#C24747', '#e2674a', '#FFCC66', '#99CC99', '#669999',
                    '#CC6699',
                    '#339966', '#666699', '#3399CC', '#33CCCC', '#996699', '#C24747', '#e2674a', '#FFCC66', '#99CC99', '#669999',
                    '#CC6699',
                    '#339966', '#666699', '#3399CC', '#33CCCC', '#996699', '#C24747', '#e2674a', '#FFCC66', '#99CC99', '#669999',
                    '#CC6699',
                    '#339966', '#666699'
                ]
            };
        },
        created() {
            let self = this;
            checkLogin().then(function (result) {
                if (result.data.code == 403) {
                    localStorage.clear()
                    self.$router.push({
                        path: '/login'
                    });
                }
            })

            this.getMenu().then(function (result) {
                self.menu = self.generateRoutes(result)
                console.log(self.menu)
            })
        },

        methods: {
            async getMenu() {
                return await menu.menuList()
            },
            jump(url) {
                this.$router.push(url)
            },
            //过滤菜单合并国际化
            generateRoutes(routes, basePath = '/', prefixTitle = []) {
                let res = []

                for (const router of routes) {

                    const data = {
                        path: router.path,
                        status: router.status,
                        title: [...prefixTitle]
                    }
                    if (router.title && router.path && !router.children) {
                        data.title = router.title
                        res.push(data)
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
        }
    }
</script>
<style scoped>
    .el-main {
        padding: 0px !important;
    }

    .welcome {
        text-align: center;
        font-size: 25px;
        margin-top: 65px;
        color: #191a23;
        font-weight: bold;
        text-shadow: -2px -2px #ccc;
    }

    .sudoku-main {
        background-image: url(../../assets/images/logo-background.svg), linear-gradient(to top, #e6e9f0 0%, #eef1f5 100%);
        width: 100%;
        height: calc(100vh - 98px);
        background-size: 100% 100%;
    }

    .sudoku {
        height: 500px;
        display: block;
        top: 10px;
        padding: 60px 100px;
        position: relative;
    }

    .sudoku-item {
        cursor: pointer;
        text-align: center;
    }

    .sudoku-item div {
        margin: 15px auto;
        height: 220px;
        width: 220px;
        line-height: 220px;
        border-radius: 220px;
        background: #fff;
        text-align: center;
        opacity: 0.85;
        filter: alpha(opacity=85);
        font-size: 22px;
        color: #e9e9eb;
        transition: .9s;
        box-shadow: 0 5px 15px -5px rgba(0, 0, 0, .5);
    }

    .sudoku-item div:hover {
        -webkit-transform: translateY(-10px);
        transform: translateY(-10px);
        transition: all 0.5s ease-in-out;
        opacity: 0.30;
        filter: alpha(opacity=30);
        font-size: 20px;
        font-weight: bold;
    }
</style>
