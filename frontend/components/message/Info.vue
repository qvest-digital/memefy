<template>
  <div class="alert alert-info alert-dismissible">
    <button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
    <h4><i class="icon fa fa-info"></i>Info</h4>
    <slot></slot>
  </div>
</template>

<script>
  export default {
    name: "message-info",
  }
</script>

<style scoped>

</style>
