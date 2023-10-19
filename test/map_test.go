package test

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/19 16:31
 * @Desc:
 */

func TestMapNil(t *testing.T) {
	var m map[string]interface{}

	m = make(map[string]interface{})

	if m == nil {
		fmt.Println("m nil")
	} else {
		fmt.Println("m not nil")
	}
}
