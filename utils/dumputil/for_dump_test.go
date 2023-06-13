package dumputil

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/10 17:37
 * @Desc:
 */

func TestForIndentSPrintf01(t *testing.T) {
	indentChar := " "
	for i := 0; i < 5; i++ {
		fmt.Println(ForIndentSPrintf(int64(i), indentChar, "我在i"))
		fmt.Println(ForIndentSPrintf(int64(i), indentChar, "我在i:%d层", i))
	}
}
func TestForIndentSPrintf02(t *testing.T) {
	indentChar := ""
	for i := 0; i < 5; i++ {
		fmt.Println(ForIndentSPrintf(int64(i), indentChar, "我在i"))
		fmt.Println(ForIndentSPrintf(int64(i), indentChar, "我在i:%d层", i))
	}
}

func TestForIndentSPrintf03(t *testing.T) {
	indentChar := "\t"
	for i := 0; i < 5; i++ {
		fmt.Println(ForIndentSPrintf(int64(i), indentChar, "我在i"))
		fmt.Println(ForIndentSPrintf(int64(i), indentChar, "我在i:%d层", i))
	}
}
