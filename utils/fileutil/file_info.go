package fileutil

import (
	"net/url"
	"path/filepath"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/14 9:55
 * @Desc:
 */

type FileInfoOfUrl struct {
	FileName string // 文件名,含后缀
	Name     string // 文件名，不含后缀
	Ext      string // 文件扩展名，不为空时含.
}

func GetFileInfoByUrl(url string) (info FileInfoOfUrl) {
	fileName := filepath.Base(url)
	
	fileInfoArr := strings.Split(fileName, ".")
	
	var name string
	if len(fileInfoArr) >= 1 {
		for i := 0; i < len(fileInfoArr)-1; i++ {
			name += fileInfoArr[i] + "."
		}
		name = strings.TrimRight(name, ".")
	}
	
	var ext string
	// ext = filepath.Ext(url) // 直接使用 fileInfoArr 数组更快些
	if len(fileInfoArr) > 1 {
		ext = "." + fileInfoArr[len(fileInfoArr)-1]
	}
	
	info = FileInfoOfUrl{
		FileName: fileName,
		Name:     name,
		Ext:      ext,
	}
	return
}

// GetFileUrlOfName 获取URL 的文件路径和名称："https://xx-1251506165.cos.ap-shanghai.myqcloud.com/e-document-import-ctl/test/0001.pdf 结果是：e-document-import-ctl/test/0001.pdf
func GetFileUrlOfName(fileURL string) (string, error) {
	u, err := url.Parse(fileURL)
	if err != nil {
		return "", err
	}
	return u.Path[1:], nil // 去掉路径前面的"/"
}

// GetFilePathOfName 获取URL 的文件路径和名称："./0001.pdf 结果是：0001.pdf
func GetFilePathOfName(filePath string) string {
	sArr := strings.Split(filePath, "/")
	return sArr[len(sArr)-1]
}
