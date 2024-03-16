package webpagedomin

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 03:09
 * @Desc:
 */

func (r *WebPageService) GetContentToMarkdown(doc *goquery.Document, title string) (content string, err error) {
	var markdownContent strings.Builder
	
	seletor := "body > div.zhuanlan-theme-container > main > div.row-content > article"
	contentSection := doc.Find(seletor)
	
	var preDataNodeAtomString Tag
	contentSection.Find("*").Each(func(i int, s *goquery.Selection) {
		dataNode := s.Get(0)
		dataNodeAtom := dataNode.DataAtom
		dataNodeAtomString := dataNodeAtom.String()
		
		switch Tag(dataNodeAtomString) {
		case TagH1:
			_, _ = fmt.Fprintln(&markdownContent, "\n")
			_, err = fmt.Fprintln(&markdownContent, "# ", s.Text())
			if err != nil {
				fmt.Println("ContentToMarkdown TagH1 err:", err)
				return
			}
			preDataNodeAtomString = Tag(dataNodeAtomString)
		case TagH2:
			_, _ = fmt.Fprintln(&markdownContent, "\n")
			_, err = fmt.Fprintln(&markdownContent, "## ", s.Text())
			if err != nil {
				fmt.Println("ContentToMarkdown TagH2 err:", err)
				return
			}
			preDataNodeAtomString = Tag(dataNodeAtomString)
		case TagH3:
			_, _ = fmt.Fprintln(&markdownContent, "\n")
			_, err = fmt.Fprintln(&markdownContent, "### ", s.Text())
			if err != nil {
				fmt.Println("ContentToMarkdown TagH3 err:", err)
				return
			}
			preDataNodeAtomString = Tag(dataNodeAtomString)
		
		case TagH4:
			_, _ = fmt.Fprintln(&markdownContent, "\n")
			_, err = fmt.Fprintln(&markdownContent, "#### ", s.Text())
			if err != nil {
				fmt.Println("ContentToMarkdown TagH4 err:", err)
				return
			}
			preDataNodeAtomString = Tag(dataNodeAtomString)
		
		case TagBr:
			preDataNodeAtomString = TagBr // 把换行标签记录起来，用于下次匹配到 p时添加换行
		case TagHr:
			_, _ = fmt.Fprintln(&markdownContent, "\n---")
		case TagP:
			if preDataNodeAtomString == TagTableTd ||
				preDataNodeAtomString == TagBlockquote {
				// 重置前一个标签
				preDataNodeAtomString = ""
				return
			}
			// 如果前面一个是换行标签，则添加换行符
			if preDataNodeAtomString == TagBr {
				_, _ = fmt.Fprintln(&markdownContent, "\n")
				// 重置前一个标签
				preDataNodeAtomString = ""
			}
			
			html, htmlerr := s.Html()
			if htmlerr != nil {
				fmt.Println("ContentToMarkdown TagP err:", err)
				return
			}
			
			// 需要处理图片了
			if strings.Contains(html, string(TagImg)) &&
				strings.Contains(html, "src") {
				fmt.Println("ContentToMarkdown TagP TagImg src return")
				return
			}
			err = r.OutputSourceHtml(&markdownContent, s, TagP)
			if err != nil {
				fmt.Println("ContentToMarkdown TagP err:", err)
				return
			}
		case TagUl:
			err = r.OutputSourceHtml(&markdownContent, s, TagUl)
			if err != nil {
				fmt.Println("ContentToMarkdown TagUl err:", err)
				return
			}
		case TagImg:
			err = r.ConvertImg(&markdownContent, s, title)
			if err != nil {
				fmt.Println("ContentToMarkdown TagUl err:", err)
				return
			}
		case TagTable:
			err = r.OutputSourceHtml(&markdownContent, s, TagTable)
			if err != nil {
				fmt.Println("ContentToMarkdown TagTable err:", err)
				return
			}
			preDataNodeAtomString = "" // 把td标签记录起来，用于下次匹配到 p时跳过
		case TagTableTd:
			preDataNodeAtomString = TagTableTd // 把td标签记录起来，用于下次匹配到 p时跳过
		case TagBlockquote:
			err = r.OutputSourceHtml(&markdownContent, s, TagBlockquote)
			if err != nil {
				fmt.Println("ContentToMarkdown TagBlockquote err:", err)
				return
			}
			preDataNodeAtomString = TagTableTd // 把blockquote标签记录起来，用于下次匹配到 p时跳过
		case TagPre:
			code := s.Find("code")
			if code.Length() > 0 {
				_, err = fmt.Fprintln(&markdownContent, "```\n", code.Text(), "```")
				if err != nil {
					fmt.Println("ContentToMarkdown pre Length > 0 err:", err)
					return
				}
			} else {
				_, err = fmt.Fprintln(&markdownContent, "```\n", s.Text(), "\n```")
				if err != nil {
					fmt.Println("ContentToMarkdown pre Length > 0 err:", err)
					return
				}
			}
		default:
			// 其他标签或者文本节点，这里不做处理
		}
	})
	
	content = markdownContent.String()
	return
}
