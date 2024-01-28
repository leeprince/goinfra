package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/27 15:38
 * @Desc:
 */
import (
	"fmt"
	"github.com/iris-contrib/blackfriday"
	"log"
	"os"
	
	"github.com/jung-kurt/gofpdf"
)

func main() {
	input, err := os.ReadFile("input.md")
	if err != nil {
		log.Fatal(err)
	}
	
	html := blackfriday.Run(input)
	
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetHeaderFunc(func() {
		pdf.Image("header.jpg", 10, 10, 30, 0, false, "", 0, "")
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(0, 10, "itgogogo.cn IT go go go，程序员编程资料站")
	})
	
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.Cell(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()))
	})
	
	pdf.AddPage()
	pdf.Write(30, string(html))
	
	err = pdf.OutputFileAndClose("output.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
