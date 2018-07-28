<template>
  <div class="card">
    <div class="card-content">
      <div class="content">
        {{ text }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    text: String,
  },
}
</script>

<style lang="scss" scoped>
</style>
