package main

import (
	"fmt"
	"sync"
	"time"
)

/*
假设你是一名 golang 维护者，
我收到了一个任务，需要添加到本地队列中，任务的结构为 Task{id string }, 如果检查到本地队列不为空，则取队列的第一个任务作为当前任务并转发到外部服务中异步处理。
我希望
1. 在收到外部服务的异步处理结果回调前，其他新任务进来，都加入本地队列中，
2. 在收到外部服务的异步处理结果回调后，才能删除当前任务，并从继续从非空的本地队列中读取第一条任务，转发到外部服务中异步处理
3. 有当前任务的概念，因为可能会因为业务逻辑的需要重发当前任务
我希望你能设计这样一个满足需求的本地队列，并避免并发时出现竞态资源问题，给出实现代码，写好注释。


需要注意的是：
外部服务实际上是发送一个 http请求到其他 http服务中，但是该 http的响应不作为任务的处理结果，任务的处理结果是是等待这个外部服务通过 http调用当前服务得到的
*/
type Task struct {
	id string
}

type LocalQueue struct {
	tasks      []*Task
	current    *Task
	processing bool
	mutex      sync.Mutex
}

func (q *LocalQueue) Add(task *Task) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	fmt.Printf("---有新任务,当前队列：%+v;当前队列数量：%+v;当前任务：%+v \n", q.tasks, len(q.tasks), q.current)
	
	q.tasks = append(q.tasks, task)
	
	fmt.Printf(">>>加入新任务：%+v,当前队列：%+v;当前队列数量：%+v;当前任务：%+v \n", task, q.tasks, len(q.tasks), q.current)
	
	// 如果当前没有任务正在处理，则立即开始处理
	if !q.processing {
		fmt.Println("---当前没有任务正在处理，立即开始处理")
		q.processing = true
		q.current = q.tasks[0]
		q.tasks = q.tasks[1:]
		
		go q.processCurrentTask()
	}
}

func (q *LocalQueue) processCurrentTask() {
	// 发送异步请求到外部服务
	// 这里假设外部服务的URL为 "http://example.com/process"
	// 并且我们使用了一个假的处理时间来模拟异步处理
	fmt.Printf("Processing task %s...\n", q.current.id)
	time.Sleep(time.Second * 10)
	fmt.Printf("Task %s processed.\n", q.current.id)
	
	q.mutex.Lock()
	defer q.mutex.Unlock()
	
	// 处理完成后，设置当前任务为nil
	fmt.Println("<<<处理完成后，设置当前任务为nil")
	q.current = nil
	
	// 如果队列不为空，则继续处理下一个任务
	if len(q.tasks) > 0 {
		fmt.Println("===队列不为空，则继续处理下一个任务")
		q.current = q.tasks[0]
		q.tasks = q.tasks[1:]
		
		go q.processCurrentTask()
	} else {
		fmt.Println("---队列为空")
		q.processing = false
	}
}

func main() {
	q := &LocalQueue{}
	
	// 添加任务到队列中
	q.Add(&Task{id: "task1"})
	q.Add(&Task{id: "task2"})
	q.Add(&Task{id: "task3"})
	
	// 报错：fatal error: all goroutines are asleep - deadlock!，是主程序中不允许空程序运行。对于调试没有影响
	select {}
}
