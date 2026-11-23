<script setup lang="ts">
import { models } from "../../../wailsjs/go/models";
import {
  setName,
  setClass,
  setLevel,
  setBackground,
  setAlignment,
  setRace,
} from "../../composables/useCharacter";

const { base } = defineProps<{ base: models.CharBase }>();

const onTextChange = async (
  setter: (value: string) => Promise<unknown>,
  e: Event,
) => {
  const value = (e.target as HTMLInputElement).value;
  setter(value);
};

const onLevelChange = async (e: Event) => {
  const level = (e.target as HTMLInputElement).valueAsNumber;
  await setLevel(level);
};
</script>

<template>
  <div>
    <label class="field">
      <input :value="base.Name" @change="onTextChange(setName, $event)" />
      <span>Name</span>
    </label>
    <div class="info">
      <label class="field">
        <input :value="base.Class" @change="onTextChange(setClass, $event)" />
        <span>Class</span>
      </label>
      <label class="field">
        <input type="number" :value="base.Level" @change="onLevelChange" />
        <span>Level</span>
      </label>
      <label class="field">
        <input :value="base.Background" @change="onTextChange(setBackground, $event)" />
        <span>Background</span>
      </label>
      <label class="field">
        <input :value="base.Race" @change="onTextChange(setRace, $event)" />
        <span>Race</span>
      </label>
      <label class="field">
        <input :value="base.Alignment" @change="onTextChange(setAlignment, $event)" />
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
