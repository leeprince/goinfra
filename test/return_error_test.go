package test

import (
	"errors"
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/18 01:34
 * @Desc:
 */

func TestReturnErr(t *testing.T) {
	err := ReturnErr01()
	fmt.Println("TestReturnErr:", err)
}

func ReturnErr01() (err error) {
	s, err := ReturnErr02()
	if err != nil {
		fmt.Println("ReturnErr02 err:", err)
		return
	}
	fmt.Println(s)
	return
}
func ReturnErr02() (s string, err error) {
	return "", errors.New("ReturnErr02")
}
