package opentracing_test_test

import (
    "fmt"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/28 上午12:23
 * @Desc:
 */

func TestRun(t *testing.T) {
    helloTo := "world"
    helloStr := fmt.Sprintf("Hello, %s!", helloTo)
    println(helloStr)
}