package concurrency

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func init() {
	println("using MAXPROC")
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
}

// TestNewPool test new goroutine pool
func TestNewPool(t *testing.T) {
	pool := NewPool(1000, 10000)
	defer pool.Release()

	iterations := 1000
	var counter uint64

	wg := sync.WaitGroup{}
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		arg := uint64(1)
		job := func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			atomic.AddUint64(&counter, arg)
		}

		pool.JobQueue <- job
	}
	wg.Wait()

	counterFinal := atomic.LoadUint64(&counter)
	if uint64(iterations) != counterFinal {
		t.Errorf("iterations %v is not equal counterFinal %v", iterations, counterFinal)
	}
}

// 测试通过协程池执行任务，并将结果全部先保留到通道
func TestWait(t *testing.T) {
	reqs := []int{}
	type data struct{}

	taskNum := len(reqs)
	jobNum := 5
	if taskNum < jobNum {
		jobNum = taskNum
	}

	pool := NewPool(jobNum, taskNum)
	defer pool.Release()
	errChan := make(chan error, taskNum)
	dataChan := make(chan []*data, taskNum)

	var wg sync.WaitGroup
	wg.Add(taskNum)

	for _, v := range reqs {
		childReq := v
		f := func() {
			defer wg.Done()

			// 业务处理流程
			dataItem, err := func(childReq int) ([]*data, error) {
				return []*data{}, nil
			}(childReq)

			if err != nil {
				errChan <- err
			} else {
				dataChan <- dataItem
			}
		}
		pool.JobQueue <- f
	}

	wg.Wait()
	// 主动关闭chan
	close(errChan)
	close(dataChan)
	//全部都有问题才返回
	if taskNum == len(errChan) {
		return
	}

	return
}
