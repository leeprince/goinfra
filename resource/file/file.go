package file

import (
	"errors"
	"io/ioutil"
	"log"
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
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
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
	if _, err = fs.Write(data); err != nil {
		return
	}
	defer fs.Close()
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
	dir, err := os.Getwd() //当前的目录
	if err != nil {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Println("can not get current path")
		}
	}
	return dir
}
