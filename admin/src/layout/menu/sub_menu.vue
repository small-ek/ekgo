<template>
  <template v-for="(item,index) in list" :key="item.id">
    <template v-if="item.children && item.status=='true'">
      <a-sub-menu :key="index">
        <template #title>
        <span>
          <MailOutlined/>
          <span>{{ item.title }}</span>
        </span>
        </template>
        <menus :list="item.children"></menus>
      </a-sub-menu>
    </template>
    <template v-else-if="item.status=='true'">
      <a-menu-item :key="index">{{ item.title }}</a-menu-item>
    </template>
  </template>
</template>

<script>
import {MailOutlined, MenuFoldOutlined, MenuUnfoldOutlined, PieChartOutlined,} from '@ant-design/icons-vue';

export default {
  components: {
    MailOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    PieChartOutlined,
  },
  emits: ["click"],
  name: "menus",
  props: {
    list: {
      type: Array,
      required: true
    },
  },
  setup(props) {
    console.log(props.list)
    return {};
  }
};
</script>
<style>
.indent {
  width: 15px;
  display: inline-block;
}
</style>
