<template>
    <div class="main-wrapper">
        
    </div>
</template>

<script setup >
import { ref } from 'vue'
const props = defineProps({
    steps: {
        type: Array,
        required: true
    }
})
const modelValue = ref(0)

</script>

<style lang="scss" scoped></style>