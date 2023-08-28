package runtimeutil

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/26 23:13
 * @Desc:
 */

func TestCallersFunc(t *testing.T) {
	var s string
	
	s = CallersFunc(0)
	fmt.Printf("CallersFunc 0: %+v \n", s)
	
	s = CallersFunc(1)
	fmt.Printf("CallersFunc 1: %+v \n", s)
	
	s = CallersFunc(2)
	fmt.Printf("CallersFunc 2: %+v \n", s)
	
	s = CallersFunc(3)
	fmt.Printf("CallersFunc 3: %+v \n", s)
	
}
