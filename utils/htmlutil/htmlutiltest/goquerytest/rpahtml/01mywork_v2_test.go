package rpahtml

import (
	"embed"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/leeprince/goinfra/utils/fileutil"
	"github.com/leeprince/goinfra/utils/stringutil"
	"gopkg.in/yaml.v3"
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

// 兼容：单人单程、单人单程-占座票、单人换乘、多人换乘
// 最后特殊的html就需要自定义处理（如<tr>祖先不存在<table>的部分 html），可以通过"golang.org/x/net/html"完成，具体请看`htmltest/rpahtml`
func TestMyWorkV2(t *testing.T) {
	filePath := "/Users/leeprince/www/go/goinfra/utils/htmlutil/htmlutiltest/goquerytest/rpahtml/"
	/*
		添加<table>包含所有原始数据，可以保证能解析出"订单概要信息的详细信息"
	*/
	// filename := "单人单程-李李李.html" // 无要求-添加<table>包含所有原始数据
	// filename := "单人单程-孙孙孙.html" // 无要求-添加<table>包含所有原始数据
	// filename := "单人单程-蔡蔡蔡.html" // 无要求-添加<table>包含所有原始数据
	// filename := "单人单程-朱朱朱.html" // 有要求-保持不添加<table>包含所有原始数据
	// filename := "单人单程-占座票-安安安.html" // 有要求-添加<table>包含所有原始数据
	// filename := "单人单程-占座票-安安安-未包含在 table内.html" // 有要求-保持不添加<table>包含所有原始数据
	// filename := "单人换乘-冯冯.html" // 有要求-添加<table>包含所有原始数据
	// filename := "单人换乘-蹇蹇蹇.html" // 有要求-添加<table>包含所有原始数据
	// filename := "多人单程-周周周.html" // 有要求-添加<table>包含所有原始数据
	// filename := "多人单程-马马.html" // 无要求-添加<table>包含所有原始数据
	// filename := "多人单程-王王王.html" // 有要求-保持不添加<table>包含所有原始数据
	filename := "多人换乘-余余余.html" // 有要求-添加<table>包含所有原始数据
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
		selection  *goquery.Selection
		orderInfo  string
		orderInfos []string
		id         string
		exist      bool
	)

	// 获取订单概要信息
	fmt.Printf("\n\n>>>>开始输出订单概要信息\n")
	selection = doc.Find("tr[aria-controls='collapseOne'][href]")
	if selection.Size() > 0 {
		selection.Each(func(i int, sl *goquery.Selection) {
			fmt.Printf("---")
			fmt.Println(sl.Html())
			fmt.Println(sl.Text())
		})
	} else {
		fmt.Println("---", "找不到tr[aria-controls='collapseOne'][href]")
	}

	fmt.Printf("\n\n>>>>开始输出订单概要信息的详细信息\n")
	selection = doc.Find("tr[aria-controls='collapseOne'][href]")
	if selection.Size() > 0 {
		selection.Each(func(i int, sl *goquery.Selection) {
			fmt.Printf(">>>")

			orderInfo = sl.Find("td:nth-child(1)").Text()
			fmt.Println("非占座时显示剩余时间或者占座时包含占座:", orderInfo)
			if strings.Contains(orderInfo, "占座") {
				fmt.Println("！！！！！！！此订单是占座订单！！！！！！！")
			} else {
				fmt.Println("！！！！！！！此订单是非占座订单！！！！！！！")
			}

			orderInfo = sl.Find("td:nth-child(2)").Text()
			fmt.Println("订单 ID:", orderInfo)

			orderInfo = sl.Find("td:nth-child(3)").Text()
			fmt.Println("乘座时间（含换乘）:", orderInfo)
			orderInfos = strings.Split(orderInfo, ">>>")
			fmt.Println("乘座时间（含换乘）转数组:", orderInfos)
			if len(orderInfos) > 1 {
				fmt.Println("！！！！！！！此订单是换乘订单！！！！！！！")
			}

			orderInfo = sl.Find("td:nth-child(4)").Text()
			fmt.Println("车次（含换乘）:", orderInfo)
			orderInfos = strings.Split(orderInfo, ">>>")
			fmt.Println("车次（含换乘）转数组:", orderInfos)

			orderInfo = sl.Find("td:nth-child(5)").Text()
			fmt.Println("出发地与目的地（含换乘）:", orderInfo)
			orderInfos = strings.Split(orderInfo, ">>>")
			fmt.Println("出发地与目的地（含换乘）转数组:", orderInfos)
			for _, info := range orderInfos {
				fromAddr, formCode, toAddr, toCode, err := parseFromToAddrAndCode(info)
				if err != nil {
					fmt.Println("发生错误", err)
					return
				}

				fmt.Printf("出发地:%s;出发地电报码:%s;目的地:%s;目的地电报码:%s\n", fromAddr, formCode, toAddr, toCode)
			}

			orderInfo = sl.Find("td:nth-child(6)").Text()
			fmt.Println("座位类型（含换乘）:", orderInfo)
			orderInfos = strings.Split(orderInfo, ">>>")
			fmt.Println("座位类型（含换乘）转数组:", orderInfos)

			orderInfo = sl.Find("td:nth-child(7)").Text()
			fmt.Println("该订单总金额:", orderInfo)
		})
	} else {
		fmt.Println("---", "找不到tr[aria-controls='collapseOne'][href]")
	}

	// 获取单程时的要求信息；或者获取多程时的单程信息和要求信息
	fmt.Printf("\n\n>>>>开始输出行程和要求信息\n")
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

	fmt.Printf("\n\n>>>>开始输出行程和要求信息的详细信息\n")

	// 用于换乘时，判断是否包含换乘信息
	moreWayReg := regexp.MustCompile(`第\d程：`)

	selection = doc.Find("td[colspan='8'][style='padding:5px;']")
	var isTransfer bool
	if selection.Size() > 0 {
		selection.Each(func(i int, sl *goquery.Selection) {
			slText := strings.Trim(sl.Text(), " ")
			fmt.Println("slText：", slText)

			if moreWayReg.MatchString(slText) {
				if !isTransfer {
					isTransfer = true
					fmt.Println("---该订单是换乘订单！")
				}
				fmt.Println("这是换乘信息，并开始输出换乘信息：", slText)
				return
			}

			requirementText := strings.Trim(sl.Find("span[class='maxRedFont'][id]").Text(), "")
			fmt.Println("这是要求信息，并开始输出要求信息：", requirementText)

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
			fmt.Printf("V1---%d \n", i)
			fmt.Println(sl.Html())

			id, exist = sl.Attr("id")
			if !exist {
				fmt.Println("当前Selection不存在id属性")
				return
			}
			fmt.Println("当前Selection的id属性:", id)
		})
	} else {
		fmt.Println("V1---", "找不到包含expectupseat、expectdownseat或expectmidseat属性的tr元素。")
	}
	// 完成乘客信息V2
	selection = doc.Find("tr[expectupseat='0'], tr[expectdownseat='0'], tr[expectmidseat='0']")
	if selection.Size() > 0 {
		selection.Each(func(i int, sl *goquery.Selection) {
			fmt.Printf("V2---%d \n", i)
			fmt.Println(sl.Html())

			id, exist = sl.Attr("id")
			if !exist {
				fmt.Println("当前Selection不存在id属性")
				return
			}
			fmt.Println("当前Selection的id属性:", id)
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
			fmt.Printf("V1---\n")

			id, exist = sl.Attr("id")
			if !exist {
				log.Println("不存在乘客Id")
				return
			}
			log.Println("乘客Id:", id)

			orderInfo = sl.Find("td:nth-child(1)").Text()
			log.Println("证件类型:", strings.TrimSpace(orderInfo))

			orderInfo = sl.Find("td:nth-child(2)").Text()
			log.Println("姓名:", strings.TrimSpace(orderInfo))

			orderInfo = sl.Find("td:nth-child(3)").Text()
			log.Println("未处理的身份证号信息:", orderInfo)
			orderInfo = stringutil.ReplaceSpace(orderInfo)
			orderInfoRune := []rune(orderInfo)
			orderInfo = string(orderInfoRune[:len(orderInfoRune)-2])
			log.Println("身份证号:", orderInfo)

			orderInfo = sl.Find("td:nth-child(4)").Text()
			log.Println("票种：成人票、小孩票:", strings.TrimSpace(orderInfo))

			orderInfo, exist = sl.Find("td:nth-child(5) input[name='seatType']").Attr("value")
			if !exist {
				log.Fatal("座位类型 !exist")
			}
			log.Println("座位类型:", strings.TrimSpace(orderInfo))

			orderInfo, exist = sl.Find("td:nth-child(6) input[name='coachNo']").Attr("value")
			if !exist {
				log.Fatal("车厢 !exist")
			}
			log.Println("车厢:", strings.TrimSpace(orderInfo))

			orderInfo, exist = sl.Find("td:nth-child(7) input[name='seatNo']").Attr("value")
			if !exist {
				log.Fatal("座位号 !exist")
			}
			log.Println("座位号:", strings.TrimSpace(orderInfo))

			orderInfo, exist = sl.Find("td:nth-child(8) input[name='ticketPrice']").Attr("value")
			if !exist {
				log.Fatal("单张票（单人一程票）的价格 !exist")
			}
			log.Println("单张票（单人一程票）的价格:", strings.TrimSpace(orderInfo))

			fmt.Printf("\n------------\n")
		})
	} else {
		fmt.Println("V1", "找不到包含expectupseat、expectdownseat或expectmidseat属性的tr元素。")
	}

	fmt.Printf("\n\n>>>>开始输出手机号信息\n")

	mobile := doc.Find("tbody .redMobile").Text()
	fmt.Println("手机号：", mobile)

}

func TestAddr01(t *testing.T) {
	data := []string{
		"崇州(CZE 15:19)->成都(CNW)",
		"青岛机场( 13:17)->潍坊北(WJK)",
		"银川(YIJ )->北京西(BXP)",
	}

	for _, d := range data {
		pattern := regexp.MustCompile(`(.+)\((\w{3}).*\)->\s*(.+?)\((\w{3}).*\)`)

		matches := pattern.FindStringSubmatch(d)
		fmt.Println(matches)
		if len(matches) != 5 {
			continue
		}
		from := strings.TrimSpace(matches[1])
		fromCode := strings.TrimSpace(matches[2])
		to := strings.TrimSpace(matches[3])
		toCode := strings.TrimSpace(matches[4])

		fmt.Printf("出发地：%s，发地电报码：%s；目的地：%s；目的地电报码：%s\n", from, fromCode, to, toCode)
	}
}

func TestAddr02(t *testing.T) {
	data := []string{
		"崇州(CZE 15:19)->成都南(CNW)",
		"青岛机场( 1317)->潍坊北(WJK)",
		"银川(YIJ )->北京西(BXP)",
	}

	for _, d := range data {
		fromAddr, formCode, toAddr, toCode, err := parseFromToAddrAndCode(d)
		if err != nil {
			fmt.Println("发生错误", err)
			continue
		}

		fmt.Printf("出发地:%s;出发地电报码:%s;目的地:%s;目的地电报码:%s\n", fromAddr, formCode, toAddr, toCode)
	}
}

//go:embed station.yaml
var stationsData embed.FS

type StationInfo struct {
	QuanPinMa string `yaml:"全拼码"`
	DianBaoMa string `yaml:"电报码"`
	JianPinMa string `yaml:"简拼码"`
}

var stations map[string]StationInfo

func init() {
	// 读取 YAML 文件
	/* stationsData.ReadFile("station.yaml") 配合 //go:embed station.yaml 使用
	//go:embed station.yaml
	var stationsData embed.FS
	*/
	data, err := stationsData.ReadFile("station.yaml")
	if err != nil {
		log.Fatalf("failed to read file %v", err)
	}

	// 解析 YAML 文件
	err = yaml.Unmarshal(data, &stations)
	if err != nil {
		log.Fatalf("failed to unmarshal data: %v", err)
	}
}

/*
	解析下面的字符串

崇州(CZE 15:19)->成都南(CNW)
青岛机场( 1317)->潍坊北(WJK)
银川(YIJ )->北京西(BXP)
*/
func parseFromToAddrAndCode(s string) (fromAddr, formCode, toAddr, toCode string, err error) {
	// 解析出出发地和目的地
	parts := strings.Split(s, "->")
	from := strings.TrimSpace(parts[0])
	to := strings.TrimSpace(parts[1])

	// 解析出出发地电报码
	getAddrAndCode := func(v string) (addr, code string) {
		leftAddrArr := strings.Split(v, "(")
		addr = leftAddrArr[0]
		leftCodeArr := strings.Split(leftAddrArr[1], " ")
		code = leftCodeArr[0]
		code = strings.TrimRight(code, ")")
		if code != "" {
			code = "-" + code
		}
		return
	}

	fromAddr, formCode = getAddrAndCode(from)
	if formCode == "" {
		formCode = stations[fromAddr].DianBaoMa
		if formCode == "" {
			err = errors.New("出发地的电报码找不到！出发地:" + fromAddr)
			return
		}
	}
	toAddr, toCode = getAddrAndCode(to)
	if toCode == "" {
		toCode = stations[toAddr].DianBaoMa
		if toCode == "" {
			err = errors.New("目的地的电报码找不到！目的地:" + toAddr)
			return
		}
	}
	return
}
