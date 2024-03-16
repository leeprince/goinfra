package application

import (
	"bytes"
	"fmt"
	"getwebpage-tomarkdown/internel/Infrastructure/config"
	"getwebpage-tomarkdown/internel/domain/webpagedomin"
	"getwebpage-tomarkdown/internel/domain/webpagedomin/webpageentity"
	"github.com/leeprince/goinfra/utils/fileutil"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 02:28
 * @Desc:
 */

// SaveWebPage 保存网页内容到本地/远程文件
func (r *Application) SaveWebPage() {
	webPageService := webpagedomin.NewWebPageService()
	for _, url := range config.C.WebPageUrlList {
		// 获取 url 内容，并转为 markdown
		webPageEntity, err := webPageService.FetchMarkdown(url)
		if err != nil {
			fmt.Println("FetchMarkdown err:", err.Error()+" url:", url)
			return
		}
		
		// 保存为本地文件
		err = r.SaveWebPageToLocal(webPageEntity)
		if err != nil {
			fmt.Println("SaveWebPageToLocal err:", err.Error()+" url:", url)
			return
		}
		
		// 保存为wordpress文章
		err = r.SaveWebPageToWrodpress(webPageEntity)
		if err != nil {
			fmt.Println("SaveWebPageToLocal err:", err.Error()+" url:", url)
			return
		}
	}
	
	return
}

func (r *Application) SaveWebPageToLocal(webPageEntity *webpageentity.MarkdownEntity) (err error) {
	if !config.C.SaveLocal.IsSave {
		return
	}
	
	title := strings.TrimSpace(webPageEntity.Title)
	fileName := fmt.Sprintf("%s.md", title)
	saveDir := config.C.SaveLocal.SaveDir
	fmt.Println("SaveWebPageToLocal fileName:", fileName, "-saveDir:", saveDir)
	
	ocntentBytes := []byte(webPageEntity.Content)
	_, err = fileutil.SaveLocalFileByIoReader(bytes.NewReader(ocntentBytes), fileName, saveDir)
	if err != nil {
		fmt.Println("SaveWebPageToLocal SaveLocalFileByIoReader err:", err)
		return
	}
	
	return
}

func (r *Application) SaveWebPageToWrodpress(webPageEntity *webpageentity.MarkdownEntity) (err error) {
	if !config.C.SaveRemoter.IsSave {
		return
	}
	
	return
}
