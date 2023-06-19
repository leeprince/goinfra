package rpahtml

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/leeprince/goinfra/utils/fileutil"
	"log"
	"strings"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/11 22:35
 * @Desc:	特殊的html就需要自定义处理，可以通过"golang.org/x/net/html"完成，具体请看`htmltest/rpahtml`
 */

// 单人单程成功,但是未能兼容单人换乘
func TestMyWorkV1(t *testing.T) {
	filePath := "/Users/leeprince/www/go/goinfra/utils/htmlutil/htmlutiltest/goquerytest/rpahtml/"
	filename := "单人单程-孙孙孙.html"
	// filename := "单人换乘-冯冯.html" // 不能兼容该 html
	fileReader, _, err := fileutil.GetFileReaderByLocalPath(filePath, filename)
	if err != nil {
		log.Fatal(err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(fileReader)
	if err != nil {
		log.Fatal(err)
	}

	var (
		orderInfo string
	)

	// 获取成功：乘客信息
	// 完成乘客信息
	orderInfo, err = doc.Find("tbody tr:nth-child(2)").Html()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("doc.Find(\"tbody tr:nth-child(2)\").Html():", strings.TrimSpace(orderInfo))
	// 证件类型
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(1)").Text()
	log.Println("证件类型:", strings.TrimSpace(orderInfo))
	// 姓名
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(2)").Text()
	log.Println("姓名:", strings.TrimSpace(orderInfo))
	// 身份证号
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(3)").Text()
	log.Println("身份证号:", strings.TrimSpace(orderInfo))
	// 票种：成人票、小孩票
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(4)").Text()
	log.Println("票种：成人票、小孩票:", strings.TrimSpace(orderInfo))
	// 座位类型
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(5) input[name='seatType']").AttrOr("value", "")
	log.Println("座位类型:", strings.TrimSpace(orderInfo))
	// 车厢
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(6) input[name='coachNo']").AttrOr("value", "")
	log.Println("车厢:", strings.TrimSpace(orderInfo))
	// 座位号
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(7) input[name='seatNo']").AttrOr("value", "")
	log.Println("座位号:", strings.TrimSpace(orderInfo))
	// 单张票（单人一程票）的价格
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(8) input[name='ticketPrice']").AttrOr("value", "")
	log.Println("单张票（单人一程票）的价格:", strings.TrimSpace(orderInfo))

	// --- 订单最后一项
	// 订单手机号
	orderInfo, err = doc.Find("tbody .redMobile").Html()
	log.Println("订单手机号:", strings.TrimSpace(orderInfo))
	// 获取成功：乘客信息-end
}
