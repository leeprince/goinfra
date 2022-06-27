package kv_test

import (
    "fmt"
    "github.com/leeprince/goinfra/kv"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/27 下午3:42
 * @Desc:
 */

func TestNewKInt32VString(t *testing.T) {
    tests := []struct {
        name string
        args kv.KInt32VString
        want kv.KInt32VString
    }{
        {
            args: kv.NewKInt32VString(0, "int32Message-0"),
        },
        {
            args: kv.NewKInt32VString(10000, "int32Message-10000"),
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var k int32
            var v string
            
            k = tt.args.Key()
            fmt.Println(k)
            
            v = tt.args.Value()
            fmt.Println(v)
        })
    }
}
