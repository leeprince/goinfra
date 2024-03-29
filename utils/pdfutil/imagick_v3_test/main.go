// Port of http://members.shaw.ca/el.supremo/MagickWand/resize.htm to Go
package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	// `https://github.com/gographics/imagick/tree/v3.5.0/examples/docker` 的示例 main.go
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
	CustomerPdfToImagesByImagickV3OfFunction()
	fmt.Printf("cost mill time: %dms\n", time.Now().UnixMilli()-startTime)
}

// 性能测试
func PerformanceTest() {
	var i int
	//maxI := 1
	// 1、2、4、8、16
	maxI := 8
	wg := sync.WaitGroup{}

	// 生产环境：务必在程序初始化时一起完成一次性初始化（加载各种Golang 的 C API 库）！无需每次请求都初始化，否则性能收到严重的影响
	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()

	for {
		i++
		if i > maxI {
			break
		}

		wg.Add(1)

		go func() {
			defer wg.Done()

			startTime := time.Now().UnixMilli()
			CustomerPdfToImagesByImagickV3OfPerformanceTest()
			fmt.Printf("cost mill time: %dms\n", time.Now().UnixMilli()-startTime)
		}()
	}
	fmt.Println("...wg.wait...")
	wg.Wait()
}

func CustomerPdfToImagesByImagickV3OfFunction() {
	//fileBytes, err := ReadFileBytesByUrl("https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-document-import-ctl/test/0001.pdf")
	//fileBytes, err := ReadFileBytesByUrl("https://kpserverprod-1251506165.cos.ap-shanghai.myqcloud.com/invoice/jammy/dependency/client_ofd.pdf")
	//fileBytes, err := ReadFileBytesByUrl("https://upload.fapiaoer.cn/e-document-import-ctl/prod/8/2023-12-27/57fc7886181b60bc_032002200611_00051075_5211b88a24f23d4029d5b34e7a2e2d53.pdf")
	//fileBytes, err := ReadFileBytesByUrl("https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/wbx/upload/FUDkFLyvTdIm4756248423cba0c700f297f15b68189d_20201844_1703840311697903.pdf") //
	//fileBytes, err := ReadFile(".", "0001.pdf")

	// pdf和图片组成的pdf
	//fileBytes, err := ReadFile(".", "0001-more-page.pdf")
	//fileBytes, err := ReadFile(".", "0001-more-page-01.pdf")

	// 读取出是乱码
	// 解决：apt-get -q -y install fonts-arphic-uming fonts-arphic-ukai fonts-noto-cjk  --no-install-recommends
	//fileBytes, err := ReadFile(".", "empty.pdf")
	//fileBytes, err := ReadFile(".", "einvoice.pdf")

	// 不清晰：金额不清晰，小数点偏移
	fileBytes, err := ReadFile(".", "prod.pdf")

	// 多页pdf：默认读取第一页
	//fileBytes, err := ReadFileBytesByUrl("https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/wbx/upload/OQgorPo0MckWeb374dae2b7e56f9ccb9a6ea1ad0d276_20201844_1703557617396552.pdf")
	//fileBytes, err := ReadFile(".", "more.pdf")
	if err != nil {
		fmt.Println("ReadFileBytesByUrl err:", err)
		return
	}
	toImageType := "jpg"

	dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.jpg", time.Now().UnixMicro())
	filePath := filepath.Join(dirPath, fileName)

	// --- imagick v3 --------------------------------
	// 生产环境：务必在程序初始化时一起完成一次性初始化（加载各种Golang 的 C API 库）！无需每次请求都初始化，否则性能收到严重的影响
	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// 设置分辨率（必须在 ReadImageBlob() 前）
	err = mw.SetResolution(100, 100)
	if err != nil {
		fmt.Println("SetResolution err:", err)
		return
	}

	err = mw.ReadImageBlob(fileBytes)
	if err != nil {
		fmt.Println("ReadImageBlob err:", err)
		return
	}

	// 对图片进行锐化处理
	/*err = mw.AdaptiveSharpenImage(0, 1)
	if err != nil {
		fmt.Println("AdaptiveSharpenImage err:", err)
		return
	}*/

	if toImageType == "jpg" {
		toImageType = "jpeg"
	}

	// 想着放大图片解决模糊
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()
	// Calculate half the size
	hWidth := uint(width * 2)
	hHeight := uint(height * 2)
	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS)
	if err != nil {
		fmt.Println("ResizeImage failed：", err)
		return
	}

	// 设置图像分辨率：导入的文件应该是图片类型时才生效。即pdf转图片的场景下：该参数无效
	/*err = mw.SetImageResolution(100, 100)
	if err != nil {
		fmt.Println("SetImageResolution err:", err)
		return
	}*/

	// 设置图像压缩质量
	err = mw.SetImageCompressionQuality(95)
	if err != nil {
		fmt.Println("SetImageCompressionQuality err:", err)
		return
	}

	/*iterations := mw.GetImageIterations()
	fmt.Println("GetImageIterations", iterations)*/

	/*length, err := mw.GetImageLength()
	fmt.Println("GetImageLength", length)
	if err != nil {
		fmt.Println("GetImageLength err:", err)
		return
	}*/

	//pix := imagick.NewPixelWand()
	//pix.SetColor("white")
	//mw.SetBackgroundColor(pix)

	// 图片的深度
	/*imageDepth := mw.GetImageDepth()
	fmt.Println("imageDepth:", imageDepth)
	err = mw.SetImageDepth(1)
	if err != nil {
		fmt.Println("SetImageDepth err:", err)
		return
	}*/

	/*w, h, x, y, err := mw.GetImagePage()
	fmt.Println("GetImagePage:", w, h, x, y, err)
	//mw.SetImagePage()

	w, h, x, y, err = mw.GetPage()
	//mw.SetPage()
	fmt.Println("GetPage:", w, h, x, y, err)*/

	// 所在的迭代数，默认是最后一页。索引所0开始
	//iteratorIndex := mw.GetIteratorIndex()
	//fmt.Println("GetIteratorIndex", iteratorIndex)
	// 设置为第一页
	mw.SetFirstIterator()
	//ok := mw.SetIteratorIndex(2)
	//if !ok {
	//	fmt.Println("SetIteratorIndex failed:", ok)
	//}

	// 激活、停用、重置或设置alpha通道。
	err = mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_REMOVE)
	if err != nil {
		fmt.Println("SetImageAlphaChannel failed:", err)
		return
	}

	// 设置需要转化为的图片格式
	err = mw.SetImageFormat(toImageType)
	if err != nil {
		fmt.Println("SetImageFormat failed:", err)
		return
	}

	// 可选：获取转化后的图片字节流
	//imageByte := mw.GetImageBlob()

	// 可选：是否在本地保存为图片
	err = mw.WriteImage(filePath)
	if err != nil {
		fmt.Println("WriteImage failed:", err)
		return
	}

	// --- imagick v3 --------------------------------
	//imageBase64 := base64.StdEncoding.EncodeToString(imageByte)
	//fmt.Println("imageBase64:", imageBase64)

	fmt.Println("successful, filepath:", filePath)

}

func CustomerPdfToImagesByImagickV3OfPerformanceTest() {
	//fileBytes, err := ReadFileBytesByUrl("https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-document-import-ctl/test/0001.pdf")
	//fileBytes, err := ReadFileBytesByUrl("https://kpserverprod-1251506165.cos.ap-shanghai.myqcloud.com/invoice/jammy/dependency/client_ofd.pdf")
	//fileBytes, err := ReadFile(".", "0001.pdf")

	// pdf和图片组成的pdf
	//fileBytes, err := ReadFile(".", "0001-more-page.pdf")
	//fileBytes, err := ReadFile(".", "0001-more-page-01.pdf")

	// 多页pdf
	fileBytes, err := ReadFileBytesByUrl("https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/wbx/upload/OQgorPo0MckWeb374dae2b7e56f9ccb9a6ea1ad0d276_20201844_1703557617396552.pdf")
	//fileBytes, err := ReadFile(".", "more.pdf")
	if err != nil {
		fmt.Println("ReadFileBytesByUrl err:", err)
		return
	}
	toImageType := "jpg"

	dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.jpg", time.Now().UnixMicro())
	filePath := filepath.Join(dirPath, fileName)

	// --- imagick v3 --------------------------------
	// 使用的时候都放到外面去初始化了
	/*imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()*/

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	// 设置分辨率
	err = mw.SetResolution(100, 100)
	if err != nil {
		fmt.Println("SetResolution err:", err)
		return
	}

	err = mw.ReadImageBlob(fileBytes)
	if err != nil {
		fmt.Println("ReadImageBlob err:", err)
		return
	}

	// 设置图像分辨率：导入的文件应该是图片类型时才生效。即pdf转图片的场景下：该参数无效
	/*err = mw.SetImageResolution(100, 100)
	if err != nil {
		fmt.Println("SetImageResolution err:", err)
		return
	}*/

	// 设置图像压缩质量。
	//mw.SetImageCompressionQuality(100)

	if toImageType == "jpg" {
		toImageType = "jpeg"
	}

	/*width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	// Calculate half the size
	hWidth := uint(width * 2)
	hHeight := uint(height * 2)

	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS)
	if err != nil {
		fmt.Println("ResizeImage failed：", err)
		return
	}*/

	/*iterations := mw.GetImageIterations()
	fmt.Println("GetImageIterations", iterations)*/

	/*length, err := mw.GetImageLength()
	fmt.Println("GetImageLength", length)
	if err != nil {
		fmt.Println("GetImageLength err:", err)
		return
	}*/

	//pix := imagick.NewPixelWand()
	//pix.SetColor("white")
	//mw.SetBackgroundColor(pix)

	// 图片的深度
	/*imageDepth := mw.GetImageDepth()
	fmt.Println("imageDepth:", imageDepth)
	err = mw.SetImageDepth(1)
	if err != nil {
		fmt.Println("SetImageDepth err:", err)
		return
	}*/

	/*w, h, x, y, err := mw.GetImagePage()
	fmt.Println("GetImagePage:", w, h, x, y, err)
	//mw.SetImagePage()

	w, h, x, y, err = mw.GetPage()
	//mw.SetPage()
	fmt.Println("GetPage:", w, h, x, y, err)*/

	// 所在的迭代数，默认是最后一页。索引所0开始
	iteratorIndex := mw.GetIteratorIndex()
	fmt.Println("GetIteratorIndex", iteratorIndex)
	// 设置为第一页
	//mw.SetFirstIterator()
	ok := mw.SetIteratorIndex(2)
	if !ok {
		fmt.Println("SetIteratorIndex failed:", ok)
	}

	// 激活、停用、重置或设置alpha通道。
	err = mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_REMOVE)
	if err != nil {
		fmt.Println("SetImageAlphaChannel failed:", err)
		return
	}

	// 设置需要转化为的图片格式
	err = mw.SetImageFormat(toImageType)
	if err != nil {
		fmt.Println("SetImageFormat failed:", err)
		return
	}

	// 可选：获取转化后的图片字节流
	//imageByte := mw.GetImageBlob()

	// 压测不需要保存
	/*// 可选：是否在本地保存为图片
	err = mw.WriteImage(filePath)
	if err != nil {
		fmt.Println("WriteImage failed:", err)
		return
	}*/

	// --- imagick v3 --------------------------------
	//imageBase64 := base64.StdEncoding.EncodeToString(imageByte)
	//fmt.Println("imageBase64:", imageBase64)

	fmt.Println("successful, filepath:", filePath)
}

func clearImagickWand(mw *imagick.MagickWand) {
	// 并不会删除在 mw.WriteImage(coverFilePath) 真实创建的图片，
	// 只是在 mw 的实例中移除图片的信息
	mw.RemoveImage()

	mw.Clear()
	mw.Destroy()
	//runtime.SetFinalizer(mw, nil)
	mw = nil
}

func ReadFileBytesByUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = errors.Errorf("resp.StatusCode != http.StatusOK")
		return nil, err
	}

	contentLength := resp.Header.Get("Content-Length")
	if contentLength == "0" {
		err = errors.Errorf("contentLength == 0")
		return nil, err
	}

	fileBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "ioutil.ReadAll")
		return nil, err
	}
	if fileLen := len(fileBytes); fileLen <= 0 {
		err = errors.Errorf("fileBytes len 0")
		return nil, err
	}

	return fileBytes, nil
}

func OfficeDockerMain() {
	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()
	var err error

	mw := imagick.NewMagickWand()

	// 官方文件
	//err = mw.ReadImage("logo:")
	// 自定义文件
	localFile := "./0001.png"
	err = mw.ReadImage(localFile)
	if err != nil {
		panic(err)
	}

	// Get original logo size
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	// Calculate half the size
	hWidth := uint(width / 2)
	hHeight := uint(height / 2)

	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS)
	if err != nil {
		panic(err)
	}

	// Set the compression quality to 95 (high quality = low compression)
	err = mw.SetImageCompressionQuality(95)
	if err != nil {
		panic(err)
	}

	dirPath := "."
	fileName := fmt.Sprintf("resize_%d.jpg", time.Now().Unix())
	filePath := filepath.Join(dirPath, fileName)
	if err = mw.WriteImage(filePath); err != nil {
		panic(err)
	}

	fmt.Printf("Wrote: %s\n", filePath)
}

var (
	FileNoExistErr = errors.New("file not exist")
)

// 读取文件
func ReadFile(filePath, filename string) (data []byte, err error) {
	fileSrc := filepath.Join(filePath, filename)
	if _, ok := CheckFileDirExist(fileSrc); !ok {
		return nil, FileNoExistErr
	}
	data, err = os.ReadFile(fileSrc)
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

func checkError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
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
