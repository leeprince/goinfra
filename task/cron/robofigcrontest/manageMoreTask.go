package robofigcrontest

import (
	"fmt"
	"log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/21 10:32
 * @Desc:	管理多个任务
 */

func manageMoreTask() {
	c := NewWithSeconds()

	// 添加一个任务
	// 每3秒执行一次
	_, err := c.AddFunc("*/3 * * * * ?", func() {
		log.Println("cron run success!")
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// 添加一个任务
	_, err = c.AddFunc("*/1 * * * * *", func() { // 可以随时添加多个定时任务
		log.Printf("cron run success -- 02")
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Start()

	select {}
}
