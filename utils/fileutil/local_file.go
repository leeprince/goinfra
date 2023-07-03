package fileutil

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

/**
 * @Author: prince.lee
 * @Date:   2022/2/14 11:28
 * @Desc:
 */

var (
	FileNoExistErr = errors.New("file not exist")
)

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
	
	flag := os.O_CREATE | os.O_RDWR
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
func ReadFile(filePath, filename string) (data []byte, err error) {
	fileSrc := filepath.Join(filePath, filename)
	if _, ok := CheckFileExist(fileSrc); !ok {
		return nil, FileNoExistErr
	}
	data, err = os.ReadFile(fileSrc)
	return
}

func ReadFileReader(filePath, filename string) (io.Reader, []byte, error) {
	fileBytes, err := ReadFile(filePath, filename)
	if err != nil {
		return nil, nil, err
	}
	
	return bytes.NewReader(fileBytes), fileBytes, nil
}

func MkdirIfNecessary(createDir string) (err error) {
	return os.MkdirAll(createDir, os.ModePerm)
}

// 跟设置的工作目录有关，相同的文件路径不同的工作目录获取的结果不一样。不明确项目所在的工作目录时，推荐使用绝对路径
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

func GetFileReaderByLocalPath(filePath, filename string) (io.Reader, []byte, error) {
	return ReadFileReader(filePath, filename)
}
