<script lang="ts" setup>
import Sidebar from "./components/sidebar/Sidebar.vue";
</script>

<template>
  <Sidebar />
  <div id="router-view">
    <RouterView />
  </div>
</template>
