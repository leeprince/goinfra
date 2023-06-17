package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/17 12:04
 * @Desc:	我通过收到了一个任务，并转发任务给其他程序异步处理，同时可以接续接受新任务。
 */

import (
	"fmt"
)

type Task struct {
	id   int
	data string
}

func main() {
	taskQueue := make(chan Task)
	resultQueue := make(chan string)

	// 启动一个 goroutine 处理任务队列
	go func() {
		for task := range taskQueue {
			// 将任务发送到另一个 goroutine 进行异步处理
			go func(t Task) {
				// 模拟异步处理
				result := t.data + " processed"
				// 将处理结果发送到结果队列
				resultQueue <- result
			}(task)
		}
	}()

	// 添加新任务到任务队列
	for i := 1; i <= 10; i++ {
		task := Task{id: i, data: fmt.Sprintf("task %d", i)}
		taskQueue <- task
	}

	// 等待异步处理结果的回调
	for i := 1; i <= 10; i++ {
		result := <-resultQueue
		fmt.Println(result)
	}
}
