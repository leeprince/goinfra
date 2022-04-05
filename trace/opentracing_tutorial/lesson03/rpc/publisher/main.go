package main

import (
    "log"
    "net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/5 下午4:11
 * @Desc:
 */

func main() {
	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		helloStr := r.FormValue("helloStr")
		println(helloStr)
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}