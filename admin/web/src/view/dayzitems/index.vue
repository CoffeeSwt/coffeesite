<template>
    <div class="main-box">
        <ContentBox>
            <StateButtonGroup ref="dayzItemCategoryButtonRef" :contents="dayzItemCategoryButtons"></StateButtonGroup>
        </ContentBox>
        <ContentBox>
            <div class="control-buttons">
                <div class="add-item-button" @click="addNewItemDialogVisible = true">
                    <el-icon :size="14">
                        <Plus />
                    </el-icon>
                    <span style="margin-left: 4px;">新增物品</span>
                </div>
                <div class="delete-item-button">
                    <el-icon :size="14">
                        <Delete />
                    </el-icon>
                    <span style="margin-left: 4px;">删除选中</span>
                </div>
            </div>
        </ContentBox>

        <ContentBox>
            <el-table ref="dayzItemsTableRef" :data="dayzItemsTableData" style="width: 100%">
                <el-table-column type="selection"></el-table-column>

                <el-table-column prop="ID" label="ID" width="40" />
                <el-table-column prop="name" label="物品名称" width="100" />
                <el-table-column prop="category" label="类别" width="200" />
                <el-table-column prop="info" label="信息" width="200" />
                <el-table-column label="图片" width="240">
                    <template #default="scope">

                    </template>
                </el-table-column>
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

        <!-- 新增物品的对话弹窗 -->
        <el-dialog v-model="addNewItemDialogVisible" width="40%" :show-close="false">
            <el-form :model="addNewItemForm" label-width="120px">
                <el-form-item required label="物品名称">
                    <el-input v-model="addNewItemForm.name" placeholder="请输入物品名称" clearable />
                </el-form-item>
                <el-form-item required label="物品类别">
                    <el-select v-model="addNewItemForm.category" placeholder="请选择物品类别" size="large">
                        <el-option v-for="item in dayzItemCategory" :key="item.value" :label="item.label"
                            :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="物品信息">
                    <el-input v-model="addNewItemForm.info" :autosize="{ minRows: 4, maxRows: 8 }" type="textarea"
                        placeholder="请输入物品信息" />
                </el-form-item>
                <el-form-item label="物品图片">
                    <div class="img-upload-form-item">

                        <el-upload action="#" list-type="picture-card" v-model:file-list="imgUploadList"
                            :auto-upload="false" accept=".png, .jpg, .jpeg">
                            <el-icon>
                                <Plus />
                            </el-icon>

                            <template #file="{ file }">
                                <img class="el-upload-list__item-thumbnail" :src="file.url" alt="" />
                                <span class="el-upload-list__item-actions">
                                    <span class="el-upload-list__item-preview" @click="handlePictureCardPreview(file)">
                                        <el-icon><zoom-in /></el-icon>
                                    </span>
                                    <span class="el-upload-list__item-delete" @click="handleRemove(file)">
                                        <el-icon>
                                            <Delete />
                                        </el-icon>
                                    </span>
                                </span>
                            </template>
                            <template #tip>
                                <div class="el-upload__tip">
                                    请上传4M以内的 png/jpeg 格式图片
                                </div>
                            </template>
                        </el-upload>
                    </div>

                </el-form-item>

            </el-form>
            <footer class="info-form-footer">
                <el-button type="default" size="large" @click="cancelAddNewItemForm">取消</el-button>
                <el-button type="primary" size="large" @click="submitAddNewItemForm">确定</el-button>
            </footer>
        </el-dialog>
        <el-dialog v-model="previewImgDialogVisible">
            <img class="prview-Img" :src="previewImgDialogUrl" alt="Preview Image" />
        </el-dialog>
    </div>
</template>
<script setup lang="js">
import { getDict } from '@/utils/dictionary'
import ContentBox from '@/components/global/ContentBox.vue'
import StateButtonGroup from '@/view/dayzitems/components/StateButtonGroup.vue';
import { Delete, Edit, Plus } from '@element-plus/icons-vue'
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

//添加物品的表单
const addNewItemDialogVisible = ref(false)
const addNewItemForm = reactive({
    name: '',
    category: '',
    info: ''
})

const resetAddNewItemForm = () => {
    Object.keys(addNewItemForm).forEach((key) => {
        if (typeof addNewItemForm[key] == 'string') {
            addNewItemForm[key] = ''
        }
        if (typeof addNewItemForm[key] == 'number') {
            addNewItemForm[key] = 0
        }
    })
}
const submitAddNewItemForm = () => {
    addNewItemDialogVisible.value = false
    resetAddNewItemForm()
}
const cancelAddNewItemForm = () => {
    addNewItemDialogVisible.value = false
    resetAddNewItemForm()
}
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
const previewImgDialogVisible = ref(false)
const previewImgDialogUrl = ref()
const imgUploadList = ref([])
const handlePictureCardPreview = (file) => {
    previewImgDialogUrl.value = file.url
    previewImgDialogVisible.value = true
}
const handleRemove = (file) => {
    const index = imgUploadList.value.indexOf(file)
    imgUploadList.value.splice(index, 1)
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

:deep(.el-dialog__header) {
    display: none;
}

.info-form-footer {
    margin: 0 auto;
    width: 40%;
    display: flex;
    justify-content: space-between;
}

.prview-Img {
    width: 100%;
    height: 100%;
}

.img-upload-form-item {
    .upload-tips {
        font-size: 14px;
        color: #858585;
    }
}
</style>
  