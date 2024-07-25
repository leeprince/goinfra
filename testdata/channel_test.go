package testdata

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/2 14:26
 * @Desc:
 */

func TestChannel(t *testing.T) {
	var ch1 chan int
	ch1 = make(chan int, 1)
	
	go func() {
		for {
			select {
			case i, ok := <-ch1:
				fmt.Printf("<-ch1 ok:%v; i:%v \n", ok, i)
			}
			
		}
	}()
	
	for i := 0; i < 5; i++ {
		ch1 <- i
	}
}

func TestChannelClose(t *testing.T) {
	var ch1 chan int
	ch1 = make(chan int, 1)
	
	go func() {
		for {
			select {
			case i, ok := <-ch1:
				fmt.Printf("<-ch1 ok:%v; i:%v \n", ok, i)
				if !ok {
					fmt.Println("<-ch1 !ok return for")
					return
				}
				fmt.Println("<-ch1 ok--------")
			}
			
		}
	}()
	
	for i := 0; i < 5; i++ {
		if i >= 3 {
			close(ch1)
			break
		}
		ch1 <- i
		
	}
	
	select {}
}

func TestChannelCloseRead(t *testing.T) {
	// 测试场景一
	stop := make(chan int) // 因为这是无缓冲，所以如果在发送数据到通道，再关闭通道之前，在无缓冲的通道上已完成同步堵塞的接收。
	// 测试场景二
	// stop := make(chan int, 2) // 关闭有缓冲的通道之后，如果通道中还有未接收完成，则val, ok := <-stop 返回value=通道中的值,并且ok=true
	
	go func() {
		for i := 0; i < 5; i++ {
			stop <- i
		}
		close(stop)
		fmt.Println("发送任务逻辑已经关闭通道，等待接收完成后，val, ok := <-stop 返回value=零值并且ok=false")
	}()
	
	for {
		val, ok := <-stop
		if !ok {
			break // 通道已关闭且无数据，退出循环
		}
		fmt.Println(val)
		time.Sleep(time.Second * 2)
	}
	
	fmt.Println("所有数据接收完毕，通道已完全清空。")
}
