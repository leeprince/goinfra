package synclock

import (
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/4 11:05
 * @Desc:
 */

func TestNewWorkerPool(t *testing.T) {
	pool := NewWorkerPool(8)

	// 提交任务到工作池
	for i := 0; i < 8; i++ {
		task := func(i int) func() {
			return func() {
				// 这里是实际的任务执行代码
				println("Processing task:", i)
			}
		}(i)
		pool.Submit(task)
	}

	// 开始执行任务
	pool.Run()

	// 等待所有任务完成
	pool.Wait()
}
