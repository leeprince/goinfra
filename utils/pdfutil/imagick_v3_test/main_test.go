package main

import (
	"gopkg.in/gographics/imagick.v3/imagick"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/20 13:54
 * @Desc:
 */

func TestCustomerPdfToImagesByImagickV3(t *testing.T) {
	CustomerPdfToImagesByImagickV3()
}

func BenchmarkCustomerPdfToImagesByImagickV3(b *testing.B) {
	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()

	for i := 0; i < b.N; i++ {
		CustomerPdfToImagesByImagickV3()
	}
}
