<template>
  <template v-if="item.status">
    <a-sub-menu v-if="item.children" :key="item.name">
      <template #title>
        <span>
<!--          <icon-font style="color: aliceblue" :type="menuInfo.meta.icon" />-->
          <span>{{ item.title }}</span>
        </span>
      </template>
      <subMenu :item="row" v-for="row in item.children" :key="row.title"/>
    </a-sub-menu>
    <a-menu-item v-else :key="item.title">
      <!--      <icon-font style="color: aliceblue" :type="menuInfo.meta.icon" />-->
      <span>{{ item.title }}</span>
    </a-menu-item>
  </template>
</template>

<script>
import {defineComponent} from 'vue'

export default defineComponent({
  name: 'subMenu',
  components: {},
  props: {
    item: {
      type: Object,
      default: () => ({})
    }
  },
  setup(props) {

  }
})
</script>

<style scoped></style>