package webpagedomin

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 03:07
 * @Desc:
 */

func (r *WebPageService) GetTitle(doc *goquery.Document) (title string, err error) {
	// Chrome seletor【推荐：简洁、易读】 || Firefox CSS路径
	// seletor := "html body.d-flex.flex-column.h-100 div.zhuanlan-theme-container main.main-content div.page-title h1.pb-1.fw-bold" // Firefox CSS路径
	seletor := "body > div.zhuanlan-theme-container > main > div.page-title > h1" // Chrome seletor
	titleSectoin := doc.Find(seletor)
	title = titleSectoin.Text()
	if title == "" {
		fmt.Println("GetTitle err:", err)
		err = errors.New("GetTitle Find Title empty!")
		return
	}
	
	return
}
