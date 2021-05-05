<template>
  <div id="menu">
    <a-menu :mode="$store.state.layout.layout== 'layout-head' ? 'horizontal' : 'inline'" :theme="$store.state.layout.theme=== 'theme-dark' || $store.state.layout.theme === 'theme-night' ? 'dark' : 'light'" v-model:openKeys="openKey" @openChange="openChange">
      <subMenu :list="list"></subMenu>
    </a-menu>
  </div>
</template>

<script>
import {useStore} from "vuex";
import subMenu from "./sub_menu.vue";
import {ref, toRaw, watch} from "vue";

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

    const openChange = function (openKeys) {
      commit("layout/updateOpenKey", {openKeys});
    };
    // const openKey = computed(() => state.layout.openKey);
    console.log(openKey)
    return {
      openChange: openChange,
      openKey: openKey,
      list: toRaw(list)
    }
  }
});
</script>


