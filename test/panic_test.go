package test

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/13 00:19
 * @Desc:
 */

func TestPanic(t *testing.T) {
	panicFunc()
}

func panicFunc() {
	// 发生 panic 的处理方式
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			
			// 断言错误类型
			reconverErr, isError := r.(error)
			if isError {
				fmt.Println("isError reconverErr:", reconverErr)
			}
			
			// 获取panic发生的位置
			_, file, line, ok := runtime.Caller(2)
			if ok {
				fmt.Printf("Panic occurred at %s:%d\n", file, line)
			}
		}
	}()
	
	// 代码中故意引发panic
	// panic("Something went wrong!") // 方式 1
	// panicFunc01() // 方式 2
	panicFunc03() // 方式 3
}

func panicFunc01() {
	// 代码中故意引发panic
	panicFunc02()
}

func panicFunc02() {
	// 代码中故意引发panic
	panic("Something went wrong!")
}

func panicFunc03() {
	// 代码中故意引发panic
	panic(errors.New("panic is errors.New"))
}
