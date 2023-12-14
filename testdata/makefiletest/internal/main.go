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

// Echo TODO: comments
func Echo() {
	fmt.Println("echo function")
}

// For TODO: comments
func For() {
	for {
	}
}

// ---------------end go vet --------------------

// --------------- errcheck --------------------

// ErrReturn TODO: comments
func ErrReturn() (string, error) {
	return "", errors.New("ErrReturn")
}

// ErrNoReturn TODO: comments
func ErrNoReturn() string {
	return ""
}

// ---------------end errcheck --------------------
