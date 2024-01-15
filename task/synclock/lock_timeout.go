package synclock

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/10 13:06
 * @Desc:
 */

type MutexWithTimeout struct {
	mu      sync.Mutex
	Timeout time.Duration
}

func NewMutexWithTimeout(timeout time.Duration) *MutexWithTimeout {
	return &MutexWithTimeout{
		mu:      sync.Mutex{},
		Timeout: timeout,
	}
}

func (m *MutexWithTimeout) TryLock() bool {
	ch := make(chan struct{}, 1)
	go func() {
		m.mu.Lock()
		ch <- struct{}{}
	}()

	for {
		select {
		case <-ch:
			return true
		case <-time.After(m.Timeout):
			fmt.Println("MutexWithTimeout Timeout:", m.Timeout)
			return false
		}
	}

}

func (m *MutexWithTimeout) Unlock() {
	m.mu.Unlock()
}
