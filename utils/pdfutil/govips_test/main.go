package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/davidbyttow/govips/v2/vips"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/19 11:17
 * @Desc:
 */

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func main() {
	// 改造 `https://github.com/davidbyttow/govips` 中的 Example usage
	//OfficeDockerMain()

	// 自定义pdf转图片
	// 功能验证
	Function()

	// 性能测试
	//PerformanceTest()
}

// 功能验证
func Function() {
	startTime := time.Now().UnixMilli()
	CustomerPdfToImagesByGovipsOfFunction()
	fmt.Printf("cost mill time: %dms\n", time.Now().UnixMilli()-startTime)
}

// 性能测试
func PerformanceTest() {
	var i int
	// 1、2、4、8、16
	maxI := 16
	wg := sync.WaitGroup{}

	vips.LoggingSettings(myLogger, vips.LogLevelError)
	vips.Startup(&vips.Config{
		ConcurrencyLevel: maxI,
		MaxCacheFiles:    0,
		MaxCacheMem:      0,
		MaxCacheSize:     0,
		ReportLeaks:      false,
		CacheTrace:       false,
		CollectStats:     false,
	})
	defer vips.Shutdown()

	for {
		i++
		if i > maxI {
			break
		}

		wg.Add(1)

		go func() {
			defer wg.Done()

			startTime := time.Now().UnixMilli()
			CustomerPdfToImagesByGovipsOfPerformanceTest()
			fmt.Printf("cost mill time: %dms\n", time.Now().UnixMilli()-startTime)
		}()
	}
	fmt.Println("...wg.wait...")
	wg.Wait()
}

func myLogger(messageDomain string, verbosity vips.LogLevel, message string) {
	var messageLevelDescription string
	switch verbosity {
	case vips.LogLevelError:
		messageLevelDescription = "error"
	case vips.LogLevelCritical:
		messageLevelDescription = "critical"
	case vips.LogLevelWarning:
		messageLevelDescription = "warning"
	case vips.LogLevelMessage:
		messageLevelDescription = "message"
	case vips.LogLevelInfo:
		messageLevelDescription = "info"
	case vips.LogLevelDebug:
		messageLevelDescription = "debug"
	}

	log.Printf("[%v.%v] %v", messageDomain, messageLevelDescription, message)
}

func CustomerPdfToImagesByGovipsOfPerformanceTest() {
	// 压力测试的时候，都放到外面去初始化了
	/*
		vips.LoggingSettings(myLogger, vips.LogLevelError)
		vips.Startup(nil)
		defer vips.Shutdown()
	*/

	image1, err := vips.NewImageFromFile("0001.pdf")
	//image1, err := vips.NewImageFromFile("0001-more-page.pdf")
	//image1, err := vips.NewImageFromFile("0001-more-page-01.pdf")
	checkError(err)

	// 完整示例
	// Rotate the picture upright and reset EXIF orientation tag
	_, _, err = image1.ExportNative()
	checkError(err)

	// 压力测试，暂时先不用保存图片到本地
	/*dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.jpg", time.Now().UnixMicro())
	filePath := filepath.Join(dirPath, fileName)
	ok, err := WriteFile(dirPath, filePath, image1bytes, false)
	if !ok {
		fmt.Println("fileutil.WriteFile !ok")
		return
	}
	checkError(err)
	fmt.Println("successful, filepath:", filePath)*/
}

func CustomerPdfToImagesByGovipsOfFunction() {
	vips.LoggingSettings(myLogger, vips.LogLevelError)
	vips.Startup(nil)
	defer vips.Shutdown()

	//image1, err := vips.NewImageFromFile("0001.pdf")

	// pdf和图片组成的pdf
	//image1, err := vips.NewImageFromFile("0001-more-page.pdf")
	//image1, err := vips.NewImageFromFile("0001-more-page-01.pdf")

	// 多页pdf
	image1, err := vips.NewImageFromFile("more.pdf")
	checkError(err)

	// 完整示例
	// Rotate the picture upright and reset EXIF orientation tag
	image1bytes, _, err := image1.ExportNative()
	checkError(err)

	dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.jpg", time.Now().UnixMicro())
	filePath := filepath.Join(dirPath, fileName)
	ok, err := WriteFile(dirPath, filePath, image1bytes, false)
	if !ok {
		fmt.Println("fileutil.WriteFile !ok")
		return
	}
	checkError(err)
	fmt.Println("successful, filepath:", filePath)
}

func OfficeDockerMain() {
	vips.Startup(nil)
	defer vips.Shutdown()

	image1, err := vips.NewImageFromFile("0001.png")
	checkError(err)

	// Rotate the picture upright and reset EXIF orientation tag
	err = image1.AutoRotate()
	checkError(err)

	image1bytes, _, err := image1.ExportNative()
	checkError(err)

	dirPath := "."
	fileName := fmt.Sprintf("autorotate_%d.jpg", time.Now().Unix())
	filePath := filepath.Join(dirPath, fileName)

	ok, err := WriteFile(dirPath, filePath, image1bytes, false)
	if !ok {
		fmt.Println("fileutil.WriteFile !ok")
		return
	}
	checkError(err)
	fmt.Println("successful, filepath:", filePath)
}

// 写入数据到文件
func WriteFile(dirPath, filename string, data []byte, isAppend bool) (ok bool, err error) {
	filePath := filepath.Join(dirPath, filename)
	if _, ok = CheckFileDirExist(filePath); !ok {
		// 创建目录
		err = os.MkdirAll(dirPath, os.ModePerm)
		if ok = os.IsNotExist(err); ok {
			err = errors.New("创建文件目录错误")
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

// 检查文件/目录是否存在
func CheckFileDirExist(filePath string) (os.FileInfo, bool) {
	finfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, false
	}
	return finfo, true
}
