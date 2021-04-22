<template>
  <a-layout id="layout" class="theme-dark layout-side">
    <!--左侧-->
    <a-layout-sider width="240" v-model:collapsed="$store.state.MenuStatus" :trigger="null" collapsible class="fixed-side">
      <div class="pear-layout-left-sider">
        <!--Logo-->
        <Logo></Logo>
        <!--菜单-->
        <Menu></Menu>
      </div>
    </a-layout-sider>

    <!--右侧-->
    <a-layout>
      <!--头部-->
      <a-layout-header style="background: #fff; padding: 0">
      <Header></Header>
      </a-layout-header>
      <!-- main区域 -->
      <a-layout-content
          :style="{ margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' }"
      >
        <router-view></router-view>
      </a-layout-content>
    </a-layout>

  </a-layout>
</template>
<script>
import Menu from "./menu/index.vue"
import Logo from "./logo/index.vue"
import Header from "./header/index.vue"
export default ({
  components: {
    Menu,
    Logo,
    Header
  },
  data() {
    return {

    }
  },
});
</script>
<style lang="less" scoped>
//移动端侧边栏遮罩层
.layout_mobile_mask {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  z-index: 998;
  width: 100%;
  height: 100vh;
  overflow: hidden;
  background: #000;
  opacity: 0.5;
}
//移动端导航栏布局
.layout_mobile {
  position: fixed !important;
  z-index: 999;
  .pear-layout-left-sider {
    right: 0 !important;
    -ms-overflow-style: none;
    overflow: -moz-scrollbars-none;
  }
  .pear-layout-left-sider::-webkit-scrollbar {
    width: 0 !important;
  }

  &.layout_collapse {
    width: 0 !important;
    min-width: 0 !important;
    max-width: 0 !important;
    * {
      display: none !important;
      width: 0 !important;
      min-width: 0 !important;
      max-width: 0 !important;
    }
    .ant-menu-item,
    .ant-menu-submenu {
      display: none !important;
      width: 0 !important;
      min-width: 0 !important;
      max-width: 0 !important;
    }
  }
}
</style>