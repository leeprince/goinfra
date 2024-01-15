package requestrate

import (
	"fmt"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"net/http"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/4 14:34
 * @Desc:
 */

func TestDidiptollbooth(t *testing.T) {
	// 创建一个限流器，每秒只允许1个请求，最大并发请求数为5
	lmt := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{
		DefaultExpirationTTL: 1 * time.Hour,
		ExpireJobInterval:    12 * time.Hour,
	})

	// 设置自定义错误消息
	lmt.SetMessage("You have reached maximum request limit.")

	// 设置自定义错误消息内容类型
	lmt.SetMessageContentType("text/plain; charset=utf-8")

	// 设置自定义错误响应函数
	lmt.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("A request was rejected")
	})

	// 使用限流器包装HTTP处理函数
	http.Handle("/", tollbooth.LimitFuncHandler(lmt, helloWorld))

	// 启动HTTP服务器
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
