<template>
  <RouterView></RouterView>
</template>
<script setup lang="ts">
import { provide } from 'vue'
import { useLanguageStore } from '@/pinia'
provide(/* 注入名 */ 'message', /* 值 */ 'hello!')
useLanguageStore().init()
</script>
<style scoped></style>
