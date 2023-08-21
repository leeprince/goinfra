package robofigcrontest

import (
	"fmt"
	"log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/21 10:01
 * @Desc: 	简单任务
 */

func simpleTask() {
	c := NewWithSeconds()

	spec := "*/3 * * * * ?" // 每3秒执行一次

	//添加一个任务
	_, err := c.AddFunc(spec, func() {
		log.Println("cron run success!")
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	c.Start()

	select {}
}
