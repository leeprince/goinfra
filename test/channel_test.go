package test

import (
	"fmt"
	"testing"
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
