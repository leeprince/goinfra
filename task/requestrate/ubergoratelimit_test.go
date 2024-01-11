package requestrate

import (
	"fmt"
	"go.uber.org/ratelimit"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/5 15:34
 * @Desc:
 */
func TestUbergoratelimit_test(t *testing.T) {
	rl := ratelimit.New(1) // 每秒10个请求
	//rl := ratelimit.New(10) // 每秒10个请求

	prev := time.Now()
	for i := 0; i < 50; i++ {
		// 调用rl.Take()，这将阻塞直到下一个请求可以进行
		now := rl.Take()
		//time.Sleep(time.Millisecond * 200)
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
