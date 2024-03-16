package webpagedomin

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 03:07
 * @Desc:
 */

func (r *WebPageService) OutputSourceHtml(markdownContent *strings.Builder, s *goquery.Selection, tag Tag) (err error) {
	html, htmlerr := s.Html()
	if htmlerr != nil {
		fmt.Println("OutputSourceHtml Html err:", htmlerr)
		return
	}
	tagHtmlDouble, ok := TagHtmlDouble[tag]
	if !ok {
		fmt.Println("TagHtmlDouble TagHtmlDouble !ok:", tag)
		return
	}
	_, err = fmt.Fprintln(markdownContent, tagHtmlDouble[0], html, tagHtmlDouble[1])
	if err != nil {
		fmt.Println("TagHtmlDouble DataNodeAtomStringUl err:", err)
		return
	}
	_, err = fmt.Fprintln(markdownContent)
	if err != nil {
		fmt.Println("TagHtmlDouble p err:", err)
		return
	}
	
	return
}
