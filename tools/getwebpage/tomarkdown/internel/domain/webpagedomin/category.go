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

func (r *WebPageService) GetCategory(doc *goquery.Document) (category string, err error) {
	seletor := "body > div.zhuanlan-theme-container > main > div.page-title > div > span:nth-child(1) > span > a" // Chrome seletor
	titleSectoin := doc.Find(seletor)
	category = titleSectoin.Text()
	if category == "" {
		fmt.Println("GetCategory err:", err)
		err = errors.New("GetCategory Find category empty!")
		return
	}
	
	return
}
