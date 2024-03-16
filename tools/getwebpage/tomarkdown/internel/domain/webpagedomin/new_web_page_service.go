package webpagedomin

import (
	"getwebpage-tomarkdown/internel/Infrastructure/config"
	"getwebpage-tomarkdown/internel/pkg/ftpclient"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 02:58
 * @Desc:
 */

type Tag string

const (
	TagH1 Tag = "h1"
	TagH2 Tag = "h2"
	TagH3 Tag = "h3"
	TagH4 Tag = "h4"
	// 把br标签记录起来，用于下次匹配到 p时添加换行
	TagBr  Tag = "br"
	TagHr  Tag = "hr"
	TagDiv Tag = "div"
	// TagP 段落，需要配置 TagHtmlDouble 原样输出。
	// 	- 如果前面一个是br标签，则添加换行符。匹配后重置
	// 	- 如果前面一个是td标签，则无需记录，因为table已经原样输出了。匹配后重置
	TagP          Tag = "p"
	TagPre        Tag = "pre"        // 代码块
	TagUl         Tag = "ul"         // 无序列表，需要配置 TagHtmlDouble 原样输出
	TagTable      Tag = "table"      // 表格，需要配置 TagHtmlDouble 原样输出
	TagTableTr    Tag = "tr"         // 表格行，表格已经包含无需处理
	TagTableTd    Tag = "td"         // 单元格，表格已经包含无需处理
	TagImg        Tag = "img"        // 图片，需要上传到指定服务器，可通过ftp、cos等方式上传
	TagBlockquote Tag = "blockquote" // 引用，需要配置 TagHtmlDouble 原样输出
)

var TagHtmlDouble = map[Tag][2]string{
	TagP:          {"<p>", "</p>"},
	TagUl:         {"<ul>", "</ul>"},
	TagTable:      {"<table>", "</table>"},
	TagBlockquote: {"<blockquote>", "</blockquote>"},
}

type WebPageService struct {
	ftpClient *ftpclient.FtpClient
}

func NewWebPageService() *WebPageService {
	ftp := config.C.FTP
	ftpConf := ftpclient.Conf{
		Host:     ftp.Conf.Host,
		Port:     ftp.Conf.Port,
		Username: ftp.Conf.Username,
		Password: ftp.Conf.Password,
	}
	if ftpConf.Host == "" {
		panic("ftpConf.Host must config")
	}
	if ftpConf.Port == "" {
		panic("ftpConf.Port must config")
	}
	if ftpConf.Username == "" {
		panic("ftpConf.Username must config")
	}
	if ftpConf.Password == "" {
		panic("ftpConf.Password must config")
	}
	
	return &WebPageService{
		ftpClient: ftpclient.NewFtpClient(ftpConf, ftp.AccessHost),
	}
}
