<template>
  <RouterView></RouterView>
</template>
<script setup lang="ts">
import { useLanguageStore } from '@/pinia'
useLanguageStore().init()
</script>
<style scoped></style>
