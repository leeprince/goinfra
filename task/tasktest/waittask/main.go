package main

import (
	"fmt"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/4 11:35
 * @Desc:
1、当前并发控制的数量为8；
2、如果有超过8的并发请求过来则开始阻塞等待1秒，如果1秒内正在进行的8个任务没有一个任务释放琐，则直接返回请求频繁，请稍候重试
*/

const MaxConcurrent = 8

var sem = make(chan struct{}, MaxConcurrent)

func processRequest(req int) {
	fmt.Printf("开始处理请求 %d\n", req)
	time.Sleep(2 * time.Second) // 模拟请求处理时间
	fmt.Printf("请求 %d 处理完毕\n", req)
}

func main() {
	for i := 1; i <= 20; i++ {
		select {
		case sem <- struct{}{}: // 尝试获取一个并发量
			go func(i int) {
				defer func() {
					<-sem // 释放一个并发量
				}()
				processRequest(i)
			}(i)
		case <-time.After(1 * time.Second): // 超时等待
			fmt.Println("请求频繁，请稍后重试：", i)
		}
	}

	/*// 等待所有请求处理完毕
	for i := 0; i < MaxConcurrent; i++ {
		sem <- struct{}{}
	}*/

	// 等待所有请求处理完毕
	// 等待后执行完毕后会报死锁的错误：fatal error: all goroutines are asleep - deadlock!（不影响测试结果）
	select {}

}
