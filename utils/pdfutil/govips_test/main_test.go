package main

import (
	"github.com/davidbyttow/govips/v2/vips"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/20 10:50
 * @Desc:
 */

func TestCustomerPdfToImagesByGovips(t *testing.T) {
	CustomerPdfToImagesByGovipsOfPerformanceTest()
}

func BenchmarkCustomerPdfToImagesByGovips(b *testing.B) {
	vips.LoggingSettings(myLogger, vips.LogLevelError)
	vips.Startup(nil)
	//defer vips.Shutdown() // 基准测试需要注释掉，否则报错。main.go 中的性能测试则不需要

	for i := 0; i < b.N; i++ {
		CustomerPdfToImagesByGovipsOfPerformanceTest()
	}
}
