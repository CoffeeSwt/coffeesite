<template>
    <div class="main-box">
        <ContentBox>
            <StateButtonGroup ref="dayzItemCategoryButtonRef" @submit="setDayzItemsTableData"></StateButtonGroup>
        </ContentBox>

        <ContentBox>
            <div class="control-buttons">
                <div class="add-item-button" @click="addNewItemDialogVisible = true">
                    <el-icon :size="14">
                        <Plus />
                    </el-icon>
                    <span style="margin-left: 4px;">新增物品</span>
                </div>
                <div class="delete-item-button" @click="deletItemsSelected">
                    <el-icon :size="14">
                        <Delete />
                    </el-icon>
                    <span style="margin-left: 4px;">删除选中</span>
                </div>
            </div>

            <el-table ref="dayzItemsTableRef" class="el-table-main" :data="dayzItemsTableData" table-layout="auto"
                @selection-change="handleSelectionChange">
                <el-table-column type="selection"></el-table-column>

                <el-table-column prop="ID" label="ID" width="40" />
                <el-table-column prop="name" label="物品名称" />
                <el-table-column prop="category" label="类别" />
                <el-table-column prop="info" label="信息" />
                <el-table-column label="图片" width="240">
                    <template #default="scope">

                    </template>
                </el-table-column>
                <el-table-column fixed="right" label="操作">
                    <template #default="scope">
                        <el-button class="row-button" :icon="Edit" type="primary" size="small" @click="editRow(scope)">
                            查看/编辑
                        </el-button>
                        <el-button class="row-button" :icon="Delete" type="danger" size="small" @click="deleteRow(scope)">
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
                <el-form-item label="最大数量">
                    <el-input-number v-model="addNewItemForm.maxQuantity" :step="10" />
                </el-form-item>
                <el-form-item label="物品信息">
                    <el-input v-model="addNewItemForm.info" :autosize="{ minRows: 4, maxRows: 8 }" type="textarea"
                        placeholder="请输入物品信息" />
                </el-form-item>
                <el-form-item label="物品图片">
                    <el-upload ref="uploadRef" :action="`${path}/dayzItemImg/uploadFile`" list-type="picture-card"
                        v-model:file-list="imgUploadList" :auto-upload="false" accept=".png, .jpg, .jpeg"
                        :headers="{ 'x-token': token }" :drag="true" v-model:data="imgUploadData">
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
import { useStore } from 'vuex'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getDict } from '@/utils/dictionary'
import ContentBox from '@/components/global/ContentBox.vue'
import StateButtonGroup from '@/view/dayzitems/components/StateButtonGroup.vue';
import { Delete, Edit, Plus } from '@element-plus/icons-vue'
import {
    ref,
    computed,
    watchEffect,
    reactive,
    onMounted,
} from 'vue'
import * as dayzItemApi from '@/api/dayz/dayzitem'
import * as dayzItemImgApi from '@/api/dayz/dayzItemImg'
const path = import.meta.env.VITE_BASE_API
const store = useStore()
const token = computed(() => store.state.user.token)

//物品类别
const dayzItemCategory = ref([])
const getLabelByIndex = (index) => {
    return dayzItemCategory.value.find((i) => i.value == index).label
}
//物品类别按钮Ref
const dayzItemCategoryButtonRef = ref(null)
watchEffect(async () => {
    dayzItemCategory.value = await getDict('dayzItemCategory')
    dayzItemCategoryButtonRef.value.setContents(dayzItemCategory.value.map(i => i.label))
})
//分页
const pageInfo = reactive({
    pageSize: 10,
    total: null,
    page: 1
})

//添加物品的表单
const addNewItemDialogVisible = ref(false)
const addNewItemFormDefault = {
    name: '',
    category: '',
    info: '',
    maxQuantity: 0
}
const addNewItemForm = reactive({
    ...addNewItemFormDefault
})

//清空表单项：新增物品的表单
const resetAddNewItemForm = () => {
    Object.keys(addNewItemForm).forEach((key) => {
        addNewItemForm[key] = addNewItemFormDefault[key]
    })
    imgUploadList.value = []
    previewImgDialogUrl.value = ''
}
//验证表单:新增物品
const checkAddNewItemForm = () => {
    if (addNewItemForm.name == '') return false
    if (addNewItemForm.category == 0) return false
    return true
}
//提交表单：新增物品的表单
const submitAddNewItemForm = async () => {
    if (!checkAddNewItemForm()) {
        ElMessage({
            message: '请填写所有必填项',
            type: 'warning',
        })
        return
    }
    uploadFile()
    await dayzItemApi.createDayzItem(addNewItemForm)
    addNewItemDialogVisible.value = false
    resetAddNewItemForm()
    setDayzItemsTableData()
}
const uploadRef = ref(null)
const uploadFile = () => {
    uploadRef.value.submit()
}

//取消表单对话框：新增物品的表单
const cancelAddNewItemForm = () => {
    const confirmCancel = () => {
        addNewItemDialogVisible.value = false
        resetAddNewItemForm()
    }
    const hasInputForm = computed(() => {
        if (Object.values(addNewItemForm).filter((i) => !(i == '' || i == 0)).length > 0) return true
        if (imgUploadList.value.length > 0) return true
        return false
    })

    if (!hasInputForm.value) { confirmCancel(); return }
    ElMessageBox.confirm(
        '确认取消吗？填写的内容不会保存',
        '',
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    ).then(() => {
        confirmCancel()
    })
}

const useQueryParm = () => {
    const category = dayzItemCategoryButtonRef.value.getActiveItem()
    const categoryIndex = category.map((label) => dayzItemCategory.value.findIndex((item) => item.label == label) + 1)
    return {
        pageSize: pageInfo.pageSize,
        page: 1,
        category: categoryIndex
    }
}

const deletItemByIDs = (ids) => {
    ElMessageBox.confirm(
        '确认删除吗？',
        '',
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    ).then(async () => {
        await dayzItemApi.deleteDayzItemByIds({ ids })
        setDayzItemsTableData()
    })
}

const dayzItemsTableRef = ref(null)
const dayzItemsTableData = ref([])
const dayzItemsTableSelection = ref([])
const setDayzItemsTableData = async () => {
    const res = await dayzItemApi.getDayzItem(useQueryParm())
    const items = res.data.items
    dayzItemsTableData.value = items.map((i) => {
        return {
            ...i,
            category: getLabelByIndex(i.category),
        }
    })
}
onMounted(() => {
    setDayzItemsTableData()
})

const handleSelectionChange = (val) => {
    dayzItemsTableSelection.value = val
}
const deletItemsSelected = () => {
    const ids = dayzItemsTableSelection.value.map(i => i.ID)
    deletItemByIDs(ids)
}
const editRow = (scope) => {
    console.log(scope)
}
const deleteRow = (scope) => {
    deletItemByIDs([scope.row.ID])
}

const imgUploadData = reactive({
    
})
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

    margin-bottom: 24px;
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

:deep(.el-upload--picture-card) {
    display: flex;
    justify-content: center;
    align-items: center;

}

:deep(.el-upload-dragger) {
    width: 100%;
    height: 100%;
    border: none;
    display: flex;
    justify-content: center;
    align-items: center;

}
</style>
  