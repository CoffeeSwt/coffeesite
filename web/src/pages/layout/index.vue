<template>
    <div class="layout">
        <header>
            <HeaderVue></HeaderVue>
        </header>
        <main>
            <RouterView></RouterView>
        </main>
        <footer></footer>
    </div>
</template>

<script setup lang="ts">
import HeaderVue from '@/components/layout/HeaderVue.vue'
import Rellax, { RellaxInstance } from 'rellax';
const rellax = ref<RellaxInstance | null>(null)

onMounted(() => {
    rellax.value = new Rellax('.rellax', {
        wrapper: '.rellax-wrapper'
    })
    // rellax.value.destroy();
})

</script>

<style scoped></style>