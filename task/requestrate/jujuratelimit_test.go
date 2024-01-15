package requestrate

import (
	"fmt"
	"github.com/juju/ratelimit"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/4 14:26
 * @Desc:
 */
func TestJujuratelimit(t *testing.T) {
	bucket := ratelimit.NewBucket(1*time.Second, 5) // 每秒填充1个令牌，桶的容量为5

	for i := 0; i < 10; i++ {
		if bucket.TakeAvailable(1) > 0 { // 尝试从桶中取出1个令牌
			fmt.Println("Request", i, "allowed")
		} else {
			fmt.Println("Request", i, "not allowed")
		}
		time.Sleep(500 * time.Millisecond) // 每500毫秒发送一个请求
	}
}
