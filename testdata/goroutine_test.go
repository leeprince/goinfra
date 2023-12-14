package testdata

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/7/21 14:49
 * @Desc:
 */

// 启动多个协程，且协程中通过通道进行通信
func TestGoroutineWaitAndChannel(t *testing.T) {
	goroutinNum := 5

	var wg sync.WaitGroup
	completeCh := make(chan int, goroutinNum)
	for i := 0; i < goroutinNum; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			completeCh <- i
			fmt.Println("time sleep ...", i)
			time.Sleep(time.Second * 5)
		}()
	}
	wg.Wait()

	// 获取协程中通过通道通知的结果方式一
	// 必须等到所有协程执行结束后关闭通道，且必须在for之前关闭。
	// 	1. 因为for 必须等到所有协程执行结束才开始 for completeCh 的
	// 	2. 需要关闭completeCh是因为要避免循环completeCh结束后出现`fatal error: all goroutines are asleep - deadlock!`。关闭后可以在for completeCh 之后结束运行
	close(completeCh)
	for ch := range completeCh {
		fmt.Println("close ch:", ch)
	}

	// 获取协程中通过通道通知的结果方式二
	/*// 启动一个协程来获取通道中数据
	go func() {
		for ch := range completeCh {
			fmt.Println("go ch:", ch)
		}
	}()*/
}
