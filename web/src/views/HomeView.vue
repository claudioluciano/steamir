<script setup lang="ts">
import SearchBar from '@/components/SearchBar.vue'
import { ref } from 'vue'
import { useResolveVanityURL } from '@/composables/useSteam'
// import GameItem from '@/components/GameItem.vue'

const searchValue = ref<string>('')
const { exec, json, loading, status } = useResolveVanityURL()

async function search () {
  console.log('searchValue', searchValue.value)

  exec(searchValue.value)

  console.log(status.value)
  console.log(json.value)
  console.log(loading.value)
}

</script>

<template>
  <div class="flex flex-col items-center justify-center w-full gap-4">
    <SearchBar
      class="w-5/6 mt-10"
      @keyup.enter="search()"
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
