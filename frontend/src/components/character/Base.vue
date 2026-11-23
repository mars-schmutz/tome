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

const onNameChange = async (e: Event) => {
  const name = (e.target as HTMLInputElement).value;
  await setName(name);
};

const onClassChange = async (e: Event) => {
  const clss = (e.target as HTMLInputElement).value;
  await setClass(clss);
};

const onLevelChange = async (e: Event) => {
  const level = (e.target as HTMLInputElement).valueAsNumber;
  await setLevel(level);
};

const onBackgroundChange = async (e: Event) => {
  const bg = (e.target as HTMLInputElement).value;
  await setBackground(bg);
};

const onRaceChange = async (e: Event) => {
  const race = (e.target as HTMLInputElement).value;
  await setRace(race);
};

const onAlignmentChange = async (e: Event) => {
  const alignment = (e.target as HTMLInputElement).value;
  await setAlignment(alignment);
};
</script>

<template>
  <div>
    <label class="field">
      <input :value="base.Name" @change="onNameChange" />
      <span>Name</span>
    </label>
    <div class="info">
      <label class="field">
        <input :value="base.Class" @change="onClassChange" />
        <span>Class</span>
      </label>
      <label class="field">
        <input type="number" :value="base.Level" @change="onLevelChange" />
        <span>Level</span>
      </label>
      <label class="field">
        <input :value="base.Background" @change="onBackgroundChange" />
        <span>Background</span>
      </label>
      <label class="field">
        <input :value="base.Race" @change="onRaceChange" />
        <span>Race</span>
      </label>
      <label class="field">
        <input :value="base.Alignment" @change="onAlignmentChange" />
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
