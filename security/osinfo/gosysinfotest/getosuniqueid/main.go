package main

import (
	"fmt"
	"github.com/leeprince/goinfra/security/osinfo"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/12 12:57
 * @Desc:
 */

func main() {
	mark, err := osinfo.GetOsUniqueIdBase64()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	
	fmt.Println("mark:", mark)
	
	select {}
}
