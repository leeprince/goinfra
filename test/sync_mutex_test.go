package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/9 17:57
 * @Desc:
 */

func TestSyncMutex(t *testing.T) {
	var syncMutex sync.Mutex
	
	syncMutex.Lock()
	fmt.Println("1 Lock")
	syncMutex.Unlock()
	fmt.Println("1 Unlock")
	
	// 未加锁情况下，解锁会抛异常
	go func() {
		time.Sleep(time.Second * 1)
		syncMutex.Lock()
		fmt.Println("2 Lock")
	}()
	go func() {
		syncMutex.Unlock()
		fmt.Println("2 Unlock")
	}()
	
	select {}
}

func TestSyncMutexUnlock(t *testing.T) {
	var syncMutexUnlock sync.Mutex
	
	syncMutexUnlock.Lock()
	fmt.Println("1 Lock")
	syncMutexUnlock.Unlock()
	fmt.Println("1 Unlock")
	
	syncMutexUnlock.Unlock()
	fmt.Println("1 Unlock")
	syncMutexUnlock.Lock()
	fmt.Println("1 Lock")
	
}

func TestSyncMutexTimeout(t *testing.T) {
	var mutex sync.Mutex
	
	mutex.Lock()
	fmt.Println("1 Lock")
	
	isUnLockSyncMutexTimeout := false
	go func() {
		select {
		case <-time.After(time.Second * 2):
			if !isUnLockSyncMutexTimeout {
				isUnLockSyncMutexTimeout = true
				mutex.Unlock()
				fmt.Println("1 After Unlock")
			}
			
		}
	}()
	time.Sleep(time.Second * 3)
	
	if !isUnLockSyncMutexTimeout {
		isUnLockSyncMutexTimeout = true
		mutex.Unlock()
		fmt.Println("1 Unlock")
	}
	
	select {}
}

var ch chan int
var mutex sync.Mutex
var mutexOfValue int
var unlockMutex bool

func TestSyncMutexTimeoutChannel(t *testing.T) {
	ch = make(chan int, 10)
	
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		
		// 应使用通道进行共享变量，而不是共享变量来通信
		// SyncMutexTimeoutChannel1(i) // 出现各种异常：可能关闭到其他刚设置的锁
		SyncMutexTimeoutChannel2(i) // 正确
	}
	
	go func() {
		for {
			select {
			case i := <-ch:
				go func(i int) {
					fmt.Println("<-ch:", i)
					time.Sleep(time.Second * 10)
				}(i)
			}
		}
	}()
	
	select {}
}

// 出现各种异常：应使用通道进行共享变量，而不是共享变量来通信
func SyncMutexTimeoutChannel1(i int) {
	go func(i int) {
		mutex.Lock()
		mutexOfValue = i
		unlockMutex = false
		fmt.Println("Lock:", i)
		
		go func(i int) {
			select {
			case <-time.After(time.Second * 2):
				if !unlockMutex {
					unlockMutex = true
					mutex.Unlock()
					fmt.Printf("time.After Unlock i:%d \n", i)
				}
			}
		}(i)
	}(i)
	
	go func(i int) {
		time.Sleep(time.Second * 1)
		
		for ii := 0; ii < 10; ii++ {
			fmt.Printf("i:%d---ii:%d---mutexOfValue:%d \n", i, ii, mutexOfValue)
			
			time.Sleep(time.Second * 1)
			
			if ii != mutexOfValue {
				fmt.Printf("ii != mutexOfValue i:%d---ii:%d---mutexOfValue:%d \n", i, ii, mutexOfValue)
				continue
			}
			fmt.Printf("ii == mutexOfValue i:%d---ii:%d---mutexOfValue:%d \n", i, ii, mutexOfValue)
			
			defer func() {
				if !unlockMutex {
					fmt.Println("Unlock:", mutexOfValue)
					
					mutex.Unlock()
					unlockMutex = true
				}
				
			}()
			ch <- mutexOfValue
			break
		}
	}(i)
}

func SyncMutexTimeoutChannel2(i int) {
	go func(i int) {
		mutex.Lock()
		mutexOfValue = i
		fmt.Println("Lock:", i)
	}(i)
	
	go func(i int) {
		time.Sleep(time.Second * 1)
		
		for ii := 0; ii < 10; ii++ {
			fmt.Printf("i:%d---ii:%d---mutexOfValue:%d \n", i, ii, mutexOfValue)
			
			time.Sleep(time.Second * 1)
			
			if ii != mutexOfValue {
				fmt.Printf("ii != mutexOfValue i:%d---ii:%d---mutexOfValue:%d \n", i, ii, mutexOfValue)
				continue
			}
			fmt.Printf("ii == mutexOfValue i:%d---ii:%d---mutexOfValue:%d \n", i, ii, mutexOfValue)
			
			defer func() {
				fmt.Println("Unlock:", mutexOfValue)
				mutex.Unlock()
			}()
			ch <- mutexOfValue
			break
		}
	}(i)
}
