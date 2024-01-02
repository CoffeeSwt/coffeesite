<template>
    <div class="main-wrapper">
        <div class="main-wrapper">
            <div class="button-item select-all" @click="selectAll">全部</div>
            <div class="button-item clear-select" @click="clearSelect">清空</div>
            <template v-for="item in contents">
                <div class="button-item" :class="{ 'active-button': activeItem.includes(item) }"
                    @click="changeActive(item)">
                    {{ item }}
                </div>
            </template>
        </div>
        <el-tooltip content="刷新" placement="top">
            <div class="submit-button" @click="submit">
                <el-icon>
                    <Refresh />
                </el-icon>
            </div>
        </el-tooltip>

    </div>
</template>
<script setup lang="js">
import { ref } from 'vue';
const props = defineProps({
    contents: {
        type: Array,
        default: []
    }
})
const emit = defineEmits(['submit'])
const activeItem = ref([])
const changeActive = (item) => {
    const i = activeItem.value.indexOf(item)
    if (i == -1) {
        activeItem.value.push(item)
    } else {
        activeItem.value.splice(i, 1)
    }
}
const getActiveItem = () => {
    return activeItem.value
}
const selectAll = () => {
    activeItem.value = props.contents
}
const clearSelect = () => {
    activeItem.value = []
}

const submit = () => {
    emit('submit', activeItem.value)
}

defineExpose({
    getActiveItem
})
</script>
  
<style lang="scss" scoped>
.main-wrapper {
    display: flex;
    height: 44px;
    border-radius: 8px;
}

.button-item {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 8px 16px;
    cursor: pointer;
    transition: all linear .25s;

    background-color: #d9ecff;
    color: #858585;
}


.button-item:first-child {
    border-radius: 8px 0 0 8px;
}

.button-item:nth-child(2) {
    border-radius: 0 8px 8px 0;
    margin-right: 8px;
}

.button-item:nth-child(3) {
    border-radius: 8px 0 0 8px;
}

.button-item:last-child {
    border-radius: 0 8px 8px 0;
    margin-right: 8px;
}

.active-button {
    background-color: #79bbff;
    color: #e6f2ff;
}

.select-all {
    background-color: #f0f9eb;
    color: #67c23a;
}

.select-all:hover {
    background-color: #67c23a;
    color: #f0f9eb;
}

.clear-select {
    background-color: #fdf6ec;
    color: #eebe77;
}

.clear-select:hover {
    background-color: #eebe77;
    color: #fdf6ec;
}

.submit-button {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 8px 16px;
    cursor: pointer;
    transition: all linear .25s;
    border-radius: 8px;

    background-color: #d9ecff;
    color: #858585;
}

.submit-button:hover {
    background-color: #79bbff;
    color: #e6f2ff;
}
</style>
  