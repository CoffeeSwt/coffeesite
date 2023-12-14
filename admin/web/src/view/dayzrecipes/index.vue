<template>
    <div class="main-box">
        <div class="centent-box">11112312</div>
    </div>
</template>
<script setup lang="js">
import { findDayzItems } from '@/api/dayzItems'
import { onMounted } from 'vue';
console.log(findDayzItems)
onMounted(async () => {
    const res = await findDayzItems({
        ID: 1
    })
    console.log(res)
})
</script>
  
<style lang="scss" scoped>
.main-box {
    padding: 24px;
    background-color: #f0f2f5;
}

.centent-box {
    padding: 24px;
    background-color: #fff;
    border-radius: 2px;
    box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
}
</style>
  