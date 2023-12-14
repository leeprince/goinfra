package testdata

import (
	"fmt"
	"sync"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/8 12:58
 * @Desc:
 */

func TestSyncMap(t *testing.T) {
	var wg sync.WaitGroup
	var m sync.Map

	// 并发写入 Map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i*10)
			fmt.Println("> Store")
		}(i)
	}
	wg.Wait()

	// 并发读取 Map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if v, ok := m.Load(i); ok {
				fmt.Printf("key=%v, value=%v\n", i, v)
			}
		}(i)
	}
	wg.Wait()
}

func TestSyncMapDelete(t *testing.T) {
	var m sync.Map

	// 添加键值对到 Map
	m.Store("key1", "value1")
	m.Store("key2", "value2")
	m.Store("key3", "value3")

	// 删除键值对
	m.Delete("key2")

	// 并发读取 Map
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key=%v, value=%v\n", key, value)
		return true
	})

	value, ok := m.LoadAndDelete("key2")
	if ok {
		fmt.Println("Deleted key2 ok value:", value)
	} else {
		fmt.Println("Deleted key2 !ok value:", value)
	}

	value, ok = m.LoadAndDelete("key3")
	if ok {
		fmt.Println("Deleted key3 ok value:", value)
	} else {
		fmt.Println("Deleted key3 !ok value:", value)
	}
}
