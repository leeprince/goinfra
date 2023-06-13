package test

import (
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)

			// 获取panic发生的位置
			_, file, line, ok := runtime.Caller(2)
			if ok {
				fmt.Printf("Panic occurred at %s:%d\n", file, line)
			}
		}
	}()

	// 代码中故意引发panic
	// panic("Something went wrong!")
	panicFunc01()
}

func panicFunc01() {
	// 代码中故意引发panic
	// panic("Something went wrong!")
	panicFunc02()
}

func panicFunc02() {
	// 代码中故意引发panic
	panic("Something went wrong!")
}
