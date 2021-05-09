<template>
  <icon-font type="icon-tuichu" />
</template>
<script>
import {createFromIconfontCN} from '@ant-design/icons-vue';
import { defineComponent } from 'vue';
const IconFont = createFromIconfontCN({
  scriptUrl: '//at.alicdn.com/t/font_8d5l8fzk5b87iudi.js',
});
export default defineComponent({
  props: {
    type: {
      type: String
    }
  },
  components: {
    IconFont,
  },
  setup(prop) {

  }
});
</script>