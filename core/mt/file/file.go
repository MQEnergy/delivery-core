package file

import (
	"github.com/MQEnergy/delivery-core/core/mt"
	mt2 "github.com/MQEnergy/delivery-core/routers/mt"
)

type File struct {
	Base *mt.Base
}

func New(appKey, appSecret string) *File {
	return &File{
		Base: mt.New(appKey, appSecret),
	}
}

// ImageUpload 图片上传 visit: https://peisong.meituan.com/tscc/docNew#%E5%9B%BE%E7%89%87%E4%B8%8A%E4%BC%A0%E6%8E%A5%E5%8F%A3
func (f *File) ImageUpload(params map[string]interface{}) (string, error) {
	return f.Base.WithRequestParams(mt2.IMAGE_UPLOAD_URL, params).Result()
}
