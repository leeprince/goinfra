package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/12 00:15
 * @Desc:
 */
import (
	"fmt"
	"github.com/elastic/go-sysinfo"
	"github.com/leeprince/goinfra/utils/dumputil"
	"log"
)

func main() {
	// 获取系统信息
	sysInfo, err := sysinfo.Host()
	if err != nil {
		log.Fatal(err)
	}
	dumputil.Println("---sysInfo:", sysInfo)
	
	memory, err := sysInfo.Memory()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("- sysInfo.Memory:%+v \n", memory)
	
	select {}
}
