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
	// filename := "单人单程-孙孙孙.html" // 兼容
	// filename := "单人换乘-冯冯.html" // 兼容
	filename := "多人单程-周周周.html" // 兼容
	// filename := "单人单程-占座票-安安安.html" // 兼容
	// filename := "多人换乘-余余余.html" // 兼容
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
