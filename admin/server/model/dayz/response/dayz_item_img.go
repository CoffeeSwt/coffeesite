package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/dayz"

type DayzItemImgResponse struct {
	File dayz.DayzItemImg `json:"file"`
}
