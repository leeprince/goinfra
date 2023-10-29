package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/10 12:52
 * @Desc:
 */

type MutexWithTimeout struct {
	mu      sync.Mutex
	timeout time.Duration
}

func (m *MutexWithTimeout) Lock() bool {
	if m.timeout == 0 {
		m.mu.Lock()
		return true
	}
	
	ch := make(chan bool, 1)
	go func() {
		m.mu.Lock()
		ch <- true
	}()
	
	select {
	case <-ch:
		return true
	case <-time.After(m.timeout):
		fmt.Println("> time.After Unlock")
		m.mu.Unlock()
		return false
	}
}

func (m *MutexWithTimeout) Unlock() {
	defer func() {
		if r := recover(); r != nil {
			// 忽略解锁错误
			fmt.Println("忽略解锁错误")
			// fmt.Println("Recovered from panic:", r)
		}
	}()
	m.mu.Unlock()
}

func main() {
	var mu MutexWithTimeout
	mu.timeout = 3 * time.Second
	
	fmt.Println("Trying to acquire the lock...")
	
	if mu.Lock() {
		fmt.Println("Lock acquired---")
		time.Sleep(4 * time.Second)
		mu.Unlock()
		fmt.Println("Lock released---")
	} else {
		fmt.Println("Failed to acquire the lock---")
	}
	
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 1)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			
			if mu.Lock() {
				fmt.Println("Lock acquired---", i)
				time.Sleep(5 * time.Second)
				mu.Unlock()
				fmt.Println("Lock released---", i)
			} else {
				fmt.Println("Failed to acquire the lock---", i)
			}
		}(i)
	}
	wg.Wait()
	
	wg1 := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			
			if mu.Lock() {
				fmt.Println("wg1 Lock acquired---", i)
			} else {
				fmt.Println("wg1 Failed to acquire the lock---", i)
			}
		}(i)
		
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			fmt.Println("wg1 Lock released...", i)
			mu.Unlock()
			fmt.Println("wg1 Lock released---", i)
		}(i)
	}
	wg1.Wait()
	
}
