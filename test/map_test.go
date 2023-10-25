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

	// 测试时
	m = make(map[string]interface{})

	if m == nil {
		fmt.Println("m nil")
	} else {
		fmt.Println("m not nil")
	}

	if len(m) == 0 {
		fmt.Println("len(m) == 0")
	} else {
		fmt.Println("len(m) != 0")
	}

	m["a"] = 1
	m["b"] = 2
	if m == nil {
		fmt.Println("m nil")
	} else {
		fmt.Println("m not nil")
	}
	if len(m) == 0 {
		fmt.Println("1 len(m) == 0")
	} else {
		fmt.Println("1 len(m) != 0")
	}
}

func TestMapExist(t *testing.T) {
	m := make(map[string]interface{})

	// 测试时
	m["a"] = 1

	fmt.Println(m["a"])

	mvalue, ok := m["a"]
	fmt.Println(mvalue, ok)
}
