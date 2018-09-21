<template>
  <div class="alert alert-warning alert-dismissible">
    <button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
    <h4><i class="icon fa fa-warning"></i>Warnung</h4>
    <slot></slot>
  </div>
</template>

<script>
  export default {
    name: "message-warning",
  }
</script>

<style scoped>

</style>
