// Port of http://members.shaw.ca/el.supremo/MagickWand/resize.htm to Go
package main

import (
	"encoding/base64"
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
	// OfficeDockerMain()
	
	// 自定义pdf转图片
	// 功能验证
	// Function()
	
	// 性能测试
	PerformanceTest()
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
	// maxI := 1
	// 1、2、4、8、16
	maxI := 16
	wg := sync.WaitGroup{}
	
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
	// fileBytes, err := ReadFileBytesByUrl("https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-document-import-ctl/test/0001.pdf")
	// fileBytes, err := ReadFileBytesByUrl("https://kpserverprod-1251506165.cos.ap-shanghai.myqcloud.com/invoice/jammy/dependency/client_ofd.pdf")
	fileBytes, err := ReadFile(".", "0001.pdf")
	// fileBytes, err := ReadFile(".", "0001-more-page.pdf")
	// fileBytes, err := ReadFile(".", "0001-more-page-01.pdf")
	if err != nil {
		fmt.Println("ReadFileBytesByUrl err:", err)
		return
	}
	toImageType := "jpg"
	
	dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.jpg", time.Now().Unix())
	filePath := filepath.Join(dirPath, fileName)
	
	// --- imagick v3 --------------------------------
	// 使用的时候都放到外面去初始化了
	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()
	
	// sourceImagePath := getSourceImageForCover(filepath.Dir(pathNoExtension))
	mw := imagick.NewMagickWand()
	// defer clearImagickWand(mw)
	
	// mw.SetResolution(192, 192)
	// mw.SetResolution(350, 350)
	// mw.SetImageResolution(350, 350)
	// mw.SetImageCompressionQuality(100)
	
	err = mw.ReadImageBlob(fileBytes)
	if err != nil {
		fmt.Println("[CreateImage] ReadImageBlob err:", err)
		return
	}
	
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
	
	// length := mw.GetImageIterations()
	// fmt.Println("length", length)
	// fmt.Println("width", mw.GetImageWidth())
	// fmt.Println("height", mw.GetImageHeight())
	
	// pix := imagick.NewPixelWand()
	// pix.SetColor("white")
	// mw.SetBackgroundColor(pix)
	
	// mw.GetImageDepth()
	// mw.SetImageDepth(16)
	
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
	
	// 转化后的图片字节流
	imageByte := mw.GetImageBlob()
	
	// 可选：是否在本地保存为图片
	err = mw.WriteImage(filePath)
	if err != nil {
		fmt.Println("[CreateImage] WriteImage failed:", err)
		return
	}
	
	// --- imagick v3 --------------------------------
	
	fmt.Println("successful, filepath:", filePath)
	
	imageBase64 := base64.StdEncoding.EncodeToString(imageByte)
	fmt.Println("imageBase64:", imageBase64)
	
	fmt.Println("successful, filepath:", filePath)
	
}

func CustomerPdfToImagesByImagickV3OfPerformanceTest() {
	// fileBytes, err := ReadFileBytesByUrl("https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-document-import-ctl/test/0001.pdf")
	// fileBytes, err := ReadFileBytesByUrl("https://kpserverprod-1251506165.cos.ap-shanghai.myqcloud.com/invoice/jammy/dependency/client_ofd.pdf")
	fileBytes, err := ReadFile(".", "0001.pdf")
	// fileBytes, err := ReadFile(".", "0001-more-page.pdf")
	// fileBytes, err := ReadFile(".", "0001-more-page-01.pdf")
	if err != nil {
		fmt.Println("ReadFileBytesByUrl err:", err)
		return
	}
	toImageType := "jpg"
	
	// 压力测试无需保存图片
	/*dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.jpg", time.Now().Unix())
	filePath := filepath.Join(dirPath, fileName)*/
	
	// --- imagick v3 --------------------------------
	// 使用的时候都放到外面去初始化了
	/*imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()*/
	
	// sourceImagePath := getSourceImageForCover(filepath.Dir(pathNoExtension))
	mw := imagick.NewMagickWand()
	// defer clearImagickWand(mw)
	
	// mw.SetResolution(192, 192)
	// mw.SetResolution(350, 350)
	// mw.SetImageResolution(350, 350)
	// mw.SetImageCompressionQuality(100)
	
	err = mw.ReadImageBlob(fileBytes)
	if err != nil {
		fmt.Println("[CreateImage] ReadImageBlob err:", err)
		return
	}
	
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
	
	// length := mw.GetImageIterations()
	// fmt.Println("length", length)
	// fmt.Println("width", mw.GetImageWidth())
	// fmt.Println("height", mw.GetImageHeight())
	
	// pix := imagick.NewPixelWand()
	// pix.SetColor("white")
	// mw.SetBackgroundColor(pix)
	
	// mw.GetImageDepth()
	// mw.SetImageDepth(16)
	
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
	
	// 转化后的图片字节流
	// 压力测试，不处理返回的图片字节流
	// imageByte := mw.GetImageBlob()
	mw.GetImageBlob()
	
	// 可选：是否在本地保存为图片
	// 压力测试，暂时先不用保存图片到本地
	/*err = mw.WriteImage(filePath)
	if err != nil {
		fmt.Println("[CreateImage] WriteImage failed:", err)
		return
	}*/
	
	// --- imagick v3 --------------------------------
	/*fmt.Println("successful, filepath:", filePath)
	
	imageBase64 := base64.StdEncoding.EncodeToString(imageByte)
	fmt.Println("imageBase64:", imageBase64)
	
	fmt.Println("successful, filepath:", filePath)*/
}

func clearImagickWand(mw *imagick.MagickWand) {
	// 并不会删除在 mw.WriteImage(coverFilePath) 真实创建的图片，
	// 只是在 mw 的实例中移除图片的信息
	mw.RemoveImage()
	
	mw.Clear()
	mw.Destroy()
	// runtime.SetFinalizer(mw, nil)
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
	// err = mw.ReadImage("logo:")
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
