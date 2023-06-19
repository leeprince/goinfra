package rpahtml

import (
	"github.com/leeprince/goinfra/utils/fileutil"
	"github.com/leeprince/goinfra/utils/htmlutil"
	"golang.org/x/net/html"
	"log"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/9 09:35
 * @Desc:
 */

// 打印结果
func TestHtmlPrint(t *testing.T) {
	filePath := "/Users/leeprince/www/go/goinfra/utils/htmlutil/htmlutiltest/htmltest/rpahtml/"
	// filename := "多人单程-王王王.html" // 有要求-保持不添加<table>包含所有原始数据
	filename := "单人单程-李李李.html" // 无要求-添加<table>包含所有原始数据
	fileReader, _, err := fileutil.GetFileReaderByLocalPath(filePath, filename)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(fileReader)
	if err != nil {
		log.Fatal(err)
	}

	htmlutil.PrintHtml(doc, 0)
}
