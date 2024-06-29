package perror_test

import (
	"errors"
	"fmt"
	"github.com/leeprince/goinfra/perror"
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
	err := perror.Success
	if errors.Is(err, &perror.BizErr{}) {
		fmt.Println("Success>>>>>>>>>>errors.Is(&perror.BizErr{})")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("Success>>>>>>>>>>not errors.Is(&perror.BizErr{})")
	}
	if errors.Is(err, perror.Success) {
		fmt.Println("Success>>>>>>>>>>errors.Is(err, code.Success)")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("Success>>>>>>>>>>not errors.Is(error, code.Success)")
	}
	if errors.As(err, &perror.BizErr{}) {
		fmt.Println("Success>>>>>>>>>>errors.As(error, &perror.BizErr{})")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("Success>>>>>>>>>>not errors.As(error, &perror.BizErr{})")
	}
	if errors.As(err, &perror.Success) {
		fmt.Println("Success>>>>>>>>>>errors.As(error, code.Success)")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("Success>>>>>>>>>>not errors.As(error, code.Success)")
	}
	
	// ---
	fmt.Println("---")
	err = perror.BizErrRequired
	if errors.Is(err, &perror.BizErr{}) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, &perror.BizErr{})")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, &perror.BizErr{})")
	}
	if errors.Is(err, perror.BizErrRequired) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, perror.BizErrRequired)")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, perror.BizErrRequired)")
	}
	if errors.As(err, &perror.BizErr{}) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.As(error, &perror.BizErr{})")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.As(error, &perror.BizErr{})")
	}
	if errors.As(err, &perror.BizErrRequired) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.As(error, perror.BizErrRequired)")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.As(error, perror.BizErrRequired)")
	}
	
	// ---
	fmt.Println("---重新设置Message")
	err = perror.BizErrRequired.SetMessage("重新设置Message")
	if errors.Is(err, &perror.BizErr{}) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, &perror.BizErr{})")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, &perror.BizErr{})")
	}
	if errors.Is(err, perror.BizErrRequired) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, perror.BizErrRequired)")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, perror.BizErrRequired)")
	}
	if errors.As(err, &perror.BizErr{}) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.As(error, &perror.BizErr{})")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.As(error, &perror.BizErr{})")
	}
	if errors.As(err, &perror.BizErrRequired) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.As(error, perror.BizErrRequired)")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.As(error, perror.BizErrRequired)")
	}
	if errors.As(err, &perror.Success) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.As(error, perror.Success)")
		fmt.Println(err)
		fmt.Println(err.GetCode())
		fmt.Println(err.GetMessage())
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.As(error, perror.Success)")
	}
	fmt.Println("----------------errors.Is()/errros.As()-end")
}

func TestErrorNew(t *testing.T) {
	fmt.Println("----------------errors.New")
	err := errors.New("ddd")
	if bizErr, ok := err.(perror.BizErr); ok {
		fmt.Println("errors.New>>>>>>>>>>err.(perror.BizErr)")
		fmt.Println(bizErr)
		fmt.Println(bizErr.GetCode())
		fmt.Println(bizErr.GetMessage())
	} else {
		fmt.Println("errors.New>>>>>>>>>>not err.(perror.BizErr)")
	}
	
	if bizErr, ok := err.(*perror.BizErr); ok {
		fmt.Println("errors.New>>>>>>>>>>err.(*perror.BizErr)")
		fmt.Println(bizErr)
		fmt.Println(bizErr.GetCode())
		fmt.Println(bizErr.GetMessage())
	} else {
		fmt.Println("errors.New>>>>>>>>>>not err.(*perror.BizErr)")
	}
	
	if errors.Is(err, &perror.BizErr{}) {
		fmt.Println("errors.New>>>>>>>>>>errors.Is(error, &perror.BizErr{})")
		fmt.Println(err)
	} else {
		fmt.Println("errors.New>>>>>>>>>>not errors.Is(error, &perror.BizErr{})")
	}
	
	if errors.As(err, &perror.BizErr{}) {
		fmt.Println("errors.New>>>>>>>>>>errors.As(error, &perror.BizErr{})")
		fmt.Println(err)
	} else {
		fmt.Println("errors.New>>>>>>>>>>not errors.As(error, &perror.BizErr{})")
	}
	
	// ---
	fmt.Println("---")
	err = perror.Success
	if bizErr, ok := err.(perror.BizErr); ok {
		fmt.Println("Success>>>>>>>>>>err.(perror.BizErr)")
		fmt.Println(bizErr)
		fmt.Println(bizErr.GetCode())
		fmt.Println(bizErr.GetMessage())
	} else {
		fmt.Println("Success>>>>>>>>>>not err.(perror.BizErr)")
	}
	
	if bizErr, ok := err.(*perror.BizErr); ok {
		fmt.Println("errors.New>>>>>>>>>>err.(*perror.BizErr)")
		fmt.Println(bizErr)
		fmt.Println(bizErr.GetCode())
		fmt.Println(bizErr.GetMessage())
	} else {
		fmt.Println("errors.New>>>>>>>>>>not err.(*perror.BizErr)")
	}
	
	if errors.Is(err, &perror.BizErr{}) {
		fmt.Println("Success>>>>>>>>>>errors.Is(error, &perror.BizErr{})")
		fmt.Println(err)
	} else {
		fmt.Println("Success>>>>>>>>>>not errors.Is(error, &perror.BizErr{})")
	}
	if errors.Is(err, perror.BizErr{}) {
		fmt.Println("Success>>>>>>>>>>errors.Is(error, perror.BizErr{})")
		fmt.Println(err)
	} else {
		fmt.Println("Success>>>>>>>>>>not errors.Is(error, perror.BizErr{})")
	}
	
	// perror.BizErr{} 必须是指针类型，否则报错：must be a pointer to an interface or to a type implementing the interface
	if errors.As(err, &perror.BizErr{}) {
		fmt.Println("Success>>>>>>>>>>errors.As(error, &perror.BizErr{})")
		fmt.Println(err)
	} else {
		fmt.Println("Success>>>>>>>>>>not errors.As(error, &perror.BizErr{})")
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
	err = errors.Unwrap(err)
	fmt.Println("0008:", err)
}

func TestBizErrFmtErrorfWrap(t *testing.T) {
	// --- 定义方式
	// ---1
	// err := perror.BizErr{} // 测试
	// ---2
	err := perror.BizErrRequired
	
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

func TestBizErrFmtErrorfWithErrorAssert(t *testing.T) {
	err := perror.BizErrRequired
	
	fmt.Println(err)
	if errors.Is(err, &perror.BizErr{}) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, &perror.BizErr{})")
		fmt.Println(err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, &perror.BizErr{})")
	}
	if errors.Is(err, perror.BizErrRequired) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, perror.BizErrRequired)")
		fmt.Println(err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, perror.BizErrRequired)")
	}
	if errors.As(err, &perror.BizErr{}) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.As(error, &perror.BizErr{})")
		fmt.Println(err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.As(error, &perror.BizErr{})")
	}
	
	fmt.Println("========")
	err.WithError(perror.BizErrFormatConvert)
	fmt.Println(err)
	if errors.Is(err, &perror.BizErr{}) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, &perror.BizErr{})")
		fmt.Println(err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, &perror.BizErr{})")
	}
	if errors.Is(err, perror.BizErrRequired) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, perror.BizErrRequired)")
		fmt.Println(err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, perror.BizErrRequired)")
	}
	if errors.Is(err, perror.BizErrFormatConvert) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, perror.BizErrFormatConvert)")
		fmt.Println(err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, perror.BizErrFormatConvert)")
	}
	if errors.As(err, &perror.BizErr{}) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.As(error, &perror.BizErr{})")
		fmt.Println(err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.As(error, &perror.BizErr{})")
	}
}

func TestBizErrFmtErrorfWrapAssert(t *testing.T) {
	var err error
	err = perror.BizErrRequired
	
	fmt.Println(err)
	if errors.Is(err, &perror.BizErr{}) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, &perror.BizErr{})")
		fmt.Println(err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, &perror.BizErr{})")
	}
	if errors.Is(err, perror.BizErrRequired) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.Is(error, perror.BizErrRequired)")
		fmt.Println(err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.Is(error, perror.BizErrRequired)")
	}
	if errors.As(err, &perror.BizErr{}) {
		fmt.Println("BizErrRequired>>>>>>>>>>errors.As(error, &perror.BizErr{})")
		fmt.Println(err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not errors.As(error, &perror.BizErr{})")
	}
	
	fmt.Println()
	fmt.Println("=========1")
	
	if bizErr, ok := err.(perror.BizErr); ok {
		fmt.Println("BizErrRequired>>>>>>>>>>err.(perror.BizErr)")
		
		fmt.Println(bizErr)
		fmt.Println(bizErr.GetError())
		err1 := errors.New("my error 0001")
		bizErr = bizErr.WithError(err1)
		fmt.Println(bizErr)
		err = bizErr.GetError()
		fmt.Println("===", err)
		
		// ---
		fmt.Println()
		fmt.Println("---")
		err = errors.Unwrap(err)
		fmt.Println("0004:", err)
		err = errors.Unwrap(err)
		fmt.Println("0005:", err)
		err = errors.Unwrap(err)
		fmt.Println("0006:", err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not err.(perror.BizErr)")
	}
	
	fmt.Println()
	fmt.Println("=========2")
	
	err = perror.BizErrLen
	if bizErr, ok := err.(perror.BizErr); ok {
		fmt.Println("BizErrRequired>>>>>>>>>>err.(perror.BizErr)")
		
		fmt.Println(bizErr)
		fmt.Println(bizErr.GetError())
		bizErr = bizErr.WithError(perror.BizErrDataEmpty)
		bizErr = bizErr.WithError(perror.BizErrTypeNoExist)
		fmt.Println(bizErr)
		err = bizErr.GetError()
		fmt.Println("===", err)
		
		// ---
		fmt.Println()
		fmt.Println("---")
		err = errors.Unwrap(err)
		fmt.Println("0004:", err)
		err = errors.Unwrap(err)
		fmt.Println("0005:", err)
		err = errors.Unwrap(err)
		fmt.Println("0006:", err)
	} else {
		fmt.Println("BizErrRequired>>>>>>>>>>not err.(perror.BizErr)")
	}
}

func TestBizErrDataParse(t *testing.T) {
	fmt.Println(perror.BizErrDataParse)
	
	perror.BizErrDataParse.SetMessage("订单数据解析失败1")
	fmt.Println(perror.BizErrDataParse)
	
	bizErrDataParse := perror.BizErrDataParse.SetMessage("订单数据解析失败2")
	fmt.Println(bizErrDataParse)
}
