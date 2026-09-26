<script setup lang="ts">
import { ref, onMounted } from "vue";
import { GetCharacter, SetName } from "../../wailsjs/go/main/CharacterApi";
import { models } from "../../wailsjs/go/models";

const character = ref<models.Character | null>(null);

const nameInput = ref("");

onMounted(async () => {
  character.value = await GetCharacter();
  nameInput.value = character.value.Base.Name;
});

const saveName = async () => {
  character.value = await SetName(nameInput.value);
};
</script>

<template>
  <div v-if="character">
    <label>
      Name
      <input v-model="nameInput" @change="saveName" />
    </label>

    <p>Saved in Go: {{ character.Base.Name || "no name yet" }}</p>
  </div>
  <p v-else>Loading</p>
</template>
