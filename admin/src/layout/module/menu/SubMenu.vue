<template>
  <template v-if="!item.hidden">
    <a-menu-item
      v-if="
        item.children &&
          item.children.length == 1 &&
          item.meta.alwaysShow != true
      "
      :key="resolvePath(item.path, true)"
      @click="handleFoldSideBar"
    >
      <router-link :to="item.path + '/' + item.children[0].path">
        <MenuIcon />
        <!-- <span>{{ item.meta.title }}</span> -->
        <span>{{ t(item.meta.i18nTitle) }}</span>
      </router-link>
    </a-menu-item>

    <!-- if item.children is not null 渲染 a-sub-menu -->
    <a-sub-menu
      @click="handleFoldSideBar"
      :key="item.path"
      v-else-if="item.children && item.children.length > 0"
    >
      <template v-slot:title>
        <span>
          <MenuIcon v-if="level === 0" />
          <span v-else><div class="indent"></div></span>
          <!-- <span>{{ item.meta.title }}</span> -->
          <span>{{ t(item.meta.i18nTitle) }}</span>
        </span>
      </template>
      <!-- 递归 item.children -->
      <sub-menu
        v-for="child in item.children"
        :key="resolvePath(child.path)"
        :item="child"
        :level="level + 1"
        :base-path="resolvePath(child.path)"
      />
    </a-sub-menu>
    <!-- if item.chilren is null 渲染 a-menu-item -->
    <a-menu-item
      @click="foldSide"
      v-bind="$attrs"
      :key="resolvePath(item.path, true)"
      v-else
    >
      <router-link :to="resolvePath(item.path, true)">
        <MenuIcon v-if="level === 0" />
        <span v-else><div class="indent"></div></span>
        <span>{{ t(item.meta.i18nTitle) }}</span>
      </router-link>
    </a-menu-item>
  </template>
</template>

<script>
import { computed } from "vue";
import path from "path";
import { useStore } from "vuex";
import * as Icons from "@ant-design/icons-vue";
import {useI18n} from "vue-i18n";
export default {
  emits: ["click"],
  name: "SubMenu",
  props: {
    item: {
      type: Object,
      required: true
    },
    basePath: {
      type: String,
      default: ""
    },
    level: {
      type: Number,
      required: true
    }
  },
  setup(props) {
    const { commit, getters } = useStore();
    const resolvePath = (routePath, single) => {
      if (/^(https?:|mailto:|tel:)/.test(routePath)) {
        return routePath;
      }
      if (single) {
        return props.basePath;
      }
      // 当处于 comp 模式下拼接相关路由
      return path.resolve(props.basePath, routePath);
    };

    const foldSide = () => {
      const isComputedMobile = computed(() => getters.isMobile);
      if (isComputedMobile.value) {
        commit("layout/UPDATE_COLLAPSED", true);
      }
    };
    const MenuIcon = Icons[(props.item.meta || {}).icon] || {};

    const { t } = useI18n()

    return {
      foldSide,
      resolvePath,
      MenuIcon,
      t
    };
  }
};
</script>
<style>
.indent {
  width: 15px;
  display: inline-block;
}
</style>
