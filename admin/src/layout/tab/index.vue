<template>
  <div id="tab" :class="[tabType]">
    <a-tabs
        hide-add
        v-model:activeKey="activeKey"
        type="editable-card"
        @edit="onEdit"
        @change="callback"
        class="tab">
      <a-tab-pane
          v-for="pane in panes"
          :key="pane.path"
          :tab="pane.title"
          :closable="pane.closable">
      </a-tab-pane>
    </a-tabs>
    <a-dropdown class="tab-tool" :placement="placement">
      <a-button>
        <template v-slot:icon>
          <DownOutlined/>
        </template>
      </a-button>
      <template v-slot:overlay>
        <a-menu>
          <a-menu-item>
            <a @click="closeAll()">关 闭 所 有</a>
          </a-menu-item>
          <a-menu-item>
            <a @click="closeOther()">关 闭 其 他</a>
          </a-menu-item>
          <a-menu-item>
            <a @click="closeCurrent()">关 闭 当 前</a>
          </a-menu-item>
        </a-menu>
      </template>
    </a-dropdown>
  </div>
</template>
<script>
import {computed, ref} from "vue";
import {useStore} from "vuex";
import {DownOutlined} from "@ant-design/icons-vue";
import {useRoute} from "vue-router";

export default {
  components: {
    DownOutlined
  },
  methods: {
    callback(key) {
      console.log("callback", key);
      this.selectTab(key);
    },
    onEdit(targetKey, action) {
      this[action](targetKey);
    },
    closeAll() {
      this.closeAllTab();
    },
    closeOther() {
      this.closeOtherTab();
    },
    closeCurrent() {
      this.closeCurrentTab();
    },
    remove(targetKey) {
      this.removeTab(targetKey);
    }
  },
  setup() {
    const {state, commit, dispatch} = useStore();
    const defaultPanes = computed(() => state.layout.panes);
    const panes = ref(defaultPanes);
    const route = useRoute();



    return {

    };
  }
};
</script>
