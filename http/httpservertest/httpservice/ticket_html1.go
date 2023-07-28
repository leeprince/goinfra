package main

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/fileutil"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/9 11:01
 * @Desc:
 */

func ticketHtml1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fileBytes, err := fileutil.ReadFile(resoureDir, "ticketHtml1.html")
	if err != nil {
		http.Error(w, "读取 html文件错误", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(fileBytes))
}
