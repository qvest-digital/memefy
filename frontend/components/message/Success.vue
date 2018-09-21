<template>
  <div class="alert alert-success alert-dismissible">
    <button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
    <h4><i class="icon fa fa-check"></i>Erfolgreich</h4>
    <slot></slot>
  </div>
</template>

<script>
  export default {
    name: "message-success",
  }
</script>

<style scoped>

</style>
