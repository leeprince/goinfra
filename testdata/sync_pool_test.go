package testdata

import (
	"fmt"
	"sync"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/8 12:44
 * @Desc:
 */

type Object struct {
	ID int
}

func TestSyncPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return &Object{}
		},
	}

	// 从对象池获取对象
	obj1 := pool.Get().(*Object)
	obj1.ID = 1
	fmt.Println("TestSyncPool 1", obj1)

	// 将对象放回对象池
	pool.Put(obj1)

	// 从对象池获取对象
	obj2 := pool.Get().(*Object)
	fmt.Println("TestSyncPool 2", obj2)

	// 清空对象池
	pool.New = nil
	pool.Put(&Object{ID: 2})
	pool.Put(&Object{ID: 3})
	pool.Put(&Object{ID: 4})
	pool.Put(&Object{ID: 5})
	pool.Put(&Object{ID: 6})

	// 对象池中的对象可能会被清除，所以需要检查是否为 nil
	for {
		obj := pool.Get()
		if obj == nil {
			break
		}
		fmt.Println("TestSyncPool for", obj)
	}
}
