package code_test

import (
    "errors"
    "fmt"
    "github.com/leeprince/goinfra/code"
    errors2 "github.com/pkg/errors"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/24 上午12:37
 * @Desc:
 */

func TestBizErr(t *testing.T) {
    fmt.Println("----------------errors.Is()/errros.As()")
    err := code.BizErrSuccess
    if errors.Is(err, &code.BizErr{}) {
        fmt.Println("BizErrSuccess>>>>>>>>>>errors.Is(&code.BizErr{})")
        fmt.Println(err.Error())
        fmt.Println(err.GetCode())
        fmt.Println(err.GetMessage())
    } else {
        fmt.Println("BizErrSuccess>>>>>>>>>>not errors.Is(&code.BizErr{})")
    }
    if errors.Is(err, code.BizErrSuccess) {
        fmt.Println("BizErrSuccess>>>>>>>>>>errors.Is(err, code.BizErrSuccess)")
        fmt.Println(err.Error())
        fmt.Println(err.GetCode())
        fmt.Println(err.GetMessage())
    } else {
        fmt.Println("BizErrSuccess>>>>>>>>>>not errors.Is(error, code.BizErrSuccess)")
    }
    if errors.As(err, &code.BizErr{}) {
        fmt.Println("BizErrSuccess>>>>>>>>>>errors.As(error, &code.BizErr{})")
        fmt.Println(err.Error())
        fmt.Println(err.GetCode())
        fmt.Println(err.GetMessage())
    } else {
        fmt.Println("BizErrSuccess>>>>>>>>>>not errors.As(error, &code.BizErr{})")
    }
    if errors.As(err, &code.BizErrSuccess) {
        fmt.Println("BizErrSuccess>>>>>>>>>>errors.As(error, code.BizErrSuccess)")
        fmt.Println(err.Error())
        fmt.Println(err.GetCode())
        fmt.Println(err.GetMessage())
    } else {
        fmt.Println("BizErrSuccess>>>>>>>>>>not errors.As(error, code.BizErrSuccess)")
    }
    
    // ---
    fmt.Println("---")
    err = code.BizErrRequired
    if errors.Is(err, &code.BizErr{}) {
        fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, &code.BizErr{})")
        fmt.Println(err.Error())
        fmt.Println(err.GetCode())
        fmt.Println(err.GetMessage())
    } else {
        fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, &code.BizErr{})")
    }
    if errors.Is(err, code.BizErrRequired) {
        fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, code.BizErrRequired)")
        fmt.Println(err.Error())
        fmt.Println(err.GetCode())
        fmt.Println(err.GetMessage())
    } else {
        fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, code.BizErrRequired)")
    }
    if errors.As(err, &code.BizErr{}) {
        fmt.Println("BizErrRequired>>>>>>>>>>errors.As(error, &code.BizErr{})")
        fmt.Println(err.Error())
        fmt.Println(err.GetCode())
        fmt.Println(err.GetMessage())
    } else {
        fmt.Println("BizErrRequired>>>>>>>>>>not errors.As(error, &code.BizErr{})")
    }
    if errors.As(err, &code.BizErrRequired) {
        fmt.Println("BizErrRequired>>>>>>>>>>errors.As(error, code.BizErrRequired)")
        fmt.Println(err.Error())
        fmt.Println(err.GetCode())
        fmt.Println(err.GetMessage())
    } else {
        fmt.Println("BizErrRequired>>>>>>>>>>not errors.As(error, code.BizErrRequired)")
    }
    fmt.Println("----------------errors.Is()/errros.As()-end")
}

func TestErrorNew(t *testing.T) {
    fmt.Println("----------------errors.New")
    err := errors.New("ddd")
    if bizErr, ok := err.(code.BizErr); ok {
        fmt.Println("errors.New>>>>>>>>>>err.(code.BizErr)")
        fmt.Println(bizErr.Error())
        fmt.Println(bizErr.GetCode())
        fmt.Println(bizErr.GetMessage())
    } else {
        fmt.Println("errors.New>>>>>>>>>>not err.(code.BizErr)")
    }
    if errors.Is(err, &code.BizErr{}) {
        fmt.Println("errors.New>>>>>>>>>>errors.Is(error, &code.BizErr{})")
        fmt.Println(err.Error())
    } else {
        fmt.Println("errors.New>>>>>>>>>>not errors.Is(error, &code.BizErr{})")
    }
    if errors.As(err, &code.BizErr{}) {
        fmt.Println("errors.New>>>>>>>>>>errors.As(error, &code.BizErr{})")
        fmt.Println(err.Error())
    } else {
        fmt.Println("errors.New>>>>>>>>>>not errors.As(error, &code.BizErr{})")
    }
    
    // ---
    fmt.Println("---")
    err = code.BizErrSuccess
    if bizErr, ok := err.(code.BizErr); ok {
        fmt.Println("BizErrSuccess>>>>>>>>>>err.(code.BizErr)")
        fmt.Println(bizErr.Error())
        fmt.Println(bizErr.GetCode())
        fmt.Println(bizErr.GetMessage())
    } else {
        fmt.Println("BizErrSuccess>>>>>>>>>>not err.(code.BizErr)")
    }
    if errors.Is(err, &code.BizErr{}) {
        fmt.Println("BizErrSuccess>>>>>>>>>>errors.Is(error, &code.BizErr{})")
        fmt.Println(err.Error())
    } else {
        fmt.Println("BizErrSuccess>>>>>>>>>>not errors.Is(error, &code.BizErr{})")
    }
    if errors.As(err, &code.BizErr{}) {
        fmt.Println("BizErrSuccess>>>>>>>>>>errors.As(error, &code.BizErr{})")
        fmt.Println(err.Error())
    } else {
        fmt.Println("BizErrSuccess>>>>>>>>>>not errors.As(error, &code.BizErr{})")
    }
    fmt.Println("----------------errors.New-end")
}

func TestFmtErrorfWrap(t *testing.T) {
    err := errors.New(">>>>01")
    fmt.Println("0001:", err)
    err = fmt.Errorf("MyErrorf001:%w", err)
    fmt.Println("0002:", err)
    err = fmt.Errorf("MyErrorf002:%w", err)
    fmt.Println("0003:", err)
    
    // ---
    fmt.Println("=========")
    err = errors.Unwrap(err)
    fmt.Println("0004:", err)
    err = errors.Unwrap(err)
    fmt.Println("0005:", err)
    err = errors.Unwrap(err)
    fmt.Println("0006:", err)
    err = errors.Unwrap(err)
    fmt.Println("0007:", err)
}

func TestErrorWrap(t *testing.T) {
    err := errors.New(">>>>01")
    fmt.Println("0001:", err)
    err = errors2.Wrap(err, "MyErrorf001")
    fmt.Println("0002:", err)
    err = errors2.Wrap(err, "MyErrorf002")
    fmt.Println("0003:", err)
    
    // ---
    fmt.Println("=========")
    err = errors.Unwrap(err)
    fmt.Println("0004:", err)
    err = errors.Unwrap(err)
    fmt.Println("0005:", err)
    err = errors.Unwrap(err)
    fmt.Println("0006:", err)
    err = errors.Unwrap(err)
    fmt.Println("0007:", err)
}

func TestBizErrFmtErrorfWrap(t *testing.T) {
    // --- 定义方式
    // ---1
    // err := code.BizErr{} // 测试
    // ---2
    err := code.BizErrRequired
    
    fmt.Println(err)
    fmt.Println(err.GetError())
    
    fmt.Println("=========1")
    
    err1 := errors.New("my error 0001")
    err = err.WithError(err1)
    fmt.Println(err)
    fmt.Println(err.GetError())
    
    fmt.Println("=========2")
    
    err2 := errors.New("my error 0002")
    err = err.WithError(err2)
    fmt.Println(err)
    fmt.Println(err.GetError())
}

func TestBizErrFmtErrorfWrapAssert(t *testing.T) {
    var err error
    err = code.BizErrRequired
    
    fmt.Println(err)
    if errors.Is(err, &code.BizErr{}) {
        fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, &code.BizErr{})")
        fmt.Println(err.Error())
    } else {
        fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, &code.BizErr{})")
    }
    if errors.Is(err, code.BizErrRequired) {
        fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, code.BizErrRequired)")
        fmt.Println(err.Error())
    } else {
        fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, code.BizErrRequired)")
    }
    if errors.As(err, &code.BizErr{}) {
        fmt.Println("BizErrRequired>>>>>>>>>>errors.As(error, &code.BizErr{})")
        fmt.Println(err.Error())
    } else {
        fmt.Println("BizErrRequired>>>>>>>>>>not errors.As(error, &code.BizErr{})")
    }
    
    fmt.Println("=========1")
    
    if bizErr, ok := err.(code.BizErr); ok {
        fmt.Println("BizErrRequired>>>>>>>>>>err.(code.BizErr)")
        
        fmt.Println(bizErr.Error())
        fmt.Println(bizErr.GetError())
        err1 := errors.New("my error 0001")
        bizErr = bizErr.WithError(err1)
        fmt.Println(bizErr)
        err = bizErr.GetError()
        fmt.Println("===", err)
        // ---
        fmt.Println("===---")
        err = errors.Unwrap(err)
        fmt.Println("0004:", err)
        err = errors.Unwrap(err)
        fmt.Println("0005:", err)
        err = errors.Unwrap(err)
        fmt.Println("0006:", err)
        err = errors.Unwrap(err)
        fmt.Println("0007:", err)
    } else {
        fmt.Println("BizErrRequired>>>>>>>>>>not err.(code.BizErr)")
    }
}