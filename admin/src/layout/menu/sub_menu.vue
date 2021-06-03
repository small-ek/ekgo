<template>
  <template v-if="row.status">
    <a-sub-menu v-if="row.children" :key="index">
      <template #title>
        <component :is="$antIcons[row.icon]"/>
        <span>
          <span>{{ row.title }}</span>
        </span>
      </template>
      <subMenu :row="item" v-for="(item,index2) in row.children" :keys="index2"/>
    </a-sub-menu>
    <a-menu-item v-else-if="row.status" :index="index2+index">
      <router-link :to="row.path">
        <component :is="$antIcons[row.icon]"/>
        <span>{{ row.title }}</span>
      </router-link>
    </a-menu-item>
  </template>
</template>

<script>
export default {
  name: 'subMenu',
  components: {},
  props: {
    row: {
      type: Object,
      default: () => ({})
    },
    index: {
      type: Number,
      default: 0
    }
  },
  setup(props) {
    return {}
  },
}
</script>

<style scoped></style>