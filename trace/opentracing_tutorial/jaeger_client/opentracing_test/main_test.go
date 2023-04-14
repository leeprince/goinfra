package main

import (
	"context"
	"github.com/leeprince/goinfra/trace/opentracing/jaeger_client"
	"os"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/14 10:09
 * @Desc:
 */

// 初始化日志
func TestMain(m *testing.M) {
	// 初始化日志
	initLog()

	// 该包下的所有TextXxx的执行入口。默认
	os.Exit(m.Run())
}

func TestInitJaeger(t *testing.T) {
	// 初始化链路跟踪
	initTrace("TestInitJaeger")
	// 注意：这里一定要关闭Trace
	defer jaegerclient.Close()

	// 开始 span 链路跟踪
	ctx := context.Background()
	spanCtx := jaegerclient.StartSpan(ctx, "TestInitJaeger#StartSpan")
	// 注意：这里一定要结束 span
	defer jaegerclient.Finish(spanCtx)
}

func TestSpanTraceLog(t *testing.T) {
	// 初始化链路跟踪
	initTrace("TestSpanTraceLog")
	// 注意：这里一定要关闭Trace
	defer jaegerclient.Close()

	// 开始 span 链路跟踪
	ctx := context.Background()
	spanCtx := jaegerclient.StartSpan(ctx, "TestSpanTraceLog#StartSpan")
	// 注意：这里一定要结束 span
	defer jaegerclient.Finish(spanCtx)

	// 通过 span 继续链路跟踪
	spanTraceLog(spanCtx)
}

func TestLocalRemoteTraceRPCFormatter(t *testing.T) {
	// 初始化链路跟踪
	initTrace("TestLocalRemoteTraceRPCFormatter")
	// 注意：这里一定要关闭Trace
	defer jaegerclient.Close()

	// 开始 span 链路跟踪
	ctx := context.Background()
	spanCtx := jaegerclient.StartSpan(ctx, "TestLocalRemoteTraceRPCFormatter#StartSpan")
	// 注意：这里一定要结束 span
	defer jaegerclient.Finish(spanCtx)

	// 通过 span 继续链路跟踪
	localRemoteTraceRPCFormatter(spanCtx)
}
func TestLocalRemoteTraceRPCFormatterV1(t *testing.T) {
	// 初始化链路跟踪
	initTrace("TestLocalRemoteTraceRPCFormatterV1")
	// 注意：这里一定要关闭Trace
	defer jaegerclient.Close()

	// 开始 span 链路跟踪
	ctx := context.Background()
	spanCtx := jaegerclient.StartSpan(ctx, "TestLocalRemoteTraceRPCFormatterV1#StartSpan")
	// 注意：这里一定要结束 span
	defer jaegerclient.Finish(spanCtx)

	// 通过 span 继续链路跟踪
	localRemoteTraceRPCFormatterV1(spanCtx)
}
