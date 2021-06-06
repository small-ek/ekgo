<template>
  <div id="menu">
    <a-menu
        :mode="$store.state.layout.layout== 'layout-head' ? 'horizontal' : 'inline'"
        :theme="$store.state.layout.theme=== 'theme-dark' || $store.state.layout.theme === 'theme-night' ? 'dark' : 'light'"
        v-model:openKeys="openKey"
        @openChange="openChange"
        v-model:selectedKeys="selectKey"
    >
      <subMenu v-for="(row,index) in list" :row="row" :index="index" @click="onFoldSide"/>
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
    var list = ref(state.routes.menu);
    const route = useRoute();
    const openKey = ref(['sub1']);

    const selectKey = ref(['用户']);
    watch(
        () => openKey,
        val => {
          console.log('openKeys', val);
        },
    );
    watch(state.routes, n => changeLayout(n));
    watch(state.layout, n => changeLayout(n));

    const changeLayout = model => {
      if (state.layout.layout === "layout-comp") {
        list.value = state.routes.subMenu;
      } else {
        list.value = state.routes.menu;
      }
    };



    const openChange = function (openKeys) {
      commit("layout/updateOpenKey", {openKeys});

    };

    const onFoldSide = () => {
      if (state.layout.isMobile) {
        commit("layout/updateCollapsed", false);
      }
    }
    return {
      onFoldSide,
      openChange,
      openKey,
      list,
      selectKey
    }
  }
});
</script>


