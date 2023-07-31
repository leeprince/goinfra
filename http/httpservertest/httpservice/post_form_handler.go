package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/9 10:12
 * @Desc:
 */

func postFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println(">>> postFormHandler")
		
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "ParseForm: "+err.Error(), http.StatusBadRequest)
			return
		}
		
		// 获取表单数据
		name := r.Form.Get("name")
		age := r.Form.Get("age")
		
		// 打印表单数据
		fmt.Println("Name:", name)
		fmt.Println("Age:", age)
		
		// 返回响应
		fmt.Fprintf(w, "Hello, %s! Your age is %s.", name, age)
	}
	
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
