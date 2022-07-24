package utils

import (
    "fmt"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/16 上午10:34
 * @Desc:
 */

func TestUniqID(t *testing.T) {
    got := UniqID()
    fmt.Println(got)
}