package pdfutil

import (
	"errors"
	"fmt"
	"github.com/h2non/bimg"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/utils/fileutil"
	"path"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/14 9:34
 * @Desc:	pdf 转图片 待解决本地环境问题：windows、linux、mac
 */

// PdfToImageV1
/*
- windows 环境依赖 pkg-config 和 libvips。

- linux 直接运行也会报错：
	#15 44.06 # github.com/h2non/bimg
	#15 44.06 /go/pkg/mod/github.com/h2non/bimg@v1.1.9/image.go:93:49: undefined: Gravity
	#15 44.06 /go/pkg/mod/github.com/h2non/bimg@v1.1.9/image.go:133:29: undefined: Watermark
	#15 44.06 /go/pkg/mod/github.com/h2non/bimg@v1.1.9/image.go:139:34: undefined: WatermarkImage
	#15 44.06 /go/pkg/mod/github.com/h2non/bimg@v1.1.9/image.go:152:26: undefined: Angle
	#15 44.06 /go/pkg/mod/github.com/h2non/bimg@v1.1.9/image.go:181:31: undefined: Interpretation
	#15 44.06 /go/pkg/mod/github.com/h2non/bimg@v1.1.9/image.go:202:27: undefined: Options
	#15 44.06 /go/pkg/mod/github.com/h2non/bimg@v1.1.9/image.go:212:29: undefined: ImageMetadata
	#15 44.06 /go/pkg/mod/github.com/h2non/bimg@v1.1.9/image.go:218:35: undefined: Interpretation
	#15 44.06 /go/pkg/mod/github.com/h2non/bimg@v1.1.9/image.go:234:25: undefined: ImageSize
	#15 44.06 /go/pkg/mod/github.com/h2non/bimg@v1.1.9/resize.go:11:27: undefined: Options
	#15 44.06 /go/pkg/mod/github.com/h2non/bimg@v1.1.9/resize.go:11:27: too many errors
*/
func PdfToImageV1(pdfUrl string) (imagePath string, err error) {
	fileInfo := fileutil.GetFileInfoByUrl(pdfUrl)
	if fileInfo.Ext != ".pdf" {
		err = errors.New("pdfurl no .pdf")
		return
	}

	filePath := fmt.Sprintf("/%s/%s/", "goinfra", consts.ENV_LOCAL)
	fileName := fmt.Sprintf("pdftojpeg_pdf_tmp-%s", fileInfo.FileName)
	fileNameOfJpeg := fmt.Sprintf("pdftojpeg_jpeg_%s", fileInfo.Name, ".jpeg")
	pathFile, err := fileutil.SaveLocalFileByUrl(pdfUrl, fileName, filePath)
	if err != nil {
		return
	}

	buffer, err := bimg.Read(pathFile)
	if err != nil {
		return
	}

	newImage, err := bimg.NewImage(buffer).Convert(bimg.JPEG)
	if err != nil {
		return
	}
	//bimg.NewImage(newImage).Type() // 获取新的图片文件类型

	imagePath = path.Join(filePath, fileNameOfJpeg)
	err = bimg.Write(imagePath, newImage)
	return
}
