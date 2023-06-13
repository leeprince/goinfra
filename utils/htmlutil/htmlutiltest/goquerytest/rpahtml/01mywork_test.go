package rpahtml

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/leeprince/goinfra/utils/fileutil"
	"log"
	"regexp"
	"strings"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/11 22:35
 * @Desc:	特殊的html就需要自定义处理，可以通过"golang.org/x/net/html"完成，具体请看`htmltest/rpahtml`
 */

// 单人单程成功,但是未能兼容单人换乘
func TestOnePersonOneWay(t *testing.T) {
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

// 兼容：单人单程、单人单程-占座票、单人换乘、多人换乘
// 最后特殊的html就需要自定义处理（如<tr>祖先不存在<table>的部分 html），可以通过"golang.org/x/net/html"完成，具体请看`htmltest/rpahtml`
func TestMyWork(t *testing.T) {
	filePath := "/Users/leeprince/www/go/goinfra/utils/htmlutil/htmlutiltest/goquerytest/rpahtml/"
	// filename := "单人单程-孙孙孙.html" // 兼容
	// filename := "单人换乘-冯冯.html" // 兼容
	// filename := "多人单程-周周周.html" // 兼容
	// filename := "单人单程-占座票-安安安.html" // 兼容
	filename := "多人换乘-余余余.html" // 兼容
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
		selection *goquery.Selection
		orderInfo string
		exist     bool
	)

	// 获取单程时的要求信息；或者获取多程时的单程信息和要求信息
	fmt.Printf("\n\n>>>>开始输出获取行程和要求信息\n")
	selection = doc.Find("td[colspan='8'][style='padding:5px;']")
	if selection.Size() > 0 {
		selection.Each(func(i int, sl *goquery.Selection) {
			fmt.Printf("V---")
			fmt.Println(sl.Html())
			fmt.Println(sl.Text())
		})
	} else {
		fmt.Println("V---", "找不到td[colspan='8'][style='padding:5px;']的td元素")
	}

	fmt.Printf("\n\n>>>>开始输出获取行程和要求信息的详细信息\n")

	// 用于换乘时，判断是否包含换乘信息
	moreWayReg := regexp.MustCompile(`第\d程：`)

	selection = doc.Find("td[colspan='8'][style='padding:5px;']")
	if selection.Size() > 0 {
		selection.Each(func(i int, sl *goquery.Selection) {
			slText := strings.Trim(sl.Text(), " ")
			fmt.Println("slText：", slText)
			if slText == "" {
				fmt.Println("单程且无任何要求信息")
				return
			}

			if moreWayReg.MatchString(slText) {
				fmt.Println("这是换乘信息，并开始输出换乘信息：", slText)
			}

			requirementText := strings.Trim(sl.Find("span[class='maxRedFont'][id]").Text(), "")
			if requirementText != "" {
				fmt.Println("这是要求信息，并开始输出要求信息：", requirementText)
			}

			fmt.Printf("\n------------\n")
		})
	} else {
		fmt.Println("V---", "找不到td[colspan='8'][style='padding:5px;']的td元素")
	}

	fmt.Printf("\n\n>>>>开始输出乘客信息\n")

	// 获取成功：乘客信息
	// 完成乘客信息V1
	selection = doc.Find("tr[expectupseat], tr[expectdownseat], tr[expectmidseat]")
	if selection.Size() > 0 {
		selection.Each(func(i int, sl *goquery.Selection) {
			fmt.Printf("V1---")
			fmt.Println(sl.Html())
		})
	} else {
		fmt.Println("V1---", "找不到包含expectupseat、expectdownseat或expectmidseat属性的tr元素。")
	}
	// 完成乘客信息V2
	selection = doc.Find("tr[expectupseat='0'], tr[expectdownseat='0'], tr[expectmidseat='0']")
	if selection.Size() > 0 {
		selection.Each(func(i int, sl *goquery.Selection) {
			fmt.Printf("V2---")
			fmt.Println(sl.Html())
		})
	} else {
		fmt.Println("V2---", "找不到包含expectupseat、expectdownseat或expectmidseat属性的tr元素。")
	}

	fmt.Printf("\n\n>>>>开始输出乘客详细信息\n")

	// 完成乘客信息V1 完成下面数据的查找
	// 证件类型
	selection = doc.Find("tr[expectupseat], tr[expectdownseat], tr[expectmidseat]")
	if selection.Size() > 0 {
		selection.Each(func(i int, sl *goquery.Selection) {
			fmt.Printf("V1---")

			orderInfo = sl.Find("td:nth-child(1)").Text()
			log.Println("证件类型:", orderInfo)
			orderInfo = sl.Find("td:nth-child(2)").Text()
			log.Println("姓名:", orderInfo)
			orderInfo = sl.Find("td:nth-child(3)").Text()
			log.Println("身份证号:", orderInfo)
			orderInfo = sl.Find("td:nth-child(4)").Text()
			log.Println("票种：成人票、小孩票:", orderInfo)
			orderInfo, exist = sl.Find("td:nth-child(5) input[name='seatType']").Attr("value")
			if !exist {
				log.Fatal("座位类型 !exist")
			}
			log.Println("座位类型:", orderInfo)
			orderInfo, exist = sl.Find("td:nth-child(6) input[name='coachNo']").Attr("value")
			if !exist {
				log.Fatal("车厢 !exist")
			}
			log.Println("车厢:", orderInfo)
			orderInfo, exist = sl.Find("td:nth-child(7) input[name='seatNo']").Attr("value")
			if !exist {
				log.Fatal("座位号 !exist")
			}
			log.Println("座位号:", orderInfo)
			orderInfo, exist = sl.Find("td:nth-child(8) input[name='ticketPrice']").Attr("value")
			if !exist {
				log.Fatal("单张票（单人一程票）的价格 !exist")
			}
			log.Println("单张票（单人一程票）的价格:", orderInfo)

			fmt.Printf("\n------------\n")
		})
	} else {
		fmt.Println("V1", "找不到包含expectupseat、expectdownseat或expectmidseat属性的tr元素。")
	}

	fmt.Printf("\n\n>>>>开始输出手机号信息\n")

	mobile := doc.Find("tbody .redMobile").Text()
	fmt.Println("手机号：", mobile)

}
