<script setup lang="ts">
import ScoreBox from "./ScoreBox.vue";
import { signed } from "../../utils/format";
import { models } from "../../../wailsjs/go/models";
import { character, setScore } from "../../composables/useCharacter";

const abilities: (keyof models.CharScores)[] = [
  "Str",
  "Dex",
  "Con",
  "Int",
  "Wis",
  "Cha",
];
</script>

<template>
  <div v-if="character">
    <ScoreBox v-for="ability in abilities" :key="ability" :name="ability" :score="character.Scores[ability]"
      :modifier="signed(character.ScoreModifiers[ability])" @scoreChange="(value) => setScore(ability, value)" />
  </div>
</template>
