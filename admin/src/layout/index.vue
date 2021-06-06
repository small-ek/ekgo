<template>
  <a-layout id="layout" :class="[$store.state.layout.theme,$store.state.layout.layout]">
    <div
        v-if="$store.state.layout.isMobile&&!$store.state.layout.collapsed"
        class="layout_mobile_mask"
        @click="onCloseSideBar"
    />
    <!--左侧-->
    <a-layout-sider
        v-if="$store.state.layout.layout != 'layout-head'"
        :width="$store.state.layout.sideWitch" v-model:collapsed="$store.state.layout.collapsed" :trigger="null" collapsible
        :collapsedWidth="$store.state.layout.collapsedWidth"
        :class="[
        $store.state.layout.fixedSide ? 'fixed-side' : '',
        $store.state.layout.isMobile && 'layout_mobile',
        $store.state.layout.collapsed && 'layout_collapse'
      ]"
    >
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
          :class="[$store.state.layout.fixedHeader ? 'fixedHeader' : '', $store.state.layout.tab ? 'muiltTab' : '']"
      >
        <!-- 选项卡页面 -->
        <router-view></router-view>
        <!-- 设置页面 -->
        <Style></Style>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script>
import Menu from "./menu/index.vue"
import Logo from "./logo/index.vue"
import Header from "./header/index.vue"
import Tab from "./tab/index.vue"
import Style from "./style/index.vue"
import {useStore} from "vuex";

export default ({
  components: {
    Menu,
    Logo,
    Header,
    Style,
    Tab
  },
  setup() {
    const {commit, state} = useStore();
    const onCloseSideBar = () => {
      // const isComputedMobile = computed(() => state.layout.isMobile);
      commit("layout/updateCollapsed", false);

    }
    const handleLayouts = () => {
      const domWidth = document.body.getBoundingClientRect().width;
      const isLayoutMobile = domWidth !== 0 && domWidth - 1 < 992;
      commit("layout/updateIsMobile", isLayoutMobile);
      if (isLayoutMobile) {
        commit("layout/updateLayout","layout-side")

        setTimeout(() => {

        }, 1000);
      }else{}
    };
    handleLayouts();
    window.addEventListener("resize", handleLayouts);

    return {
      onCloseSideBar
    }
  }
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