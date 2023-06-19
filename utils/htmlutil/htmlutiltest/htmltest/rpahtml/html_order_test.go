package rpahtml

import (
	"github.com/leeprince/goinfra/utils/fileutil"
	"golang.org/x/net/html"
	"log"
	"strings"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/9 09:35
 * @Desc:
 */

// 成功
func TestHtmlOrder(t *testing.T) {
	filePath := "/Users/leeprince/www/go/goinfra/utils/htmlutil/htmlutiltest/goquerytest/rpahtml/"
	filename := "多人单程-王王王.html" // 有要求
	// filename := "单人单程-李李李.html" // 无要求-添加<table>包含所有原始数据
	fileReader, _, err := fileutil.GetFileReaderByLocalPath(filePath, filename)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(fileReader)
	if err != nil {
		log.Fatal(err)
	}

	HtmlOrder(doc)
}

var (
	fullOrderBaseInfo               string
	fullOrderBaseInfoAfter01OrderId string
)

func HtmlOrder(n *html.Node) string {
	// 未能解析的完整订单数据
	if n.Type == html.TextNode && strings.Contains(n.Data, "||||") {
		log.Println("未能解析的完整订单数据：", n.Data)
		fullOrderBaseInfo = n.Data
	}

	// 只有解析过完整订单数据后，才需要会继续解析
	// 订单ID
	if n.Type == html.ElementNode && strings.Contains(n.Data, "span") {
		for _, attr := range n.Attr {
			if attr.Key == "id" && strings.Contains(attr.Val, "SelectSeatInfoTip_") {
				log.Println("未解析 orderId：", attr.Val)
				fullOrderArr := strings.Split(attr.Val, "_")
				if len(fullOrderArr) == 2 {
					fullOrderBaseInfoAfter01OrderId = fullOrderArr[1]
					log.Println("orderId：", fullOrderBaseInfoAfter01OrderId)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		HtmlOrder(c)
		if fullOrderBaseInfo != "" && fullOrderBaseInfoAfter01OrderId != "" {
			return ""
		}
	}

	return ""
}
