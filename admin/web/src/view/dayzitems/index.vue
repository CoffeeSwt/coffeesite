<template>
    <div class="main-box">
        <ContentBox>
            <StateButtonGroup ref="dayzItemCategoryButtonRef" :contents="dayzItemCategoryButtons"></StateButtonGroup>
        </ContentBox>
        <ContentBox>
            <div class="control-buttons">
                <div class="add-item-button">新增物品</div>
                <div class="delete-item-button">删除选中</div>
            </div>

        </ContentBox>
        <ContentBox>
            <el-table ref="dayzItemsTableRef" :data="tableData" style="width: 100%">
                <el-table-column type="selection"></el-table-column>

                <el-table-column prop="ID" label="ID" width="40" />
                <el-table-column prop="name" label="字段名称" width="100" />
                <el-table-column prop="path" label="页面路径" width="200" />
                <el-table-column prop="zh" label="中文" :formatter="(row) => sliceString(row.zh, 20)" />
                <el-table-column prop="en" label="英文" :formatter="(row) => sliceString(row.en, 40)" />
                <el-table-column prop="ja" label="日文" :formatter="(row) => sliceString(row.ja, 20)" />

                <el-table-column fixed="right" label="操作" width="240">
                    <template #default="scope">
                        <el-button class="row-button" :icon="Edit" type="primary" size="small"
                            @click="editRow(scope.$index)">
                            查看/编辑
                        </el-button>
                        <el-button class="row-button" :icon="Delete" type="danger" size="small"
                            @click="deleteRow(scope.$index)">
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
        </ContentBox>
    </div>
</template>
<script setup lang="js">
import { getDict } from '@/utils/dictionary'
import ContentBox from '@/components/global/ContentBox.vue'
import StateButtonGroup from '@/view/dayzitems/components/StateButtonGroup.vue';
import {
    ref,
    computed,
    watchEffect,
    reactive
} from 'vue'

//物品类别
const dayzItemCategory = ref([])
watchEffect(async () => {
    dayzItemCategory.value = await getDict('dayzItemCategory')
})
//物品类别按钮内容
const dayzItemCategoryButtons = computed(() => {
    return dayzItemCategory.value.map(i => i.label)
})
//物品类别按钮Ref
const dayzItemCategoryButtonRef = ref(null)

const pageInfo = reactive({
    pageSize: 10,
    total: null,
    page: 1
})

const queryParm = computed(() => {
    const category = dayzItemCategoryButtonRef.value.getActiveItem()
    return {
        pageSize: pageInfo.pageSize,
        page: 1
    }
})
const dayzItemsTableRef = ref(null)
const dayzItemsTableData = ref([])
watchEffect(() => {

})
const editRow = (index) => {
    console.log(index)
}
const deleteRow = (index) => {
    console.log(index)
}


</script>
  
<style lang="scss" scoped>
.main-box {
    padding: 24px;
    background-color: #f0f2f5;
}

.add-item-button {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 8px 16px;
    cursor: pointer;
    transition: all linear .25s;
    border-radius: 8px;

    background-color: #d9ecff;
    color: #858585;

    margin-right: 8px;
}

.control-buttons {
    display: flex;
    height: 44px;
    border-radius: 8px;

}

.add-item-button:hover {
    background-color: #79bbff;
    color: #e6f2ff;
}

.delete-item-button {
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 8px 16px;
    cursor: pointer;
    transition: all linear .25s;
    border-radius: 8px;

    background-color: #fef0f0;
    color: #f56c6c;
}

.delete-item-button:hover {
    background-color: #f56c6c;
    color: #fef0f0;
}
</style>
  