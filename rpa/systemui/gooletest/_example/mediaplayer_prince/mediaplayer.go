//go:build windows
// +build windows

package main

import (
	"log"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func main() {
	ole.CoInitialize(0)

	unknown, err := oleutil.CreateObject("WMPlayer.OCX") // 成功
	//unknown, err := oleutil.CreateObject("Windows Media Player") // 失败
	//unknown, err := oleutil.CreateObject("17208") // 失败
	if err != nil {
		log.Fatal(err)
	}
	log.Println("unknown:", unknown)

	wmp := unknown.MustQueryInterface(ole.IID_IDispatch)
	log.Println("wmp:", wmp)

	collection := oleutil.MustGetProperty(wmp, "MediaCollection").ToIDispatch()
	log.Println("collection:", collection)

	list := oleutil.MustCallMethod(collection, "getAll").ToIDispatch()
	log.Println("list:", list)

	count := int(oleutil.MustGetProperty(list, "count").Val)
	log.Println("count:", count)

	log.Println("-------------------------------------------------------")

	for i := 0; i < count; i++ {
		log.Println("\n\n>>>>")

		item := oleutil.MustGetProperty(list, "item", i).ToIDispatch()
		log.Println("item:", item)

		name := oleutil.MustGetProperty(item, "name").ToString()
		log.Println("name:", name)

		sourceURL := oleutil.MustGetProperty(item, "sourceURL").ToString()
		log.Println("sourceURL:", sourceURL)

		// 添加仅为测试，正常删除。
		//break
	}
}
