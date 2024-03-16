package webpagedomin

import (
	"fmt"
	"getwebpage-tomarkdown/internel/domain/webpagedomin/webpageentity"
	"github.com/PuerkitoBio/goquery"
	"io"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 03:10
 * @Desc:
 */

func (r *WebPageService) ConvertMarkdown(httpBody io.Reader) (entity *webpageentity.MarkdownEntity, err error) {
	doc, err := goquery.NewDocumentFromReader(httpBody)
	if err != nil {
		err = fmt.Errorf("failed to parse HTML: %w", err)
		return
	}
	
	// 获取标题
	title, err := r.GetTitle(doc)
	if err != nil {
		err = fmt.Errorf("GetTitle err: %w", err)
		return
	}
	fmt.Println("title:", title)
	
	// 获取标题
	category, err := r.GetCategory(doc)
	if err != nil {
		err = fmt.Errorf("GetCategory err: %w", err)
		return
	}
	fmt.Println("category:", category)
	
	// 获取标签
	tag, err := r.GetTag(doc)
	if err != nil {
		err = fmt.Errorf("GetTag err: %w", err)
		return
	}
	fmt.Println("tag:", tag)
	
	// 将正文区域的HTML转换为Markdown
	content, err := r.GetContentToMarkdown(doc, title)
	if err != nil {
		err = fmt.Errorf("GetContentToMarkdown err: %w", err)
		return
	}
	fmt.Println("content:", len(content))
	
	entity = &webpageentity.MarkdownEntity{
		Title:    title,
		Category: category,
		Tag:      tag,
		Content:  content,
	}
	fmt.Println("ConvertMarkdown complete")
	
	return
}
