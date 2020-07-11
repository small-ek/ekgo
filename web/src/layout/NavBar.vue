<template>
    <nav class="nav-bar" v-if="this.$Config.theme.multi_label==true">
        <ScrollPane class="nav-bar-scroll">
            <router-link :to="{path:item.path}" @click.prevent="jump(item.path)" class="nav-bar-tag" v-for="(item,index) in tagsList" :key="index" :class="{'active': isActive(item.path)}">
                <i class="point"></i>
                {{ item.title }}
                <div v-show="showTags" class="close-box">
                    <span class='el-icon-close' @click.prevent.stop='closeTags(index)'></span>
                </div>
            </router-link>
        </ScrollPane>
        <div class="tags-close-box">
            <el-dropdown @command="handleTags">
                <el-button size="mini" icon="el-icon-arrow-down" style="height: 32px;font-weight: bold;border: 0px;padding: 0px 10px;color:black;font-size: 13px;">
                </el-button>
                <el-dropdown-menu size="small" slot="dropdown">
                    <el-dropdown-item command="other">{{$t('close_other')}}</el-dropdown-item>
                    <el-dropdown-item command="all">{{$t('close_all')}}</el-dropdown-item>
                </el-dropdown-menu>
            </el-dropdown>
        </div>
    </nav>
</template>
<script>
    import ScrollPane from '../components/ScrollPane/Index.vue'

    export default {
        data() {
            return {
                tagsList: []
            }
        },
        components: {
            ScrollPane
        },
        methods: {
            jump(url) {
                this.$router.push(url);
            },
            isActive(path) {
                return path === this.$route.fullPath;
            },
            /**
             * 关闭单个标签
             * @param index
             */
            closeTags(index) {
                const delItem = this.tagsList.splice(index, 1)[0];
                const item = this.tagsList[index] ? this.tagsList[index] : this.tagsList[index - 1];

                if (item) {
                    delItem.path === this.$route.fullPath && this.$router.push(item.path);
                } else {
                    this.$router.push('/');
                }
            },
            /**
             * 关闭全部标签
             */
            closeAll() {
                this.tagsList = [];
                this.$router.push('/');
            },
            /**
             * 关闭其他标签
             */
            closeOther() {
                const curItem = this.tagsList.filter(item => {
                    return item.path === this.$route.fullPath;
                })
                this.tagsList = curItem;
            },
            /**
             * 设置标签
             * @param route
             */
            setTags(route) {
                const isExist = this.tagsList.some(item => {
                    return item.path === route.fullPath;
                })
                if (!isExist) {
                    if (this.tagsList.length >= 8) {
                        this.tagsList.shift();
                    }
                    if (route.matched[0]['path'] != "/redirect") {

                        this.tagsList.push({
                            title: route.meta.title,
                            path: route.fullPath,
                            name: route.matched[1].components.default.name
                        })

                    }
                }
            },
            handleTags(command) {
                command === 'other' ? this.closeOther() : this.closeAll();
            }
        },
        computed: {
            showTags() {
                return this.tagsList.length > 0;
            }
        },
        watch: {
            $route(newValue, oldValue) {
                this.setTags(newValue);
            }
        },
        created() {
            this.setTags(this.$route);
        }
    }
</script>

<style lang="less">
    .nav-bar-scroll {
        width: calc(100% - 42px);
        position: relative;
    }

    .tags-close-box {
        position: absolute;
        right: 5px;
        top: 4px;
        box-sizing: border-box;
        padding-top: 1px;
        text-align: center;
        // width: 110px;
        height: 30px;
        background: #fff;
        box-shadow: -3px 0 15px 3px rgba(0, 0, 0, .1);
        z-index: 9;

    }

    .nav-bar {
        box-shadow: 0 1px 4px rgba(0, 21, 41, .08);
        position: relative;
        margin-top: 0px;
        height: 43px;
        width: 100%;
        z-index: 2;
        background: #fff;
        white-space: nowrap;
        transition: left 0.3s ease-in-out;

        .nav-bar-tag {
            height: 32px;
            color: black;
            margin: 6px 2px;
            padding-left: 6px;
            padding-right: 6px;
            font-size: 12px;
            line-height: 32px;
            border: 1px solid #ebeef5;
            border-radius: 3px;
            display: inline-block;
            padding: 0px 15px;

            .close-box {
                display: inline-block;
                height: 100%;
                margin-left: 6px;
            }

            &:hover {
                opacity: 0.85;
            }

            .point {
                display: none;
                width: 8px;
                height: 8px;
                border-radius: 50%;
                margin: auto 0;
                position: relative;
                background: #fff;
            }

            .el-icon-close {
                font-weight: bolder;
                margin-left: 6px;

                &:hover {
                    background: #ff0000;
                    color: white;
                    border-radius: 10px;
                }
            }
        }

        .active {
            color: white;
            background: #5677FC;
        }

        .scroll-container {
            height: 38px;
        }
    }
</style>
