<script setup lang="ts">
import { onMounted } from "vue";
import {
  character,
  characters,
  error,
  pick,
  loadList,
  createNew,
  setName,
} from "../composables/useCharacter";

onMounted(loadList);

const onNameChange = async (e: Event) => {
  const name = (e.target as HTMLInputElement).value;
  await setName(name);
};
</script>

<template>
  <div>
    <button v-for="char in characters" :key="char.ID" @click="pick(char.ID)">
      {{ char.Name || "(unnamed)" }}
    </button>
    <button @click="createNew">New Character</button>
  </div>
  <div v-if="character">
    <label>
      Name
      <input :value="character.Base.Name" @change="onNameChange" />
    </label>

    <p>Saved in Go: {{ character.Base.Name || "no name yet" }}</p>
  </div>
  <p v-else>Select or create a character</p>
  <p v-if="error">{{ error }}</p>
</template>
