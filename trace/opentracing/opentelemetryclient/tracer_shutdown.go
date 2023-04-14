package opentelemetryclient

import (
	"context"
	"github.com/leeprince/goinfra/plog"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/15 下午5:42
 * @Desc:
 */

// 当应用程序退出时，干净地关闭并刷新 opentelemetry
func Shutdown(ctx context.Context, timeout time.Duration) {
	// 关闭应用程序时不要使其挂起，设置超时时间
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if err := TracerProvider().Shutdown(timeoutCtx); err != nil {
		plog.Fatal(err)
	}
}
