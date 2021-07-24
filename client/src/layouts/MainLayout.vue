<template lang="pug">
q-layout(view="lHh Lpr lFf")
  q-header.header()
    q-toolbar.bg-white.text-primary
      q-btn(flat round icon="las la-swatchbook" color="green-5")
  q-page-container
    router-view
</template>

<script>
import { defineComponent, ref } from "vue"

export default defineComponent({
  name: "MainLayout",

  components: {},

  setup() {
    return {}
  },
})
</script>

<style lang="scss">
.header {
  border-bottom: 1px solid $light-green-1;
}
</style>
