package webpagedomin

import (
	"bytes"
	"errors"
	"fmt"
	"getwebpage-tomarkdown/internel/Infrastructure/config"
	"getwebpage-tomarkdown/internel/pkg/imagewaterhander"
	"github.com/PuerkitoBio/goquery"
	"image"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 12:09
 * @Desc:
 */

func (r *WebPageService) ConvertImg(markdownContent *strings.Builder, s *goquery.Selection, title string) (err error) {
	imageURL, exist := s.Attr("src")
	if !exist {
		fmt.Println("ConvertImg err:", err)
		err = errors.New("ConvertImg err: src is null")
		return
	}
	
	fmt.Println("src imageURL:", imageURL)
	if !strings.HasPrefix(imageURL, "http") {
		imageURL = fmt.Sprintf("%s%s", config.C.FTP.TargetImgHost, imageURL)
	}
	fmt.Println("full src imageURL:", imageURL)
	
	// 发送GET请求
	resp, err := http.Get(imageURL)
	if err != nil {
		fmt.Println("UploadImage Error fetching image:", err)
		return
	}
	
	defer resp.Body.Close()
	
	// 检查HTTP状态码是否为200 OK
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("UploadImage Failed to fetch image with status code: %d\n", resp.StatusCode)
		return
	}
	
	// 读取响应体
	imgDataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("UploadImage Error reading image data:", err)
		return
	}
	
	// 上传文件
	urlParse, err := url.Parse(imageURL)
	if err != nil {
		fmt.Println("UploadImage Parse err:", err)
		return
	}
	urlPathList := strings.Split(urlParse.Path, "/")
	fileName := urlPathList[len(urlPathList)-1]
	saveDir := filepath.Join(config.C.SaveImagePathPrefix, title)
	filePath := filepath.Join(saveDir, fileName)
	fmt.Println("UploadImage saveDir:", saveDir, "filePath:", filePath)
	
	// 判断是否需要保存到本地，如果保存则先将图片保存; 只有保存本地才会进行处理水印
	if config.C.SaveLocal.IsSave {
		saveLocalDir := filepath.Join(config.C.SaveLocal.SaveDir, saveDir)
		fmt.Println("UploadImage SaveLocal  saveLocalDir:", saveLocalDir)
		
		// 解码图片
		fileReader := bytes.NewReader(imgDataBytes)
		img, format, decodeerr := image.Decode(fileReader)
		if decodeerr != nil {
			err = decodeerr
			fmt.Println("Decode err:", err)
			return
		}
		
		err = r.imageWaterHander.ImageWatermarkProcess(img, imagewaterhander.ImageFormat(format), saveLocalDir, fileName)
		if err != nil {
			fmt.Println("UploadImage ImageWatermarkProcess err:", err)
			return
		}
		
		localFile := filepath.Join(saveLocalDir, fileName)
		fmt.Println("UploadImage SaveLocal  localFile:", localFile)
		imgDataBytes, err = os.ReadFile(localFile)
		if err != nil {
			fmt.Println("Open localFile err:", err)
			return
		}
	}
	
	// 重新保存到其他服务器远程服务
	accessImgUrl, err := r.ftpClient.UploadImage(imgDataBytes, saveDir, filePath)
	if err != nil {
		fmt.Println("UploadImage err:", err)
		return
	}
	markdownContent.WriteString(fmt.Sprintf("<p><img src='%s'></p>\n\n", accessImgUrl))
	
	return
}
