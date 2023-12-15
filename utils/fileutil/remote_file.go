package fileutil

import (
	"bytes"
	"fmt"
	errors2 "github.com/pkg/errors"
	"io"
	"net/http"
	"os"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/1 11:07
 * @Desc:	url 文件：如：cos 文件、oos 文件
 */

func ReadFileBytesByUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors2.Errorf("resp.StatusCode != http.StatusOK")
		return nil, err
	}

	contentLength := resp.Header.Get("Content-Length")
	if contentLength == "0" {
		err = errors2.Errorf("contentLength == 0")
		return nil, err
	}

	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors2.Wrap(err, "ioutil.ReadAll")
		return nil, err
	}
	if fileLen := len(fileBytes); fileLen <= 0 {
		err = errors2.Errorf("fileBytes len 0")
		return nil, err
	}

	return fileBytes, nil
}

func GetFileByteSizeByUrl(url string) (byteSize int64, err error) {
	resp, err := http.Head(url)
	if err != nil {
		return
	}

	return resp.ContentLength, nil
}

func ReadFileReaderByUrl(url string) (io.Reader, []byte, error) {
	fileBytes, err := ReadFileBytesByUrl(url)
	if err != nil {
		return nil, nil, err
	}

	return bytes.NewReader(fileBytes), fileBytes, nil
}

// 将io.Reader类型的数据转换为.File类型的数据
func WriteFileByIoReader(data io.Reader, fileName string, filePath ...string) (pathFile string, err error) {
	path := "tmp"
	if len(filePath) > 0 {
		path = filePath[0]
	}
	pathFile = fmt.Sprintf("%s%s", path, fileName) // 本地环境中当前项目是在F:盘中，所以使用根路径/即指向的根路径就是F:盘

	// 创建文件
	if _, err := os.Stat(pathFile); err != nil {
		if err = os.MkdirAll(path, 0777); err != nil {
			return "", err
		}
	}

	tmpPathFile, err := os.Create(pathFile)
	if err != nil {
		return
	}
	defer func() {
		tmpPathFile.Close()
		// os.Remove(pathFile) // 外部自行删除
	}()
	copyInt, err := io.Copy(tmpPathFile, data)
	if err != nil {
		err = errors2.Errorf("contentLength ParseInt")
		return
	}
	if copyInt == 0 {
		err = errors2.Errorf("io.Copy 0")
		return
	}

	return
}

// 将io.Reader类型的数据转换为.File类型的数据
func SaveLocalFileByIoReader(data io.Reader, fileName string, filePath ...string) (pathFile string, err error) {
	return WriteFileByIoReader(data, fileName, filePath...)
}

func WriteFileByUrl(url, fileName string, filePath ...string) (pathFile string, err error) {
	readerData, _, err := ReadFileReaderByUrl(url)
	if err != nil {
		err = errors2.Wrap(err, "ReadFileReaderByUrl")
		return
	}

	return SaveLocalFileByIoReader(readerData, fileName, filePath...)
}

func SaveLocalFileByUrl(url, fileName string, filePath ...string) (pathFile string, err error) {
	return WriteFileByUrl(url, fileName, filePath...)
}
