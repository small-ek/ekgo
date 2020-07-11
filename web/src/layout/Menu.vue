<template>
    <div v-bind:class="{isHorizontal:isHorizontal}">
        <template v-for="(item,key) in list">
            <template v-if="item.children && item.status=='true'">
                <el-submenu @click="hideMenu" :index="key+'-'" :key="item.path">
                    <template slot="title">
                        <i :class="item.icon||'el-icon-menu'"></i>
                        <span slot="title">{{ item.title }}</span>
                    </template>
                    <menus :list="item.children"></menus>
                </el-submenu>
            </template>
            <template v-else-if="item.status=='true'">
                <el-menu-item :index="item.path" :key="item.path" @click="hideMenu">
                    <i :class="item.icon"></i>
                    <span slot="title">{{ item.title }}</span>
                </el-menu-item>
            </template>
        </template>
    </div>
</template>
<script>
    export default {
        name: 'menus',
        props: {
            list: Object
        },
        data() {
            return {
                isHorizontal: false
            };
        },
        created() {
            this.$Config.theme.layout == "horizontal" ? this.isHorizontal = true : ''
        },
        computed: {},
        methods: {
            hideMenu() {
                this.$bus.$emit('switchMenu', 'left');
            }
        }
    }
</script>
<style>
    .isHorizontal > li {
        float: left;
        height: 50px;
        line-height: 50px;
        margin: 0;
        border-bottom: 2px solid transparent;
        color: #909399;
    }

    .isHorizontal .is-active {
        border-bottom: 2px solid #5677FC;
    }

    .isHorizontal .el-icon-arrow-down {
        position: static;
        vertical-align: middle;
        margin-left: 8px;
        margin-top: -3px;
        right: 1px;
    }
</style>
