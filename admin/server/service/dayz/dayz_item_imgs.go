package dayz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dayz"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
)

type DayzItemImgsUploadAndDownloadService struct{}

func (e *DayzItemImgsUploadAndDownloadService) Upload(file dayz.DayzItemImg) error {
	return global.GVA_DB.Create(&file).Error
}

func (e *DayzItemImgsUploadAndDownloadService) FindFile(id uint) (error, dayz.DayzItemImg) {
	var file dayz.DayzItemImg
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return err, file
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.ExaFileUploadAndDownload
//@return: err error

//func (e *DayzItemImgsUploadAndDownloadService) DeleteFile(file dayz.DayzItemImg) (err error) {
//	var fileFromDb dayz.DayzItemImg
//	err, fileFromDb = e.FindFile(file.ID)
//	if err != nil {
//		return
//	}
//	oss := upload.NewOss()
//	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
//		return errors.New("文件删除失败")
//	}
//	err = global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
//	return err
//}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (e *DayzItemImgsUploadAndDownloadService) GetFileRecordInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&example.ExaFileUploadAndDownload{})
	var fileLists []example.ExaFileUploadAndDownload
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

//func (e *DayzItemImgsUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (err error, file example.ExaFileUploadAndDownload) {
//	oss := upload.NewOss()
//	filePath, key, uploadErr := oss.UploadFile(header)
//	if uploadErr != nil {
//		panic(err)
//	}
//	if noSave == "0" {
//		s := strings.Split(header.Filename, ".")
//		f := example.ExaFileUploadAndDownload{
//			Url:  filePath,
//			Name: header.Filename,
//			Tag:  s[len(s)-1],
//			Key:  key,
//		}
//		return e.Upload(f), f
//	}
//	return
//}
