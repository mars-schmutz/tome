<script setup lang="ts">
import type { Component } from "vue";

defineProps<{
  title: string;
  path: string;
  icon: Component;
}>();
</script>

<template>
  <RouterLink :to="path">
    <component :is="icon" />
    <span>{{ title }}</span>
  </RouterLink>
</template>

<style scoped>
a {
  color: white;
  text-decoration: none;
  font-size: 1.5rem;
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 0.25rem 1rem;
}
</style>
