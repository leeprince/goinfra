package synclock

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/4 11:02
 * @Desc:
 */

import (
	"sync"
)

type WorkerPool struct {
	concurrency int
	taskQueue   chan func()
	wg          sync.WaitGroup
}

func NewWorkerPool(concurrency int) *WorkerPool {
	return &WorkerPool{
		concurrency: concurrency,
		taskQueue:   make(chan func(), concurrency),
	}
}

func (p *WorkerPool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.worker()
	}
}

func (p *WorkerPool) Submit(task func()) {
	p.wg.Add(1)
	p.taskQueue <- task
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
}

func (p *WorkerPool) worker() {
	for task := range p.taskQueue {
		func() {
			defer p.wg.Done()
			task()
		}()
	}
}
