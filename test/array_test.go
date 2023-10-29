package test

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/dumputil"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/23 09:30
 * @Desc:
 */

func TestTwoArray(t *testing.T) {
	type person struct {
		i int
	}

	var personList []person
	var personListList [][]person

	for i := 0; i <= 20; i++ {
		personInfo := person{
			i: i,
		}

		personList = append(personList, personInfo)
		if i != 0 && i%5 == 0 {
			personListList = append(personListList, personList)

			// 初始化下一个数组
			// personList = nil
			personList = []person{}
		}
	}

	dumputil.Println("personListList:%+v", personListList)
}

// 因为golang的函数、方法的所有参数都是传值，所以就是是引用类型类型，需要方法中修改到外部变量时需要传递变量地址
func TestRequestParamHaveArray(t *testing.T) {
	var strArray []string
	var strArrayQute []string

	RequestParamHaveArray(strArray)
	fmt.Println(strArray)

	RequestParamHaveArrayQuote(&strArrayQute)
	fmt.Println(strArrayQute)

	fmt.Println("-------------------")
	str2Array := []string{}
	str2ArrayQute := make([]string, 0)

	RequestParamHaveArray(str2Array)
	fmt.Println(str2Array)

	RequestParamHaveArrayQuote(&str2ArrayQute)
	fmt.Println(str2ArrayQute)
	fmt.Println("-------------------")

	var str3Array []string
	str3ArrayQute := make([]string, 0)

	RequestParamHaveArray(str3Array)
	fmt.Println(str3Array)

	RequestParamHaveArrayQuote(&str3ArrayQute)
	fmt.Println(str3ArrayQute)
}

func RequestParamHaveArray(strArray []string) {
	strArray = append(strArray, "aaa")
}

func RequestParamHaveArrayQuote(strArray *[]string) {
	*strArray = append(*strArray, "aaa")
}

func TestArrCutting(t *testing.T) {
	a1 := 1
	a2 := 2
	a3 := 3
	arr := []*int{&a1, &a2, &a3}

	fmt.Println(arr)

	arrNew := arr[:2]
	fmt.Println(arrNew)
	fmt.Println(arr)

	fmt.Println(arr[:2])
	fmt.Println(arr)
}

type ReadInvoiceListReq struct {
	OrderSnInvoiceIdList []*OrderSnInvoiceId `json:"order_sn_invoice_id_list,omitempty"`
	HasRed               string              `json:"has_red,omitempty"`          // 红票标记
	ReimburseStatus      string              `json:"reimburse_status,omitempty"` // 报账状态
}

// 订单编号&发票id
type OrderSnInvoiceId struct {
	OrderSn   string `json:"order_sn,omitempty"`   // 订单编号
	InvoiceId string `json:"invoice_id,omitempty"` // 发票id
}

// 需要注意：指针变量的赋值
func TestReadInvoiceListReq(t *testing.T) {
	req := &ReadInvoiceListReq{
		OrderSnInvoiceIdList: []*OrderSnInvoiceId{
			{
				OrderSn:   "01",
				InvoiceId: "01",
			},
			{
				OrderSn:   "02",
				InvoiceId: "02",
			},
			{
				OrderSn:   "03",
				InvoiceId: "03",
			},
			{
				OrderSn:   "04",
				InvoiceId: "04",
			},
			{
				OrderSn:   "05",
				InvoiceId: "05",
			},
			{
				OrderSn:   "06",
				InvoiceId: "06",
			},
			{
				OrderSn:   "07",
				InvoiceId: "07",
			},
			{
				OrderSn:   "08",
				InvoiceId: "08",
			},
			{
				OrderSn:   "09",
				InvoiceId: "09",
			},
			{
				OrderSn:   "10",
				InvoiceId: "10",
			},
		},
		HasRed:          "HasRed",
		ReimburseStatus: "ReimburseStatus",
	}

	ReadInvoiceList(req)
}

func ReadInvoiceList(req *ReadInvoiceListReq) {
	fmt.Println(req)

	fmt.Println(req.OrderSnInvoiceIdList)
	fmt.Println(req.OrderSnInvoiceIdList[1:])
	fmt.Println(req.OrderSnInvoiceIdList)

	fmt.Println("-------------1")
	// 错误赋值
	newReq := req
	newReq.OrderSnInvoiceIdList = req.OrderSnInvoiceIdList[1:]
	newReq.ReimburseStatus = "ReimburseStatus1"
	fmt.Println(newReq)
	fmt.Println(req)
	for i, id := range req.OrderSnInvoiceIdList {
		fmt.Println(i, ":", id)
	}

	fmt.Println("-------------2")
	// 错误赋值
	newReq2 := req
	newReqIdList2 := req.OrderSnInvoiceIdList[1:]
	newReq2.OrderSnInvoiceIdList = newReqIdList2
	newReq2.ReimburseStatus = "ReimburseStatus2"
	fmt.Println(newReq2)
	fmt.Println(req)
	for i, id := range req.OrderSnInvoiceIdList {
		fmt.Println(i, ":", id)
	}

	fmt.Println("-------------3")
	// 错误赋值
	var newReq3 *ReadInvoiceListReq
	newReq3 = req
	newReqIdList3 := req.OrderSnInvoiceIdList[1:]
	newReq3.OrderSnInvoiceIdList = newReqIdList3
	fmt.Println(newReq3)
	fmt.Println(req)
	for i, id := range req.OrderSnInvoiceIdList {
		fmt.Println(i, ":", id)
	}

	fmt.Println("-------------4")

	// 正确赋值
	var newReq4 ReadInvoiceListReq
	newReq4 = *req
	newReqIdList4 := req.OrderSnInvoiceIdList[1:]
	newReq4.OrderSnInvoiceIdList = newReqIdList4
	fmt.Println(newReq4)
	fmt.Println(req)
	for i, id := range req.OrderSnInvoiceIdList {
		fmt.Println(i, ":", id)
	}
	fmt.Println("-------------5")

}

func modify(array []int) {
	array[0] = 10 // 对入参slice的元素修改会影响原始数据
}

func add(array []int) {
	array = append(array, 100)
}

func addPointer(array *[]int) {
	*array = append(*array, 100)
}

// 【推荐】不使用slice作为函数入参
/*
slice在作为函数入参时，函数内对slice的修改可能会影响原始数据。
优化：数组作为函数入参，而不是slice
*/
func TestArrayModify(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}

	modify(array)
	fmt.Println(array) // output：[10 2 3 4 5]

	add(array)
	fmt.Println(array) // output：[10 2 3 4 5]

	addPointer(&array)
	fmt.Println(array) // output：[10 2 3 4 5 100]

}
