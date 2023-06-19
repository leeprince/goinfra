package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/17 12:14
 * @Desc:
 */

/*
我通过收到了一个任务，并转发任务给其他程序异步处理，同时可以接续接受新任务。等到异步程序处理回调结果后，删除当前任务，并将异步程序处理回调结果前添加的任务继续让异步程序处理.
请给出设计思路和实现代码，写好注释。

需要注意的是：
 1. 收到异步程序处理回调结果前，可以继续接受新任务
 2. 收到异步程序处理回调结果后，异步程序才能处理下一个任务
*/

/*
设计思路：
1. 创建一个任务队列，用于存储待处理的任务。
2. 创建一个结果队列，用于存储异步程序处理回调结果。
3. 创建一个协程，用于处理任务队列中的任务。
4. 在协程中，从任务队列中取出一个任务，将其发送给异步程序进行处理等待异步程序处理回调结果。
5. 当异步程序处理回调结果后，将结果存储到结果队列中，并删除当前任务。
6. 在协程中，如果任务队列不为空，则继续处理下一个任务，否则等待新任务的到来。
7. 在主程序中，接收新任务并将其添加到任务队列中。

*/

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	id int
}

type Result struct {
	id int
}

func main() {
	ProceeV1()

	time.Sleep(time.Second * 1)
	fmt.Println("------------")

	ProceeV2()
}

// 通过 sync.WaitGroup 实现
func ProceeV1() {
	taskQueue := make(chan Task, 10)
	resultQueue := make(chan Result, 10)
	var wg sync.WaitGroup

	// 创建一个协程，用于处理任务队列中的任务
	go func() {
		for {
			// 从任务队列中取出一个任务
			task := <-taskQueue
			fmt.Printf("Processing task %d\n", task.id)

			// 将任务发送给异步程序进行处理，并等待异步程序处理回调
			wg.Add(1)
			go func(t Task) {
				defer wg.Done()
				// 模拟异步程序处理任务
				result := Result{id: t.id}
				resultQueue <- result
			}(task)

			// 当异步程序处理回调结果后，将结果存储到结果队列中，并删除当前任务
			wg.Wait()
		}
	}()

	go func() {
		for {
			// 如果结果队列中有结果
			select {
			case result := <-resultQueue:
				fmt.Printf("Task %d processed, result: %v\n", result.id, result)
			default:
			}
		}
	}()

	// 接收新任务并将其添加到任务队列中
	for i := 1; i <= 5; i++ {
		task := Task{id: i}
		taskQueue <- task
		fmt.Printf("New task %d added\n", task.id)
	}

	// 等待任务处理完成
	select {}
}

// 通过 select + channel 实现
func ProceeV2() {
	taskQueue := make(chan Task, 10)
	resultQueue := make(chan Result, 10)

	// 创建一个协程，用于处理任务队列中的任务
	go func() {
		for {
			select {
			case task := <-taskQueue:
				fmt.Printf("Processing task %d\n", task.id)

				// 将任务发送给异步程序进行处理，并等待异步程序处理回调
				resultChan := make(chan Result)
				go func(t Task) {
					// 模拟异步程序处理任务
					result := Result{id: t.id}
					resultChan <- result
					resultQueue <- result
				}(task)

				// 等待异步程序处理回调结果后，将结果存储到结果队列中，并删除当前任务
				<-resultChan
			default:
				// 如果任务队列为空，则等待新任务的到来
			}
		}
	}()

	go func() {
		for {
			// 如果结果队列中有结果
			select {
			case result := <-resultQueue:
				fmt.Printf("Task %d processed, result: %v\n", result.id, result)
			default:
			}
		}
	}()

	// 接收新任务并将其添加到任务队列中
	for i := 1; i <= 5; i++ {
		task := Task{id: i}
		taskQueue <- task
		fmt.Printf("New task %d added\n", task.id)
	}

	// 等待任务处理完成
	select {}

	fmt.Println("All tasks processed")
}
