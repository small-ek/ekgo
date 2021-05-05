<template>
  <div id="logo">
    <span v-if="$store.state.layout.collapsed"><img class="image" src="@/assets/image/logo.png" /></span>
    <div v-else><span class="title">Ekgo Admin</span></div>
  </div>
</template>

<script>

export default ({
  components: {
  },
  data() {
    return {

    }
  },
});
</script>


