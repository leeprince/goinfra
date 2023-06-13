package rpahtml

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/leeprince/goinfra/utils/fileutil"
	"log"
	"strings"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/9 09:35
 * @Desc:	部分成功、部分失败（失败部分是html格式比较特殊，实际上是不规范的html）
 *				特殊的html就需要自定义处理，可以通过"golang.org/x/net/html"完成，具体请看`htmltest/rpahtml`
 */

// 获取失败
func TestGOQueryOrder01(t *testing.T) {
	filePath := "/Users/leeprince/www/go/goinfra/utils/htmlutil/htmlutiltest/goquerytest/rpahtml/"
	filename := "单人单程-孙孙孙.html"
	fileReader, _, err := fileutil.GetFileReaderByLocalPath(filePath, filename)
	if err != nil {
		log.Fatal(err)
	}

	var (
		orderId string
	)
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(fileReader)
	if err != nil {
		log.Fatal(err)
	}

	orderId = doc.Has("td[name='orderNumberTd']").Text()
	log.Println("orderId:", orderId)

	orderId = doc.Find("td[name='orderNumberTd']").Text()
	log.Println("orderId:", orderId)

	orderId = doc.Find("td[style=' word-wrap:break-word; word-break:break-all; width:250px;max-width:250px;  ']").Text()
	log.Println("orderId:", orderId)

	orderId = doc.Find("td[style='word-wrap:break-word; word-break:break-all; width:250px;max-width:250px;']").Text()
	log.Println("orderId:", orderId)

}

// 获取失败
func TestGOQueryOrder02(t *testing.T) {
	html := `0||||<tr data-toggle='collapse'  href='#20514293' aria-expanded='false' aria-controls='collapseOne' style='height:80px;' class=''><td >1</td><td style=' word-wrap:break-word; word-break:break-all; width:250px;max-width:250px;  ' name='orderNumberTd'>FeiZ1905328274246370080</td><td class='maxRedFont'>2023-06-03</td><td class='maxRedFont'>G1072</td> <td class='maxFont'>青岛机场( 13:17)->潍坊北(WJK)</td><td  class='maxRedFont'>二等座</td><td>47.0</td></tr>`
	// html := `<html><body>0||||<tr data-toggle='collapse'  href='#20514293' aria-expanded='false' aria-controls='collapseOne' style='height:80px;' class=''><td >1</td><td style=' word-wrap:break-word; word-break:break-all; width:250px;max-width:250px;  ' name='orderNumberTd'>FeiZ1905328274246370080</td><td class='maxRedFont'>2023-06-03</td><td class='maxRedFont'>G1072</td> <td class='maxFont'>青岛机场( 13:17)->潍坊北(WJK)</td><td  class='maxRedFont'>二等座</td><td>47.0</td></tr></body></html>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	orderNumber := doc.Find("td[name='orderNumberTd']").Text()
	trainNumber := doc.Find("td:nth-child(4)").Text()
	route := doc.Find("td:nth-child(5)").Text()
	seat := doc.Find("td:nth-child(6)").Text()
	price := doc.Find("td:nth-child(7)").Text()

	fmt.Printf("订单号=%s、车次=%s、出发地和目的地=%s、坐席=%s、总价格=%s\n", strings.TrimSpace(orderNumber), strings.TrimSpace(trainNumber), strings.TrimSpace(route), strings.TrimSpace(seat), strings.TrimSpace(price))
}

// 部分成功、部分失败
func TestGOQueryOrder03(t *testing.T) {
	filePath := "/Users/leeprince/www/go/goinfra/utils/htmlutil/htmlutiltest/goquerytest/rpahtml/"
	filename := "单人单程-孙孙孙.html"
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

	// 获取成功：获取所有内容
	orderInfo, err = doc.Find("table").Parent().Html()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("doc.Find(\"table\").Parent().Html():", orderInfo)

	// 获取成功：body所有内容
	orderInfo, err = doc.Find("body").Html()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("doc.Find(\"body\").Html():", orderInfo)

	orderInfo = doc.Find("body").Text()
	log.Println("doc.Find(\"body\").Text():", orderInfo)

	orderInfo = doc.Find("body").Contents().Text()
	log.Println("doc.Find(\"body\").Contents().Text():", orderInfo)
	// 获取成功：body所有内容-end

	// 获取成功：大概信息
	orderInfo, err = doc.Find("table").Html()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("doc.Find(\"table\").Html():", orderInfo)

	orderInfo = doc.Find("table").Text()
	log.Println("doc.Find(\"table\").Text():", orderInfo)
	// 获取成功：大概信息 - end

	// 获取失败: 获取订单（车次、订单ID）失败
	orderInfo = doc.Find("table").PrevAll().Text()
	log.Println("doc.Find(\"table\").PrevAll().Text():", orderInfo)

	// 获取失败: 获取订单（车次、订单ID）失败
	orderInfo, err = doc.Find("table").PrevAll().Html()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("doc.Find(\"table\").PrevAll().Html():", orderInfo)

	// 获取失败: 获取订单（车次、订单ID）失败
	orderInfo, err = doc.Find("table").Siblings().Prev().Html()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("doc.Find(\"table\").Siblings().Prev().Html():", orderInfo)

	// 获取成功：乘客信息
	// 完成乘客信息
	orderInfo, err = doc.Find("tbody tr:nth-child(2)").Html()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("doc.Find(\"tbody tr:nth-child(2)\").Html():", orderInfo)
	// 证件类型
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(1)").Text()
	log.Println("doc.Find(\"tbody tr:nth-child(2) td:nth-child(1)\").Text():", orderInfo)
	// 获取成功：乘客信息-end
}

// 成功
func TestGOQueryOrder04(t *testing.T) {
	filePath := "/Users/leeprince/www/go/goinfra/utils/htmlutil/htmlutiltest/goquerytest/rpahtml/"
	filename := "单人单程-孙孙孙.html"
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
	log.Println("doc.Find(\"tbody tr:nth-child(2)\").Html():", orderInfo)
	// 证件类型
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(1)").Text()
	log.Println("证件类型:", orderInfo)
	// 姓名
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(2)").Text()
	log.Println("姓名:", orderInfo)
	// 身份证号
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(3)").Text()
	log.Println("身份证号:", orderInfo)
	// 票种：成人票、小孩票
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(4)").Text()
	log.Println("票种：成人票、小孩票:", orderInfo)
	// 座位类型
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(5) input[name='seatType']").AttrOr("value", "")
	log.Println("座位类型:", orderInfo)
	// 车厢
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(6) input[name='coachNo']").AttrOr("value", "")
	log.Println("车厢:", orderInfo)
	// 座位号
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(7) input[name='seatNo']").AttrOr("value", "")
	log.Println("座位号:", orderInfo)
	// 单张票（单人一程票）的价格
	orderInfo = doc.Find("tbody tr:nth-child(2) td:nth-child(8) input[name='ticketPrice']").AttrOr("value", "")
	log.Println("单张票（单人一程票）的价格:", orderInfo)

	// --- 订单最后一项
	// 订单手机号
	orderInfo, err = doc.Find("tbody .redMobile").Html()
	log.Println("订单手机号:", orderInfo)
	// 获取成功：乘客信息-end

}
