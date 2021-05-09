<template>
  <template v-if="row.status">
    <a-sub-menu v-if="row.children" :key="row.name">
      <template #title>
        <component :is="$antIcons[row.icon]"/>
        <span>
          <span>{{ row.title }}</span>
        </span>
      </template>
      <subMenu :row="item" v-for="item in row.children" :key="row.title"/>
    </a-sub-menu>
    <a-menu-item v-else :key="row.title">
      <component :is="$antIcons[row.icon]"/>
      <span>{{ row.title }}</span>
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
    }
  },
  setup(props) {

    return {}
  },
}
</script>

<style scoped></style>