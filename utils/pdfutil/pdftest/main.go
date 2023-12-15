package main

import (
	"fmt"
	errors2 "github.com/pkg/errors"
	"gopkg.in/gographics/imagick.v2/imagick"
	"io"
	"net/http"
	"path/filepath"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/15 14:12
 * @Desc:
 */

func main() {
	//pdfBytes, err := ReadFileBytesByUrl("https://kpserverdev-1251506165.cos.ap-shanghai.myqcloud.com/e-document-import-ctl/test/0001.pdf")
	pdfBytes, err := ReadFileBytesByUrl("https://kpserverprod-1251506165.cos.ap-shanghai.myqcloud.com/invoice/jammy/dependency/client_ofd.pdf")
	if err != nil {
		fmt.Println("ReadFileBytesByUrl err:", err)
		return
	}

	dirPath := "."
	fileName := fmt.Sprintf("pdf_to_jpg_%d.jpg", time.Now().Unix())
	filePath := filepath.Join(dirPath, fileName)

	_, ok := CreateImage(pdfBytes, "jpg", filePath)
	if !ok {
		fmt.Println("CreateImage !ok")
		return
	}

	fmt.Println("successful, filepath:", filePath)
}

func CreateImage(data []byte, toImageType string, coverFilePath string) ([]byte, bool) {
	//sourceImagePath := getSourceImageForCover(filepath.Dir(pathNoExtension))
	mw := imagick.NewMagickWand()
	defer clearImagickWand(mw)

	mw.SetResolution(192, 192)
	//mw.SetResolution(350, 350)
	//mw.SetImageResolution(350, 350)
	//mw.SetImageCompressionQuality(100)
	err := mw.ReadImageBlob(data)
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

	err = mw.WriteImage(coverFilePath)
	if err != nil {
		fmt.Println("[CreateImage] WriteImage failed:", err)
		return nil, false
	}
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
