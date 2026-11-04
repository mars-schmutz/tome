<script setup lang="ts">
import Base from "../components/character/Base.vue";
import { onMounted } from "vue";
import {
  character,
  characters,
  error,
  pick,
  loadList,
  createNew,
} from "../composables/useCharacter";

onMounted(loadList);
</script>

<template>
  <div>
    <button v-for="char in characters" :key="char.ID" @click="pick(char.ID)">
      {{ char.Name || "(unnamed)" }}
    </button>
    <button @click="createNew">New Character</button>
  </div>
  <div v-if="character">
    <Base />
    <p>Saved in Go: {{ character.Base.Name || "no name yet" }}</p>
  </div>
  <p v-else>Select or create a character</p>
  <p v-if="error">{{ error }}</p>
</template>
