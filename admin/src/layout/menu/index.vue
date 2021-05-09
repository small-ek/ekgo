<template>
  <div id="menu">
    <a-menu :mode="$store.state.layout.layout== 'layout-head' ? 'horizontal' : 'inline'" :theme="$store.state.layout.theme=== 'theme-dark' || $store.state.layout.theme === 'theme-night' ? 'dark' : 'light'" v-model:openKeys="openKey" @openChange="openChange">
      <subMenu :item="item" v-for="item in list"/>
    </a-menu>
  </div>
</template>

<script>
import {useStore} from "vuex";
import subMenu from "./sub_menu.vue";
import {ref, watch} from "vue";

export default ({
  components: {subMenu},
  setup() {
    const {state, commit} = useStore();

    const list = state.routes.routesList;
    const openKey = ref(['sub1']);

    watch(
        () => openKey,
        val => {
          console.log('openKeys', val);
        },
    );

    const rootPath = ref("");
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
      rootPath,
      list
    }
  }
});
</script>


