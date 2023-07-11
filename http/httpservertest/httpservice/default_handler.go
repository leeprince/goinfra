package main

import (
	"fmt"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/9 15:17
 * @Desc:
 */

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	
	// fmt.Fprintf(w, "时间：%s \n", timeutil.DataTimeMicrosecond())
	// fmt.Fprintf(w, "Hello, World!")
	
	fmt.Fprint(w, "0123458789012345878901234587")
}
