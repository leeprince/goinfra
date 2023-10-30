package mutexlock

import (
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

func (m *MutexWithTimeout) Lock() bool {
	if m.Timeout == 0 {
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
	case <-time.After(m.Timeout):
		m.mu.Unlock()
		return false
	}
}

func (m *MutexWithTimeout) Unlock() {
	m.mu.Unlock()
}
