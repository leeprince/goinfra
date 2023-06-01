package fileutil

import "os"

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

// 检查文件是否存在
func CheckFileExist(filePath string) (os.FileInfo, bool) {
	finfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, false
	}
	return finfo, true
}
