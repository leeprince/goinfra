package requestrate

import (
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/4 13:59
 * @Desc:
 */

func TestXtimerate(t *testing.T) {
	//limiter := rate.NewLimiter(1, 5) // 每秒产生1个令牌，桶的容量为5
	limiter := rate.NewLimiter(1, 5) // 每秒产生1个令牌，桶的容量为5

	for i := 0; i < 10; i++ {
		if limiter.Allow() { // 检查当前是否允许请求
			fmt.Println("------------------------------- Request", i, "allowed")
			//fmt.Println(">", limiter)
		} else {
			fmt.Println("------------------------------- Request", i, "not allowed")
		}
		time.Sleep(500 * time.Millisecond) // 每500毫秒发送一个请求
	}
}

func TestXtimerateV1(t *testing.T) {
	//limiter := rate.NewLimiter(1, 9) // 每秒产生2个令牌，桶的容量为9
	limiter := rate.NewLimiter(5, 9) // 每秒产生2个令牌，桶的容量为9

	for i := 0; i < 10; i++ {
		go func(i int) {
			//limiter := rate.NewLimiter(1, 5) // 每秒产生1个令牌，桶的容量为5
			if limiter.Allow() { // 检查当前是否允许请求
				//fmt.Println("------------------------------- Request", i, "allowed")
				//fmt.Println(">", limiter)
			} else {
				fmt.Println("------------------------------- Request", i, "not allowed")
			}
			time.Sleep(100 * time.Millisecond) // 每500毫秒发送一个请求
		}(i)
	}
}
