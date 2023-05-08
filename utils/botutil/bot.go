package botutil

import (
	"bytes"
	"encoding/json"
	"github.com/leeprince/goinfra/http/httpcli"
	"github.com/leeprince/goinfra/utils/sliceutil"
	"github.com/pkg/errors"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/16 18:13
 * @Desc:
 */

type BotContentType string

const (
	BOT_CONTENTTYPE_TEXT     BotContentType = "text"
	BOT_CONTENTTYPE_MARKDOWN BotContentType = "markdown"
)

type TextBody struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type MarkdownBody struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

// contextType
//
//	contextType=BOT_CONTENTTYPE_TEXT,则content的内容自动以中文逗号分割（不支持换行）
//	contextType=BOT_CONTENTTYPE_MARKDOWN,则content的内容以换行符分割（支持换行）
func SendQYWXBot(url string, contextType BotContentType, title string, contents []string) (respBobyBytes []byte, resp *http.Response, err error) {
	if url == "" {
		err = errors.New("SendQYWXBot url 必填")
		return
	}
	if title == "" {
		err = errors.New("SendQYWXBot title 必填")
		return
	}

	if !sliceutil.InString(string(contextType), []string{
		string(BOT_CONTENTTYPE_TEXT),
		string(BOT_CONTENTTYPE_MARKDOWN),
	}) {
		err = errors.New("SendQYWXBot contextType 暂不支持")
		return
	}

	// 组装默认格式
	var sendContentBuf *bytes.Buffer
	if contextType == BOT_CONTENTTYPE_TEXT {
		sendContentBuf = bytes.NewBufferString(title)
		sendContentBuf.WriteString("。")
	} else {
		sendContentBuf = bytes.NewBufferString("<font color=\"warning\">")
		sendContentBuf.WriteString(title)
		sendContentBuf.WriteString("</font>\n")
	}
	for i, content := range contents {
		if contextType == BOT_CONTENTTYPE_TEXT {
			sendContentBuf.WriteString(content)
			if i != len(contents)-1 {
				sendContentBuf.WriteString("，")
			}
			continue
		}
		sendContentBuf.WriteString("- ")
		sendContentBuf.WriteString(content)
		if i != len(contents)-1 {
			sendContentBuf.WriteString("\n")
		}
	}
	sendContent := sendContentBuf.String()

	var sendBodyByte []byte
	if contextType == BOT_CONTENTTYPE_TEXT {
		sendBody := TextBody{
			Msgtype: string(BOT_CONTENTTYPE_TEXT),
			Text: struct {
				Content string `json:"content"`
			}{
				Content: sendContent,
			},
		}

		sendBodyByte, err = json.Marshal(sendBody)
		if err != nil {
			return
		}
	} else {
		sendBody := MarkdownBody{
			Msgtype: string(BOT_CONTENTTYPE_MARKDOWN),
			Markdown: struct {
				Content string `json:"content"`
			}{
				Content: sendContent,
			},
		}

		sendBodyByte, err = json.Marshal(sendBody)
		if err != nil {
			return
		}
	}

	// 发送数据
	return httpcli.NewHttpClient().
		WithURL(url).
		WithMethod(http.MethodPost).
		WithBody(sendBodyByte).
		Do()
}
