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
	m.mu.Unlock()
}

func main() {
	var mu MutexWithTimeout
	mu.timeout = 3 * time.Second
	
	fmt.Println("Trying to acquire the lock...")
	
	if mu.Lock() {
		fmt.Println("Lock acquired---")
		time.Sleep(5 * time.Second)
		mu.Unlock()
		fmt.Println("Lock released---")
	} else {
		fmt.Println("Failed to acquire the lock---")
	}
	
	go func() {
		for i := 0; i < 3; i++ {
			if mu.Lock() {
				fmt.Println("Lock acquired---", i)
				time.Sleep(3 * time.Second)
				mu.Unlock()
				fmt.Println("Lock released---", i)
			} else {
				fmt.Println("Failed to acquire the lock---", i)
			}
		}
		
	}()
	
	select {}
}
