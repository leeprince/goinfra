package webpagedomin

import (
	"fmt"
	"getwebpage-tomarkdown/internel/domain/webpagedomin/webpageentity"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 02:58
 * @Desc:
 */

func (r *WebPageService) FetchMarkdown(url string) (webPageEntity *webpageentity.MarkdownEntity, err error) {
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("failed to fetch URL: %w", err)
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
		return
	}
	
	httpBody := resp.Body
	webPageEntity, err = r.ConvertMarkdown(httpBody)
	if err != nil {
		err = fmt.Errorf("ConvertMarkdown err : %s", err.Error())
		return
	}
	
	return
}
