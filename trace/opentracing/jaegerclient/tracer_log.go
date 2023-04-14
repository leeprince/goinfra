package jaegerclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/leeprince/goinfra/plog"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/12 下午11:47
 * @Desc:   日志：Plog + 分布式链路追踪
 */

func Plog(ctx context.Context, level plog.Level, args ...interface{}) {
	plog.SetCustomerTempLoggerPackage(packageName)

	if level == plog.ErrorLevel {
		LogError(ctx, errors.New(fmt.Sprint(args...)))
	}
	plog.LogID(TraceID(ctx)).WithContext(ctx).Log(plog.PLevel(level), args...)
}

func Plogf(ctx context.Context, level plog.Level, format string, args ...interface{}) {
	plog.SetCustomerTempLoggerPackage(packageName)

	if level == plog.ErrorLevel {
		LogError(ctx, errors.New(fmt.Sprintf(format, args...)))
	}
	plog.LogID(TraceID(ctx)).WithContext(ctx).Logf(plog.PLevel(level), format, args...)
}

// --- 基于 Plog、Plogf 的实现
func PlogInfo(ctx context.Context, args ...interface{}) {
	Plog(ctx, plog.InfoLevel, args...)
}

func PlogInfof(ctx context.Context, format string, args ...interface{}) {
	Plogf(ctx, plog.InfoLevel, format, args...)
}

func PlogError(ctx context.Context, args ...interface{}) {
	Plog(ctx, plog.ErrorLevel, args...)
}

func PlogErrorf(ctx context.Context, format string, args ...interface{}) {
	Plogf(ctx, plog.ErrorLevel, format, args...)
}

// --- 基于 Plog、Plogf 的实现 -end
