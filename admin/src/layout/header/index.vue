<template>
  <!-- 框架顶部菜单区域 -->
  <div id="header" :class="[isMobile && 'mobile_header']">
    <template v-if="layout !== 'layout-head'">
      <!-- 左侧菜单功能项 -->
      <div class="prev-menu">
        <!-- 左侧缩进功能键 -->
        <div class="menu-item" @click="trigger()">

          <AlignLeftOutlined v-if="collapsed"/>
          <!-- 左侧缩进功能键盘 -->
          <AlignRightOutlined v-else/>
        </div>
        <div class="menu-item" @click="refresh">
          <!-- 刷新当前页面路由 -->
          <ReloadOutlined v-if="routerActive"/>
          <LoadingOutlined v-else/>
        </div>
      </div>
    </template>

    <template v-else>
      <div class="head-logo">
        <Logo></Logo>
      </div>
      <div class="head-menu">
        <Menu></Menu>
      </div>
    </template>

    <!-- 实现综合布局方式 -->
    <div v-if="layout == 'layout-comp'" class="comp-menu">
      <template :key="index" v-for="(route, index) in routes">
        <router-link
            :to="toPath(route)"
            class="menu-item"
            :class="[active === route.path ? 'is-active' : '']"
        >
          <span>{{ route.meta.title }}</span>
        </router-link>
      </template>
    </div>

    <!-- 右侧菜单功能项 基本公用 -->
    <div class="next-menu">
      <div class="menu-item" v-if="!fullscreen" @click="full(1)">
        <ExpandOutlined/>
      </div>
      <div class="menu-item" v-else @click="full(2)">
        <CompressOutlined/>
      </div>
      <!--语言包-->
      <a-dropdown class="locale-item">
        <GlobalOutlined/>
        <template #overlay>
          <a-menu @click="toggleLang" :selectedKeys="selectedKeys">
            <a-menu-item key="zh-CN"> 简体中文</a-menu-item>
            <a-menu-item key="en-US"> English</a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
      <!--头像-->
      <a-dropdown class="avatar-item">
        <a-avatar
            src="https://portrait.gitee.com/uploads/avatars/user/1611/4835367_Jmysy_1578975358.png"
        ></a-avatar>
        <template #overlay>
          <a-menu class="avatar-dropdown">
            <a-menu-item key="0">
              <router-link to="/account/center">个人中心</router-link>
            </a-menu-item>
            <a-menu-item key="1">
              <a
                  target="_blank"
                  rel="noopener noreferrer"
                  href="http://www.taobao.com/"
              >系统设置</a
              >
            </a-menu-item>
            <a-menu-divider/>
            <a-menu-item key="3">
              <a-menu-item @click="logOut"> 注销登录</a-menu-item>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
      <div v-if="!isMobile" class="menu-item" @click="setting()">
        <!-- 主题设置隐显键 -->
        <MoreOutlined/>
      </div>
    </div>
  </div>
</template>
<script>
import {computed, ref, unref, watch} from "vue";
import {useStore} from 'vuex'
import Menu from "../menu/index.vue";
import Logo from "../logo/index.vue";
import {useRoute} from "vue-router";
import {AlignLeftOutlined, AlignRightOutlined, BellOutlined, CompressOutlined, ExpandOutlined, GlobalOutlined, LoadingOutlined, MoreOutlined, ReloadOutlined,} from "@ant-design/icons-vue";

export default {
  components: {
    AlignLeftOutlined,
    AlignRightOutlined,
    MoreOutlined,
    ExpandOutlined,
    CompressOutlined,
    ReloadOutlined,
    GlobalOutlined,
    Menu,
    Logo,
    BellOutlined,
    LoadingOutlined,
  },

  methods: {
    full: function (num) {
      num = num || 1;
      num = num * 1;
      var docElm = document.documentElement;
      switch (num) {
        case 1:
          if (docElm.requestFullscreen) {
            docElm.requestFullscreen();
          } else if (docElm.mozRequestFullScreen) {
            docElm.mozRequestFullScreen();
          } else if (docElm.webkitRequestFullScreen) {
            docElm.webkitRequestFullScreen();
          } else if (docElm.msRequestFullscreen) {
            docElm.msRequestFullscreen();
          }
          break;
        case 2:
          if (document.exitFullscreen) {
            document.exitFullscreen();
          } else if (document.mozCancelFullScreen) {
            document.mozCancelFullScreen();
          } else if (document.webkitCancelFullScreen) {
            document.webkitCancelFullScreen();
          } else if (document.msExitFullscreen) {
            document.msExitFullscreen();
          }
          break;
      }
      this.updateFullscreen();
    },
  },
  setup() {
    const {state, commit, dispatch} = useStore()
    const layout = computed(() => state.layout.layout);
    const collapsed = computed(() => state.layout.collapsed)
    const fullscreen = computed(() => state.layout.fullscreen);
    const menuModel = "";
    const theme = "";
    const $route = useRoute();
    const active = ref($route.matched[0].path);
    const isMobile = false;
    const routerActive = computed(() => state.layout.routerActive);

    watch(
        computed(() => $route.fullPath),
        () => {
          active.value = $route.matched[0].path;
        }
    );

    const toPath = (route) => {
      let {redirect, children, path} = route;
      if (redirect) {
        return redirect;
      }
      while (children && children[0]) {
        // path = _path.resolve(path, children[0].path);
        children = children[0].children;
      }
      return path;
    };

    const routes = "";

    const refresh = async () => {
      commit("layout/updateRouterActive");
      setTimeout(() => {
        commit("layout/updateRouterActive");
      }, 500);
    };

    const logOut = async (e) => {
      await dispatch("user/logout");
      window.location.reload();
    };

    const store = useStore();
    const defaultLang = "";
    const selectedKeys = ref([unref(defaultLang)]);
    const toggleLang = async ({key}) => {
      selectedKeys.value = [key];
      // await loadLocaleMessages(i18n, key);
      await store.dispatch("app/setLanguage", key);
    };

    return {
      isMobile,
      layout,
      collapsed,
      fullscreen,
      setting: () => commit("layout/updateSettingOpen"),
      updateFullscreen: () => commit("layout/updateFullscreen"),
      trigger: () => commit("layout/updateCollapsed"),
      menuModel,
      routerActive,
      theme,
      refresh,
      routes,
      active,
      toPath,
      logOut,
      toggleLang,
      selectedKeys,
    };
  },
};
</script>
<style lang="less" scoped>
.mobile_header {
  padding-right: 0px !important;
}
</style>
