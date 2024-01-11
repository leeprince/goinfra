package synclock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/4 11:04
 * @Desc:
 */

func TestMutexWithTimeout_Lock(t *testing.T) {
	mu := NewMutexWithTimeout(time.Second * 1)

	fmt.Println("Trying to acquire the lock...")

	if mu.TryLock() {
		fmt.Println("TryLock acquired---")
		time.Sleep(4 * time.Second)
		mu.Unlock()
		fmt.Println("TryLock released---")
	} else {
		fmt.Println("Failed to acquire the lock---")
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		//time.Sleep(time.Second * 1)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			if mu.TryLock() {
				fmt.Println("TryLock acquired---", i)
				time.Sleep(time.Millisecond * 500)
				mu.Unlock()
				fmt.Println("TryLock released---", i)
			} else {
				fmt.Println("Failed to acquire the lock---", i)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("---------------")

	wg1 := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()

			if mu.TryLock() {
				fmt.Println("wg1 TryLock acquired---", i)
				mu.Unlock()
			} else {
				fmt.Println("wg1 Failed to acquire the lock---", i)
			}
		}(i)
	}
	wg1.Wait()
}
