package gostreaming

/*
 * @Date: 2020-07-08 18:09:25
 * @LastEditors: aiden.deng(Zhenpeng Deng)
 * @LastEditTime: 2020-07-08 21:14:06
 */

import "sync"

// WindowsCache用于在的窗口滑动中多个窗口的数据缓存与共享。Thread Safe。
type WindowsCache struct {
	m  map[string]interface{}
	mu sync.RWMutex
}

func NewWindowsCache() *WindowsCache {
	return &WindowsCache{
		m: make(map[string]interface{}),
	}
}

func (wc *WindowsCache) Get(key string) (interface{}, bool) {
	wc.mu.RLock()
	defer wc.mu.RUnlock()
	value, ok := wc.m[key]
	if !ok {
		return nil, false
	}
	return value, true
}

func (wc *WindowsCache) Set(key string, value interface{}) {
	wc.mu.Lock()
	wc.m[key] = value
	wc.mu.Unlock()
}
