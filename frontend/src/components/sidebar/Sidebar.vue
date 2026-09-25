<script setup lang="ts">
import { House, WandSparkles } from "@lucide/vue";
import NavLink from "./NavLink.vue";
</script>

<template>
  <nav>
    <NavLink title="Home" path="/" :icon="House" />
    <NavLink title="Spells" path="/spells" :icon="WandSparkles" />
  </nav>
</template>

<style scoped>
nav {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
  border-right: 1px solid white;
  padding: 0 0.5rem;
}
</style>
