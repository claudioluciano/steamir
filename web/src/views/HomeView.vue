<script setup lang="ts">
import SearchBar from '@/components/SearchBar.vue'
import { ref } from 'vue'
// import GameItem from '@/components/GameItem.vue'

const searchValue = ref<string>('')

</script>

<template>
  <div class="flex flex-col items-center justify-center w-full gap-4">
    <SearchBar
      class="w-11/12 mt-10 sm:w-9/12 md:w-6/12 lg:w-5/12"
      v-model="searchValue"
    />

    <!-- <div class="flex flex-wrap items-center justify-center w-5/6 gap-3">
      <GameItem
        v-for="i in 13"
        :key="i"
      />
    </div> -->
  </div>
</template>

<style scoped></style>
