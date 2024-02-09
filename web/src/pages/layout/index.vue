<template>
    <div class="layout" flex flex-col w-full h-screen overflow-hidden >
        <header bg-gray-200>
            <HeaderVue></HeaderVue>
        </header>
        <main grow bg-gray-300>
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