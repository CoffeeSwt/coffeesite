<template>
    <main ref="wrapper" class="wrapper" w-full h-screen relative overflow-hidden>
        <div relative>
            <section ref="section1" h-screen bg-blue-100>
                123123
            </section>
            <section ref="section2" h-screen>
                <header h-24 flex>
                    <div class="w-3/5" h-full bg-red-300></div>
                </header>
                <main></main>
            </section>
            <section h-screen bg-gray-400>12312312</section>
            <section h-screen bg-gray-500>3123123</section>
        </div>
    </main>
</template>

<script setup lang="ts">
import type { Ref } from 'vue';

//
import BScroll from '@better-scroll/core'
import MouseWheel from '@better-scroll/mouse-wheel'
import ScrollBar from '@better-scroll/scroll-bar'
BScroll.use(MouseWheel)
BScroll.use(ScrollBar)
const wrapper: Ref<HTMLDivElement | null> = ref(null)

onMounted(() => {
    new BScroll('.wrapper', {
        wheel: true,
        scrollbar: false,
        bounce: false,
        probeType: 2,
        mouseWheel: {
            speed: 30,
            invert: false,
            easeTime: 300
        }
    })
})

</script>

<style scoped></style>