<script setup lang="ts">
import { character, setName } from "../../composables/useCharacter";

const onNameChange = async (e: Event) => {
  const name = (e.target as HTMLInputElement).value;
  await setName(name);
};
</script>

<template>
  <div v-if="character">
    <label class="field">
      <input :value="character.Base.Name" @change="onNameChange" />
      <span>Name</span>
    </label>
    <div class="info">
      <label class="field">
        <input :value="character.Base.Class" />
        <span>Class</span>
      </label>
      <label class="field">
        <input :value="character.Base.Level" />
        <span>Level</span>
      </label>
      <label class="field">
        <input :value="character.Base.Background" />
        <span>Background</span>
      </label>
      <label class="field">
        <input :value="character.Base.Race" />
        <span>Race</span>
      </label>
      <label class="field">
        <input :value="character.Base.Alignment" />
        <span>Alignment</span>
      </label>
    </div>
  </div>
</template>

<style scoped>
div {
  display: flex;
  flex-wrap: wrap;
  gap: 1rem;
}

.info {
  display: flex;
  gap: 0.25rem;
}

.field {
  display: flex;
  flex-direction: column;
}

.field span {
  font-size: 0.7rem;
  text-transform: uppercase;
}

.field input {
  border: none;
  border-bottom: 1px solid currentColor;
}
</style>
