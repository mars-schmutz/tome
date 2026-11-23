<script setup lang="ts">
import { models } from "../../../wailsjs/go/models";
import { signed } from "../../utils/format";
import { abilities, type Ability } from "../../utils/character";
import { setSaveProficiency } from "../../composables/useCharacter";

const { saves } = defineProps<{ saves: models.CharSaves }>();

const onProficiencyChange = (ability: Ability, e: Event) => {
  const proficient = (e.target as HTMLInputElement).checked;
  setSaveProficiency(ability, proficient);
};
</script>

<template>
  <div class="saves">
    <label v-for="ability in abilities" :key="ability" class="save">
      <input
        type="checkbox"
        :checked="saves[ability].Proficient"
        @change="onProficiencyChange(ability, $event)"
      />
      <span class="bonus">{{ signed(saves[ability].Bonus) }}</span>
      <span>{{ ability }}</span>
    </label>
  </div>
</template>

<style scoped>
.saves {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.save {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.bonus {
  min-width: 2ch;
  text-align: right;
}
</style>
