package htmlutil

import (
	"github.com/leeprince/goinfra/utils/dumputil"
	"golang.org/x/net/html"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/10 18:17
 * @Desc:
 */

func PrintHtml(n *html.Node, forNum int64) {
	// --- 打印输出
	var typeString string
	switch n.Type {
	case html.ErrorNode:
		typeString = "ErrorNode"
	case html.TextNode:
		typeString = "TextNode"
	case html.DocumentNode:
		typeString = "DocumentNode"
	case html.ElementNode:
		typeString = "ElementNode"
	case html.CommentNode:
		typeString = "CommentNode"
	case html.DoctypeNode:
		typeString = "DoctypeNode"
	case html.RawNode:
		typeString = "RawNode"
	default:
		typeString = ">>>error!"
	}
	dumputil.ForIndentPrintf(forNum,
		"",
		"n.Type:%v;typeString:%s;n.Data:%s \n",
		n.Type, typeString, n.Data,
	)

	if n.Attr != nil {
		for _, attribute := range n.Attr {
			dumputil.ForIndentPrintf(forNum,
				"",
				"n.Attr Key:%s; Val:%s;\n",
				attribute.Key, attribute.Val,
			)
		}
	}
	// --- 打印输出-end

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// if forNum == 2 {
		// 	return
		// }

		forNum++
		PrintHtml(c, forNum)
	}

	return
}
