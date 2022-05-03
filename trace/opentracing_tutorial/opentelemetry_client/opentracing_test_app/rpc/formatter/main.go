package main

import (
	"fmt"
    "github.com/leeprince/goinfra/plog"
    "log"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/5 下午4:12
 * @Desc:
 */
func main() {
    http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
        helloTo := r.FormValue("helloTo")
        helloStr := fmt.Sprintf("Hello, %s!", helloTo)
        w.Write([]byte(helloStr))
        
        plog.Info("HandleFunc /format end")
    })
    
    log.Fatal(http.ListenAndServe(":8201", nil))
}