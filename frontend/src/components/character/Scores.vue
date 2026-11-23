<script setup lang="ts">
import ScoreBox from "./ScoreBox.vue";
import { signed } from "../../utils/format";
import { abilities } from "../../utils/character";
import { character, setScore } from "../../composables/useCharacter";
</script>

<template>
  <div v-if="character">
    <ScoreBox v-for="ability in abilities" :key="ability" :name="ability" :score="character.Scores[ability]"
      :modifier="signed(character.ScoreModifiers[ability])" @scoreChange="(value) => setScore(ability, value)" />
  </div>
</template>
