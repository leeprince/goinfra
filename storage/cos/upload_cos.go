package cos

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/leeprince/goinfra/utils/fileutil"
	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"os"
	"runtime"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/24 10:27
 * @Desc:
 */

/*
cosClient:resource/cos/cos.go
base64File:是不包含文件格式描述的，如："data:image/png;base64,"
filePath: 文件路径，起始符和结束符存在目录分割符
fileName: 无目录分割符
*/
func (r *CosClient) UploadBase64FileToCos(base64File string, filePath, fileName string) (url string, err error) {
	// base64转bytes
	fileBytes, errn := base64.StdEncoding.DecodeString(base64File)
	if err != nil {
		return "", errors.Wrap(errn, "文件解码错误")
	}

	// 创建临时文件：上传到本地临时路径下，方便上传cos
	tmpPath := fmt.Sprintf("%s%s", "tmp", filePath)
	if runtime.GOOS == "windows" {
		tmpPath = fmt.Sprintf("F:%s%s", string(os.PathSeparator), tmpPath) // windows 本地环境中当前项目是在F:盘中，所以使用根路径/即指向的根路径就是F:盘
	}
	tmpPathFile := fmt.Sprintf("%s%s", tmpPath, fileName)

	if _, err = fileutil.WriteFile(tmpPath, fileName, fileBytes, false); err != nil {
		return "", errors.Wrap(err, "文件写入错误")
	}

	// 上传cos
	cosFileKey := fmt.Sprintf("%s%s", filePath, fileName)
	opt := cos.MultiUploadOptions{PartSize: 5, ThreadPoolSize: 1}
	completeMultipartUploadResult, _, errn := client.Object.Upload(context.Background(), cosFileKey, tmpPathFile, &opt)
	if errn != nil {
		return "", errors.Wrap(errn, "上传cos失败")
	}

	// 删除临时文件
	// err = os.Remove(tmpPathFile)
	if err != nil {
		return "", errors.Wrap(err, "删除临时文件失败")
	}

	url = completeMultipartUploadResult.Location

	return
}

// UploadCos
// 	cosName 起始符不应该有目录分隔符；结束符不能是目录分隔符
func (r *CosClient) UploadCos(fileContentBytes []byte, cosName string, customeAccessHost string, opts *cos.ObjectPutOptions) (cosUrl string, err error) {
	fileReader := bytes.NewReader(fileContentBytes)

	// 自定义访问域名：做反向代理
	if customeAccessHost != "" {
		cosUrl = fmt.Sprintf("%s/%s", customeAccessHost, cosName)
	}

	// 上传cos
	putResponse, err := client.Object.Put(context.Background(), cosName, fileReader, opts)
	if err != nil {
		fmt.Println("Object.Put err:", err)
		return
	}
	if putResponse.StatusCode != http.StatusOK {
		fmt.Println("putResponse.StatusCode !StatusOK")
		return
	}

	if customeAccessHost == "" {
		cosUrl = fmt.Sprintf("https://%s/%s", client.BaseURL.BucketURL.Host, cosName)
	}

	return
}

// GetPresignedURL 获取预先签署Url: 将上传文件到腾讯云COS后的资源，创建一个带有自定义过期时间的预签名下载URL。
// 	使用腾讯云提供的 SDK 或 API 动态生成一个带有下载权限和过期时间的预签名 URL，该 URL 可以用于下载操作。在生成预签名 URL 时，可以指定 response-content-disposition 参数，设置其值为 attachment; filename=image.jpg 这样的形式，强制浏览器下载而非显示图片。
// 	cosName 起始符不应该有目录分隔符；结束符不能是目录分隔符
func (r *CosClient) GetPresignedURL(cosName string, expired time.Duration) (presigndUrl string, err error) {
	/*opt := cos.PresignedURLOptions{
		Query:      nil,
		Header:     nil,
		SignMerged: false,
	}*/
	resp, err := client.Object.GetPresignedURL(context.Background(), http.MethodGet, cosName, r.secretID, r.secretKey, expired, nil)
	if err != nil {
		fmt.Println("Object.Put err:", err)
		return
	}
	//fmt.Println("resp.String()", resp.String())

	presigndUrl = resp.String()

	return
}
