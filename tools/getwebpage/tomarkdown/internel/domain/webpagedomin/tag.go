package webpagedomin

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 13:20
 * @Desc:
 */

func (r *WebPageService) GetTag(doc *goquery.Document) (tag string, err error) {
	seletor := "body > div.zhuanlan-theme-container > main > div.page-title > div > span:nth-child(2) > span > a" // Chrome seletor
	titleSectoin := doc.Find(seletor)
	tag = titleSectoin.Text()
	if tag == "" {
		fmt.Println("GetTag err:", err)
		err = errors.New("GetTag Find tag empty!")
		return
	}
	
	return
}
