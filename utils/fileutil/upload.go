package fileutil

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"os"
	"runtime"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/24 10:27
 * @Desc:
 */

/*
cosClient:resource/cos/cos.go
base64File:是不包含文件格式描述的，如："data:image/png;base64,"
filePath: 文件路径，起始符和结束符有存在目录分割符
fileName: 无目录分割符
*/
func UploadBase64FileToCos(cosClient *cos.Client, base64File string, filePath, fileName string) (url string, err error) {
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
	if err = CreateFileAndWrite(tmpPath, tmpPathFile, fileBytes); err != nil {
		return "", errors.Wrap(err, "文件写入错误")
	}

	// 上传cos
	cosFileKey := fmt.Sprintf("%s%s", filePath, fileName)
	opt := cos.MultiUploadOptions{PartSize: 5, ThreadPoolSize: 1}
	completeMultipartUploadResult, _, errn := cosClient.Object.Upload(context.Background(), cosFileKey, tmpPathFile, &opt)
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

// 创建指定路径的文件，并写入数据
func CreateFileAndWrite(path string, pathFile string, fileBytes []byte) (err error) {
	// 创建文件
	if _, err = os.Stat(pathFile); err != nil {
		if err = os.MkdirAll(path, 0777); err != nil {
			return errors.Wrap(err, "创建文件路径错误")
		}
	}

	f, errn := os.Create(pathFile)
	if errn != nil {
		return errors.Wrap(errn, "创建文件错误")
	}
	defer f.Close()
	// 写入数据
	_, err = f.Write(fileBytes)
	if err != nil {
		return errors.Wrap(errn, "写入文件错误")
	}
	return nil
}
