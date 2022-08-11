<script setup lang="ts">
import ArrowRight from 'virtual:icons/mdi/arrow-right'

defineProps<{modelValue: string}>()
const emit = defineEmits(['update:modelValue'])

function input (event: Event) {
  const target = event.target as HTMLInputElement
  emit('update:modelValue', target.value)
}
</script>

<template>
  <div class="relative w-full">
    <input
      type="search"
      :value="modelValue"
      @input="input"
      placeholder="Search"
      class="w-full p-4 outline-none"
      required
    >
    <button
      type="button"
      class="absolute inset-y-0 right-0 flex items-center pr-3"
    >
      <ArrowRight />
    </button>
  </div>
</template>

<style scoped>
</style>
