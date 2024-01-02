package main

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/model"
	"log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/2 16:05
 * @Desc:
 */

func main() {
	// 打开PDF文件
	pdfReader, _, err := model.NewPdfReaderFromFile("input.pdf", nil)
	if err != nil {
		log.Fatalf("Could not create reader: %v\n", err)
	}

	// 获取PDF的总页数
	totalPage, err := pdfReader.GetNumPages()
	if err != nil {
		log.Fatalf("GetNumPages: %v\n", err)
	}

	// 遍历每一页并将其转换为图像
	for i := 1; i <= totalPage; i++ {
		// 从PDF中提取页面
		page, err := pdfReader.GetPage(i)
		if err != nil {
			log.Fatalf("Could not retrieve page: %v\n", err)
		}

		// 将页面转换为图像
		img, err := page.ToImage(nil)
		if err != nil {
			log.Fatalf("Could not convert page to image: %v\n", err)
		}

		// 将图像保存为PNG文件
		err = img.WriteToFile(fmt.Sprintf("output_%d.png", i))
		if err != nil {
			log.Fatalf("Could not write image to file: %v\n", err)
		}
	}
}
