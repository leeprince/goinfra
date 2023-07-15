package main

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/idutil"
	"github.com/leeprince/goinfra/utils/stringutil"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/9 15:17
 * @Desc:
 */

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	
	// 默认响应字符串
	/*fmt.Fprintf(w, "时间：%s \n", timeutil.DataTimeMicrosecond())
	fmt.Fprintf(w, "Hello, World!")*/
	
	// 默认响应指定长度字符串
	fmt.Fprint(w, stringutil.FillCharRight(idutil.UniqIDV3(), '0', 28))
}
