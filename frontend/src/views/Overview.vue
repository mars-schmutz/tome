<script setup lang="ts">
const charName = defineModel();
</script>

<template>
  <div>
    <input v-model="charName" />
    <div>
      <input />
      <input />
      <input />
    </div>
  </div>
</template>
