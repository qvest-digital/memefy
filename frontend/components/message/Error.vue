<template>
  <div class="alert alert-danger alert-dismissible">
    <button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
    <h4><i class="icon fa fa-ban"></i>Fehler!</h4>
    <slot></slot>
  </div>
</template>

<script>
  export default {
    name: "message-error",
  }
</script>

<style scoped>

</style>
