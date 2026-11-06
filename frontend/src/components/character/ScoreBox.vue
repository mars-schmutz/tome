<script setup lang="ts">
const props = defineProps<{
  name: string;
  score: number;
  modifier: string;
}>();

const emit = defineEmits<{ scoreChange: [value: number] }>();

const onScoreChange = (e: Event) => {
  const num = (e.target as HTMLInputElement).valueAsNumber;
  if (Number.isNaN(num)) return;
  emit("scoreChange", num);
};
</script>

<template>
  <div>
    <label>
      <span>{{ name }}</span>
      <input type="number" :value="score" @change="onScoreChange" />
      <span>{{ modifier }}</span>
    </label>
  </div>
</template>
