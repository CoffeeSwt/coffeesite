<template>
    <div class="main-box">
        <div class="centent-box">
            1111
        </div>
        <div class="centent-box">
            <div class="header-buttons">
                <el-button class="primary-button" type="primary" @click="onAddItem">添加物品</el-button>
            </div>

            <el-table :data="tableData" style="width: 100%">
                <el-table-column fixed prop="name" label="物品名称"  />
                <el-table-column fixed="right" label="Operations" width="120">
                    <template #default="scope">
                        <el-button link type="primary" size="small" @click.prevent="deleteRow(scope.$index)">
                            Remove
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>
<script setup lang="js">
import { reactive, ref } from 'vue'
const tableData = ref([])
const onAddItem = () => {
    tableData.value.push({
        ID: 1,
        name: 'Tom',
        maxCount: 1
    })
}
const deleteRow = (index) => {
    tableData.value.splice(index, 1)
}
</script>
  
<style lang="scss" scoped>
.main-box {
    padding: 24px;
    background-color: #f0f2f5;
}

.centent-box {
    margin-bottom: 12px;
    padding: 24px;
    background-color: #fff;
    border-radius: 2px;
    box-shadow: 0 0 7px 1px rgba(0, 0, 0, 0.03);
}

.header-buttons {
    margin-bottom: 12px;
}
</style>
  