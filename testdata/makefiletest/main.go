package main

import (
	"errors"
	"fmt"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/12/13 9:55
 * @Desc:
 */

func main() {
	// --------------- go vet --------------------
	Echo()
	For()
	// ---------------end go vet --------------------

	// --------------- errcheck --------------------
	// errcheck 不通过
	ErrReturn()
	// errcheck 通过
	_, _ = ErrReturn()

	ErrNoReturn()
	// ---------------end errcheck --------------------
}

// --------------- go vet --------------------

// Echo 输出内容
func Echo() {
	fmt.Println("echo function")
}

// For TODO: comments
// 无限循环
func For() {
	for {
	}
}

// ---------------end go vet --------------------

// --------------- errcheck --------------------

// ErrReturn ErrRetur 返回错误
func ErrReturn() (string, error) {
	return "", errors.New("ErrReturn")
}

// ErrNoReturn 不返回错误
func ErrNoReturn() string {
	return ""
}

// ---------------end errcheck --------------------
