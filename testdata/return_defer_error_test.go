package testdata

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

func TestReturnDeferError(t *testing.T) {
	ReturnDeferError()
	ReturnDeferError1()
}

func ReturnDeferError() (err error) {
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("ReturnDeferError")
	return
}

func ReturnDeferError1() error {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("ReturnDeferError1")
	return err
}
