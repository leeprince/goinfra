package main

// ocr 客户端示例；服务端则需要部署，参考：https://github.com/otiai10/ocrserver

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("ocr_test.png")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}
