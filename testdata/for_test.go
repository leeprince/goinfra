package testdata

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/11/21 14:01
 * @Desc:
 */

func TestForBreak(t *testing.T) {
	parent := []int{1, 2, 3, 4, 5}
	child := []int{1, 2}
	for _, p := range parent {
		var isEq bool
		for _, c := range child {
			if p == c {
				fmt.Println("p==c", c)
				isEq = true
				break
			}
		}
		if isEq {
			continue
		}
		fmt.Println("p!=c", p)
	}
}
