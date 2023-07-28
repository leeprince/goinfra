package main

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/fileutil"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/9 10:14
 * @Desc:
 */

func OperateHttp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fileBytes, err := fileutil.ReadFile(resoureDir, "OperateHttp.html")
		if err != nil {
			http.Error(w, "读取 html文件错误", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(fileBytes))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
