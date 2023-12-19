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
	"time"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	// `https://github.com/gographics/imagick/tree/v3.5.0/examples/docker` 的示例 main.go
	//OfficeDockerMain()

	// 自定义pdf转图片
	startTime := time.Now().UnixMilli()
	CustomerPdfToImagesByImagickV3()
	fmt.Printf("cost mill time: %dms\n", time.Now().UnixMilli()-startTime)
}

func CustomerPdfToImagesByImagickV3() {
	//pdfBytes, err := ReadFileBytesByUrl("https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-document-import-ctl/test/0001.pdf")
	//pdfBytes, err := ReadFileBytesByUrl("https://kpserverprod-1251506165.cos.ap-shanghai.myqcloud.com/invoice/jammy/dependency/client_ofd.pdf")
	//pdfBytes, err := ReadFile(".", "0001.pdf")
	//pdfBytes, err := ReadFile(".", "0001-more-page.pdf")
	pdfBytes, err := ReadFile(".", "0001-more-page-01.pdf")
	if err != nil {
		fmt.Println("ReadFileBytesByUrl err:", err)
		return
	}

	dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.jpg", time.Now().Unix())
	filePath := filepath.Join(dirPath, fileName)

	//imageByte, ok := CreateImage(pdfBytes, "jpg", filePath)
	imageByte, ok := CreateImage(pdfBytes, "jpg")
	if !ok {
		fmt.Println("CreateImage !ok")
		return
	}

	ok, err = WriteFile(dirPath, filePath, imageByte, false)
	if !ok {
		fmt.Println("fileutil.WriteFile !ok")
		return
	}
	checkError(err)
	fmt.Println("successful, filepath:", filePath)
}

func CreateImage(data []byte, toImageType string) ([]byte, bool) {
	//func CreateImage(data []byte, toImageType string, coverFilePath string) ([]byte, bool) {
	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()
	var err error

	//sourceImagePath := getSourceImageForCover(filepath.Dir(pathNoExtension))
	mw := imagick.NewMagickWand()
	//defer clearImagickWand(mw)

	//mw.SetResolution(192, 192)
	//mw.SetResolution(350, 350)
	//mw.SetImageResolution(350, 350)
	//mw.SetImageCompressionQuality(100)
	err = mw.ReadImageBlob(data)
	if err != nil {
		fmt.Println("[CreateImage] ReadImageBlob err:", err)
		return nil, false
	}

	if toImageType == "jpg" {
		toImageType = "jpeg"
	}

	/*
		width := mw.GetImageWidth()
		height := mw.GetImageHeight()

		// Calculate half the size
		hWidth := uint(width * 2)
		hHeight := uint(height * 2)

		// Resize the image using the Lanczos filter
		// The blur factor is a float, where > 1 is blurry, < 1 is sharp
		err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
		if err != nil {
			log.ErrorS(seq, "[CreateImage] ResizeImage failed[%v]", err)
			return nil, false
		}*/

	//length := mw.GetImageIterations()
	//fmt.Println("length", length)
	//fmt.Println("width", mw.GetImageWidth())
	//fmt.Println("height", mw.GetImageHeight())

	//pix := imagick.NewPixelWand()
	//pix.SetColor("white")
	//mw.SetBackgroundColor(pix)
	mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_REMOVE)
	//mw.SetImageFormat("jpeg")
	mw.SetImageFormat(toImageType)

	//mw.GetImageDepth()
	//mw.SetImageDepth(16)

	/*err = mw.WriteImage(coverFilePath)
	if err != nil {
		fmt.Println("[CreateImage] WriteImage failed:", err)
		return nil, false
	}*/

	content := mw.GetImageBlob()

	return content, true
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
