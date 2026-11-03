<script setup lang="ts">
import { ref, onMounted } from "vue";
import {
  CreateCharacter,
  ListCharacters,
  SelectCharacter,
  SetName,
} from "../../wailsjs/go/main/CharacterApi";
import { models } from "../../wailsjs/go/models";

const character = ref<models.Character | null>(null);
const characters = ref<models.CharacterSummary[]>([]);
const error = ref<string>("");

onMounted(async () => {
  try {
    characters.value = await ListCharacters();
  } catch (err: unknown) {
    if (typeof err === "string") {
      error.value = err;
    } else {
      console.log("Unexpected error: ", err);
    }
  }
});

const pick = async (id: number) => {
  try {
    character.value = await SelectCharacter(id);
  } catch (err: unknown) {
    if (typeof err === "string") {
      error.value = err;
    } else {
      console.log("Unexpected error: ", err);
    }
  }
};

const createNew = async () => {
  try {
    character.value = await CreateCharacter();
    characters.value = await ListCharacters();
  } catch (err: unknown) {
    if (typeof err === "string") {
      error.value = err;
    } else {
      console.log("Unexpected error: ", err);
    }
  }
};

const saveName = async (e: Event) => {
  const name = (e.target as HTMLInputElement).value;
  try {
    character.value = await SetName(name);
    characters.value = await ListCharacters();
  } catch (err: unknown) {
    if (typeof err === "string") {
      error.value = err;
    } else {
      console.log("Unexpected error: ", err);
    }
  }
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
      <input :value="character.Base.Name" @change="saveName" />
    </label>

    <p>Saved in Go: {{ character.Base.Name || "no name yet" }}</p>
  </div>
  <p v-else>Select or create a character</p>
  <p v-if="error">{{ error }}</p>
</template>
