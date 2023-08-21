package robofigcrontest

import (
	"fmt"
	"log"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/21 10:37
 * @Desc:	定时任务在运行中，需要注意的情况
 * 				- 默认上次任务没运行完，下次任务依然会运行（任务运行在goroutine里相互不干扰）
 * 				- 支持上次任务未执行完，下次任务不启动
 */

func taskRunning() {
	// 默认上次任务没运行完，下次任务依然会运行（任务运行在goroutine里相互不干扰）
	taskRunningDefault()

	//  支持上次任务未执行完，下次任务不启动
	taskRunningSkipIfStillRunning()
}

//默认上次任务没运行完，下次任务依然会运行（任务运行在goroutine里相互不干扰）
func taskRunningDefault() {
	c := NewWithSeconds()

	// 多个定时任务相当于多个协程，使用共享变量一定要注意并发更新的问题。正常处理方法：1.使用sync包加锁；2.使用通道通信自定义锁机制
	taskId := 0

	//添加一个任务
	// 每3秒执行一次
	_, err := c.AddFunc("*/1 * * * * ?", func() {
		log.Println(taskId, "-time.Sleep...")
		time.Sleep(time.Second * 5)
		log.Println(taskId, "-time.Sleep end")

		log.Println(taskId, "-cron run success!")
		taskId++
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	c.Start()

	select {}
}

// 支持上次任务未执行完，下次任务不启动
func taskRunningSkipIfStillRunning() {
	c := NewWithSecondsWithChain()

	// 多个定时任务相当于多个协程，使用共享变量一定要注意并发更新的问题。正常处理方法：1.使用sync包加锁；2.使用通道通信自定义锁机制
	taskId := 0

	//添加一个任务
	// 每3秒执行一次
	_, err := c.AddFunc("*/1 * * * * ?", func() {
		log.Println(taskId, "-time.Sleep...")
		time.Sleep(time.Second * 5)
		log.Println(taskId, "-time.Sleep end")

		log.Println(taskId, "-cron run success!")
		taskId++
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	c.Start()

	select {}
}
