package webpagedomin

import (
	"bytes"
	"errors"
	"fmt"
	"getwebpage-tomarkdown/internel/Infrastructure/config"
	"github.com/PuerkitoBio/goquery"
	"github.com/leeprince/goinfra/utils/fileutil"
	"io"
	"net/http"
	"net/url"
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
	title = strings.TrimSpace(title)
	remoteFileDir := fmt.Sprintf("/%s/%s", config.C.SaveImagePathPrefix, title)
	remoteFilePath := fmt.Sprintf("%s/%s", remoteFileDir, fileName)
	fmt.Println("UploadImage remoteFileDir:", remoteFileDir, "remoteFilePath:", remoteFilePath)
	
	// 重新保存到其他服务器远程服务
	accessImgUrl, err := r.ftpClient.UploadImage(imgDataBytes, remoteFileDir, remoteFilePath)
	if err != nil {
		fmt.Println("UploadImage err:", err)
		return
	}
	markdownContent.WriteString(fmt.Sprintf("<p><img src='%s'></p>\n\n", accessImgUrl))
	
	// 判断是否需要保存到本地，如果保存则先将图片保存
	if config.C.SaveLocal.IsSave {
		saveDir := fmt.Sprintf("%s/%s/", config.C.SaveLocal.SaveDir, config.C.SaveImagePathPrefix)
		fmt.Println("ConvertImg SaveLocalFileByIoReader:", fileName, "-saveDir:", saveDir)
		
		_, err = fileutil.SaveLocalFileByIoReader(bytes.NewReader(imgDataBytes), fileName, saveDir)
		if err != nil {
			fmt.Println("UploadImage SaveLocalFileByIoReader err:", err)
			return
		}
	}
	
	return
}
