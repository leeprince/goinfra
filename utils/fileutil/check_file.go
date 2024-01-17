package fileutil

import (
	"os"
	"path/filepath"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/1 11:06
 * @Desc:
 */

// byteSize：单位：字节
// isThan：是否超过限制
func CheckFileSizeByUrl(url string, byteSize int64) (isLimit bool, err error) {
	fileByteSize, err := GetFileByteSizeByUrl(url)
	if err != nil {
		return
	}

	if fileByteSize > byteSize {
		return true, nil
	}

	return false, nil
}

// 检查文件/目录是否存在
func CheckFileDirExist(filePath string) (os.FileInfo, bool) {
	finfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, false
	}
	return finfo, true
}

func GetFileExtName(filePath string) string {
	ext := filepath.Ext(filePath)
	return ext
}

func CheckFileExtName(filePath string, exts []string) bool {
	ext := GetFileExtName(filePath)
	for _, i2 := range exts {
		if i2 == ext {
			return true
		}
	}
	return false
}

func inSliceString(v string, dest []string) bool {
	for _, i2 := range dest {
		if i2 == v {
			return true
		}
	}
	return false
}

func IsPDF(url string) bool {
	ext := filepath.Ext(url)
	return strings.EqualFold(ext, ".pdf")
}
