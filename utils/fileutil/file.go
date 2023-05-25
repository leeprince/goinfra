package fileutil

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

/**
 * @Author: prince.lee
 * @Date:   2022/2/14 11:28
 * @Desc:
 */

var (
	FileNoExistErr = errors.New("file not exist")
)

var osType string
var osFilePath string

func init() {
	osType = runtime.GOOS
	if os.IsPathSeparator('\\') { // 前边的判断是否是系统的分隔符
		osFilePath = "\\"
	} else {
		osFilePath = "/"
	}
}

// 检查文件是否存在
func CheckFileExist(filePath string) (os.FileInfo, bool) {
	finfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, false
	}
	return finfo, true
}

// 写入数据到文件
func WriteFile(dirPath, filename string, data []byte, isAppend bool) (ok bool, err error) {
	filePath := filepath.Join(dirPath, filename)
	if _, ok = CheckFileExist(filePath); !ok {
		// 创建目录
		err = os.MkdirAll(dirPath, os.ModePerm)
		if ok = os.IsNotExist(err); ok {
			return
		}
	}

	flag := os.O_CREATE | os.O_WRONLY
	if isAppend {
		flag = flag | os.O_APPEND
	}
	fs, fErr := os.OpenFile(filePath, flag, 0666)
	if fErr != nil {
		err = fErr
		return
	}
	defer fs.Close()

	// 创建带有缓冲区的Writer对象
	writer := bufio.NewWriter(fs)
	// 写入数据
	if _, err = writer.Write(data); err != nil {
		return
	}
	// 自动添加换行符
	if isAppend {
		if _, err = writer.Write([]byte("\n")); err != nil {
			return
		}
	}

	// 刷新缓冲区
	writer.Flush()

	ok = true
	return
}

// 读取文件
func ReadFile(filePath, file string) (data []byte, err error) {
	fileSrc := filepath.Join(filePath, file)
	if _, ok := CheckFileExist(fileSrc); !ok {
		return nil, FileNoExistErr
	}
	data, err = ioutil.ReadFile(fileSrc)
	return
}

func MkdirIfNecessary(createDir string) (err error) {
	return os.MkdirAll(createDir, os.ModePerm)
}

func GetCurrentPath() string {
	dir, err := os.Getwd() // 当前的目录
	if err != nil {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Println("can not get current path")
		}
	}
	return dir
}

func GetFileBytesByUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
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

func GetFileReaderByUrl(url string) (io.Reader, []byte, error) {
	fileBytes, err := GetFileBytesByUrl(url)
	if err != nil {
		return nil, nil, err
	}

	return bytes.NewReader(fileBytes), fileBytes, nil
}

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
