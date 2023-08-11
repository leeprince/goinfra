package main

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/fileutil"
	"github.com/leeprince/goinfra/utils/idutil"
	"github.com/leeprince/goinfra/utils/stringutil"
	"log"
	"net/http"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/29 22:05
 * @Desc:
 */

// 读取文件内容返回
func SupplierHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println(">>> SupplierHandler")
		
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "ParseForm: "+err.Error(), http.StatusBadRequest)
			return
		}
		var data struct {
			Optype            string `json:"optype"`
			ExistOrderNumbers string `json:"ExistOrderNumbers"`
		}
		data.Optype = r.Form.Get("optype")
		data.ExistOrderNumbers = r.Form.Get("ExistOrderNumbers")
		
		if data.Optype == "BookingOrderForAli" {
			// SupplierHandler.txt 的内容参考：fileData/BookingOrderForAli.txt.dt
			fileData, err := fileutil.ReadFile(dataDir, "BookingOrderForAli.txt")
			if err != nil {
				http.Error(w, "fileutil.ReadFile: "+err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprint(w, string(fileData))
			return
		}
		
		if data.Optype == "ConfirmSetBookSuc" {
			// SupplierHandler.txt 的内容参考：fileData/ConfirmSetBookSuc.txt.dt
			fileData, err := fileutil.ReadFile(dataDir, "ConfirmSetBookSuc.txt")
			if err != nil {
				http.Error(w, "fileutil.ReadFile: "+err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprint(w, string(fileData))
			return
		}
	}
	
	// 默认响应指定长度字符串
	time.Sleep((time.Millisecond * 800))
	fmt.Fprint(w, stringutil.FillCharRight(idutil.UniqIDV3(), '0', 28))
	return
}
