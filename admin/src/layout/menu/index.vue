<template>
  <div id="menu">
    <a-menu
        :mode="$store.state.layout.layout== 'layout-head' ? 'horizontal' : 'inline'"
        :theme="$store.state.layout.theme=== 'theme-dark' || $store.state.layout.theme === 'theme-night' ? 'dark' : 'light'"
        v-model:openKeys="openKey"
        @openChange="openChange"
        v-model:selectedKeys="selectKey"
    >
      <subMenu v-for="(row,index) in list" :row="row" :index="index"/>
    </a-menu>
  </div>
</template>

<script>
import {useStore} from "vuex";
import subMenu from "./sub_menu.vue";
import {useRoute} from "vue-router";
import {ref, watch} from "vue";

export default ({
  components: {subMenu},
  setup() {
    const {state, commit} = useStore();
    const list = state.routes.menu;
    const route = useRoute();
    const openKey = ref(['sub1']);

    const selectKey = ref(['用户']);
    watch(
        () => openKey,
        val => {
          console.log('openKeys', val);
        },
    );
    const changeLayout = model => {
      console.log(model)
      if (model.layout === "layout-comp") {
        console.log(route)
        let topPath = route.matched[0].path;
        console.log(topPath)
        // menu.value = state.routes.routers.find(r => r.path === topPath).children;
        // console.log(menu)
      }
    };

    watch(state.layout, n => changeLayout(n));
    const openChange = function (openKeys) {
      commit("layout/updateOpenKey", {openKeys});
    };

    const foldSide = () => {
      const isComputedMobile = computed(() => state.layout.isMobile);
      if (isComputedMobile.value) {
        commit("layout/updateCollapsed", true);
      }
    };
    return {
      foldSide,
      openChange,
      openKey,
      list,
      selectKey
    }
  }
});
</script>


