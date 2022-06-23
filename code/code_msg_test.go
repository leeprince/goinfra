package code_test

import (
    "errors"
    "fmt"
    "github.com/leeprince/goinfra/code"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/24 上午12:37
 * @Desc:
 */

func TestError(t *testing.T) {
    err := code.BizErrSuccess
    
    if errors.Is(err, code.BizErr{}) {
        fmt.Println(">>>>>>>>>>errors.Is(err, code.BizErr{})")
        fmt.Println(err.Error())
        fmt.Println(err.GetCode())
        fmt.Println(err.GetMessage())
    } else {
        fmt.Println(">>>>>>>>>>not errors.Is(err, code.BizErr{})")
    }
    
    err = code.BizErrRequired
    if errors.Is(err, code.BizErr{}) {
        fmt.Println(">>>>>>>>>>errors.Is(err, code.BizErr{})")
        fmt.Println(err.Error())
        fmt.Println(err.GetCode())
        fmt.Println(err.GetMessage())
    } else {
        fmt.Println(">>>>>>>>>>not errors.Is(err, code.BizErr{})")
    }
}