<template>
    <div class="main-wrapper">
        <div class="main-wrapper">
            <div class="button-item select-all" :class="{ 'active-button-select-all': allItemsSelected }"
                @click="selectAll">全部
            </div>
            <template v-for="item in contents">
                <div class="button-item" :class="{ 'active-button': activeItem.includes(item) }"
                    @click="changeActive(item)">
                    {{ item }}
                </div>
            </template>
        </div>
        <div class="input-form">
            <el-input @clear="handleClear" v-model="inputs" placeholder="请输入物品名称（支持模糊）" size="large" clearable
                class="input">

            </el-input>
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
import { Search } from '@element-plus/icons-vue'
import { computed, ref } from 'vue';
const contents = ref([])
const inputs = ref('')
const emit = defineEmits(['submit'])
const activeItem = ref([])
const allItemsSelected = computed(() => {
    let selectAll = true
    contents.value.forEach((item) => {
        if (!activeItem.value.includes(item)) selectAll = false
    })
    return selectAll
})
const changeActive = (item) => {
    //判断是否处于全选状态
    if (allItemsSelected.value) clearAll()
    //判断是否已经选中
    const i = activeItem.value.indexOf(item)
    if (i == -1) {
        activeItem.value.push(item)
    } else {
        activeItem.value.splice(i, 1)
        if (activeItem.value.length == 0) selectAll()
    }
    emit('submit')
}
const getActiveItem = () => {
    return activeItem.value
}
const selectAll = () => {
    activeItem.value = contents.value.map(i => i)
}
const clearAll = () => {
    activeItem.value = []
}

const setContents = (content) => {
    contents.value = content
    selectAll()
}

const handleClear = () => {
    emit('submit')
}

const submit = () => {
    emit('submit')
}
const getInputName = () => {
    return inputs.value
}

defineExpose({
    getActiveItem,
    selectAll,
    setContents,
    getInputName,
})
</script>
  
<style lang="scss" scoped>
.main-wrapper {
    display: flex;
    height: 44px;
    border-radius: 8px;
}

.input-form {
    align-self: center;
    margin: 0 8px;

    .input {
        width: 320px;
        height: 44px;
    }
}

.button-item {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 8px 16px;
    cursor: pointer;
    transition: background-color linear .25s;

    background-color: #d9ecff;
    color: #858585;
}

.button-item:first-child {
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

.active-button-select-all {
    background-color: #67c23a;
    color: #f0f9eb;
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
  