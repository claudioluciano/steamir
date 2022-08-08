<script setup lang="ts">
defineProps<{modelValue: string}>()
const emit = defineEmits(['update:modelValue'])

function input (event: Event) {
  const target = event.target as HTMLInputElement
  emit('update:modelValue', target.value)
}
</script>

<template>
  <input
    :value="modelValue"
    @input="input"
    placeholder="Search"
    class="border rounded p-2"
  >
</template>

<style scoped></style>
