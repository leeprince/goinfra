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

	// 生产环境：务必在程序初始化时一起完成一次性初始化（加载各种Golang 的 C API 库）！无需每次请求都初始化，否则性能收到严重的影响
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

	// 生产环境：务必在程序初始化时一起完成一次性初始化（加载各种Golang 的 C API 库）！无需每次请求都初始化，否则性能收到严重的影响
	vips.Startup(nil)
	defer vips.Shutdown()

	//image1, err := vips.NewImageFromFile("0001.pdf")

	// pdf和图片组成的pdf
	//image1, err := vips.NewImageFromFile("0001-more-page.pdf")
	//image1, err := vips.NewImageFromFile("0001-more-page-01.pdf")

	// 多页pdf
	//image1, err := vips.NewImageFromFile("more.pdf")

	// 不清晰：金额不清晰，字体不清晰
	//image1, err := vips.NewImageFromFile("prod.pdf")

	/*
		ImportParams 结构体中的各个参数的作用如下：
		    AutoRotate：这是一个布尔参数，如果设置为 true，那么在加载图像时，govips 将会自动根据图像的 EXIF 数据进行旋转，以确保图像的方向正确。
		    FailOnError：这是一个布尔参数，如果设置为 true，那么在加载图像时，如果遇到任何错误，govips 将会立即停止加载并返回错误。如果设置为 false，那么在遇到错误时，govips 将会尽可能地加载图像。
		    Page：这是一个整数参数，用于指定要加载的 PDF 或多页 TIFF 文件的页码。页码从 0 开始计数。
		    NumPages：这是一个整数参数，用于指定要加载的 PDF 或多页 TIFF 文件的页数。
		    Density：这是一个整数参数，用于指定加载 PDF 文件时的 DPI（每英寸点数）。DPI 越高，转换后的图像质量越好。
		    JpegShrinkFactor这是一个整数参数，用于指定加载 JPEG 文件时的缩小因子。例如，如果设置为 2，那么加载的 JPEG 图像的大小将会是原始大小的一半。
		    HeifThumbnail：这是一个布尔参数，如果设置为 true，那么在加载 HEIF 文件时，govips 将会尝试加载缩略图，而不是完整的图像。
		    SvgUnlimited：这是一个布尔参数，如果设置为 true，那么在加载 SVG 文件时，govips 将不会限制 SVG 的大小。如果设置为 false，那么 SVG 的大小将会被限制在 10,000 x 10,000 像素以内。
		以上就是 ImportParams 结构体中各个参数的作用。希望这个解释对你有所帮助。
	*/
	importParams := vips.NewImportParams()
	// 解决：金额不清晰，字体不清晰的问题
	importParams.Density.Set(100)
	image1, err := vips.LoadImageFromFile("prod.pdf", importParams)
	checkError(err)

	// 图片放大：尝试解决金额不清晰，字体不清晰的问题，但是失败了。最终的解决办法是`importParams.Density.Set(100)`
	err = image1.Resize(1.5, vips.KernelLanczos3)
	checkError(err)

	// 锐化图片：锐化图片是一种图像处理技术，主要用于增强图像的细节，使图像看起来更清晰。它的主要作用是提高图像的对比度，尤其是在图像的边缘和颜色变化的地方。
	/*Sigma：这个参数控制锐化的强度。值越大，锐化效果越强。通常，这个值需要根据图像的具体情况和需求进行调整。
	X1：这个参数控制锐化的阈值。只有当像素的亮度差异超过这个阈值时，才会进行锐化。这可以防止对图像的平滑区域进行过度的锐化。
	M2：这个参数控制锐化的最大亮度增加量。这可以防止对图像的亮部进行过度的锐化。*/
	/*err = image1.Sharpen(5.0, 1.0, 1.0)
	checkError(err)*/

	// 完整示例
	// 简单快速生成
	/*// Rotate the picture upright and reset EXIF orientation tag
	image1bytes, _, err := image1.ExportNative()
	checkError(err)*/

	ep := vips.NewJpegExportParams()
	ep.StripMetadata = true
	ep.Quality = 85
	ep.Interlace = true
	ep.OptimizeCoding = true
	ep.SubsampleMode = vips.VipsForeignSubsampleAuto
	ep.TrellisQuant = true
	ep.OvershootDeringing = true
	ep.OptimizeScans = true
	ep.QuantTable = 5
	image1bytes, _, err := image1.ExportJpeg(ep)
	dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.jpg", time.Now().UnixMicro())

	/*ep := vips.NewPngExportParams()
	ep.StripMetadata = true
	ep.Compression = 0
	ep.Filter = vips.PngFilterAll
	ep.Quality = 95
	ep.Interlace = false
	ep.Palette = false
	image1bytes, _, err := image1.ExportPng(ep)
	dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.png", time.Now().UnixMicro())*/

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
